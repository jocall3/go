```go
package events

import (
	"time"

	"github.com/google/uuid"
)

// This file defines the event structs that model the lifecycle of a payment.
// These events represent an explicit state machine, where each event is an
// immutable fact that drives the payment from one state to the next. This
// event-sourced approach ensures full auditability and replayability, which are
// critical for a financial system. Each event is a granular, self-contained
// record of a specific business occurrence.

// Event is the interface that all system events must implement.
// It provides a standard way to access common event metadata.
type Event interface {
	// Header returns the common event header.
	Header() *EventHeader
	// EventType returns a string identifier for the event type.
	EventType() string
}

// EventHeader contains common metadata for all events. It is intended to be
// embedded in specific event structs.
type EventHeader struct {
	EventID     uuid.UUID `json:"eventId"`
	EventType   string    `json:"eventType"`
	AggregateID uuid.UUID `json:"aggregateId"` // The ID of the entity this event pertains to, e.g., PaymentID.
	Version     int       `json:"version"`     // The version of the aggregate after this event is applied.
	Timestamp   time.Time `json:"timestamp"`   // The UTC timestamp when the event was created.
}

// Header returns the event header itself.
func (h *EventHeader) Header() *EventHeader {
	return h
}

// Amount represents a monetary value in its smallest unit (e.g., cents).
// This avoids floating-point arithmetic errors, a critical invariant for financial safety.
type Amount struct {
	Value    int64  `json:"value"`    // The amount in the smallest currency unit (e.g., cents for USD).
	Currency string `json:"currency"` // The ISO 4217 currency code (e.g., "USD", "EUR").
}

// --- Payment Event Constants ---

const (
	PaymentInitiatedEvent                 = "payment.initiated"
	PaymentValidationSucceededEvent       = "payment.validation.succeeded"
	PaymentValidationFailedEvent          = "payment.validation.failed"
	FundsReservationInitiatedEvent        = "payment.funds.reservation.initiated"
	FundsReservedEvent                    = "payment.funds.reserved"
	FundsReservationFailedEvent           = "payment.funds.reservation.failed"
	CreditTransferInitiatedEvent          = "payment.credit.transfer.initiated"
	CreditTransferSucceededEvent          = "payment.credit.transfer.succeeded"
	CreditTransferFailedEvent             = "payment.credit.transfer.failed"
	FundsReservationReleaseInitiatedEvent = "payment.funds.reservation.release.initiated"
	FundsReservationReleasedEvent         = "payment.funds.reservation.released"
	PaymentSettledEvent                   = "payment.settled"
	PaymentFailedEvent                    = "payment.failed"
)

// --- Payment Lifecycle Events ---

// PaymentInitiated is the first event in a payment's lifecycle.
// It signifies that a request to make a payment has been received and recorded.
// This event contains all the initial data required to process the payment.
type PaymentInitiated struct {
	EventHeader
	PaymentID         uuid.UUID `json:"paymentId"`
	DebtorAccountID   uuid.UUID `json:"debtorAccountId"`
	CreditorAccountID uuid.UUID `json:"creditorAccountId"`
	Amount            Amount    `json:"amount"`
	IdempotencyKey    string    `json:"idempotencyKey"`
	Reference         string    `json:"reference"`
}

// EventType returns the constant type for PaymentInitiated.
func (e PaymentInitiated) EventType() string {
	return PaymentInitiatedEvent
}

// PaymentValidationSucceeded indicates that the initial payment data has passed
// all preliminary business rule checks (e.g., valid accounts, positive amount).
type PaymentValidationSucceeded struct {
	EventHeader
}

// EventType returns the constant type for PaymentValidationSucceeded.
func (e PaymentValidationSucceeded) EventType() string {
	return PaymentValidationSucceededEvent
}

// PaymentValidationFailed indicates that the initial payment data failed
// preliminary business rule checks. This is a terminal state for the payment.
type PaymentValidationFailed struct {
	EventHeader
	Reason string `json:"reason"`
}

// EventType returns the constant type for PaymentValidationFailed.
func (e PaymentValidationFailed) EventType() string {
	return PaymentValidationFailedEvent
}

// FundsReservationInitiated indicates the system is attempting to reserve
// funds from the debtor's account. This is a transient state.
type FundsReservationInitiated struct {
	EventHeader
}

// EventType returns the constant type for FundsReservationInitiated.
func (e FundsReservationInitiated) EventType() string {
	return FundsReservationInitiatedEvent
}

// FundsReserved indicates that the required amount has been successfully
// held in the debtor's account, ensuring it is available for settlement.
type FundsReserved struct {
	EventHeader
	ReservationID uuid.UUID `json:"reservationId"`
}

// EventType returns the constant type for FundsReserved.
func (e FundsReserved) EventType() string {
	return FundsReservedEvent
}

// FundsReservationFailed indicates a failure to reserve funds, typically
// due to insufficient balance. This is a terminal state for the payment.
type FundsReservationFailed struct {
	EventHeader
	Reason string `json:"reason"` // e.g., "INSUFFICIENT_FUNDS"
}

// EventType returns the constant type for FundsReservationFailed.
func (e FundsReservationFailed) EventType() string {
	return FundsReservationFailedEvent
}

// CreditTransferInitiated indicates the system is attempting to credit
// the creditor's account. This is a transient state.
type CreditTransferInitiated struct {
	EventHeader
}

// EventType returns the constant type for CreditTransferInitiated.
func (e CreditTransferInitiated) EventType() string {
	return CreditTransferInitiatedEvent
}

// CreditTransferSucceeded indicates the creditor's account has been successfully
// credited. The payment is now committed and moving towards final settlement.
type CreditTransferSucceeded struct {
	EventHeader
	CreditTransactionID uuid.UUID `json:"creditTransactionId"`
}

// EventType returns the constant type for CreditTransferSucceeded.
func (e CreditTransferSucceeded) EventType() string {
	return CreditTransferSucceededEvent
}

// CreditTransferFailed indicates a failure to credit the creditor's account.
// This triggers a compensating action to release the reserved funds.
type CreditTransferFailed struct {
	EventHeader
	Reason string `json:"reason"` // e.g., "CREDITOR_ACCOUNT_CLOSED", "REGULATORY_BLOCK"
}

// EventType returns the constant type for CreditTransferFailed.
func (e CreditTransferFailed) EventType() string {
	return CreditTransferFailedEvent
}

// FundsReservationReleaseInitiated indicates the system is attempting to
// release a previously made funds reservation as a compensating action.
type FundsReservationReleaseInitiated struct {
	EventHeader
	ReservationID uuid.UUID `json:"reservationId"`
}

// EventType returns the constant type for FundsReservationReleaseInitiated.
func (e FundsReservationReleaseInitiated) EventType() string {
	return FundsReservationReleaseInitiatedEvent
}

// FundsReservationReleased indicates a previously made funds reservation
// has been successfully released. The payment is now moving to a failed state.
type FundsReservationReleased struct {
	EventHeader
	ReservationID uuid.UUID `json:"reservationId"`
}

// EventType returns the constant type for FundsReservationReleased.
func (e FundsReservationReleased) EventType() string {
	return FundsReservationReleasedEvent
}

// PaymentSettled is the final success event in a payment's lifecycle.
// It confirms that funds have been debited from the debtor (converting the
// reservation to a final debit) and credited to the creditor, finalizing the transaction.
type PaymentSettled struct {
	EventHeader
	DebitTransactionID  uuid.UUID `json:"debitTransactionId"`
	CreditTransactionID uuid.UUID `json:"creditTransactionId"`
}

// EventType returns the constant type for PaymentSettled.
func (e PaymentSettled) EventType() string {
	return PaymentSettledEvent
}

// PaymentFailed is a terminal event indicating that the payment could not be
// completed. This event is published after all necessary compensating actions
// (like releasing funds) have been successfully executed.
type PaymentFailed struct {
	EventHeader
	FinalReason string `json:"finalReason"`
}

// EventType returns the constant type for PaymentFailed.
func (e PaymentFailed) EventType() string {
	return PaymentFailedEvent
}
### END_OF_FILE_COMPLETED ###
```