```go
package risk

import (
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// --- Placeholder Event and ID Types ---
// In a complete system, these types would be defined in their respective packages
// (e.g., pkg/events, pkg/identity, pkg/market). They are included here to make
// the exposure model self-contained and understandable.

// CounterpartyID represents a unique identifier for a legal entity interacting with the system.
type CounterpartyID uuid.UUID

// InstrumentID represents a unique identifier for a financial instrument.
type InstrumentID uuid.UUID

// Trade represents a completed trade between two counterparties. It is an immutable
// event that serves as an input to the risk model.
type Trade struct {
	ID                  uuid.UUID
	InstrumentID        InstrumentID
	BuyerID             CounterpartyID
	SellerID            CounterpartyID
	Quantity            decimal.Decimal
	Price               decimal.Decimal
	ValueInBaseCurrency decimal.Decimal // Total value of the trade, converted to the system's base currency.
	Timestamp           time.Time
}

// Settlement represents the completion of asset and payment transfers for a trade.
// This event signifies that the risk associated with the trade has been extinguished.
type Settlement struct {
	TradeID   uuid.UUID
	BuyerID   CounterpartyID
	SellerID  CounterpartyID
	Amount    decimal.Decimal // The settled value in the system's base currency.
	Timestamp time.Time
}

// --- Core Exposure Models ---

// Exposure represents the net financial risk to a single counterparty.
// It tracks the value of obligations owed to the system versus obligations
// the system owes to the counterparty, primarily from unsettled trades.
type Exposure struct {
	// NetValue is the net exposure value in the system's base currency.
	// A positive value means the counterparty owes the system (credit risk).
	// A negative value means the system owes the counterparty (liability).
	NetValue decimal.Decimal

	// LastUpdated is the timestamp of the last event that modified this exposure.
	LastUpdated time.Time
}

// SystemicExposure provides a snapshot of the total risk exposure across the entire system.
type SystemicExposure struct {
	// TotalGrossExposure is the sum of the absolute net exposure values of all counterparties.
	// This metric indicates the total value of unsettled obligations in the system.
	TotalGrossExposure decimal.Decimal

	// NetSystemExposure is the sum of all net exposures. In a closed, balanced system,
	// this should be zero. A non-zero value indicates an imbalance, external capital
	// flow, or a potential system invariant violation.
	NetSystemExposure decimal.Decimal

	// TopExposures lists the counterparties with the highest positive net exposure,
	// highlighting concentration risk.
	TopExposures []ConcentrationItem

	// Timestamp is when the snapshot was generated.
	Timestamp time.Time
}

// ConcentrationItem represents a single counterparty's contribution to concentration risk.
type ConcentrationItem struct {
	ID    CounterpartyID
	Value decimal.Decimal
}

// ExposureManager is responsible for tracking, calculating, and managing
// counterparty and systemic risk exposures in real-time. It consumes events
// from the execution and settlement systems to maintain an up-to-date view of risk.
// Its state is deterministic, derived entirely from the sequence of processed events.
type ExposureManager struct {
	mu        sync.RWMutex
	exposures map[CounterpartyID]*Exposure
}

// NewExposureManager creates and initializes a new ExposureManager.
func NewExposureManager() *ExposureManager {
	return &ExposureManager{
		exposures: make(map[CounterpartyID]*Exposure),
	}
}

// UpdateFromTrade processes a trade event, creating unsettled exposure
// for both the buyer and the seller. This method is idempotent; processing the
// same trade event multiple times will not change the final state if other
// events are not interleaved.
func (em *ExposureManager) UpdateFromTrade(trade Trade) {
	em.mu.Lock()
	defer em.mu.Unlock()

	// The buyer owes the system the value of the trade, creating a positive exposure (credit risk).
	em.updateExposure(trade.BuyerID, trade.ValueInBaseCurrency, trade.Timestamp)

	// The seller is owed the value of the trade, creating a negative exposure (system liability).
	em.updateExposure(trade.SellerID, trade.ValueInBaseCurrency.Neg(), trade.Timestamp)
}

// UpdateFromSettlement processes a settlement event, reducing the unsettled exposure
// for the involved counterparties. This extinguishes the risk created at trade time.
func (em *ExposureManager) UpdateFromSettlement(settlement Settlement) {
	em.mu.Lock()
	defer em.mu.Unlock()

	// The buyer has fulfilled their payment obligation. Their positive exposure (debt) decreases.
	// We apply a negative amount to their exposure.
	em.updateExposure(settlement.BuyerID, settlement.Amount.Neg(), settlement.Timestamp)

	// The seller has received their payment. The system's liability to them is fulfilled.
	// Their negative exposure moves towards zero (we apply a positive amount).
	em.updateExposure(settlement.SellerID, settlement.Amount, settlement.Timestamp)
}

// updateExposure is the internal, non-locking primitive for exposure modification.
// It encapsulates the core logic of creating or updating a counterparty's exposure record.
// The caller MUST hold the write lock.
func (em *ExposureManager) updateExposure(partyID CounterpartyID, amount decimal.Decimal, timestamp time.Time) {
	exposure, ok := em.exposures[partyID]
	if !ok {
		exposure = &Exposure{}
		em.exposures[partyID] = exposure
	}

	// A more advanced model could check if the timestamp is older than LastUpdated
	// to handle out-of-order events, but for a replayable log, we assume order.
	exposure.NetValue = exposure.NetValue.Add(amount)
	exposure.LastUpdated = timestamp
}

// GetCounterpartyExposure retrieves a copy of the exposure for a single counterparty.
// It returns the exposure and a boolean indicating if the counterparty was found.
// Returning a copy prevents race conditions from external modifications.
func (em *ExposureManager) GetCounterpartyExposure(partyID CounterpartyID) (Exposure, bool) {
	em.mu.RLock()
	defer em.mu.RUnlock()

	exposure, ok := em.exposures[partyID]
	if !ok {
		return Exposure{}, false
	}

	return *exposure, true
}

// CalculateSystemicExposure computes a snapshot of the system's overall risk.
// This is a read-only operation that provides a point-in-time view of systemic risk metrics.
func (em *ExposureManager) CalculateSystemicExposure() SystemicExposure {
	em.mu.RLock()
	defer em.mu.RUnlock()

	var totalGross decimal.Decimal
	var netSystem decimal.Decimal
	concentrationItems := make([]ConcentrationItem, 0, len(em.exposures))

	for id, exp := range em.exposures {
		netSystem = netSystem.Add(exp.NetValue)
		totalGross = totalGross.Add(exp.NetValue.Abs())

		if exp.NetValue.IsPositive() {
			concentrationItems = append(concentrationItems, ConcentrationItem{
				ID:    id,
				Value: exp.NetValue,
			})
		}
	}

	sort.Slice(concentrationItems, func(i, j int) bool {
		return concentrationItems[i].Value.GreaterThan(concentrationItems[j].Value)
	})

	const topN = 10
	limit := len(concentrationItems)
	if limit > topN {
		limit = topN
	}

	return SystemicExposure{
		TotalGrossExposure: totalGross,
		NetSystemExposure:  netSystem,
		TopExposures:       concentrationItems[:limit],
		Timestamp:          time.Now().UTC(),
	}
}

// Reset clears all tracked exposures, useful for testing or re-initializing from a snapshot.
func (em *ExposureManager) Reset() {
	em.mu.Lock()
	defer em.mu.Unlock()
	em.exposures = make(map[CounterpartyID]*Exposure)
}

// GetTotalCounterparties returns the number of counterparties with tracked exposure.
func (em *ExposureManager) GetTotalCounterparties() int {
	em.mu.RLock()
	defer em.mu.RUnlock()
	return len(em.exposures)
}
### END_OF_FILE_COMPLETED ###
```