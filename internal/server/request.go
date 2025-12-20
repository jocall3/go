```go
package server

// request.go defines the request payload structs for all API endpoints.
// These structs include validation tags to ensure the integrity of incoming data.
// The validation tags are compatible with libraries like go-playground/validator.

// CreateAccountRequest defines the payload for creating a new financial account.
type CreateAccountRequest struct {
	// OwnerID is the unique identifier for the entity (user, organization) that owns this account.
	// It must be a valid UUIDv4.
	OwnerID string `json:"owner_id" binding:"required" validate:"required,uuid4"`

	// AccountType specifies the purpose or category of the account (e.g., "TRADING", "SETTLEMENT").
	// The value must be one of the predefined, system-supported account types.
	AccountType string `json:"account_type" binding:"required" validate:"required,alpha,uppercase"`

	// Currency is the ISO 4217 currency code (e.g., "USD") or a supported crypto asset code (e.g., "BTC").
	Currency string `json:"currency" binding:"required" validate:"required,alphanum,uppercase,min=3,max=10"`

	// AllowNegativeBalance determines if the account can have a negative balance.
	// This is a critical risk parameter and defaults to false.
	AllowNegativeBalance bool `json:"allow_negative_balance" validate:"omitempty,boolean"`
}

// CreateInternalTransferRequest defines the payload for transferring funds between two internal accounts.
type CreateInternalTransferRequest struct {
	// IdempotencyKey is a unique client-generated identifier (UUIDv4) to prevent duplicate requests.
	IdempotencyKey string `json:"idempotency_key" binding:"required" validate:"required,uuid4"`

	// SourceAccountID is the unique identifier of the account from which funds will be debited.
	SourceAccountID string `json:"source_account_id" binding:"required" validate:"required,uuid4"`

	// DestinationAccountID is the unique identifier of the account to which funds will be credited.
	DestinationAccountID string `json:"destination_account_id" binding:"required" validate:"required,uuid4,nefield=SourceAccountID"`

	// Amount is the quantity of the asset to transfer, represented as a string to preserve precision.
	// It must be a positive decimal value. A custom validator is recommended.
	Amount string `json:"amount" binding:"required" validate:"required,numeric"`

	// Currency is the ISO 4217 currency code or crypto asset code for the transfer.
	// It must match the currency of both the source and destination accounts.
	Currency string `json:"currency" binding:"required" validate:"required,alphanum,uppercase,min=3,max=10"`
}

// CreateDepositRequest defines the payload for an external deposit into an internal account.
type CreateDepositRequest struct {
	// IdempotencyKey is a unique client-generated identifier (UUIDv4) to prevent duplicate requests.
	IdempotencyKey string `json:"idempotency_key" binding:"required" validate:"required,uuid4"`

	// DestinationAccountID is the unique identifier of the account to which funds will be credited.
	DestinationAccountID string `json:"destination_account_id" binding:"required" validate:"required,uuid4"`

	// Amount is the quantity of the asset being deposited, represented as a string to preserve precision.
	// It must be a positive decimal value. A custom validator is recommended.
	Amount string `json:"amount" binding:"required" validate:"required,numeric"`

	// Currency is the ISO 4217 currency code or crypto asset code for the deposit.
	Currency string `json:"currency" binding:"required" validate:"required,alphanum,uppercase,min=3,max=10"`

	// ExternalTransactionID is a reference to the transaction on the external network (e.g., blockchain transaction hash, bank wire reference).
	ExternalTransactionID string `json:"external_transaction_id" binding:"required" validate:"required,printascii,min=1,max=256"`
}

// CreateWithdrawalRequest defines the payload for a withdrawal from an internal account to an external destination.
type CreateWithdrawalRequest struct {
	// IdempotencyKey is a unique client-generated identifier (UUIDv4) to prevent duplicate requests.
	IdempotencyKey string `json:"idempotency_key" binding:"required" validate:"required,uuid4"`

	// SourceAccountID is the unique identifier of the account from which funds will be debited.
	SourceAccountID string `json:"source_account_id" binding:"required" validate:"required,uuid4"`

	// Amount is the quantity of the asset to withdraw, represented as a string to preserve precision.
	// It must be a positive decimal value. A custom validator is recommended.
	Amount string `json:"amount" binding:"required" validate:"required,numeric"`

	// Currency is the ISO 4217 currency code or crypto asset code for the withdrawal.
	Currency string `json:"currency" binding:"required" validate:"required,alphanum,uppercase,min=3,max=10"`

	// DestinationAddress holds the external address details for the withdrawal.
	// For crypto, this would be a blockchain address. For fiat, it could be bank account details.
	// The structure and validation of this field would be currency-specific.
	DestinationAddress string `json:"destination_address" binding:"required" validate:"required,printascii,min=1,max=256"`
}

// CreateOrderRequest defines the payload for submitting a new trading order.
type CreateOrderRequest struct {
	// IdempotencyKey is a unique client-generated identifier (UUIDv4) to prevent duplicate requests.
	IdempotencyKey string `json:"idempotency_key" binding:"required" validate:"required,uuid4"`

	// AccountID is the unique identifier of the trading account placing the order.
	AccountID string `json:"account_id" binding:"required" validate:"required,uuid4"`

	// InstrumentID is the identifier for the trading pair (e.g., "BTC-USD").
	InstrumentID string `json:"instrument_id" binding:"required" validate:"required,alphanum,uppercase,min=3,max=20"`

	// Side specifies whether the order is to "BUY" or "SELL".
	Side string `json:"side" binding:"required" validate:"required,oneof=BUY SELL"`

	// Type specifies the order type, e.g., "LIMIT" or "MARKET".
	Type string `json:"type" binding:"required" validate:"required,oneof=LIMIT MARKET"`

	// Quantity is the amount of the base asset to buy or sell, represented as a string for precision.
	// It must be a positive decimal value and conform to the instrument's lot size rules.
	Quantity string `json:"quantity" binding:"required" validate:"required,numeric"`

	// Price is the limit price for a "LIMIT" order, represented as a string for precision.
	// It is required for LIMIT orders and must be a positive decimal value conforming to the instrument's tick size rules.
	// This field should be omitted for "MARKET" orders.
	Price string `json:"price,omitempty" validate:"omitempty,numeric"`

	// TimeInForce specifies how long the order remains in effect, e.g., "GTC" (Good-Til-Canceled),
	// "IOC" (Immediate-Or-Cancel), "FOK" (Fill-Or-Kill). Defaults to "GTC" if not provided.
	TimeInForce string `json:"time_in_force,omitempty" validate:"omitempty,oneof=GTC IOC FOK"`
}

// CancelOrderRequest defines the payload for cancelling an existing order.
type CancelOrderRequest struct {
	// OrderID is the unique identifier of the order to be cancelled.
	OrderID string `json:"order_id" binding:"required" validate:"required,uuid4"`

	// AccountID is the unique identifier of the account that placed the order.
	// This is required for authorization and to prevent cancelling orders of other accounts.
	AccountID string `json:"account_id" binding:"required" validate:"required,uuid4"`
}

// AdminCreateInstrumentRequest defines the payload for an admin to create a new tradable instrument.
type AdminCreateInstrumentRequest struct {
	// ID is the unique identifier for the new instrument (e.g., "BTC-USD").
	ID string `json:"id" binding:"required" validate:"required,alphanum,uppercase,min=3,max=20"`

	// BaseCurrency is the currency being traded (e.g., "BTC" in "BTC-USD").
	BaseCurrency string `json:"base_currency" binding:"required" validate:"required,alphanum,uppercase,min=3,max=10"`

	// QuoteCurrency is the currency in which the price is denominated (e.g., "USD" in "BTC-USD").
	QuoteCurrency string `json:"quote_currency" binding:"required" validate:"required,alphanum,uppercase,min=3,max=10,nefield=BaseCurrency"`

	// TickSize is the smallest possible price movement, represented as a string for precision.
	TickSize string `json:"tick_size" binding:"required" validate:"required,numeric"`

	// LotSize is the smallest possible order quantity increment, represented as a string for precision.
	LotSize string `json:"lot_size" binding:"required" validate:"required,numeric"`
}

// AdminUpdateRiskLimitsRequest defines the payload for an admin to update system-wide or per-instrument risk limits.
type AdminUpdateRiskLimitsRequest struct {
	// Scope indicates the level at which the limit applies ("SYSTEM", "INSTRUMENT", "ACCOUNT").
	Scope string `json:"scope" binding:"required" validate:"required,oneof=SYSTEM INSTRUMENT ACCOUNT"`

	// TargetID is the identifier for the scope.
	// e.g., an instrument ID for "INSTRUMENT" scope, or an account ID for "ACCOUNT" scope.
	// Can be empty for "SYSTEM" scope.
	TargetID string `json:"target_id,omitempty" validate:"omitempty,alphanum,uppercase,min=3,max=36"`

	// MaxPositionSize is the maximum net position allowed, represented as a string for precision.
	MaxPositionSize string `json:"max_position_size" binding:"required" validate:"required,numeric"`

	// MaxOrderValue is the maximum value of a single order in quote currency, represented as a string.
	MaxOrderValue string `json:"max_order_value" binding:"required" validate:"required,numeric"`
}

```