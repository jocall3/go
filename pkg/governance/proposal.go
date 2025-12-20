```go
// Copyright (c) 2024. All rights reserved.
// This file is part of the Go-based financial infrastructure project.
//
// Defines the `Proposal` struct and its types, representing a proposed change
// to a system rule, such as updating a risk limit or a fee. This is a
// foundational element of the system's governance layer, ensuring that all
// changes are transparent, auditable, and subject to a formal lifecycle.

package governance

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ProposalID is a unique identifier for a governance proposal.
// It uses a UUID to ensure global uniqueness.
type ProposalID uuid.UUID

// NewProposalID generates a new, random ProposalID.
func NewProposalID() ProposalID {
	return ProposalID(uuid.New())
}

// String returns the string representation of the ProposalID.
func (id ProposalID) String() string {
	return uuid.UUID(id).String()
}

// ParseProposalID parses a string into a ProposalID.
func ParseProposalID(s string) (ProposalID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return ProposalID{}, fmt.Errorf("failed to parse proposal ID: %w", err)
	}
	return ProposalID(id), nil
}

// ProposalType defines the category of a governance proposal.
// This allows the system to correctly interpret the proposal's payload and apply the change.
type ProposalType string

const (
	// ProposalTypeUnknown is the default zero value, indicating an uninitialized or invalid type.
	ProposalTypeUnknown ProposalType = "UNKNOWN"
	// ProposalTypeUpdateRiskParameter changes a risk management parameter (e.g., collateral factor, liquidation threshold).
	ProposalTypeUpdateRiskParameter ProposalType = "UPDATE_RISK_PARAMETER"
	// ProposalTypeUpdateFee changes a system fee (e.g., trading fee, withdrawal fee).
	ProposalTypeUpdateFee ProposalType = "UPDATE_FEE"
	// ProposalTypeAddAsset lists a new asset for trading or collateral.
	ProposalTypeAddAsset ProposalType = "ADD_ASSET"
	// ProposalTypeRemoveAsset delists an existing asset.
	ProposalTypeRemoveAsset ProposalType = "REMOVE_ASSET"
	// ProposalTypeSystemUpgrade proposes a software upgrade or a change to a core system parameter.
	ProposalTypeSystemUpgrade ProposalType = "SYSTEM_UPGRADE"
	// ProposalTypeEmergencyPause triggers a system-wide halt in response to a critical event.
	ProposalTypeEmergencyPause ProposalType = "EMERGENCY_PAUSE"
	// ProposalTypeText is a non-binding proposal for community signaling or discussion.
	ProposalTypeText ProposalType = "TEXT"
)

// IsValid checks if the proposal type is a defined and supported constant.
func (pt ProposalType) IsValid() bool {
	switch pt {
	case ProposalTypeUpdateRiskParameter,
		ProposalTypeUpdateFee,
		ProposalTypeAddAsset,
		ProposalTypeRemoveAsset,
		ProposalTypeSystemUpgrade,
		ProposalTypeEmergencyPause,
		ProposalTypeText:
		return true
	}
	return false
}

// ProposalState represents the current status of a proposal in its lifecycle.
// The state transitions are strictly controlled to ensure a deterministic process.
type ProposalState string

const (
	// ProposalStatePending is for proposals that have been submitted but are not yet active for voting.
	ProposalStatePending ProposalState = "PENDING"
	// ProposalStateActive is for proposals currently in their voting period.
	ProposalStateActive ProposalState = "ACTIVE"
	// ProposalStateSucceeded is for proposals that have passed the voting threshold and are awaiting execution.
	ProposalStateSucceeded ProposalState = "SUCCEEDED"
	// ProposalStateFailed is for proposals that did not meet the voting threshold by the end of the voting period.
	ProposalStateFailed ProposalState = "FAILED"
	// ProposalStateExecuted is for proposals that have been successfully enacted on-chain or in the system.
	ProposalStateExecuted ProposalState = "EXECUTED"
	// ProposalStateExecutionFailed is for proposals that passed voting but failed during the execution step.
	ProposalStateExecutionFailed ProposalState = "EXECUTION_FAILED"
	// ProposalStateExpired is for proposals that succeeded but were not executed within the grace period.
	ProposalStateExpired ProposalState = "EXPIRED"
	// ProposalStateVetoed is for proposals that have been rejected by a veto power, overriding the voting process.
	ProposalStateVetoed ProposalState = "VETOED"
)

// IsValid checks if the proposal state is a defined and supported constant.
func (ps ProposalState) IsValid() bool {
	switch ps {
	case ProposalStatePending,
		ProposalStateActive,
		ProposalStateSucceeded,
		ProposalStateFailed,
		ProposalStateExecuted,
		ProposalStateExecutionFailed,
		ProposalStateExpired,
		ProposalStateVetoed:
		return true
	}
	return false
}

// VoteTally holds the current voting results for a proposal.
// It uses high-precision decimals to prevent floating-point errors in financial calculations.
type VoteTally struct {
	// VotesFor is the total weight of votes in favor of the proposal.
	VotesFor decimal.Decimal `json:"votes_for"`
	// VotesAgainst is the total weight of votes against the proposal.
	VotesAgainst decimal.Decimal `json:"votes_against"`
	// VotesAbstain is the total weight of votes to abstain.
	VotesAbstain decimal.Decimal `json:"votes_abstain"`
	// TotalVotingPower is the total voting power snapshotted when the proposal became active.
	TotalVotingPower decimal.Decimal `json:"total_voting_power"`
}

// NewVoteTally creates an initialized VoteTally struct with zeroed vote counts.
func NewVoteTally(totalPower decimal.Decimal) *VoteTally {
	return &VoteTally{
		VotesFor:         decimal.Zero,
		VotesAgainst:     decimal.Zero,
		VotesAbstain:     decimal.Zero,
		TotalVotingPower: totalPower,
	}
}

// Proposal represents a formal proposal for a change to the system.
// It encapsulates all data related to a proposal's lifecycle, from submission
// to final resolution, ensuring full auditability.
type Proposal struct {
	// ID is the unique identifier for the proposal.
	ID ProposalID `json:"id"`
	// Type indicates the nature of the proposed change.
	Type ProposalType `json:"type"`
	// State is the current status in the proposal lifecycle.
	State ProposalState `json:"state"`
	// Title is a short, human-readable title for the proposal.
	Title string `json:"title"`
	// Description provides a detailed explanation of the proposal's purpose and impact.
	Description string `json:"description"`
	// Proposer is the identifier of the entity that submitted the proposal (e.g., a public key or account ID).
	Proposer string `json:"proposer"`
	// Payload contains the specific, machine-readable details of the proposed change.
	// It is a JSON object whose structure depends on the ProposalType.
	Payload json.RawMessage `json:"payload"`
	// Timestamps for key lifecycle events. All times are in UTC.
	CreatedAt      time.Time `json:"created_at"`
	VotingStartsAt time.Time `json:"voting_starts_at"`
	VotingEndsAt   time.Time `json:"voting_ends_at"`
	ExecutedAt     *time.Time `json:"executed_at,omitempty"`
	// Tally holds the current vote counts.
	Tally VoteTally `json:"tally"`
	// ExecutionHash is a reference (e.g., transaction hash) to the on-chain or system-level execution of the proposal.
	ExecutionHash string `json:"execution_hash,omitempty"`
	// ExecutionError stores any error message if the proposal execution failed.
	ExecutionError string `json:"execution_error,omitempty"`
}

// NewProposal creates and validates a new governance proposal.
// The proposal starts in the PENDING state.
func NewProposal(
	id ProposalID,
	proposalType ProposalType,
	title, description, proposer string,
	payload json.RawMessage,
	votingStartsAt, votingEndsAt time.Time,
	initialTotalPower decimal.Decimal,
) (*Proposal, error) {
	if !proposalType.IsValid() {
		return nil, fmt.Errorf("invalid proposal type: %s", proposalType)
	}
	if title == "" || description == "" || proposer == "" {
		return nil, fmt.Errorf("title, description, and proposer cannot be empty")
	}
	if votingStartsAt.IsZero() || votingEndsAt.IsZero() || !votingEndsAt.After(votingStartsAt) {
		return nil, fmt.Errorf("invalid voting period: start=%v, end=%v", votingStartsAt, votingEndsAt)
	}

	return &Proposal{
		ID:             id,
		Type:           proposalType,
		State:          ProposalStatePending,
		Title:          title,
		Description:    description,
		Proposer:       proposer,
		Payload:        payload,
		CreatedAt:      time.Now().UTC(),
		VotingStartsAt: votingStartsAt.UTC(),
		VotingEndsAt:   votingEndsAt.UTC(),
		Tally:          *NewVoteTally(initialTotalPower),
	}, nil
}

// IsActive checks if the proposal is currently in its voting period at a given time.
func (p *Proposal) IsActive(currentTime time.Time) bool {
	return p.State == ProposalStateActive &&
		!currentTime.Before(p.VotingStartsAt) &&
		currentTime.Before(p.VotingEndsAt)
}

// CanTransitionTo checks if a state transition is valid according to the proposal state machine.
// This enforces deterministic behavior and prevents invalid state changes.
func (p *Proposal) CanTransitionTo(newState ProposalState) bool {
	switch p.State {
	case ProposalStatePending:
		return newState == ProposalStateActive || newState == ProposalStateVetoed
	case ProposalStateActive:
		return newState == ProposalStateSucceeded || newState == ProposalStateFailed || newState == ProposalStateExpired || newState == ProposalStateVetoed
	case ProposalStateSucceeded:
		return newState == ProposalStateExecuted || newState == ProposalStateExecutionFailed || newState == ProposalStateExpired
	case ProposalStateFailed, ProposalStateExecuted, ProposalStateExecutionFailed, ProposalStateExpired, ProposalStateVetoed:
		return false // Terminal states cannot transition further.
	default:
		return false
	}
}

// SetState attempts to transition the proposal to a new state.
// It returns an error if the transition is not allowed by the state machine.
func (p *Proposal) SetState(newState ProposalState) error {
	if !p.CanTransitionTo(newState) {
		return fmt.Errorf("invalid state transition from %s to %s for proposal %s", p.State, newState, p.ID)
	}
	p.State = newState
	return nil
}

```