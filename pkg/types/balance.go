```go
package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/shopspring/decimal"
)

// Balance represents a monetary value in its smallest, indivisible unit (e.g., satoshis, cents).
// It uses a signed 64-bit integer (int64) to avoid floating-point inaccuracies,
// a critical requirement for capital safety.
//
// The choice of int64 provides a balance between high performance for core ledger operations
// and a sufficiently large range for most financial applications (approx +/- 9.2 quintillion).
//
// All arithmetic operations are checked for overflow/underflow to ensure deterministic behavior
// and fail-closed semantics, preventing silent, catastrophic errors.
type Balance int64

// Constants for common Balance values and error messages.
var (
	// ZeroBalance is a Balance with a value of zero.
	ZeroBalance = Balance(0)

	// ErrBalanceOverflow is returned when an arithmetic operation exceeds the maximum int64 value.
	ErrBalanceOverflow = errors.New("balance operation resulted in overflow")
	// ErrBalanceUnderflow is returned when an arithmetic operation goes below the minimum int64 value.
	ErrBalanceUnderflow = errors.New("balance operation resulted in underflow")
	// ErrDivisionByZero is returned when a division operation attempts to divide by zero.
	ErrDivisionByZero = errors.New("division by zero")
	// ErrInvalidBalanceString is returned when parsing a string that is not a valid number.
	ErrInvalidBalanceString = errors.New("invalid balance string representation")
	// ErrPrecisionLoss is returned when a conversion from a less precise type (like float64)
	// would result in a loss of monetary value.
	ErrPrecisionLoss = errors.New("conversion would result in precision loss")
)

// NewBalance creates a new Balance from an int64 value.
func NewBalance(value int64) Balance {
	return Balance(value)
}

// FromString creates a Balance from a decimal string representation, scaled by a given precision.
// It uses a high-precision decimal library for safe parsing and scaling.
// Example: FromString("1.23", 2) results in a Balance of 123.
// Example: FromString("100", 0) results in a Balance of 100.
func FromString(value string, precision uint8) (Balance, error) {
	if value == "" {
		return ZeroBalance, nil
	}

	dec, err := decimal.NewFromString(value)
	if err != nil {
		return ZeroBalance, fmt.Errorf("%w: %v", ErrInvalidBalanceString, err)
	}

	scaler := decimal.New(1, int32(precision))
	scaledValue := dec.Mul(scaler)

	if !scaledValue.IsInteger() {
		return ZeroBalance, fmt.Errorf("%w: string '%s' has more than %d decimal places", ErrPrecisionLoss, value, precision)
	}

	if !scaledValue.FitsInInt64() {
		return ZeroBalance, ErrBalanceOverflow
	}

	return Balance(scaledValue.IntPart()), nil
}

// FromFloat64 creates a Balance from a float64, scaled by a given precision.
// This function is inherently unsafe due to the nature of floating-point numbers
// and should be used with extreme caution, typically only at system boundaries.
// It is strongly recommended to use FromString whenever possible.
// The function will return an error if the conversion results in a loss of precision.
func FromFloat64(value float64, precision uint8) (Balance, error) {
	// Use decimal library for robust conversion and checking.
	dec := decimal.NewFromFloat(value)
	scaler := decimal.New(1, int32(precision))
	scaledValue := dec.Mul(scaler)

	// Check if the result of scaling is an integer. If not, precision is lost.
	if !scaledValue.IsInteger() {
		return ZeroBalance, fmt.Errorf("%w: float '%f' cannot be represented with precision %d", ErrPrecisionLoss, value, precision)
	}

	if !scaledValue.FitsInInt64() {
		return ZeroBalance, ErrBalanceOverflow
	}

	return Balance(scaledValue.IntPart()), nil
}

// Int64 returns the raw int64 value of the Balance.
func (b Balance) Int64() int64 {
	return int64(b)
}

// String implements the fmt.Stringer interface.
// It returns the string representation of the raw integer value.
// For a formatted decimal string, use the Format() method.
func (b Balance) String() string {
	return strconv.FormatInt(int64(b), 10)
}

// Format returns a decimal string representation of the balance with a given precision.
// Example: a balance of 12345 with precision 2 would be formatted as "123.45".
func (b Balance) Format(precision uint8) string {
	dec := decimal.NewFromInt(int64(b))
	scaler := decimal.New(1, int32(precision))
	return dec.Div(scaler).String()
}

// Add safely adds two Balances, checking for overflow.
func (b Balance) Add(other Balance) (Balance, error) {
	// Overflow occurs if the signs of the operands are the same,
	// but the sign of the result is different.
	if (other > 0 && b > math.MaxInt64-other) || (other < 0 && b < math.MinInt64-other) {
		return ZeroBalance, ErrBalanceOverflow
	}
	return b + other, nil
}

// Sub safely subtracts another Balance, checking for underflow/overflow.
func (b Balance) Sub(other Balance) (Balance, error) {
	// Overflow occurs if the signs of the operands are different,
	// and the sign of the result is different from the first operand's sign.
	if (other > 0 && b < math.MinInt64+other) || (other < 0 && b > math.MaxInt64+other) {
		return ZeroBalance, ErrBalanceUnderflow
	}
	return b - other, nil
}

// Mul safely multiplies the Balance by a scalar, checking for overflow.
func (b Balance) Mul(scalar int64) (Balance, error) {
	if b == 0 || scalar == 0 {
		return ZeroBalance, nil
	}

	result := b * Balance(scalar)
	// A simple and effective way to check for multiplication overflow
	// is to divide the result by one of the original numbers and see if we get the other back.
	if int64(b) != int64(result)/scalar {
		return ZeroBalance, ErrBalanceOverflow
	}
	return result, nil
}

// Div safely divides the Balance by a scalar, checking for division by zero.
// The result is truncated towards zero (integer division).
func (b Balance) Div(scalar int64) (Balance, error) {
	if scalar == 0 {
		return ZeroBalance, ErrDivisionByZero
	}
	// The only overflow case for division is math.MinInt64 / -1
	if b == math.MinInt64 && scalar == -1 {
		return ZeroBalance, ErrBalanceOverflow
	}
	return b / Balance(scalar), nil
}

// Cmp compares two balances.
// It returns:
//
//	-1 if b < other
//	 0 if b == other
//	+1 if b > other
func (b Balance) Cmp(other Balance) int {
	if b < other {
		return -1
	}
	if b > other {
		return 1
	}
	return 0
}

// Equal checks if two Balances are equal.
func (b Balance) Equal(other Balance) bool {
	return b == other
}

// GreaterThan checks if the Balance is greater than another.
func (b Balance) GreaterThan(other Balance) bool {
	return b > other
}

// LessThan checks if the Balance is less than another.
func (b Balance) LessThan(other Balance) bool {
	return b < other
}

// IsZero checks if the Balance is zero.
func (b Balance) IsZero() bool {
	return b == 0
}

// IsPositive checks if the Balance is positive (greater than zero).
func (b Balance) IsPositive() bool {
	return b > 0
}

// IsNegative checks if the Balance is negative (less than zero).
func (b Balance) IsNegative() bool {
	return b < 0
}

// Abs returns the absolute value of the Balance.
// It returns an error on overflow, which only occurs when b is math.MinInt64.
func (b Balance) Abs() (Balance, error) {
	if b == math.MinInt64 {
		return ZeroBalance, ErrBalanceOverflow
	}
	if b < 0 {
		return -b, nil
	}
	return b, nil
}

// MarshalJSON implements the json.Marshaler interface.
// It marshals the Balance as a JSON string to prevent precision loss in JavaScript clients,
// which may interpret large integers as floating-point numbers.
func (b Balance) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatInt(int64(b), 10))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It can unmarshal from a JSON string or a JSON number, providing flexibility
// for different API clients.
func (b *Balance) UnmarshalJSON(data []byte) error {
	// First, try to unmarshal as a string.
	var strValue string
	if err := json.Unmarshal(data, &strValue); err == nil {
		val, err := strconv.ParseInt(strValue, 10, 64)
		if err != nil {
			return fmt.Errorf("cannot unmarshal balance from string '%s': %w", strValue, err)
		}
		*b = Balance(val)
		return nil
	}

	// If string unmarshaling fails, try as a raw number.
	var intValue int64
	if err := json.Unmarshal(data, &intValue); err != nil {
		return fmt.Errorf("cannot unmarshal balance from '%s': %w", string(data), err)
	}

	*b = Balance(intValue)
	return nil
}

```