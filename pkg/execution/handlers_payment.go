```go
package execution

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/Tender-Services/bridge/pkg/domain/account"
	"github.com/Tender-Services/bridge/pkg/domain/payment"
	"github.com/Tender-Services/bridge/pkg/events"
	"github.com/Tender-Services/bridge/pkg/instrument"
	"github.com/Tender-Services/bridge/pkg/money"
	"github.com/Tender-Services/bridge/pkg/storage"
	"github.com/Tender-Services/bridge/pkg/system"
)

// InitiatePaymentCommand represents the intent to create a new payment.
// It contains all the necessary information to initiate a payment flow.
type InitiatePaymentCommand struct {
	IdempotencyKey    string
	DebtorAccountID   uuid.UUID
	CreditorAccountID uuid.UUID
	Amount            money.Money
	Instrument        instrument.Instrument
	Reference         string
}

// CancelPaymentCommand represents the intent to cancel an existing payment.
type CancelPaymentCommand struct {
	PaymentID uuid.UUID
	Reason    string
}

// ApprovePaymentCommand represents an internal action to approve a payment,
// typically after risk assessment.
type ApprovePaymentCommand struct {
	PaymentID uuid.UUID
}

// FailPaymentCommand represents an action to fail a payment,
// from any system that detects a terminal issue (e.g., risk, settlement).
type FailPaymentCommand struct {
	PaymentID   uuid.UUID
	Reason      string
	FailureCode string
}

// CompletePaymentCommand represents an internal action to complete a payment,
// typically after settlement confirmation.
type CompletePaymentCommand struct {
	PaymentID uuid.UUID
}

// PaymentCommandHandler handles commands related to the payment lifecycle.
// It orchestrates the business logic, validation, and state changes for payments,
// acting as the sole writer for the Payment aggregate.
type PaymentCommandHandler struct {
	paymentRepo    payment.Repository
	accountRepo    account.Repository
	eventPublisher events.Publisher
	clock          system.Clock
	unitOfWork     storage.UnitOfWork
}

// NewPaymentCommandHandler creates a new PaymentCommandHandler.
func NewPaymentCommandHandler(
	paymentRepo payment.Repository,
	accountRepo account.Repository,
	eventPublisher events.Publisher,
	clock system.Clock,
	unitOfWork storage.UnitOfWork,
) *PaymentCommandHandler {
	return &PaymentCommandHandler{
		paymentRepo:    paymentRepo,
		accountRepo:    accountRepo,
		eventPublisher: eventPublisher,
		clock:          clock,
		unitOfWork:     unitOfWork,
	}
}

// HandleInitiatePayment processes the InitiatePaymentCommand.
// It performs validation, creates a new payment aggregate, persists it,
// and publishes a PaymentInitiated event. This entire process is transactional.
func (h *PaymentCommandHandler) HandleInitiatePayment(ctx context.Context, cmd InitiatePaymentCommand) (*payment.Payment, error) {
	var p *payment.Payment

	// The entire operation is atomic, managed by the Unit of Work.
	// This ensures that saving the payment and publishing the event either both succeed or both fail.
	err := h.unitOfWork.Execute(ctx, func(store storage.RepositoryProvider) error {
		paymentRepo := store.PaymentRepository()
		accountRepo := store.AccountRepository()

		// 1. Idempotency Check
		existingPayment, err := paymentRepo.FindByIdempotencyKey(ctx, cmd.IdempotencyKey)
		if err != nil && err != payment.ErrPaymentNotFound {
			return fmt.Errorf("failed to check for idempotency key: %w", err)
		}
		if existingPayment != nil {
			p = existingPayment
			return nil // Command already processed, return success without re-processing.
		}

		// 2. Validation
		if err := h.validateInitiatePayment(ctx, accountRepo, cmd); err != nil {
			return fmt.Errorf("payment initiation validation failed: %w", err)
		}

		// 3. Create Payment Aggregate
		newPayment, err := payment.NewPayment(
			uuid.New(), // Generate a new UUID for the payment
			cmd.IdempotencyKey,
			cmd.DebtorAccountID,
			cmd.CreditorAccountID,
			cmd.Amount,
			cmd.Instrument,
			cmd.Reference,
			h.clock.Now(),
		)
		if err != nil {
			return fmt.Errorf("failed to create new payment: %w", err)
		}

		// 4. Persist
		if err := paymentRepo.Save(ctx, newPayment); err != nil {
			return fmt.Errorf("failed to save payment: %w", err)
		}

		// 5. Publish Event
		event, err := events.NewPaymentInitiated(newPayment, h.clock.Now())
		if err != nil {
			return fmt.Errorf("failed to create PaymentInitiated event: %w", err)
		}
		if err := h.eventPublisher.Publish(ctx, event); err != nil {
			return fmt.Errorf("failed to publish PaymentInitiated event: %w", err)
		}

		p = newPayment
		return nil
	})

	return p, err
}

// validateInitiatePayment performs business rule checks before creating a payment.
func (h *PaymentCommandHandler) validateInitiatePayment(ctx context.Context, accountRepo account.Repository, cmd InitiatePaymentCommand) error {
	if cmd.IdempotencyKey == "" {
		return fmt.Errorf("idempotency key is required")
	}
	if cmd.DebtorAccountID == cmd.CreditorAccountID {
		return fmt.Errorf("debtor and creditor accounts cannot be the same")
	}

	debtor, err := accountRepo.FindByID(ctx, cmd.DebtorAccountID)
	if err != nil {
		return fmt.Errorf("debtor account lookup failed: %w", err)
	}
	if !debtor.IsActive() {
		return fmt.Errorf("debtor account %s is not active", cmd.DebtorAccountID)
	}

	creditor, err := accountRepo.FindByID(ctx, cmd.CreditorAccountID)
	if err != nil {
		return fmt.Errorf("creditor account lookup failed: %w", err)
	}
	if !creditor.IsActive() {
		return fmt.Errorf("creditor account %s is not active", cmd.CreditorAccountID)
	}

	if !debtor.AllowsInstrument(cmd.Instrument) {
		return fmt.Errorf("debtor account %s does not allow instrument %s", cmd.DebtorAccountID, cmd.Instrument.Code)
	}
	if !creditor.AllowsInstrument(cmd.Instrument) {
		return fmt.Errorf("creditor account %s does not allow instrument %s", cmd.CreditorAccountID, cmd.Instrument.Code)
	}

	// Note: We do not check for sufficient funds here. That is the responsibility of the Risk
	// and Settlement systems. The execution layer's job is to validate the *intent* and record it.
	return nil
}

// HandleCancelPayment processes the CancelPaymentCommand.
func (h *PaymentCommandHandler) HandleCancelPayment(ctx context.Context, cmd CancelPaymentCommand) (*payment.Payment, error) {
	var p *payment.Payment
	err := h.unitOfWork.Execute(ctx, func(store storage.RepositoryProvider) error {
		paymentRepo := store.PaymentRepository()

		targetPayment, err := paymentRepo.FindByID(ctx, cmd.PaymentID)
		if err != nil {
			return fmt.Errorf("failed to find payment %s: %w", cmd.PaymentID, err)
		}

		if err := targetPayment.Cancel(cmd.Reason, h.clock.Now()); err != nil {
			return err
		}

		if err := paymentRepo.Update(ctx, targetPayment); err != nil {
			return fmt.Errorf("failed to update cancelled payment: %w", err)
		}

		event, err := events.NewPaymentCancelled(targetPayment, h.clock.Now())
		if err != nil {
			return fmt.Errorf("failed to create PaymentCancelled event: %w", err)
		}
		if err := h.eventPublisher.Publish(ctx, event); err != nil {
			return fmt.Errorf("failed to publish PaymentCancelled event: %w", err)
		}

		p = targetPayment
		return nil
	})
	return p, err
}

// HandleApprovePayment processes the ApprovePaymentCommand, typically triggered by an internal system.
func (h *PaymentCommandHandler) HandleApprovePayment(ctx context.Context, cmd ApprovePaymentCommand) (*payment.Payment, error) {
	var p *payment.Payment
	err := h.unitOfWork.Execute(ctx, func(store storage.RepositoryProvider) error {
		paymentRepo := store.PaymentRepository()

		targetPayment, err := paymentRepo.FindByID(ctx, cmd.PaymentID)
		if err != nil {
			return fmt.Errorf("failed to find payment %s: %w", cmd.PaymentID, err)
		}

		if err := targetPayment.Approve(h.clock.Now()); err != nil {
			return err
		}

		if err := paymentRepo.Update(ctx, targetPayment); err != nil {
			return fmt.Errorf("failed to update approved payment: %w", err)
		}

		event, err := events.NewPaymentApproved(targetPayment, h.clock.Now())
		if err != nil {
			return fmt.Errorf("failed to create PaymentApproved event: %w", err)
		}
		if err := h.eventPublisher.Publish(ctx, event); err != nil {
			return fmt.Errorf("failed to publish PaymentApproved event: %w", err)
		}

		p = targetPayment
		return nil
	})
	return p, err
}

// HandleFailPayment processes the FailPaymentCommand.
func (h *PaymentCommandHandler) HandleFailPayment(ctx context.Context, cmd FailPaymentCommand) (*payment.Payment, error) {
	var p *payment.Payment
	err := h.unitOfWork.Execute(ctx, func(store storage.RepositoryProvider) error {
		paymentRepo := store.PaymentRepository()

		targetPayment, err := paymentRepo.FindByID(ctx, cmd.PaymentID)
		if err != nil {
			return fmt.Errorf("failed to find payment %s: %w", cmd.PaymentID, err)
		}

		if err := targetPayment.Fail(cmd.Reason, cmd.FailureCode, h.clock.Now()); err != nil {
			return err
		}

		if err := paymentRepo.Update(ctx, targetPayment); err != nil {
			return fmt.Errorf("failed to update failed payment: %w", err)
		}

		event, err := events.NewPaymentFailed(targetPayment, h.clock.Now())
		if err != nil {
			return fmt.Errorf("failed to create PaymentFailed event: %w", err)
		}
		if err := h.eventPublisher.Publish(ctx, event); err != nil {
			return fmt.Errorf("failed to publish PaymentFailed event: %w", err)
		}

		p = targetPayment
		return nil
	})
	return p, err
}

// HandleCompletePayment processes the CompletePaymentCommand.
func (h *PaymentCommandHandler) HandleCompletePayment(ctx context.Context, cmd CompletePaymentCommand) (*payment.Payment, error) {
	var p *payment.Payment
	err := h.unitOfWork.Execute(ctx, func(store storage.RepositoryProvider) error {
		paymentRepo := store.PaymentRepository()

		targetPayment, err := paymentRepo.FindByID(ctx, cmd.PaymentID)
		if err != nil {
			return fmt.Errorf("failed to find payment %s: %w", cmd.PaymentID, err)
		}

		if err := targetPayment.Complete(h.clock.Now()); err != nil {
			return err
		}

		if err := paymentRepo.Update(ctx, targetPayment); err != nil {
			return fmt.Errorf("failed to update completed payment: %w", err)
		}

		event, err := events.NewPaymentCompleted(targetPayment, h.clock.Now())
		if err != nil {
			return fmt.Errorf("failed to create PaymentCompleted event: %w", err)
		}
		if err := h.eventPublisher.Publish(ctx, event); err != nil {
			return fmt.Errorf("failed to publish PaymentCompleted event: %w", err)
		}

		p = targetPayment
		return nil
	})
	return p, err
}
### END_OF_FILE_COMPLETED ###
```