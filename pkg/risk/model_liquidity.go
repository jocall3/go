```go
package risk

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/your-org/your-project/pkg/asset"
)

// LiquidityStatus represents the overall liquidity state of the system for a given asset.
type LiquidityStatus string

const (
	// StatusNormal indicates sufficient liquidity to meet all obligations comfortably.
	StatusNormal LiquidityStatus = "NORMAL"
	// StatusWarning indicates liquidity is approaching a cautionary threshold.
	// May trigger alerts to operators to arrange for more liquidity.
	StatusWarning LiquidityStatus = "WARNING"
	// StatusCritical indicates a significant liquidity shortfall is projected.
	// May trigger automated actions like pausing new trade submissions.
	StatusCritical LiquidityStatus = "CRITICAL"
	// StatusHalt indicates liquidity is below the minimum required level to guarantee settlement.
	// This is a fail-safe state that halts settlement processing to prevent defaults.
	StatusHalt LiquidityStatus = "HALT"
)

// LiquidityThresholds defines the boundaries for different liquidity statuses.
// Values represent the required coverage ratio (AvailableBalance / ProjectedObligation).
// For example, a Warning level of 1.5 means liquidity must be at least 150% of obligations.
// Invariant: Halt <= Critical <= Warning.
type LiquidityThresholds struct {
	Halt     decimal.Decimal // e.g., 1.0 (100%) - cannot drop below obligations
	Critical decimal.Decimal // e.g., 1.2 (120%)
	Warning  decimal.Decimal // e.g., 1.5 (150%)
}

// Validate checks if the threshold values are logically consistent.
func (t LiquidityThresholds) Validate() error {
	if t.Halt.GreaterThan(t.Critical) {
		return fmt.Errorf("halt threshold (%s) cannot be greater than critical threshold (%s)", t.Halt, t.Critical)
	}
	if t.Critical.GreaterThan(t.Warning) {
		return fmt.Errorf("critical threshold (%s) cannot be greater than warning threshold (%s)", t.Critical, t.Warning)
	}
	if t.Halt.IsNegative() || t.Critical.IsNegative() || t.Warning.IsNegative() {
		return fmt.Errorf("thresholds cannot be negative")
	}
	return nil
}

// SettlementAccount represents a pool of funds for a specific asset used for settlement.
type SettlementAccount struct {
	ID        uuid.UUID
	AssetID   asset.ID
	Balance   decimal.Decimal
	UpdatedAt time.Time
}

// ProjectedObligation represents the net settlement amount for an asset in a future cycle.
type ProjectedObligation struct {
	AssetID asset.ID
	// NetAmount represents a net outflow. It must be non-negative.
	// Net inflows are treated as a zero obligation for liquidity risk purposes.
	NetAmount  decimal.Decimal
	SettleTime time.Time
}

// LiquidityPosition represents the calculated liquidity state for a single asset.
type LiquidityPosition struct {
	AssetID             asset.ID
	AvailableBalance    decimal.Decimal
	ProjectedObligation decimal.Decimal
	NetPosition         decimal.Decimal // AvailableBalance - ProjectedObligation
	CoverageRatio       decimal.Decimal // AvailableBalance / ProjectedObligation
	Status              LiquidityStatus
	LastAssessedAt      time.Time
}

// LiquidityModel monitors system-wide liquidity across all assets.
// It is responsible for assessing if the system has sufficient funds to meet settlement obligations.
// The model is safe for concurrent use.
type LiquidityModel struct {
	mu                sync.RWMutex
	defaultThresholds LiquidityThresholds
	assetThresholds   map[asset.ID]LiquidityThresholds
	accounts          map[asset.ID]*SettlementAccount
	positions         map[asset.ID]*LiquidityPosition
}

// NewLiquidityModel creates a new liquidity model.
// It requires default thresholds that apply to any asset without specific overrides.
func NewLiquidityModel(defaultThresholds LiquidityThresholds) (*LiquidityModel, error) {
	if err := defaultThresholds.Validate(); err != nil {
		return nil, fmt.Errorf("invalid default thresholds: %w", err)
	}
	return &LiquidityModel{
		defaultThresholds: defaultThresholds,
		assetThresholds:   make(map[asset.ID]LiquidityThresholds),
		accounts:          make(map[asset.ID]*SettlementAccount),
		positions:         make(map[asset.ID]*LiquidityPosition),
	}, nil
}

// SetAssetThresholds allows setting specific, overriding thresholds for a given asset.
func (m *LiquidityModel) SetAssetThresholds(assetID asset.ID, thresholds LiquidityThresholds) error {
	if err := thresholds.Validate(); err != nil {
		return fmt.Errorf("invalid thresholds for asset %s: %w", assetID, err)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.assetThresholds[assetID] = thresholds
	return nil
}

// UpdateAccountBalance updates the balance for a specific settlement account.
// This is typically called upon receiving funds from liquidity providers or after settlement payout.
func (m *LiquidityModel) UpdateAccountBalance(assetID asset.ID, newBalance decimal.Decimal) error {
	if newBalance.IsNegative() {
		return fmt.Errorf("account balance for asset %s cannot be negative: %s", assetID, newBalance)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	account, ok := m.accounts[assetID]
	if !ok {
		account = &SettlementAccount{
			ID:      uuid.New(),
			AssetID: assetID,
		}
		m.accounts[assetID] = account
	}

	account.Balance = newBalance
	account.UpdatedAt = time.Now().UTC()

	return nil
}

// AssessLiquidity evaluates the current liquidity against a set of projected obligations.
// It updates the LiquidityPosition for each affected asset and returns the most severe system-wide status.
// This should be called before committing to a new settlement cycle.
func (m *LiquidityModel) AssessLiquidity(obligations []ProjectedObligation) (LiquidityStatus, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	obligationsMap := make(map[asset.ID]decimal.Decimal)
	for _, ob := range obligations {
		if ob.NetAmount.IsNegative() {
			return "", fmt.Errorf("projected obligation for asset %s cannot be negative: %s", ob.AssetID, ob.NetAmount)
		}
		obligationsMap[ob.AssetID] = obligationsMap[ob.AssetID].Add(ob.NetAmount)
	}

	// Assess all assets with known accounts, even if they have no new obligations.
	for assetID, account := range m.accounts {
		obligationAmount := obligationsMap[assetID] // Defaults to zero if not in map
		m.assessAsset(assetID, account.Balance, obligationAmount)
	}

	// Assess all assets with new obligations, even if they have no pre-existing account.
	// This handles the case where an obligation exists for an asset with a zero balance.
	for assetID, obligationAmount := range obligationsMap {
		if _, ok := m.accounts[assetID]; !ok {
			m.assessAsset(assetID, decimal.Zero, obligationAmount)
		}
	}

	return m.getSystemStatus(), nil
}

// assessAsset is an internal helper to calculate and store the position for a single asset.
// It must be called under a write lock.
func (m *LiquidityModel) assessAsset(assetID asset.ID, balance, obligation decimal.Decimal) {
	thresholds := m.getThresholdsForAsset(assetID)
	position := m.calculatePosition(assetID, balance, obligation, thresholds)
	m.positions[assetID] = position
}

// GetPosition returns the current liquidity position for a specific asset.
func (m *LiquidityModel) GetPosition(assetID asset.ID) (LiquidityPosition, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	position, ok := m.positions[assetID]
	if !ok {
		return LiquidityPosition{}, false
	}
	return *position, true
}

// GetSystemStatus determines the most severe liquidity status across all monitored assets.
// The order of severity is Halt > Critical > Warning > Normal.
func (m *LiquidityModel) GetSystemStatus() LiquidityStatus {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.getSystemStatus()
}

// getSystemStatus is the non-locking version for internal use.
func (m *LiquidityModel) getSystemStatus() LiquidityStatus {
	overallStatus := StatusNormal
	for _, position := range m.positions {
		switch position.Status {
		case StatusHalt:
			return StatusHalt // Halt is the most severe, return immediately.
		case StatusCritical:
			overallStatus = StatusCritical
		case StatusWarning:
			if overallStatus != StatusCritical {
				overallStatus = StatusWarning
			}
		}
	}
	return overallStatus
}

// getThresholdsForAsset retrieves the specific thresholds for an asset, falling back to defaults.
// It must be called under a lock.
func (m *LiquidityModel) getThresholdsForAsset(assetID asset.ID) LiquidityThresholds {
	if thresholds, ok := m.assetThresholds[assetID]; ok {
		return thresholds
	}
	return m.defaultThresholds
}

// calculatePosition is a pure function that computes the liquidity position for a single asset.
func (m *LiquidityModel) calculatePosition(assetID asset.ID, balance, obligation decimal.Decimal, thresholds LiquidityThresholds) *LiquidityPosition {
	pos := &LiquidityPosition{
		AssetID:             assetID,
		AvailableBalance:    balance,
		ProjectedObligation: obligation,
		NetPosition:         balance.Sub(obligation),
		LastAssessedAt:      time.Now().UTC(),
	}

	// If there is no obligation, liquidity is by definition sufficient.
	// This also handles net inflows (represented as zero obligation).
	if obligation.IsZero() || obligation.IsNegative() {
		pos.Status = StatusNormal
		// Coverage ratio is conceptually infinite, represent as zero or a special value if needed.
		// For simplicity, we leave it as zero.
		pos.CoverageRatio = decimal.Zero
		return pos
	}

	// Calculate coverage ratio: Available / Obligation
	pos.CoverageRatio = balance.Div(obligation)

	// Determine status based on thresholds
	if pos.CoverageRatio.LessThan(thresholds.Halt) {
		pos.Status = StatusHalt
	} else if pos.CoverageRatio.LessThan(thresholds.Critical) {
		pos.Status = StatusCritical
	} else if pos.CoverageRatio.LessThan(thresholds.Warning) {
		pos.Status = StatusWarning
	} else {
		pos.Status = StatusNormal
	}

	return pos
}
### END_OF_FILE_COMPLETED ###
```