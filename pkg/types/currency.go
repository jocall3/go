```go
// Copyright (c) 2024. The Bridge Project. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Currency represents an ISO 4217 currency code.
// It is a distinct string type to enforce type safety for currency operations.
type Currency string

// Constants for major ISO 4217 currency codes.
// This provides compile-time safety and autocompletion for common currencies.
const (
	USD Currency = "USD" // United States Dollar
	EUR Currency = "EUR" // Euro
	JPY Currency = "JPY" // Japanese Yen
	GBP Currency = "GBP" // British Pound Sterling
	AUD Currency = "AUD" // Australian Dollar
	CAD Currency = "CAD" // Canadian Dollar
	CHF Currency = "CHF" // Swiss Franc
	CNY Currency = "CNY" // Chinese Yuan
	HKD Currency = "HKD" // Hong Kong Dollar
	NZD Currency = "NZD" // New Zealand Dollar
	SEK Currency = "SEK" // Swedish Krona
	KRW Currency = "KRW" // South Korean Won
	SGD Currency = "SGD" // Singapore Dollar
	NOK Currency = "NOK" // Norwegian Krone
	MXN Currency = "MXN" // Mexican Peso
	INR Currency = "INR" // Indian Rupee
	RUB Currency = "RUB" // Russian Ruble
	ZAR Currency = "ZAR" // South African Rand
	BRL Currency = "BRL" // Brazilian Real
)

// validCurrencies holds the set of all valid ISO 4217 currency codes.
// Using a map with an empty struct provides an efficient O(1) lookup for validation.
var validCurrencies = make(map[Currency]struct{})

// init populates the set of valid currencies at package initialization.
// This ensures the validation map is ready at program start, avoiding runtime setup costs.
// The list is based on ISO 4217 and can be expanded as needed.
// Source: https://en.wikipedia.org/wiki/ISO_4217
func init() {
	codes := []Currency{
		"AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN",
		"BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL",
		"BSD", "BTN", "BWP", "BYN", "BZD", "CAD", "CDF", "CHF", "CLP", "CNY",
		"COP", "CRC", "CUC", "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD",
		"EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GGP", "GHS",
		"GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF",
		"IDR", "ILS", "IMP", "INR", "IQD", "IRR", "ISK", "JEP", "JMD", "JOD",
		"JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRW", "KWD", "KYD", "KZT",
		"LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA", "MKD",
		"MMK", "MNT", "MOP", "MRU", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN",
		"NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK",
		"PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR",
		"SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLL", "SOS", "SRD",
		"STN", "SVC", "SYP", "SZL", "THB", "TJS", "TMT", "TND", "TOP", "TRY",
		"TTD", "TWD", "TZS", "UAH", "UGX", "USD", "UYU", "UZS", "VES",
		"VND", "VUV", "WST", "XAF", "XCD", "XDR", "XOF", "XPF", "YER",
		"ZAR", "ZMW", "ZWL",
	}
	for _, code := range codes {
		validCurrencies[code] = struct{}{}
	}
}

// NewCurrency creates and validates a new Currency instance from a string.
// It enforces that only valid ISO 4217 codes can be used to create a Currency type.
// The input string is trimmed of whitespace and converted to uppercase.
func NewCurrency(code string) (Currency, error) {
	c := Currency(strings.ToUpper(strings.TrimSpace(code)))
	if !c.IsValid() {
		return "", fmt.Errorf("invalid currency code: %q", code)
	}
	return c, nil
}

// IsValid checks if the currency code is a valid and supported ISO 4217 code.
func (c Currency) IsValid() bool {
	_, ok := validCurrencies[c]
	return ok
}

// String implements the fmt.Stringer interface, returning the string representation
// of the currency code.
func (c Currency) String() string {
	return string(c)
}

// MarshalJSON implements the json.Marshaler interface.
// It ensures the currency is marshaled as a valid JSON string, and it fails
// if the currency code is not a valid ISO 4217 code.
func (c Currency) MarshalJSON() ([]byte, error) {
	if !c.IsValid() {
		return nil, fmt.Errorf("cannot marshal invalid currency code: %q", c)
	}
	return json.Marshal(string(c))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It validates the currency code during unmarshaling, ensuring that only valid
// currency codes can be deserialized into the Currency type.
func (c *Currency) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("currency code must be a string, but got %s", data)
	}

	newCurrency, err := NewCurrency(s)
	if err != nil {
		return err // The error from NewCurrency is already descriptive.
	}

	*c = newCurrency
	return nil
}

```