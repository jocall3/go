```go
package governance

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// ProposalState represents the lifecycle state of a governance proposal.
// The states follow a deterministic progression, ensuring auditability and predictability.
type ProposalState int

const (
	// Proposed is the initial state of a proposal after submission.
	// It is pending the start of the voting period.
	Proposed ProposalState = iota

	// Voting is the state where the proposal is open for votes.
	Voting

	// Succeeded is the state when a proposal has passed the voting thresholds
	// but is waiting for the enactment delay period to pass.
	Succeeded

	// Failed is a terminal state for a proposal that did not meet the
	// required quorum or pass thresholds.
	Failed

	// Enacted is a terminal state for a proposal whose changes have been
	// successfully applied to the system.
	Enacted

	// Vetoed is a terminal state for a proposal that was explicitly rejected
	// by a safety or administrative override. This provides a fail-safe mechanism.
	Vetoed
)

// String provides a human-readable representation of the ProposalState.
func (s ProposalState) String() string {
	switch s {
	case Proposed:
		return "Proposed"
	case Voting:
		return "Voting"
	case Succeeded:
		return "Succeeded"
	case Failed:
		return "Failed"
	case Enacted:
		return "Enacted"
	case Vetoed:
		return "Vetoed"
	default:
		return fmt.Sprintf("Unknown(%d)", s)
	}
}

var (
	ErrProposalNotFound      = errors.New("proposal not found")
	ErrInvalidProposalState  = errors.New("operation not allowed in current proposal state")
	ErrVotingPeriodNotActive = errors.New("voting period is not active")
	ErrVoterAlreadyVoted     = errors.New("voter has already voted on this proposal")
	ErrProposalExists        = errors.New("proposal with this ID already exists")
)

// Change defines a proposed modification to the system.
// In a real-world system, this would likely be an interface with different concrete types
// for various parameter changes (e.g., fee changes, risk limit adjustments),
// allowing for type-safe application of changes.
type Change struct {
	// Parameter identifies the system parameter to be changed (e.g., "risk.max_leverage").
	Parameter string
	// NewValue represents the proposed new value for the parameter.
	NewValue string
}

// Proposal represents a single governance proposal, containing its state,
// content, and voting tally.
type Proposal struct {
	ID              string
	State           ProposalState
	Description     string
	Change          Change
	SubmitTime      time.Time
	VotingStartTime time.Time
	VotingEndTime   time.Time
	EnactmentTime   time.Time
	VotesFor        uint64
	VotesAgainst    uint64
	VoterRecords    map[string]bool // Key: VoterID, Value: InFavor
	mu              sync.RWMutex    // Protects the internal state of this proposal
}

// Vote represents a single vote cast on a proposal.
type Vote struct {
	VoterID    string
	ProposalID string
	InFavor    bool
	Timestamp  time.Time
}

// Config holds the parameters that define the rules of the governance process.
type Config struct {
	// VotingPeriod is the duration for which a proposal is open for voting.
	VotingPeriod time.Duration
	// EnactmentDelay is the "cool-down" period after a proposal succeeds before
	// its changes are applied. This allows operators to prepare or intervene.
	EnactmentDelay time.Duration
	// QuorumThreshold is the minimum percentage of total voting power that must
	// participate for a vote to be considered valid (e.g., 0.40 for 40%).
	QuorumThreshold float64
	// PassThreshold is the minimum percentage of 'yes' votes (of the total votes cast)
	// required for a proposal to pass, assuming quorum is met (e.g., 0.66 for 66%).
	PassThreshold float64
	// TotalVotingPower represents the total number of possible votes in the system.
	// This is the denominator for the quorum calculation.
	TotalVotingPower uint64
}

// Clock is an interface for time-related operations, allowing for deterministic testing.
type Clock interface {
	Now() time.Time
}

// SystemClock is a concrete implementation of the Clock interface using the system's time.
type SystemClock struct{}

// Now returns the current system time.
func (c SystemClock) Now() time.Time {
	return time.Now().UTC()
}

// ChangeApplier is an interface for applying enacted governance proposals.
// This decouples the governance engine from the specific system components it modifies,
// adhering to the principle of separation of concerns.
type ChangeApplier interface {
	Apply(change Change) error
}

// Logger is an interface for structured logging, crucial for creating an immutable audit trail.
type Logger interface {
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

// Engine is the state machine that manages the lifecycle of governance proposals.
// It is the core of the automated governance system.
type Engine struct {
	proposals map[string]*Proposal
	config    Config
	clock     Clock
	applier   ChangeApplier
	logger    Logger
	mu        sync.RWMutex // Protects the proposals map
}

// NewEngine creates and initializes a new governance Engine.
func NewEngine(config Config, clock Clock, applier ChangeApplier, logger Logger) *Engine {
	return &Engine{
		proposals: make(map[string]*Proposal),
		config:    config,
		clock:     clock,
		applier:   applier,
		logger:    logger,
	}
}

// SubmitProposal creates a new proposal and adds it to the engine.
// The proposal ID must be unique.
func (e *Engine) SubmitProposal(id, description string, change Change) (*Proposal, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, exists := e.proposals[id]; exists {
		return nil, fmt.Errorf("%w: %s", ErrProposalExists, id)
	}

	now := e.clock.Now()
	votingStartTime := now
	votingEndTime := votingStartTime.Add(e.config.VotingPeriod)
	enactmentTime := votingEndTime.Add(e.config.EnactmentDelay)

	p := &Proposal{
		ID:              id,
		State:           Proposed,
		Description:     description,
		Change:          change,
		SubmitTime:      now,
		VotingStartTime: votingStartTime,
		VotingEndTime:   votingEndTime,
		EnactmentTime:   enactmentTime,
		VoterRecords:    make(map[string]bool),
	}

	e.proposals[id] = p
	e.logger.Info("Proposal submitted", "proposalID", p.ID, "state", p.State.String(), "votingEndTime", p.VotingEndTime)
	return p, nil
}

// CastVote records a vote for a specific proposal.
// It enforces several invariants: the proposal must exist, be in the 'Voting' state,
// and the voter must not have already voted.
func (e *Engine) CastVote(vote Vote) error {
	p, exists := e.GetProposal(vote.ProposalID)
	if !exists {
		return ErrProposalNotFound
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.State != Voting {
		return fmt.Errorf("%w: proposal is in state %s", ErrInvalidProposalState, p.State.String())
	}

	now := e.clock.Now()
	if now.Before(p.VotingStartTime) || now.After(p.VotingEndTime) {
		return ErrVotingPeriodNotActive
	}

	if _, voted := p.VoterRecords[vote.VoterID]; voted {
		return ErrVoterAlreadyVoted
	}

	p.VoterRecords[vote.VoterID] = vote.InFavor
	if vote.InFavor {
		p.VotesFor++
	} else {
		p.VotesAgainst++
	}

	e.logger.Info("Vote cast", "proposalID", p.ID, "voterID", vote.VoterID, "inFavor", vote.InFavor)
	return nil
}

// Tick advances the state of all proposals based on the current time.
// This method is designed to be called periodically (e.g., via a time.Ticker).
// It is the heart of the state machine, ensuring proposals move through their lifecycle.
func (e *Engine) Tick() {
	now := e.clock.Now()

	// Create a snapshot of proposals to check to avoid holding the engine lock
	// during individual proposal processing.
	e.mu.RLock()
	proposalsToCheck := make([]*Proposal, 0, len(e.proposals))
	for _, p := range e.proposals {
		proposalsToCheck = append(proposalsToCheck, p)
	}
	e.mu.RUnlock()

	for _, p := range proposalsToCheck {
		p.mu.Lock()
		// State transitions are idempotent and based on the current state.
		switch p.State {
		case Proposed:
			if !now.Before(p.VotingStartTime) {
				e.transitionToVoting(p)
			}
		case Voting:
			if !now.Before(p.VotingEndTime) {
				e.transitionToFinished(p)
			}
		case Succeeded:
			if !now.Before(p.EnactmentTime) {
				e.transitionToEnacted(p)
			}
		}
		p.mu.Unlock()
	}
}

// GetProposal retrieves a proposal by its ID in a thread-safe manner.
func (e *Engine) GetProposal(id string) (*Proposal, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	p, exists := e.proposals[id]
	return p, exists
}

// transitionToVoting moves a proposal from Proposed to Voting.
func (e *Engine) transitionToVoting(p *Proposal) {
	p.State = Voting
	e.logger.Info("Proposal state changed", "proposalID", p.ID, "oldState", Proposed.String(), "newState", p.State.String())
}

// transitionToFinished evaluates a completed vote and moves the proposal to Succeeded or Failed.
// This is where the core governance rules (quorum, threshold) are enforced.
func (e *Engine) transitionToFinished(p *Proposal) {
	totalVotes := p.VotesFor + p.VotesAgainst

	// Invariant: TotalVotingPower must be non-zero to avoid division by zero.
	if e.config.TotalVotingPower == 0 {
		p.State = Failed
		e.logger.Error("TotalVotingPower is zero, cannot calculate quorum. Proposal failed.", "proposalID", p.ID)
		return
	}

	participation := float64(totalVotes) / float64(e.config.TotalVotingPower)
	if participation < e.config.QuorumThreshold {
		p.State = Failed
		e.logger.Info("Proposal failed: quorum not met", "proposalID", p.ID, "participation", participation, "required", e.config.QuorumThreshold)
		return
	}

	// Invariant: totalVotes must be non-zero if quorum is met (and quorum > 0).
	if totalVotes == 0 {
		p.State = Failed
		e.logger.Info("Proposal failed: quorum met but no votes cast", "proposalID", p.ID)
		return
	}

	passRate := float64(p.VotesFor) / float64(totalVotes)
	if passRate >= e.config.PassThreshold {
		p.State = Succeeded
		e.logger.Info("Proposal succeeded", "proposalID", p.ID, "passRate", passRate, "required", e.config.PassThreshold, "enactmentTime", p.EnactmentTime)
	} else {
		p.State = Failed
		e.logger.Info("Proposal failed: pass threshold not met", "proposalID", p.ID, "passRate", passRate, "required", e.config.PassThreshold)
	}
}

// transitionToEnacted applies the change of a Succeeded proposal.
// This is a critical step where governance translates into system change.
// It embodies the "fail-closed" principle: if application fails, the proposal is not enacted.
func (e *Engine) transitionToEnacted(p *Proposal) {
	e.logger.Info("Attempting to enact proposal", "proposalID", p.ID, "change", p.Change)
	if err := e.applier.Apply(p.Change); err != nil {
		// Fail-closed: If the change cannot be applied, the proposal is marked as Failed.
		// This prevents the system from entering an inconsistent state.
		// An alert/monitoring system should catch this for manual intervention.
		p.State = Failed
		e.logger.Error("Failed to enact proposal", "proposalID", p.ID, "error", err)
	} else {
		p.State = Enacted
		e.logger.Info("Proposal enacted successfully", "proposalID", p.ID, "newState", p.State.String())
	}
}

```