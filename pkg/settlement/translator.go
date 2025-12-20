```go
package settlement

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"financial-bridge/pkg/events"
	"financial-bridge/pkg/ledger"
)

// TransactionIDNamespace is a constant UUID used as a namespace for generating
// deterministic transaction IDs from event IDs. Using a fixed namespace ensures
// that the same event will always produce a transaction with the same ID, which
// is critical for idempotency in event replay scenarios.
var TransactionIDNamespace = uuid.Must(uuid.Parse("7c2c8558-6a7b-4c3d-8b3a-2b5e8d9e7f8a"))

// Translator is responsible for converting business-level events into the
// fundamental, double-entry accounting transactions understood by the ledger.
// This struct is currently stateless, but provides a clear point for extension
// or dependency injection if translation logic becomes more complex (e.g.,
// requiring configuration for account mapping).
type Translator struct{}

// NewTranslator creates a new instance of a Translator.
func NewTranslator() *Translator {
	return &Translator{}
}

// TranslatePaymentSettled converts a PaymentSettled event into a double-entry
// ledger.Transaction. The function is pure and deterministic: for a given event,
// it will always produce the same transaction.
//
// The generated transaction ID is a UUIDv5 derived from the TransactionIDNamespace
// and the event's unique PaymentID. This ensures that if the settlement service
// processes the same event multiple times, it will generate the exact same
// transaction, which the ledger can then handle idempotently (i.e., apply it
// only once).
func (t *Translator) TranslatePaymentSettled(event events.PaymentSettled) (ledger.Transaction, error) {
	if err := validatePaymentSettledEvent(event); err != nil {
		return ledger.Transaction{}, fmt.Errorf("invalid PaymentSettled event: %w", err)
	}

	// Generate a deterministic transaction ID from the payment ID.
	transactionID := uuid.NewSHA1(TransactionIDNamespace, []byte(event.PaymentID.String()))

	debitEntry := ledger.Entry{
		AccountID: event.DebitAccountID,
		Amount:    event.Amount,
		Currency:  event.Currency,
		Direction: ledger.Debit,
	}

	creditEntry := ledger.Entry{
		AccountID: event.CreditAccountID,
		Amount:    event.Amount,
		Currency:  event.Currency,
		Direction: ledger.Credit,
	}

	transaction := ledger.Transaction{
		ID:            transactionID,
		Entries:       []ledger.Entry{debitEntry, creditEntry},
		EffectiveDate: event.SettledAt,
		Metadata: map[string]string{
			"event_id":       event.EventID.String(),
			"payment_id":     event.PaymentID.String(),
			"correlation_id": event.CorrelationID.String(),
			"source_event":   "PaymentSettled",
			"translator":     "settlement.Translator",
		},
	}

	// The ledger package is responsible for validating the transaction's internal
	// consistency (e.g., debits must equal credits). This translator's sole
	// responsibility is the pure transformation from event to transaction.
	return transaction, nil
}

// validatePaymentSettledEvent performs basic sanity checks on the event data
// before attempting to translate it. This ensures that the translator does not
// produce malformed transactions from invalid events.
func validatePaymentSettledEvent(event events.PaymentSettled) error {
	if event.EventID == uuid.Nil {
		return fmt.Errorf("event ID cannot be nil")
	}
	if event.PaymentID == uuid.Nil {
		return fmt.Errorf("payment ID cannot be nil")
	}
	if event.DebitAccountID == "" {
		return fmt.Errorf("debit account ID cannot be empty")
	}
	if event.CreditAccountID == "" {
		return fmt.Errorf("credit account ID cannot be empty")
	}
	if event.DebitAccountID == event.CreditAccountID {
		return fmt.Errorf("debit and credit accounts cannot be the same: %s", event.DebitAccountID)
	}
	if event.Amount.IsNegative() || event.Amount.IsZero() {
		return fmt.Errorf("settlement amount must be positive, got %s", event.Amount.String())
	}
	if event.Currency == "" {
		return fmt.Errorf("currency cannot be empty")
	}
	if event.SettledAt.IsZero() {
		return fmt.Errorf("settled_at timestamp cannot be zero")
	}
	return nil
}
