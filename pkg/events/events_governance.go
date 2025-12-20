```go
// Copyright (c) 2024. All rights reserved.
// This file is part of the Go-based financial infrastructure project.
//
// Defines event structs for governance actions, like `PolicyUpdated` or
// `ProposalEnacted`, ensuring changes to system rules are themselves auditable events.
// These events form the immutable log of all governance decisions and their execution,
// which is critical for regulatory compliance, system auditability, and deterministic state replication.

package events

import (
	"time"

	"github.com/google/uuid"
)

// =================================================================================
// NOTE: Base Event Types
//
// The types `EventType` and `EventHeader` are fundamental to the event system.
// In a complete project structure, they would be defined in a central file like
// `pkg/events/events.go` to be used across all event categories.
// They are defined here to make this file self-contained and compilable.
// =================================================================================

// EventType is a string alias for the type of an event, providing a clear,
// human-readable identifier for different event categories.
type EventType string

// EventHeader contains common fields for all system events.
// It is embedded in specific event structs to ensure consistency in event
// metadata, which is crucial for event sourcing, logging, and auditing.
type EventHeader struct {
	// EventID is a unique identifier for this specific event instance.
	EventID uuid.UUID `json:"eventId"`
	// Timestamp is the UTC time at which the event was created.
	Timestamp time.Time `json:"timestamp"`
	// EventType indicates the specific type of event.
	EventType EventType `json:"eventType"`
	// Version is the schema version of the event payload.
	Version int `json:"version"`
}

// =================================================================================
// Governance Event Type Constants
// =================================================================================

const (
	// PolicyUpdatedEventType is triggered when a system policy is changed.
	PolicyUpdatedEventType EventType = "governance.policy.updated"
	// ProposalCreatedEventType is triggered when a new governance proposal is submitted.
	ProposalCreatedEventType EventType = "governance.proposal.created"
	// ProposalVotedOnEventType is triggered when a vote is cast on a proposal.
	ProposalVotedOnEventType EventType = "governance.proposal.voted_on"
	// ProposalEnactedEventType is triggered when a proposal passes and its changes are applied.
	ProposalEnactedEventType EventType = "governance.proposal.enacted"
	// ProposalRejectedEventType is triggered when a proposal fails to pass.
	ProposalRejectedEventType EventType = "governance.proposal.rejected"
	// ParameterChangedEventType is triggered for each individual system parameter modification.
	ParameterChangedEventType EventType = "governance.parameter.changed"
)

// =================================================================================
// Governance-Specific Data Structures
// =================================================================================

// ProposedChange represents a single, atomic change within a governance proposal.
// It provides a structured format for defining modifications to system parameters.
type ProposedChange struct {
	// Target is the system component or module being changed (e.g., "risk.parameters", "system.fees").
	Target string `json:"target"`
	// Parameter is the specific configuration key to be modified (e.g., "max_leverage", "taker_fee_bps").
	Parameter string `json:"parameter"`
	// NewValue is the proposed new value, represented as a string for universal compatibility.
	// The consuming system is responsible for parsing this value into the correct type.
	NewValue string `json:"newValue"`
}

// VoteOption represents the choice made by a voter in a proposal.
type VoteOption string

const (
	VoteFor     VoteOption = "FOR"
	VoteAgainst VoteOption = "AGAINST"
	VoteAbstain VoteOption = "ABSTAIN"
)

// VoteTally represents the final count of votes for a proposal.
// Using strings for vote counts supports arbitrarily large numbers, common in
// systems with token-weighted voting.
type VoteTally struct {
	ForVotes     string `json:"forVotes"`
	AgainstVotes string `json:"againstVotes"`
	AbstainVotes string `json:"abstainVotes"`
	TotalVotes   string `json:"totalVotes"`
}

// =================================================================================
// Governance Event Definitions
// =================================================================================

// PolicyUpdated event is published when a system policy is changed.
// This provides an auditable record of all policy modifications.
type PolicyUpdated struct {
	EventHeader
	PolicyID           string    `json:"policyId"`          // A unique identifier for the policy being updated.
	PreviousVersion    string    `json:"previousVersion"`   // The version/hash of the policy before the update.
	NewVersion         string    `json:"newVersion"`        // The version/hash of the new policy.
	ChangeDescription  string    `json:"changeDescription"` // A human-readable description of the change.
	EffectiveTimestamp time.Time `json:"effectiveTimestamp"`// When the new policy takes effect.
	AuthorizedBy       string    `json:"authorizedBy"`      // The authority for the change (e.g., ProposalID, AdminID).
}

// ProposalCreated event is published when a new governance proposal is submitted.
type ProposalCreated struct {
	EventHeader
	ProposalID        uuid.UUID        `json:"proposalId"`
	Proposer          string           `json:"proposer"` // Identifier for the entity that submitted the proposal.
	Title             string           `json:"title"`
	Description       string           `json:"description"`
	VotingPeriodStart time.Time        `json:"votingPeriodStart"`
	VotingPeriodEnd   time.Time        `json:"votingPeriodEnd"`
	ProposedChanges   []ProposedChange `json:"proposedChanges"`
}

// ProposalVotedOn event is published when a vote is cast on a proposal.
type ProposalVotedOn struct {
	EventHeader
	ProposalID  uuid.UUID  `json:"proposalId"`
	Voter       string     `json:"voter"`       // Identifier for the voting entity.
	Option      VoteOption `json:"option"`      // The choice made by the voter.
	VotingPower string     `json:"votingPower"` // The weight of the vote, as a string to handle large numbers.
}

// ProposalEnacted event is published when a proposal passes and its changes are applied.
// This is a critical event marking the successful completion of a governance cycle.
type ProposalEnacted struct {
	EventHeader
	ProposalID         uuid.UUID        `json:"proposalId"`
	EnactmentTimestamp time.Time        `json:"enactmentTimestamp"`
	Outcome            VoteTally        `json:"outcome"`
	AppliedChanges     []ProposedChange `json:"appliedChanges"` // A record of the exact changes that were executed.
}

// ProposalRejected event is published when a proposal fails to pass.
type ProposalRejected struct {
	EventHeader
	ProposalID         uuid.UUID `json:"proposalId"`
	RejectionTimestamp time.Time `json:"rejectionTimestamp"`
	Outcome            VoteTally `json:"outcome"`
	Reason             string    `json:"reason"` // e.g., "Failed to meet quorum", "Majority voted against".
}

// ParameterChanged event is published for each individual system parameter modification.
// This provides a granular audit trail and is often triggered by a ProposalEnacted event.
type ParameterChanged struct {
	EventHeader
	ParameterName string `json:"parameterName"` // The fully-qualified name of the parameter (e.g., "risk.max_leverage").
	PreviousValue string `json:"previousValue"` // The value before the change.
	NewValue      string `json:"newValue"`      // The value after the change.
	ChangeContext string `json:"changeContext"` // Reference to the source of the change (e.g., "proposal:a1b2c3d4...").
}

// =================================================================================
// Event Constructor Functions
//
// These functions ensure that events are created with consistent and valid
// metadata, such as a unique event ID, timestamp, and correct event type.
// =================================================================================

// newEventHeader creates a standard header for a new event.
func newEventHeader(eventType EventType) EventHeader {
	return EventHeader{
		EventID:   uuid.New(),
		Timestamp: time.Now().UTC(),
		EventType: eventType,
		Version:   1,
	}
}

// NewPolicyUpdated creates a new PolicyUpdated event.
func NewPolicyUpdated(policyID, prevVersion, newVersion, desc, authorizedBy string, effectiveAt time.Time) *PolicyUpdated {
	return &PolicyUpdated{
		EventHeader:        newEventHeader(PolicyUpdatedEventType),
		PolicyID:           policyID,
		PreviousVersion:    prevVersion,
		NewVersion:         newVersion,
		ChangeDescription:  desc,
		EffectiveTimestamp: effectiveAt,
		AuthorizedBy:       authorizedBy,
	}
}

// NewProposalCreated creates a new ProposalCreated event.
func NewProposalCreated(proposer, title, desc string, start, end time.Time, changes []ProposedChange) *ProposalCreated {
	return &ProposalCreated{
		EventHeader:       newEventHeader(ProposalCreatedEventType),
		ProposalID:        uuid.New(),
		Proposer:          proposer,
		Title:             title,
		Description:       desc,
		VotingPeriodStart: start,
		VotingPeriodEnd:   end,
		ProposedChanges:   changes,
	}
}

// NewProposalVotedOn creates a new ProposalVotedOn event.
func NewProposalVotedOn(proposalID uuid.UUID, voter string, option VoteOption, votingPower string) *ProposalVotedOn {
	return &ProposalVotedOn{
		EventHeader: newEventHeader(ProposalVotedOnEventType),
		ProposalID:  proposalID,
		Voter:       voter,
		Option:      option,
		VotingPower: votingPower,
	}
}

// NewProposalEnacted creates a new ProposalEnacted event.
func NewProposalEnacted(proposalID uuid.UUID, outcome VoteTally, appliedChanges []ProposedChange) *ProposalEnacted {
	return &ProposalEnacted{
		EventHeader:        newEventHeader(ProposalEnactedEventType),
		ProposalID:         proposalID,
		EnactmentTimestamp: time.Now().UTC(),
		Outcome:            outcome,
		AppliedChanges:     appliedChanges,
	}
}

// NewProposalRejected creates a new ProposalRejected event.
func NewProposalRejected(proposalID uuid.UUID, outcome VoteTally, reason string) *ProposalRejected {
	return &ProposalRejected{
		EventHeader:        newEventHeader(ProposalRejectedEventType),
		ProposalID:         proposalID,
		RejectionTimestamp: time.Now().UTC(),
		Outcome:            outcome,
		Reason:             reason,
	}
}

// NewParameterChanged creates a new ParameterChanged event.
func NewParameterChanged(name, prevValue, newValue, context string) *ParameterChanged {
	return &ParameterChanged{
		EventHeader:   newEventHeader(ParameterChangedEventType),
		ParameterName: name,
		PreviousValue: prevValue,
		NewValue:      newValue,
		ChangeContext: context,
	}
}


```