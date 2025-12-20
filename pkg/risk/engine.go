```go
package risk

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/shopspring/decimal"
	"golang.org/x/exp/slog"

	"github.com/your-org/your-project/pkg/events"
	"github.com/your-org/your-project/pkg/instrument"
	"github.com/your-org/your-project/pkg/portfolio"
)

// Engine is the core risk management component. It subscribes to the event stream
// to update its internal risk models (e.g., exposure) in real-time. It embodies
// the 'price and bound risk continuously' principle by checking limits after every
// state change and halting the system if an invariant is violated.
type Engine struct {
	log *slog.Logger

	// Configuration
	limits *LimitManager

	// State
	mu              sync.RWMutex
	exposureManager *ExposureManager
	systemHalted    bool
	haltReason      string

	// Communication
	eventSubscriber events.Subscriber
	eventPublisher  events.Publisher
	eventChan       <-chan events.Event
	shutdown        chan struct{}
	wg              sync.WaitGroup
}

// Config holds the configuration for the risk engine.
type Config struct {
	Logger          *slog.Logger
	Limits          LimitConfig
	EventSubscriber events.Subscriber
	EventPublisher  events.Publisher
}

// Validate checks if the configuration is valid.
func (c *Config) Validate() error {
	if c.Logger == nil {
		return errors.New("logger is required")
	}
	if c.EventSubscriber == nil {
		return errors.New("event subscriber is required")
	}
	if c.EventPublisher == nil {
		return errors.New("event publisher is required")
	}
	return nil
}

// NewEngine creates and initializes a new risk Engine.
func NewEngine(cfg Config) (*Engine, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid risk engine config: %w", err)
	}

	limitManager, err := NewLimitManager(cfg.Limits)
	if err != nil {
		return nil, fmt.Errorf("failed to create limit manager: %w", err)
	}

	// Subscribe to relevant events. The risk engine is interested in anything
	// that changes financial positions.
	// A wildcard or specific topic subscription would be handled by the event bus implementation.
	eventChan, err := cfg.EventSubscriber.Subscribe(events.TopicTrades)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to trade events: %w", err)
	}

	return &Engine{
		log:             cfg.Logger.With(slog.String("component", "risk_engine")),
		limits:          limitManager,
		exposureManager: NewExposureManager(),
		systemHalted:    false,
		eventSubscriber: cfg.EventSubscriber,
		eventPublisher:  cfg.EventPublisher,
		eventChan:       eventChan,
		shutdown:        make(chan struct{}),
	}, nil
}

// Run starts the risk engine's main processing loop.
// It blocks until the context is canceled.
func (e *Engine) Run(ctx context.Context) error {
	e.log.Info("Starting risk engine")
	e.wg.Add(1)
	defer e.wg.Done()

	for {
		select {
		case event, ok := <-e.eventChan:
			if !ok {
				e.log.Info("Event channel closed, risk engine shutting down")
				return nil
			}
			e.processEvent(ctx, event)
		case <-ctx.Done():
			e.log.Info("Context canceled, risk engine shutting down")
			return ctx.Err()
		case <-e.shutdown:
			e.log.Info("Shutdown signal received, risk engine stopping")
			return nil
		}
	}
}

// Stop gracefully shuts down the risk engine.
func (e *Engine) Stop() {
	e.log.Info("Stopping risk engine")
	close(e.shutdown)
	e.wg.Wait()
	e.log.Info("Risk engine stopped")
}

// processEvent is the main event router. It locks the engine for the duration
// of processing to ensure state transitions are atomic.
func (e *Engine) processEvent(ctx context.Context, event events.Event) {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Invariant: Do not process any events that can alter risk if the system is halted.
	if e.systemHalted {
		e.log.Warn("System is halted, skipping event processing", "event_type", fmt.Sprintf("%T", event))
		return
	}

	e.log.Debug("Processing event", "event_type", fmt.Sprintf("%T", event))

	var err error
	switch ev := event.(type) {
	case *events.TradeExecuted:
		err = e.handleTradeExecuted(ev)
	default:
		// Not an event we are concerned with for risk calculations.
		e.log.Debug("Ignoring irrelevant event type", "type", fmt.Sprintf("%T", ev))
	}

	if err != nil {
		// A processing error is a critical failure. It implies a bug or corrupt data.
		// The only safe action is to halt.
		e.log.Error("Failed to process event, halting system", "error", err, "event", event)
		e.haltSystem(ctx, fmt.Sprintf("event processing failed: %v", err))
		return
	}

	// After every state change, check all limits. This embodies continuous risk bounding.
	e.checkAllLimits(ctx)
}

// handleTradeExecuted updates exposures based on a trade.
// This function assumes it's called within a locked section.
func (e *Engine) handleTradeExecuted(event *events.TradeExecuted) error {
	if err := event.Validate(); err != nil {
		return fmt.Errorf("invalid TradeExecuted event: %w", err)
	}

	baseAsset := event.Instrument.Base()
	quoteAsset := event.Instrument.Quote()
	quoteAmount := event.Price.Mul(event.Quantity)

	// Update buyer's exposure
	// + exposure to base asset
	// - exposure to quote asset
	e.exposureManager.UpdateExposure(event.BuyerID, baseAsset, event.Quantity)
	e.exposureManager.UpdateExposure(event.BuyerID, quoteAsset, quoteAmount.Neg())

	// Update seller's exposure
	// - exposure to base asset
	// + exposure to quote asset
	e.exposureManager.UpdateExposure(event.SellerID, baseAsset, event.Quantity.Neg())
	e.exposureManager.UpdateExposure(event.SellerID, quoteAsset, quoteAmount)

	e.log.Info("Updated exposures for trade",
		"trade_id", event.ID,
		"instrument", event.Instrument.Symbol(),
		"buyer", event.BuyerID,
		"seller", event.SellerID,
	)

	return nil
}

// checkAllLimits evaluates all current exposures against configured limits.
// If any limit is breached, it triggers a system halt.
// This function assumes it's called within a locked section.
func (e *Engine) checkAllLimits(ctx context.Context) {
	// Check global system-level invariants first.
	// For a matched-principal exchange, the net exposure for any asset across all
	// participants must be zero. A non-zero value indicates an internal inconsistency.
	if err := e.limits.CheckSystemNetExposure(e.exposureManager); err != nil {
		e.log.Error("System net exposure invariant violated, halting system", "error", err)
		e.haltSystem(ctx, fmt.Sprintf("system net exposure invariant violated: %v", err))
		return
	}

	// Check other limits, e.g., per-participant or gross exposure limits.
	// (This is where more complex limit checks would be added)
}

// haltSystem puts the system into a safe, non-trading state.
// This is the primary "fail-closed" mechanism. The action is idempotent.
// This function assumes it's called within a locked section.
func (e *Engine) haltSystem(ctx context.Context, reason string) {
	if e.systemHalted {
		return // Already halted
	}

	e.systemHalted = true
	e.haltReason = reason
	e.log.Error("RISK ENGINE HALTING SYSTEM", "reason", reason)

	haltEvent := &events.SystemHalt{
		Timestamp: time.Now().UTC(),
		Reason:    reason,
	}

	if err := e.eventPublisher.Publish(ctx, haltEvent); err != nil {
		// This is a critical failure. If we can't even announce the halt,
		// the system is in a dangerously undefined state. Log with max severity.
		e.log.Error("CRITICAL: FAILED TO PUBLISH SYSTEM HALT EVENT", "error", err)
	}
}

// IsHalted returns true if the risk engine has halted the system.
func (e *Engine) IsHalted() (bool, string) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.systemHalted, e.haltReason
}

// GetExposure returns the current exposure for a given participant and asset.
func (e *Engine) GetExposure(participantID portfolio.ParticipantID, asset instrument.Asset) (decimal.Decimal, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if e.exposureManager == nil {
		return decimal.Zero, errors.New("exposure manager not initialized")
	}
	return e.exposureManager.GetExposure(participantID, asset), nil
}

// --- Helper Components ---

// ExposureManager tracks exposure for each participant and asset.
// It is not thread-safe; synchronization is handled by the Engine.
type ExposureManager struct {
	// exposures maps: ParticipantID -> AssetID -> Net Exposure
	exposures map[portfolio.ParticipantID]map[instrument.Asset]decimal.Decimal
}

// NewExposureManager creates a new exposure manager.
func NewExposureManager() *ExposureManager {
	return &ExposureManager{
		exposures: make(map[portfolio.ParticipantID]map[instrument.Asset]decimal.Decimal),
	}
}

// UpdateExposure adjusts the exposure for a participant and asset by a given delta.
func (em *ExposureManager) UpdateExposure(participantID portfolio.ParticipantID, asset instrument.Asset, delta decimal.Decimal) {
	if _, ok := em.exposures[participantID]; !ok {
		em.exposures[participantID] = make(map[instrument.Asset]decimal.Decimal)
	}
	currentExposure := em.exposures[participantID][asset]
	em.exposures[participantID][asset] = currentExposure.Add(delta)
}

// GetExposure retrieves the exposure for a participant and asset.
func (em *ExposureManager) GetExposure(participantID portfolio.ParticipantID, asset instrument.Asset) decimal.Decimal {
	if pExposures, ok := em.exposures[participantID]; ok {
		if exposure, ok := pExposures[asset]; ok {
			return exposure
		}
	}
	return decimal.Zero
}

// LimitManager holds and checks risk limits.
// It is not thread-safe; synchronization is handled by the Engine.
type LimitManager struct {
	// systemNetExposureLimits defines the maximum allowed net exposure for the entire
	// system for a given asset. For a fully-funded, matched-principal system, this
	// should always be zero. We can allow a small epsilon for dust.
	systemNetExposureLimits map[instrument.Asset]decimal.Decimal
}

// LimitConfig defines the risk limits for the engine.
type LimitConfig struct {
	// GlobalNetExposureLimits maps an asset to its maximum allowed net system exposure.
	// A value of "0" means the system must be perfectly balanced for that asset.
	GlobalNetExposureLimits map[instrument.Asset]string
}

// NewLimitManager creates a new limit manager from a config.
func NewLimitManager(cfg LimitConfig) (*LimitManager, error) {
	limits := make(map[instrument.Asset]decimal.Decimal)
	for asset, limitStr := range cfg.GlobalNetExposureLimits {
		limit, err := decimal.NewFromString(limitStr)
		if err != nil {
			return nil, fmt.Errorf("invalid decimal limit for asset %s: %w", asset, err)
		}
		if limit.IsNegative() {
			return nil, fmt.Errorf("limit for asset %s cannot be negative", asset)
		}
		limits[asset] = limit
	}
	return &LimitManager{
		systemNetExposureLimits: limits,
	}, nil
}

// CheckSystemNetExposure verifies that the total net exposure for each asset
// across all participants is within the configured limits. This is a powerful
// invariant check for system integrity.
func (lm *LimitManager) CheckSystemNetExposure(em *ExposureManager) error {
	netSystemExposure := make(map[instrument.Asset]decimal.Decimal)

	// Calculate the total net exposure for each asset across all participants.
	for _, participantExposures := range em.exposures {
		for asset, exposure := range participantExposures {
			currentTotal := netSystemExposure[asset]
			netSystemExposure[asset] = currentTotal.Add(exposure)
		}
	}

	// Check each asset's total net exposure against its limit.
	for asset, totalExposure := range netSystemExposure {
		limit, ok := lm.systemNetExposureLimits[asset]
		if !ok {
			// If no limit is configured, we assume no risk is allowed.
			// This is a "fail-closed" approach to configuration.
			limit = decimal.Zero
		}

		// The system's net position in any asset should not exceed the limit.
		if totalExposure.Abs().Cmp(limit) > 0 {
			return fmt.Errorf("asset %s: net exposure %s exceeds limit %s", asset, totalExposure.String(), limit.String())
		}
	}

	return nil
}
### END_OF_FILE_COMPLETED ###
```