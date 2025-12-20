```go
package governance

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Change represents a single modification to a parameter within a Ruleset.
// It is the fundamental unit of a governance proposal.
type Change struct {
	// Key is a dot-separated path to the parameter being changed.
	// e.g., "risk.max_position_size"
	Key string
	// Value is the new value for the parameter, represented as a string.
	// The system is responsible for parsing this into the correct type.
	Value string
}

// Proposal is a collection of changes that are proposed, voted on, and
// enacted together as an atomic unit.
type Proposal struct {
	ID      uuid.UUID
	Changes []Change
}

// RiskParameters defines the risk limits for the entire system.
// These values are critical for maintaining capital safety.
type RiskParameters struct {
	// MaxPositionSize is the maximum value of a single position an account can hold.
	MaxPositionSize decimal.Decimal
	// CollateralizationRatio is the minimum ratio of collateral value to debt value
	// before an account is subject to liquidation. (e.g., 1.5 means 150%).
	CollateralizationRatio decimal.Decimal
	// LiquidationPenalty is the percentage fee applied to collateral during a liquidation event,
	// rewarding the liquidator and covering system risk. (e.g., 0.05 means 5%).
	LiquidationPenalty decimal.Decimal
}

// AssetInfo defines the properties of a single supported asset.
type AssetInfo struct {
	Symbol       string
	Decimals     int32
	IsCollateral bool
	IsTradable   bool
}

// SystemParameters defines global operational parameters.
type SystemParameters struct {
	// TransactionFeeBPS is the fee charged on trades, in basis points (1 BPS = 0.01%).
	TransactionFeeBPS uint64
	// MaxOpenOrdersPerAccount is the maximum number of open orders an account can have.
	MaxOpenOrdersPerAccount uint32
	// SystemHalt is a master switch to pause most system activities (e.g., new orders)
	// in case of an emergency, without stopping the entire system.
	SystemHalt bool
}

// Ruleset is a versioned, immutable snapshot of all system governance parameters.
// It represents the single source of truth for system policies at a given time.
// A new Ruleset is created and activated only when a governance proposal is enacted.
type Ruleset struct {
	Version        uint64
	ActivationTime time.Time
	ProposalID     uuid.UUID // The proposal that enacted this ruleset.

	Risk   RiskParameters
	Assets map[string]AssetInfo
	System SystemParameters
}

// clone creates a deep copy of the Ruleset. This is essential for the atomic
// update process, allowing a new version to be built and validated without
// affecting the currently active ruleset.
func (r *Ruleset) clone() *Ruleset {
	newAssets := make(map[string]AssetInfo, len(r.Assets))
	for k, v := range r.Assets {
		newAssets[k] = v // AssetInfo is a struct of simple types, so a direct copy is sufficient.
	}

	return &Ruleset{
		Version:        r.Version,
		ActivationTime: r.ActivationTime,
		ProposalID:     r.ProposalID,
		Risk:           r.Risk,   // RiskParameters is a struct of value types.
		Assets:         newAssets,
		System:         r.System, // SystemParameters is a struct of value types.
	}
}

// Rulebook provides a thread-safe, atomically updatable container for the active Ruleset.
// It ensures that the entire system always references a consistent and valid set of rules.
// It acts as the live, in-memory representation of the system's governance state.
type Rulebook struct {
	mu      sync.RWMutex
	current *Ruleset
	history map[uint64]*Ruleset // For auditability and historical queries.
}

// NewRulebook creates a new Rulebook, initialized with a genesis ruleset.
// The genesis ruleset represents the initial state of the system and must have version 1.
func NewRulebook(genesisRuleset *Ruleset) (*Rulebook, error) {
	if genesisRuleset == nil {
		return nil, fmt.Errorf("genesis ruleset cannot be nil")
	}
	if genesisRuleset.Version != 1 {
		return nil, fmt.Errorf("genesis ruleset must have version 1, got %d", genesisRuleset.Version)
	}

	history := make(map[uint64]*Ruleset)
	history[genesisRuleset.Version] = genesisRuleset

	return &Rulebook{
		current: genesisRuleset,
		history: history,
	}, nil
}

// GetCurrent returns a pointer to the currently active Ruleset.
// This is a read-only operation and is safe for concurrent access from many goroutines.
// The returned Ruleset is immutable and should NOT be modified by callers.
func (rb *Rulebook) GetCurrent() *Ruleset {
	rb.mu.RLock()
	defer rb.mu.RUnlock()
	return rb.current
}

// GetByVersion retrieves a specific version of the Ruleset from history.
// Returns nil if the version is not found. This is useful for auditing and
// replaying historical state.
func (rb *Rulebook) GetByVersion(version uint64) *Ruleset {
	rb.mu.RLock()
	defer rb.mu.RUnlock()
	// The history map itself is only written to under a write lock,
	// so reading it here is safe.
	return rb.history[version]
}

// ApplyUpdate validates a proposal and, if valid, creates a new Ruleset,
// and atomically swaps it in as the current one. This is the sole entrypoint
// for changing the system's governance state. The operation is all-or-nothing;
// a single invalid change in a proposal will cause the entire update to be rejected.
func (rb *Rulebook) ApplyUpdate(proposal *Proposal) (*Ruleset, error) {
	if proposal == nil {
		return nil, fmt.Errorf("proposal cannot be nil")
	}

	rb.mu.Lock()
	defer rb.mu.Unlock()

	// Create a deep copy of the current ruleset to work on.
	// This prevents partial updates from affecting the live system if validation fails.
	newRuleset := rb.current.clone()
	newRuleset.Version++
	newRuleset.ProposalID = proposal.ID

	// Apply changes from the proposal to the new ruleset.
	for _, change := range proposal.Changes {
		if err := applyChange(newRuleset, change); err != nil {
			return nil, fmt.Errorf("failed to apply change '%s=%s': %w", change.Key, change.Value, err)
		}
	}

	// After all changes are applied, perform a holistic validation of the new ruleset.
	// This ensures the combination of parameters is sane and self-consistent.
	if err := validateRuleset(newRuleset); err != nil {
		return nil, fmt.Errorf("new ruleset failed validation: %w", err)
	}

	// The new ruleset is valid. Atomically make it the current one.
	newRuleset.ActivationTime = time.Now().UTC()
	rb.current = newRuleset
	rb.history[newRuleset.Version] = newRuleset

	return newRuleset, nil
}

// applyChange modifies the given ruleset based on a single Change object.
// This function contains the specific logic for each configurable parameter.
// It acts as a strict, explicit gatekeeper for what can be changed, preventing
// arbitrary state modifications.
func applyChange(rs *Ruleset, change Change) error {
	// An explicit switch is deterministic, easy to audit, and safer than reflection.
	switch change.Key {
	// Risk Parameters
	case "risk.max_position_size":
		val, err := decimal.NewFromString(change.Value)
		if err != nil {
			return fmt.Errorf("invalid decimal value for %s: %w", change.Key, err)
		}
		rs.Risk.MaxPositionSize = val
	case "risk.collateralization_ratio":
		val, err := decimal.NewFromString(change.Value)
		if err != nil {
			return fmt.Errorf("invalid decimal value for %s: %w", change.Key, err)
		}
		rs.Risk.CollateralizationRatio = val
	case "risk.liquidation_penalty":
		val, err := decimal.NewFromString(change.Value)
		if err != nil {
			return fmt.Errorf("invalid decimal value for %s: %w", change.Key, err)
		}
		rs.Risk.LiquidationPenalty = val

	// System Parameters
	case "system.transaction_fee_bps":
		var val uint64
		if _, err := fmt.Sscanf(change.Value, "%d", &val); err != nil {
			return fmt.Errorf("invalid integer value for %s: %w", change.Key, err)
		}
		rs.System.TransactionFeeBPS = val
	case "system.max_open_orders_per_account":
		var val uint32
		if _, err := fmt.Sscanf(change.Value, "%d", &val); err != nil {
			return fmt.Errorf("invalid integer value for %s: %w", change.Key, err)
		}
		rs.System.MaxOpenOrdersPerAccount = val
	case "system.halt":
		var val bool
		if _, err := fmt.Sscanf(change.Value, "%t", &val); err != nil {
			return fmt.Errorf("invalid boolean value for %s: %w", change.Key, err)
		}
		rs.System.SystemHalt = val

	// Note: Asset parameter changes would be more complex, likely involving
	// dedicated keys like "asset.add", "asset.update", "asset.remove" with
	// structured (e.g., JSON) values. This is omitted for brevity but would
	// follow the same explicit, validated pattern.

	default:
		return fmt.Errorf("unknown ruleset key: %s", change.Key)
	}
	return nil
}

// validateRuleset performs sanity checks on a newly proposed ruleset.
// This is a critical safety invariant that prevents the system from entering
// a financially unsound or inconsistent state.
func validateRuleset(rs *Ruleset) error {
	// Risk parameter invariants
	if rs.Risk.CollateralizationRatio.LessThanOrEqual(decimal.NewFromInt(1)) {
		return fmt.Errorf("collateralization ratio must be greater than 1.0, got %s", rs.Risk.CollateralizationRatio)
	}
	if rs.Risk.LiquidationPenalty.IsNegative() || rs.Risk.LiquidationPenalty.GreaterThan(decimal.NewFromFloat(0.5)) {
		return fmt.Errorf("liquidation penalty must be between 0.0 and 0.5, got %s", rs.Risk.LiquidationPenalty)
	}
	if rs.Risk.MaxPositionSize.IsNegative() {
		return fmt.Errorf("max position size cannot be negative, got %s", rs.Risk.MaxPositionSize)
	}

	// System parameter invariants
	if rs.System.MaxOpenOrdersPerAccount == 0 {
		return fmt.Errorf("max open orders per account must be greater than 0")
	}

	// Asset invariants
	if len(rs.Assets) == 0 {
		return fmt.Errorf("ruleset must define at least one asset")
	}

	// Cross-parameter invariants can be added here. For example, ensuring
	// that an asset used in a default collateral pool is actually marked IsCollateral.

	return nil
}
### END_OF_FILE_COMPLETED ###
```