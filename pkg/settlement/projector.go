```go
package settlement

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	// In a real system, these would point to actual packages within the project.
	// They are defined conceptually in comments at the end of the file for clarity.
	"your_project/pkg/events"
	"your_project/pkg/ledger"
	"your_project/pkg/money"
)

// EventHandler defines the function signature for translating a specific event
// into a ledger transaction. If an event does not result in a transaction,
// it should return (nil, nil). An error indicates a deterministic failure
// in translation, which is considered a critical, system-halting event.
type EventHandler func(ctx context.Context, event events.Event) (*ledger.Transaction, error)

// Projector is the core of the settlement layer's "slow path".
// It consumes events from a durable log in a strict, sequential order.
// For each event, it deterministically projects a corresponding balanced
// ledger transaction. The Projector is the sole writer to the ledger,
// ensuring that the ledger's state is a pure function of the event history.
// This sequential, deterministic process is fundamental to the system's
// auditability and correctness. It is designed to be "boring" to regulators
// and "slow" to attackers.
type Projector struct {
	ledger      ledger.Ledger
	eventSource events.Source
	handlers    map[events.Type]EventHandler
	logger      *log.Logger
}

// NewProjector creates and initializes a new Projector.
// It requires a ledger to write to and an event source to read from.
func NewProjector(lg ledger.Ledger, src events.Source, logger *log.Logger) *Projector {
	return &Projector{
		ledger:      lg,
		eventSource: src,
		handlers:    make(map[events.Type]EventHandler),
		logger:      logger,
	}
}

// RegisterHandler associates an event type with a specific translation function.
// This should be called during initialization to build the projector's logic.
// Attempting to register a handler for an already-registered event type will panic,
// as this indicates a developer error during setup.
func (p *Projector) RegisterHandler(eventType events.Type, handler EventHandler) {
	if _, exists := p.handlers[eventType]; exists {
		panic(fmt.Sprintf("settlement.Projector: handler already registered for event type %s", eventType))
	}
	p.handlers[eventType] = handler
}

// Run starts the projector's main processing loop.
// It will continuously fetch events and process them one by one.
// The loop will run until the provided context is canceled.
// Any error encountered during the processing of an event is considered
// a fatal system error, as it implies a breakdown in deterministic state
// transition. In such a case, the projector will log the error and halt.
// This "fail-closed" semantic is crucial for maintaining capital safety.
func (p *Projector) Run(ctx context.Context) error {
	p.logger.Println("Settlement Projector started. Awaiting events...")
	defer p.logger.Println("Settlement Projector stopped.")

	for {
		// The context is checked at the top of the loop to ensure timely shutdown.
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Non-blocking check, proceed to fetch event.
		}

		// Fetch the next event. This call should block until an event is available
		// or the context is canceled.
		event, err := p.eventSource.Next(ctx)
		if err != nil {
			// If the error is due to context cancellation, it's a clean shutdown.
			if err == context.Canceled || err == context.DeadlineExceeded {
				return err
			}
			// Any other error from the event source is considered fatal.
			p.logger.Printf("CRITICAL: Failed to fetch next event: %v. Halting.", err)
			return fmt.Errorf("unrecoverable error from event source: %w", err)
		}

		p.logger.Printf("Processing event ID %s, Type %s", event.ID(), event.Type())

		// Translate the event into a transaction.
		tx, err := p.translate(ctx, event)
		if err != nil {
			// A translation error is a programming error or a corrupted event.
			// It violates the deterministic projection guarantee. Halt immediately.
			p.logger.Printf("CRITICAL: Failed to translate event %s: %v. Halting.", event.ID(), err)
			return fmt.Errorf("unrecoverable translation error for event %s: %w", event.ID(), err)
		}

		// Some events may not produce a transaction. This is a valid outcome.
		if tx == nil {
			p.logger.Printf("Event %s did not produce a transaction. Continuing.", event.ID())
			continue
		}

		// Invariant: All generated transactions MUST be balanced.
		// A failure here is a critical bug in an event handler.
		if !tx.IsBalanced() {
			p.logger.Printf("CRITICAL: Handler for event type %s produced an unbalanced transaction for event %s. Halting.", event.Type(), event.ID())
			return fmt.Errorf("unrecoverable error: unbalanced transaction generated for event %s", event.ID())
		}

		// Post the balanced, deterministic transaction to the ledger.
		err = p.ledger.PostTransaction(ctx, tx)
		if err != nil {
			// If the ledger rejects a valid, balanced transaction, it implies a
			// severe state inconsistency (e.g., violation of a ledger-level invariant
			// like account existence or non-negative balance). This is unrecoverable.
			p.logger.Printf("CRITICAL: Ledger rejected transaction for event %s: %v. Halting.", event.ID(), err)
			return fmt.Errorf("unrecoverable error: ledger rejected transaction for event %s: %w", event.ID(), err)
		}

		p.logger.Printf("Successfully posted transaction %s for event %s", tx.ID, event.ID())
	}
}

// translate finds the appropriate handler for the event and executes it.
func (p *Projector) translate(ctx context.Context, event events.Event) (*ledger.Transaction, error) {
	handler, found := p.handlers[event.Type()]
	if !found {
		// Receiving an event for which no handler is registered is a critical failure.
		// It means the system's event vocabulary has diverged from its settlement logic.
		return nil, fmt.Errorf("no handler registered for event type %s", event.Type())
	}
	return handler(ctx, event)
}

// --- Event Handlers ---
// The following are examples of event handlers. In a real application, these would
// be more complex and could live in their own files, registered with the projector
// during application setup.

// HandleDepositCompleted translates a deposit completion event into a transaction
// that moves funds from a bank's omnibus account to a user's account.
func HandleDepositCompleted(ctx context.Context, event events.Event) (*ledger.Transaction, error) {
	payload, ok := event.Payload().(events.DepositCompletedPayload)
	if !ok {
		return nil, fmt.Errorf("invalid payload type for DepositCompleted event")
	}

	// Invariant: Amount must be positive.
	if payload.Amount.IsNegative() || payload.Amount.IsZero() {
		return nil, fmt.Errorf("deposit amount must be positive, got %s", payload.Amount)
	}

	tx := &ledger.Transaction{
		ID:            uuid.New(),
		Timestamp:     time.Now().UTC(),
		CorrelationID: event.ID().String(),
		Entries: []ledger.Entry{
			// Debit the external settlement account (e.g., funds received from bank).
			{
				AccountID: ledger.ExternalSettlementAccountID,
				Amount:    payload.Amount,
				Direction: ledger.Debit,
			},
			// Credit the user's internal account.
			{
				AccountID: payload.AccountID,
				Amount:    payload.Amount,
				Direction: ledger.Credit,
			},
		},
	}

	return tx, nil
}

// HandleWithdrawalCompleted translates a withdrawal completion event into a transaction
// that moves funds from a user's account to an external settlement account.
func HandleWithdrawalCompleted(ctx context.Context, event events.Event) (*ledger.Transaction, error) {
	payload, ok := event.Payload().(events.WithdrawalCompletedPayload)
	if !ok {
		return nil, fmt.Errorf("invalid payload type for WithdrawalCompleted event")
	}

	// Invariant: Amount must be positive.
	if payload.Amount.IsNegative() || payload.Amount.IsZero() {
		return nil, fmt.Errorf("withdrawal amount must be positive, got %s", payload.Amount)
	}

	tx := &ledger.Transaction{
		ID:            uuid.New(),
		Timestamp:     time.Now().UTC(),
		CorrelationID: event.ID().String(),
		Entries: []ledger.Entry{
			// Debit the user's internal account.
			{
				AccountID: payload.AccountID,
				Amount:    payload.Amount,
				Direction: ledger.Debit,
			},
			// Credit the external settlement account (e.g., funds sent to bank).
			{
				AccountID: ledger.ExternalSettlementAccountID,
				Amount:    payload.Amount,
				Direction: ledger.Credit,
			},
		},
	}

	return tx, nil
}

// HandleInternalTransferCompleted translates a transfer between two internal users.
func HandleInternalTransferCompleted(ctx context.Context, event events.Event) (*ledger.Transaction, error) {
	payload, ok := event.Payload().(events.InternalTransferCompletedPayload)
	if !ok {
		return nil, fmt.Errorf("invalid payload type for InternalTransferCompleted event")
	}

	// Invariant: Amount must be positive.
	if payload.Amount.IsNegative() || payload.Amount.IsZero() {
		return nil, fmt.Errorf("transfer amount must be positive, got %s", payload.Amount)
	}

	// Invariant: Source and destination accounts must be different.
	if payload.FromAccountID == payload.ToAccountID {
		return nil, fmt.Errorf("source and destination accounts cannot be the same")
	}

	tx := &ledger.Transaction{
		ID:            uuid.New(),
		Timestamp:     time.Now().UTC(),
		CorrelationID: event.ID().String(),
		Entries: []ledger.Entry{
			// Debit the sender's account.
			{
				AccountID: payload.FromAccountID,
				Amount:    payload.Amount,
				Direction: ledger.Debit,
			},
			// Credit the receiver's account.
			{
				AccountID: payload.ToAccountID,
				Amount:    payload.Amount,
				Direction: ledger.Credit,
			},
		},
	}

	return tx, nil
}

// HandleFeeCharged translates a fee event into a transaction that moves funds
// from a user's account to the company's revenue account.
func HandleFeeCharged(ctx context.Context, event events.Event) (*ledger.Transaction, error) {
	payload, ok := event.Payload().(events.FeeChargedPayload)
	if !ok {
		return nil, fmt.Errorf("invalid payload type for FeeCharged event")
	}

	// Invariant: Fee amount must be positive.
	if payload.Amount.IsNegative() || payload.Amount.IsZero() {
		return nil, fmt.Errorf("fee amount must be positive, got %s", payload.Amount)
	}

	tx := &ledger.Transaction{
		ID:            uuid.New(),
		Timestamp:     time.Now().UTC(),
		CorrelationID: event.ID().String(),
		Entries: []ledger.Entry{
			// Debit the user's account for the fee.
			{
				AccountID: payload.AccountID,
				Amount:    payload.Amount,
				Direction: ledger.Debit,
			},
			// Credit the company's revenue account.
			{
				AccountID: ledger.RevenueAccountID,
				Amount:    payload.Amount,
				Direction: ledger.Credit,
			},
		},
	}

	return tx, nil
}

/*
--- Conceptual Package Definitions ---

These comments illustrate the expected structure of the packages this file depends on.

// In pkg/events/source.go
package events
type Source interface {
	Next(ctx context.Context) (Event, error)
}

// In pkg/events/event.go
package events
type Type string
type Event interface {
	ID() uuid.UUID
	Type() Type
	Timestamp() time.Time
	Payload() interface{}
}

// In pkg/events/payloads.go
package events
type DepositCompletedPayload struct {
	AccountID ledger.AccountID
	Amount    money.Money
}
type WithdrawalCompletedPayload struct {
	AccountID ledger.AccountID
	Amount    money.Money
}
type InternalTransferCompletedPayload struct {
	FromAccountID ledger.AccountID
	ToAccountID   ledger.AccountID
	Amount        money.Money
}
type FeeChargedPayload struct {
	AccountID ledger.AccountID
	Amount    money.Money
	Reason    string
}

// In pkg/ledger/ledger.go
package ledger
type Ledger interface {
	PostTransaction(ctx context.Context, tx *Transaction) error
}
var (
	ExternalSettlementAccountID = AccountID("settlement:external:omnibus")
	RevenueAccountID            = AccountID("revenue:internal:fees")
)

// In pkg/ledger/transaction.go
package ledger
type Transaction struct {
	ID            uuid.UUID
	Timestamp     time.Time
	CorrelationID string // Links back to the event ID
	Entries       []Entry
}
func (t *Transaction) IsBalanced() bool { ... }

// In pkg/ledger/entry.go
package ledger
type AccountID string
type Direction int
const (
	Debit Direction = iota
	Credit
)
type Entry struct {
	AccountID AccountID
	Amount    money.Money
	Direction Direction
}

// In pkg/money/money.go
package money
// A struct that safely handles monetary values, avoiding floating-point arithmetic.
// Typically implemented with an int64 representing the smallest currency unit (e.g., cents).
type Money struct { ... }
func (m Money) IsNegative() bool { ... }
func (m Money) IsZero() bool { ... }
func (m Money) String() string { ... }

*/
### END_OF_FILE_COMPLETED ###
```