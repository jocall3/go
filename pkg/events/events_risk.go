```go
// Copyright (c) 2024. All rights reserved.
// This file is part of the Go-based financial infrastructure project.
//
// Defines event structs related to risk, such as `RiskCheckFailed` and
// `LimitBreached`. These provide a crucial audit trail for why certain
// actions were blocked.

package events

import (
	"time"

	"github.com/google/uuid"
)

// ========================================================================================
// NOTE: The following types are common across all events and would typically reside in a
// central `pkg/events/events.go` file. They are included here to make this file's
// code complete and understandable in isolation.
// ========================================================================================

// EventType is a string representing the type of an event, typically in a
// namespaced format (e.g., "risk.limit.breached").
type EventType string

// EventHeader contains common fields for all system events, providing essential
// metadata for tracing, ordering, and identification.
type EventHeader struct {
	ID        uuid.UUID `json:"id"`         // Unique identifier for the event instance.
	Type      EventType `json:"type"`       // The type of the event.
	Timestamp time.Time `json:"timestamp"`  // UTC timestamp of when the event was generated.
	Source    string    `json:"source"`     // The service or component that generated the event.
	Version   string    `json:"version"`    // The version of the event schema.
	TraceID   uuid.UUID `json:"trace_id"`   // ID for tracing a request across multiple services.
}

// NewEventHeader creates a new EventHeader with a new UUID and current timestamp.
// The traceID should be propagated from the originating request or context.
func NewEventHeader(source string, eventType EventType, traceID uuid.UUID) EventHeader {
	return EventHeader{
		ID:        uuid.New(),
		Type:      eventType,
		Timestamp: time.Now().UTC(),
		Source:    source,
		Version:   "1.0.0", // Default schema version
		TraceID:   traceID,
	}
}

// ========================================================================================
// Risk Event Type Constants
// ========================================================================================

const (
	// RiskCheckFailedType is the event type for a generic risk check failure.
	RiskCheckFailedType EventType = "risk.check.failed"
	// LimitBreachedType is the event type for when a specific risk limit is breached.
	LimitBreachedType EventType = "risk.limit.breached"
)

// ========================================================================================
// RiskCheckFailed Event
// ========================================================================================

// RiskCheckFailed is published when a proposed action (e.g., an order) is
// rejected by the risk management system for reasons other than a simple
// limit breach. This provides a detailed audit trail of risk decisions.
type RiskCheckFailed struct {
	EventHeader
	ActionID    string                 `json:"action_id"`      // ID of the action that was checked (e.g., OrderID, WithdrawalID).
	ActionType  string                 `json:"action_type"`    // Type of action (e.g., "PlaceOrder", "RequestWithdrawal").
	AccountID   string                 `json:"account_id"`     // Account associated with the action.
	CheckType   string                 `json:"check_type"`     // The specific risk check that failed (e.g., "CreditCheck", "VelocityCheck").
	Reason      string                 `json:"reason"`         // Human-readable reason for failure.
	FailureData map[string]interface{} `json:"failure_data"`   // Structured data about the failure (e.g., {"required_margin": "1000.50", "available_collateral": "500.25"}).
}

// RiskCheckFailedPayload contains the specific data for a RiskCheckFailed event.
type RiskCheckFailedPayload struct {
	TraceID     uuid.UUID
	ActionID    string
	ActionType  string
	AccountID   string
	CheckType   string
	Reason      string
	FailureData map[string]interface{}
}

// NewRiskCheckFailed creates a new RiskCheckFailed event.
// The `source` identifies the risk engine or service that performed the check.
func NewRiskCheckFailed(source string, payload RiskCheckFailedPayload) RiskCheckFailed {
	return RiskCheckFailed{
		EventHeader: NewEventHeader(source, RiskCheckFailedType, payload.TraceID),
		ActionID:    payload.ActionID,
		ActionType:  payload.ActionType,
		AccountID:   payload.AccountID,
		CheckType:   payload.CheckType,
		Reason:      payload.Reason,
		FailureData: payload.FailureData,
	}
}

// ========================================================================================
// LimitBreached Event
// ========================================================================================

// LimitBreached is published when an action is blocked because it would exceed a
// pre-defined risk limit. This is a critical event for monitoring and alerting on
// account or system-wide risk thresholds.
type LimitBreached struct {
	EventHeader
	ActionID      string `json:"action_id,omitempty"`      // ID of the action that caused the breach (e.g., OrderID).
	AccountID     string `json:"account_id,omitempty"`     // Account that breached the limit.
	PartyID       string `json:"party_id,omitempty"`       // Party/entity that breached the limit.
	LimitType     string `json:"limit_type"`               // e.g., "PositionLimit", "ExposureLimit", "OrderRate".
	LimitScope    string `json:"limit_scope"`              // e.g., "Instrument", "Account", "Market".
	LimitValue    string `json:"limit_value"`              // The configured limit value (as a string for precision).
	BreachedValue string `json:"breached_value"`           // The value that caused the breach (as a string for precision).
	InstrumentID  string `json:"instrument_id,omitempty"`  // Instrument related to the limit, if applicable.
}

// LimitBreachedPayload contains the specific data for a LimitBreached event.
type LimitBreachedPayload struct {
	TraceID       uuid.UUID
	ActionID      string
	AccountID     string
	PartyID       string
	LimitType     string
	LimitScope    string
	LimitValue    string
	BreachedValue string
	InstrumentID  string
}

// NewLimitBreached creates a new LimitBreached event.
// The `source` identifies the risk engine or service that detected the breach.
func NewLimitBreached(source string, payload LimitBreachedPayload) LimitBreached {
	return LimitBreached{
		EventHeader:   NewEventHeader(source, LimitBreachedType, payload.TraceID),
		ActionID:      payload.ActionID,
		AccountID:     payload.AccountID,
		PartyID:       payload.PartyID,
		LimitType:     payload.LimitType,
		LimitScope:    payload.LimitScope,
		LimitValue:    payload.LimitValue,
		BreachedValue: payload.BreachedValue,
		InstrumentID:  payload.InstrumentID,
	}
}

```