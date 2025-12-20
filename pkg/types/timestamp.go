```go
// Copyright (c) 2024. All rights reserved.
// This file is part of the Go-based financial infrastructure project.
//
// The system is designed to bridge the gap between traditional banking and modern
// crypto exchanges, emphasizing bank-level invariants with exchange-level execution speed.
//
// File: pkg/types/timestamp.go
// Description: Provides standardized functions for handling timestamps, ensuring
// consistent precision and formatting throughout the system. This is critical for
// auditability, deterministic replay of events, and unambiguous ordering.

package types

import (
	"encoding/json"
	"fmt"
	"time"
)

// TimestampFormat defines the standard string representation for all timestamps.
// It uses RFC3339 with nanosecond precision and UTC timezone (Zulu time),
// ensuring maximum precision and unambiguous representation across all system components.
// This strict format is non-negotiable for creating a deterministic and auditable
// event log.
// Example: "2006-01-02T15:04:05.999999999Z"
const TimestampFormat = time.RFC3339Nano

// Timestamp is a custom time type that enforces UTC and a standard format.
// It wraps the standard time.Time but provides custom marshaling and unmarshaling
// to ensure consistency for serialization (e.g., JSON APIs, database records),
// logging, and auditing. By using a dedicated type, we prevent accidental use
// of non-standardized time representations.
type Timestamp time.Time

// Now returns the current time as a Timestamp in UTC.
// This function is the sole recommended way to get the current time within the system
// to enforce UTC usage, which is essential for global financial systems to avoid
// timezone-related ambiguity and errors.
func Now() Timestamp {
	return Timestamp(time.Now().UTC())
}

// FromTime converts a standard time.Time object to a Timestamp, ensuring it is in UTC.
// This provides a safe way to integrate with external libraries that may use the
// standard time.Time type.
func FromTime(t time.Time) Timestamp {
	return Timestamp(t.UTC())
}

// ParseTimestamp parses a string formatted according to TimestampFormat into a Timestamp.
// It returns an error if the string is not in the required, exact format, enforcing
// strict data validation at the system's boundaries.
func ParseTimestamp(s string) (Timestamp, error) {
	t, err := time.Parse(TimestampFormat, s)
	if err != nil {
		return Timestamp{}, fmt.Errorf("failed to parse timestamp string %q: %w", s, err)
	}
	return Timestamp(t), nil
}

// FromUnixNano creates a Timestamp from a Unix timestamp given in nanoseconds.
// The resulting Timestamp is in UTC. This is useful for high-performance contexts
// and for interoperating with systems that use Unix epoch time.
func FromUnixNano(ns int64) Timestamp {
	return Timestamp(time.Unix(0, ns).UTC())
}

// Time returns the underlying time.Time object.
// This allows for interoperability with the standard library's time functions.
func (t Timestamp) Time() time.Time {
	return time.Time(t)
}

// String implements the fmt.Stringer interface.
// It returns the timestamp formatted according to the standard TimestampFormat.
func (t Timestamp) String() string {
	return t.Time().Format(TimestampFormat)
}

// MarshalJSON implements the json.Marshaler interface.
// It serializes the Timestamp into a JSON string using the standard format.
// This ensures that all API responses and event logs have consistent timestamp formatting.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It deserializes a JSON string in the standard format into a Timestamp.
// It will fail if the input is not a string or does not match TimestampFormat,
// enforcing fail-closed semantics on data ingestion.
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("timestamp should be a string, got %s: %w", data, err)
	}

	parsedTime, err := time.Parse(TimestampFormat, s)
	if err != nil {
		return fmt.Errorf("failed to parse time string %q with format %q: %w", s, TimestampFormat, err)
	}

	*t = Timestamp(parsedTime)
	return nil
}

// Before reports whether the timestamp t is before u.
func (t Timestamp) Before(u Timestamp) bool {
	return t.Time().Before(u.Time())
}

// After reports whether the timestamp t is after u.
func (t Timestamp) After(u Timestamp) bool {
	return t.Time().After(u.Time())
}

// Equal reports whether t and u represent the same time instant.
// It is nanosecond-sensitive.
func (t Timestamp) Equal(u Timestamp) bool {
	return t.Time().Equal(u.Time())
}

// IsZero reports whether t represents the zero time instant,
// January 1, year 1, 00:00:00 UTC.
func (t Timestamp) IsZero() bool {
	return t.Time().IsZero()
}

// UnixNano returns t as a Unix time, the number of nanoseconds elapsed
// since January 1, 1970 UTC.
func (t Timestamp) UnixNano() int64 {
	return t.Time().UnixNano()
}

// Add returns the timestamp t+d.
func (t Timestamp) Add(d time.Duration) Timestamp {
	return Timestamp(t.Time().Add(d))
}

// Sub returns the duration t-u.
func (t Timestamp) Sub(u Timestamp) time.Duration {
	return t.Time().Sub(u.Time())
}
### END_OF_FILE_COMPLETED ###
```