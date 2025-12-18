// Package marketplace contains the core domain entities and business logic
// for the marketplace context of the application.
package marketplace

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// --- Product Enums ---

// ProductCondition represents the physical state of a product.
// It's a typed string to enforce a limited set of valid values.
type ProductCondition string

const (
	ConditionNew         ProductCondition = "NEW"
	ConditionUsedLikeNew ProductCondition = "USED_LIKE_NEW"
	ConditionUsedGood    ProductCondition = "USED_GOOD"
	ConditionUsedFair    ProductCondition = "USED_FAIR"
	ConditionRefurbished ProductCondition = "REFURBISHED"
)

var allProductConditions = map[ProductCondition]struct{}{
	ConditionNew:         {},
	ConditionUsedLikeNew: {},
	ConditionUsedGood:    {},
	ConditionUsedFair:    {},
	ConditionRefurbished: {},
}

// IsValid checks if the product condition is a defined constant.
func (pc ProductCondition) IsValid() bool {
	_, ok := allProductConditions[pc]
	return ok
}

// String returns the string representation of the ProductCondition.
func (pc ProductCondition) String() string {
	return string(pc)
}

// Scan implements the sql.Scanner interface, allowing the type to be read from a database.
func (pc *ProductCondition) Scan(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("invalid type for ProductCondition")
	}
	*pc = ProductCondition(s)
	if !pc.IsValid() {
		return fmt.Errorf("invalid product condition: %s", s)
	}
	return nil
}

// Value implements the driver.Valuer interface, allowing the type to be written to a database.
func (pc ProductCondition) Value() (driver.Value, error) {
	if !pc.IsValid() {
		return nil, fmt.Errorf("invalid product condition: %s", pc)
	}
	return string(pc), nil
}

// ProductStatus represents the availability status of a product in the marketplace.
type ProductStatus string

const (
	StatusAvailable ProductStatus = "AVAILABLE"
	StatusReserved  ProductStatus = "RESERVED" // An offer has been accepted, pending transaction
	StatusSold      ProductStatus = "SOLD"
	StatusDelisted  ProductStatus = "DELISTED" // Removed by the owner
)

var allProductStatuses = map[ProductStatus]struct{}{
	StatusAvailable: {},
	StatusReserved:  {},
	StatusSold:      {},
	StatusDelisted:  {},
}

// IsValid checks if the product status is a defined constant.
func (ps ProductStatus) IsValid() bool {
	_, ok := allProductStatuses[ps]
	return ok
}

// String returns the string representation of the ProductStatus.
func (ps ProductStatus) String() string {
	return string(ps)
}

// Scan implements the sql.Scanner interface.
func (ps *ProductStatus) Scan(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("invalid type for ProductStatus")
	}
	*ps = ProductStatus(s)
	if !ps.IsValid() {
		return fmt.Errorf("invalid product status: %s", s)
	}
	return nil
}

// Value implements the driver.Valuer interface.
func (ps ProductStatus) Value() (driver.Value, error) {
	if !ps.IsValid() {
		return nil, fmt.Errorf("invalid product status: %s", ps)
	}
	return string(ps), nil
}

// --- Product Entity ---

// Product represents an item listed for sale in the marketplace.
// It is the aggregate root for the product context.
type Product struct {
	ID          uuid.UUID        `json:"id" db:"id"`
	OwnerID     uuid.UUID        `json:"owner_id" db:"owner_id"`
	Name        string           `json:"name" db:"name"`
	Description string           `json:"description" db:"description"`
	Category    string           `json:"category" db:"category"`
	Condition   ProductCondition `json:"condition" db:"condition"`
	Price       int64            `json:"price" db:"price"`             // Price in the smallest currency unit (e.g., cents)
	Currency    string           `json:"currency" db:"currency"`       // ISO 4217 currency code (e.g., "USD")
	Images      []string         `json:"images" db:"images"`           // URLs to product images
	Status      ProductStatus    `json:"status" db:"status"`
	Version     int              `json:"-" db:"version"`               // For optimistic locking
	CreatedAt   time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at" db:"updated_at"`
}

// NewProduct creates a new Product with default values, ensuring it's in a valid state.
func NewProduct(ownerID uuid.UUID, name, description, category string, condition ProductCondition, price int64, currency string, images []string) (*Product, error) {
	now := time.Now().UTC().Truncate(time.Microsecond)
	p := &Product{
		ID:          uuid.New(),
		OwnerID:     ownerID,
		Name:        name,
		Description: description,
		Category:    category,
		Condition:   condition,
		Price:       price,
		Currency:    strings.ToUpper(currency),
		Images:      images,
		Status:      StatusAvailable,
		Version:     1,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := p.Validate(); err != nil {
		return nil, fmt.Errorf("product validation failed: %w", err)
	}

	return p, nil
}

// Validate checks if the product's fields meet the business rules.
func (p *Product) Validate() error {
	if p.ID == uuid.Nil {
		return errors.New("product ID cannot be nil")
	}
	if p.OwnerID == uuid.Nil {
		return errors.New("owner ID cannot be nil")
	}
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("product name cannot be empty")
	}
	if len(p.Name) > 100 {
		return errors.New("product name cannot exceed 100 characters")
	}
	if len(p.Description) > 5000 {
		return errors.New("product description cannot exceed 5000 characters")
	}
	if strings.TrimSpace(p.Category) == "" {
		return errors.New("product category cannot be empty")
	}
	if !p.Condition.IsValid() {
		return fmt.Errorf("invalid product condition: %s", p.Condition)
	}
	if p.Price < 0 {
		return errors.New("product price cannot be negative")
	}
	if len(p.Currency) != 3 {
		return errors.New("currency code must be 3 characters long (ISO 4217)")
	}
	if !p.Status.IsValid() {
		return fmt.Errorf("invalid product status: %s", p.Status)
	}
	if p.CreatedAt.IsZero() || p.UpdatedAt.IsZero() {
		return errors.New("timestamps must be set")
	}
	return nil
}

// --- Offer Enums ---

// OfferStatus represents the state of an offer made on a product.
type OfferStatus string

const (
	OfferStatusPending   OfferStatus = "PENDING"   // Offer made, awaiting owner response
	OfferStatusAccepted  OfferStatus = "ACCEPTED"  // Owner accepted the offer
	OfferStatusRejected  OfferStatus = "REJECTED"  // Owner rejected the offer
	OfferStatusRetracted OfferStatus = "RETRACTED" // Offerer withdrew the offer
)

var allOfferStatuses = map[OfferStatus]struct{}{
	OfferStatusPending:   {},
	OfferStatusAccepted:  {},
	OfferStatusRejected:  {},
	OfferStatusRetracted: {},
}

// IsValid checks if the offer status is a defined constant.
func (os OfferStatus) IsValid() bool {
	_, ok := allOfferStatuses[os]
	return ok
}

// String returns the string representation of the OfferStatus.
func (os OfferStatus) String() string {
	return string(os)
}

// Scan implements the sql.Scanner interface.
func (os *OfferStatus) Scan(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("invalid type for OfferStatus")
	}
	*os = OfferStatus(s)
	if !os.IsValid() {
		return fmt.Errorf("invalid offer status: %s", s)
	}
	return nil
}

// Value implements the driver.Valuer interface.
func (os OfferStatus) Value() (driver.Value, error) {
	if !os.IsValid() {
		return nil, fmt.Errorf("invalid offer status: %s", os)
	}
	return string(os), nil
}

// --- Offer Entity ---

// Offer represents a bid made by a user for a specific product.
type Offer struct {
	ID        uuid.UUID   `json:"id" db:"id"`
	ProductID uuid.UUID   `json:"product_id" db:"product_id"`
	OffererID uuid.UUID   `json:"offerer_id" db:"offerer_id"`
	OwnerID   uuid.UUID   `json:"owner_id" db:"owner_id"` // Denormalized for easier authorization checks
	Amount    int64       `json:"amount" db:"amount"`     // Amount in the smallest currency unit
	Currency  string      `json:"currency" db:"currency"` // ISO 4217 currency code
	Message   string      `json:"message" db:"message"`
	Status    OfferStatus `json:"status" db:"status"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
}

// NewOffer creates a new Offer with default values, ensuring it's in a valid state.
func NewOffer(productID, offererID, ownerID uuid.UUID, amount int64, currency, message string) (*Offer, error) {
	now := time.Now().UTC().Truncate(time.Microsecond)
	o := &Offer{
		ID:        uuid.New(),
		ProductID: productID,
		OffererID: offererID,
		OwnerID:   ownerID,
		Amount:    amount,
		Currency:  strings.ToUpper(currency),
		Message:   message,
		Status:    OfferStatusPending,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := o.Validate(); err != nil {
		return nil, fmt.Errorf("offer validation failed: %w", err)
	}

	return o, nil
}

// Validate checks if the offer's fields meet the business rules.
func (o *Offer) Validate() error {
	if o.ID == uuid.Nil {
		return errors.New("offer ID cannot be nil")
	}
	if o.ProductID == uuid.Nil {
		return errors.New("product ID cannot be nil")
	}
	if o.OffererID == uuid.Nil {
		return errors.New("offerer ID cannot be nil")
	}
	if o.OwnerID == uuid.Nil {
		return errors.New("owner ID cannot be nil")
	}
	if o.OffererID == o.OwnerID {
		return errors.New("offerer cannot make an offer on their own product")
	}
	if o.Amount <= 0 {
		return errors.New("offer amount must be positive")
	}
	if len(o.Currency) != 3 {
		return errors.New("currency code must be 3 characters long (ISO 4217)")
	}
	if len(o.Message) > 1000 {
		return errors.New("offer message cannot exceed 1000 characters")
	}
	if !o.Status.IsValid() {
		return fmt.Errorf("invalid offer status: %s", o.Status)
	}
	if o.CreatedAt.IsZero() || o.UpdatedAt.IsZero() {
		return errors.New("timestamps must be set")
	}
	return nil
}

// Note on persistence of []string:
// The `db:"images"` tag on `Product.Images` of type `[]string` requires a database driver
// that supports this conversion (e.g., via pq for PostgreSQL arrays).
// If using a standard JSON/JSONB column, the repository layer would be responsible for
// marshalling this slice into JSON before writing and unmarshalling it after reading.
// A helper type like the one below can be used in the persistence layer DTOs.
/*
// StringSlice is a helper type for marshalling/unmarshalling []string from/to a JSON db column.
type StringSlice []string

// Scan implements the sql.Scanner interface for []string.
func (s *StringSlice) Scan(src interface{}) error {
	if src == nil {
		*s = nil
		return nil
	}
	b, ok := src.([]byte)
	if !ok {
		return errors.New("scan source was not []byte")
	}
	return json.Unmarshal(b, s)
}

// Value implements the driver.Valuer interface for []string.
func (s StringSlice) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}
	return json.Marshal(s)
}
*/