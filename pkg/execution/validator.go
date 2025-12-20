```go
// Copyright (c) 2024. Your project name. All rights reserved.
//
// THIS CODE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package execution

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/your-org/your-project/pkg/command"
	"github.com/your-org/your-project/pkg/model/instrument"
	"github.com/your-org/your-project/pkg/model/order"
)

// Sentinel errors for common validation failures. These are exported to allow
// consumers of the validator to programmatically check for specific error types.
var (
	ErrNilCommand             = fmt.Errorf("command cannot be nil")
	ErrInvalidCommandID       = fmt.Errorf("command ID must be a valid non-nil UUID")
	ErrInvalidTimestamp       = fmt.Errorf("command timestamp must be a positive value")
	ErrUnknownCommandType     = fmt.Errorf("unknown command type")
	ErrInvalidOrderID         = fmt.Errorf("order ID must be a valid non-nil UUID")
	ErrInvalidAccountID       = fmt.Errorf("account ID must be a valid non-nil UUID")
	ErrInvalidInstrumentID    = fmt.Errorf("instrument ID must be a valid non-nil UUID")
	ErrInvalidOrderSide       = fmt.Errorf("invalid order side")
	ErrInvalidOrderType       = fmt.Errorf("invalid order type")
	ErrInvalidTimeInForce     = fmt.Errorf("invalid time in force")
	ErrInvalidQuantity        = fmt.Errorf("quantity must be positive")
	ErrInvalidPrice           = fmt.Errorf("price must be positive for LIMIT orders")
	ErrPriceNotAllowed        = fmt.Errorf("price must not be set for MARKET orders")
	ErrInvalidAmount          = fmt.Errorf("amount must be a finite number")
	ErrInvalidCurrency        = fmt.Errorf("currency code is invalid")
	ErrInvalidAsset           = fmt.Errorf("asset code is invalid")
	ErrInvalidInstrumentType  = fmt.Errorf("invalid instrument type")
	ErrInvalidLotSize         = fmt.Errorf("lot size must be positive")
	ErrInvalidTickSize        = fmt.Errorf("tick size must be positive")
)

// Validator defines the interface for command validation.
// It ensures that a command is well-formed and internally consistent
// before it is passed to the execution engine. This is a stateless
// validation step. State-dependent validation (e.g., checking if an
// account exists or has sufficient funds) is handled by the engine itself
// to ensure atomicity of read-modify-write cycles on the system's state.
type Validator interface {
	Validate(cmd command.Command) error
}

// CommandValidator provides stateless validation for commands.
// It is safe for concurrent use as it holds no state.
type CommandValidator struct{}

// NewCommandValidator creates a new instance of CommandValidator.
func NewCommandValidator() *CommandValidator {
	return &CommandValidator{}
}

// Validate checks the given command for structural and logical integrity.
// It uses a type switch to delegate to specific validation methods for each
// known command type, enforcing invariants early in the processing pipeline.
func (v *CommandValidator) Validate(cmd command.Command) error {
	if cmd == nil {
		return ErrNilCommand
	}

	if err := v.validateBase(cmd.Base()); err != nil {
		return err
	}

	switch c := cmd.(type) {
	case *command.SubmitOrderCommand:
		return v.validateSubmitOrder(c)
	case *command.CancelOrderCommand:
		return v.validateCancelOrder(c)
	case *command.CreateAccountCommand:
		return v.validateCreateAccount(c)
	case *command.AdjustBalanceCommand:
		return v.validateAdjustBalance(c)
	case *command.CreateInstrumentCommand:
		return v.validateCreateInstrument(c)
	default:
		return fmt.Errorf("%w: %T", ErrUnknownCommandType, c)
	}
}

// validateBase checks the fields common to all commands.
func (v *CommandValidator) validateBase(base *command.BaseCommand) error {
	if base.ID == uuid.Nil {
		return ErrInvalidCommandID
	}
	if base.Timestamp.IsZero() || base.Timestamp.UnixNano() <= 0 {
		return ErrInvalidTimestamp
	}
	return nil
}

// validateSubmitOrder validates a SubmitOrderCommand.
func (v *CommandValidator) validateSubmitOrder(cmd *command.SubmitOrderCommand) error {
	if cmd.OrderID == uuid.Nil {
		return ErrInvalidOrderID
	}
	if cmd.AccountID == uuid.Nil {
		return ErrInvalidAccountID
	}
	if cmd.InstrumentID == uuid.Nil {
		return ErrInvalidInstrumentID
	}
	if !cmd.Side.IsValid() {
		return fmt.Errorf("%w: %s", ErrInvalidOrderSide, cmd.Side)
	}
	if !cmd.Type.IsValid() {
		return fmt.Errorf("%w: %s", ErrInvalidOrderType, cmd.Type)
	}
	if !cmd.TimeInForce.IsValid() {
		return fmt.Errorf("%w: %s", ErrInvalidTimeInForce, cmd.TimeInForce)
	}
	if cmd.Quantity.IsNegative() || cmd.Quantity.IsZero() {
		return fmt.Errorf("%w: %s", ErrInvalidQuantity, cmd.Quantity)
	}

	switch cmd.Type {
	case order.TypeLimit:
		if cmd.Price.IsNegative() || cmd.Price.IsZero() {
			return fmt.Errorf("%w: %s", ErrInvalidPrice, cmd.Price)
		}
	case order.TypeMarket:
		if !cmd.Price.IsZero() {
			return fmt.Errorf("%w: price was %s", ErrPriceNotAllowed, cmd.Price)
		}
	}

	return nil
}

// validateCancelOrder validates a CancelOrderCommand.
func (v *CommandValidator) validateCancelOrder(cmd *command.CancelOrderCommand) error {
	if cmd.OrderID == uuid.Nil {
		return ErrInvalidOrderID
	}
	if cmd.AccountID == uuid.Nil {
		return ErrInvalidAccountID
	}
	return nil
}

// validateCreateAccount validates a CreateAccountCommand.
func (v *CommandValidator) validateCreateAccount(cmd *command.CreateAccountCommand) error {
	if cmd.AccountID == uuid.Nil {
		return ErrInvalidAccountID
	}
	// A basic currency code validation (e.g., 3 letters, uppercase).
	// A more robust implementation would check against a pre-configured list of supported currencies.
	if len(cmd.Currency) != 3 {
		return fmt.Errorf("%w: %s", ErrInvalidCurrency, cmd.Currency)
	}
	return nil
}

// validateAdjustBalance validates an AdjustBalanceCommand.
func (v *CommandValidator) validateAdjustBalance(cmd *command.AdjustBalanceCommand) error {
	if cmd.AccountID == uuid.Nil {
		return ErrInvalidAccountID
	}
	if len(cmd.Currency) != 3 {
		return fmt.Errorf("%w: %s", ErrInvalidCurrency, cmd.Currency)
	}
	// Amount must be a finite number. Zero is allowed for audit/reconciliation entries.
	if cmd.Amount.IsNaN() || cmd.Amount.IsInf() {
		return fmt.Errorf("%w: %s", ErrInvalidAmount, cmd.Amount)
	}
	return nil
}

// validateCreateInstrument validates a CreateInstrumentCommand.
func (v *CommandValidator) validateCreateInstrument(cmd *command.CreateInstrumentCommand) error {
	if cmd.InstrumentID == uuid.Nil {
		return ErrInvalidInstrumentID
	}
	if cmd.BaseAsset == "" {
		return fmt.Errorf("%w: base asset cannot be empty", ErrInvalidAsset)
	}
	if cmd.QuoteAsset == "" {
		return fmt.Errorf("%w: quote asset cannot be empty", ErrInvalidAsset)
	}
	if cmd.BaseAsset == cmd.QuoteAsset {
		return fmt.Errorf("base and quote assets cannot be the same: %s", cmd.BaseAsset)
	}
	if !cmd.Type.IsValid() {
		return fmt.Errorf("%w: %s", ErrInvalidInstrumentType, cmd.Type)
	}
	if cmd.LotSize.IsNegative() || cmd.LotSize.IsZero() {
		return fmt.Errorf("%w: %s", ErrInvalidLotSize, cmd.LotSize)
	}
	if cmd.TickSize.IsNegative() || cmd.TickSize.IsZero() {
		return fmt.Errorf("%w: %s", ErrInvalidTickSize, cmd.TickSize)
	}

	// Ensure precision does not exceed system limits.
	if cmd.TickSize.Exponent() < -decimal.MaxExponent {
		return fmt.Errorf("tick size precision is too high: %s", cmd.TickSize)
	}
	if cmd.LotSize.Exponent() < -decimal.MaxExponent {
		return fmt.Errorf("lot size precision is too high: %s", cmd.LotSize)
	}
	return nil
}
### END_OF_FILE_COMPLETED ###
```