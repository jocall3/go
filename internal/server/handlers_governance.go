```go
package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.comcom/gorilla/mux"
)

// For the purpose of this file, we'll define the necessary types and interfaces
// that would typically live in other packages (e.g., internal/governance).
// This makes the handler code self-contained and understandable.

// ProposalStatus represents the state of a governance proposal.
type ProposalStatus string

const (
	StatusPending   ProposalStatus = "pending"   // Proposal submitted, not yet active for voting.
	StatusActive    ProposalStatus = "active"    // Proposal is open for voting.
	StatusPassed    ProposalStatus = "passed"    // Voting period ended, quorum and threshold met.
	StatusFailed    ProposalStatus = "failed"    // Voting period ended, quorum or threshold not met.
	StatusExecuted  ProposalStatus = "executed"  // Proposal has been successfully implemented.
	StatusRejected  ProposalStatus = "rejected"  // Proposal was rejected before voting (e.g., invalid).
	StatusCancelled ProposalStatus = "cancelled" // Proposal was cancelled by its author.
)

// ProposalType defines the category of change a proposal introduces.
type ProposalType string

const (
	TypeParameterChange ProposalType = "parameter_change" // e.g., changing a risk limit, fee, etc.
	TypeSystemUpgrade   ProposalType = "system_upgrade"   // e.g., deploying new contract logic.
	TypeText            ProposalType = "text"             // A non-binding proposal to signal intent.
)

// Proposal represents a governance proposal to change a system parameter or rule.
type Proposal struct {
	ID          uuid.UUID       `json:"id"`
	Proposer    string          `json:"proposer"` // Identifier for the proposing entity.
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Type        ProposalType    `json:"type"`
	Payload     json.RawMessage `json:"payload"` // The actual change data, e.g., {"parameter": "max_leverage", "value": "10"}.
	Status      ProposalStatus  `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	VotingStart time.Time       `json:"voting_start"`
	VotingEnd   time.Time       `json:"voting_end"`
	ExecutedAt  *time.Time      `json:"executed_at,omitempty"`
}

// VoteOption represents the choice a voter can make.
type VoteOption string

const (
	VoteYes     VoteOption = "yes"
	VoteNo      VoteOption = "no"
	VoteAbstain VoteOption = "abstain"
)

// Vote represents a single vote cast on a proposal.
type Vote struct {
	ProposalID uuid.UUID  `json:"proposal_id"`
	VoterID    string     `json:"voter_id"` // Identifier for the voting entity.
	Option     VoteOption `json:"option"`
	Weight     uint64     `json:"weight"` // Voting power at the time of casting.
	CreatedAt  time.Time  `json:"created_at"`
}

// ProposalStore defines the interface for interacting with the governance proposal data layer.
// This allows us to mock the storage for testing.
type ProposalStore interface {
	CreateProposal(p *Proposal) error
	GetProposal(id uuid.UUID) (*Proposal, error)
	ListProposals(status ProposalStatus, limit, offset int) ([]*Proposal, error)
	UpdateProposalStatus(id uuid.UUID, status ProposalStatus) error
	CastVote(v *Vote) error
	GetVotes(proposalID uuid.UUID) ([]*Vote, error)
}

// GovernanceHandler holds dependencies for governance-related HTTP handlers.
type GovernanceHandler struct {
	store  ProposalStore
	logger *log.Logger
	// In a real system, this would also include a service for executing proposals.
	// executor governance.Executor
}

// NewGovernanceHandler creates a new GovernanceHandler with its dependencies.
func NewGovernanceHandler(store ProposalStore, logger *log.Logger) *GovernanceHandler {
	return &GovernanceHandler{
		store:  store,
		logger: logger,
	}
}

// RegisterGovernanceRoutes registers the governance API routes with a mux router.
func (h *GovernanceHandler) RegisterGovernanceRoutes(router *mux.Router) {
	// Note: In a real system, these routes would be protected by authentication and authorization middleware.
	router.HandleFunc("/governance/proposals", h.handleCreateProposal).Methods("POST")
	router.HandleFunc("/governance/proposals", h.handleListProposals).Methods("GET")
	router.HandleFunc("/governance/proposals/{id}", h.handleGetProposal).Methods("GET")
	router.HandleFunc("/governance/proposals/{id}/votes", h.handleCastVote).Methods("POST")
	router.HandleFunc("/governance/proposals/{id}/votes", h.handleGetVotes).Methods("GET")
	// The execute endpoint is highly sensitive and should have strict access controls.
	router.HandleFunc("/governance/proposals/{id}/execute", h.handleExecuteProposal).Methods("POST")
}

// handleCreateProposal handles the creation of a new governance proposal.
func (h *GovernanceHandler) handleCreateProposal(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Proposer    string          `json:"proposer"`
		Title       string          `json:"title"`
		Description string          `json:"description"`
		Type        ProposalType    `json:"type"`
		Payload     json.RawMessage `json:"payload"`
		VotingHours int             `json:"voting_hours"` // Duration of the voting period in hours.
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Basic validation
	if req.Title == "" || req.Description == "" || req.Proposer == "" || req.Type == "" {
		writeError(w, http.StatusBadRequest, "proposer, title, description, and type are required")
		return
	}
	if req.VotingHours <= 0 {
		req.VotingHours = 72 // Default to 3 days
	}

	// In a real system, we would validate the proposer's identity and permissions here.
	// We would also validate the payload based on the proposal type.

	now := time.Now().UTC()
	proposal := &Proposal{
		ID:          uuid.New(),
		Proposer:    req.Proposer,
		Title:       req.Title,
		Description: req.Description,
		Type:        req.Type,
		Payload:     req.Payload,
		Status:      StatusActive, // For simplicity, let's make it active immediately. A real system might have a pending/review state.
		CreatedAt:   now,
		VotingStart: now,
		VotingEnd:   now.Add(time.Duration(req.VotingHours) * time.Hour),
	}

	if err := h.store.CreateProposal(proposal); err != nil {
		h.logger.Printf("Error creating proposal: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to create proposal")
		return
	}

	h.logger.Printf("Created new governance proposal %s", proposal.ID)
	writeJSON(w, http.StatusCreated, proposal)
}

// handleListProposals retrieves a list of proposals, optionally filtered by status.
func (h *GovernanceHandler) handleListProposals(w http.ResponseWriter, r *http.Request) {
	status := ProposalStatus(r.URL.Query().Get("status"))
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 100 // Default limit
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0 // Default offset
	}

	proposals, err := h.store.ListProposals(status, limit, offset)
	if err != nil {
		h.logger.Printf("Error listing proposals: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to retrieve proposals")
		return
	}

	writeJSON(w, http.StatusOK, proposals)
}

// handleGetProposal retrieves a single proposal by its ID.
func (h *GovernanceHandler) handleGetProposal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		writeError(w, http.StatusBadRequest, "proposal ID is required")
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid proposal ID format")
		return
	}

	proposal, err := h.store.GetProposal(id)
	if err != nil {
		// This could be a DB error or not found. We should distinguish.
		// For now, assume not found is a common case.
		h.logger.Printf("Could not find proposal %s: %v", id, err)
		writeError(w, http.StatusNotFound, fmt.Sprintf("proposal with ID %s not found", id))
		return
	}

	writeJSON(w, http.StatusOK, proposal)
}

// handleCastVote handles casting a vote on an active proposal.
func (h *GovernanceHandler) handleCastVote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		writeError(w, http.StatusBadRequest, "proposal ID is required")
		return
	}

	proposalID, err := uuid.Parse(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid proposal ID format")
		return
	}

	var req struct {
		VoterID string     `json:"voter_id"`
		Option  VoteOption `json:"option"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.VoterID == "" {
		writeError(w, http.StatusBadRequest, "voter_id is required")
		return
	}
	if req.Option != VoteYes && req.Option != VoteNo && req.Option != VoteAbstain {
		writeError(w, http.StatusBadRequest, "invalid vote option")
		return
	}

	// 1. Fetch the proposal to ensure it's active.
	proposal, err := h.store.GetProposal(proposalID)
	if err != nil {
		writeError(w, http.StatusNotFound, "proposal not found")
		return
	}

	// 2. Check if the proposal is in the active voting period.
	now := time.Now().UTC()
	if proposal.Status != StatusActive || now.Before(proposal.VotingStart) || now.After(proposal.VotingEnd) {
		writeError(w, http.StatusForbidden, "proposal is not active for voting")
		return
	}

	// 3. In a real system, determine the voter's weight.
	// This could be based on token holdings, reputation, etc., at the time of proposal creation.
	// For this example, we'll use a fixed weight.
	voterWeight := uint64(1) // Placeholder

	vote := &Vote{
		ProposalID: proposalID,
		VoterID:    req.VoterID,
		Option:     req.Option,
		Weight:     voterWeight,
		CreatedAt:  now,
	}

	if err := h.store.CastVote(vote); err != nil {
		h.logger.Printf("Error casting vote for proposal %s: %v", proposalID, err)
		// This could be a unique constraint violation if the user already voted.
		writeError(w, http.StatusInternalServerError, "failed to cast vote")
		return
	}

	h.logger.Printf("Voter %s cast vote '%s' on proposal %s", req.VoterID, req.Option, proposalID)
	writeJSON(w, http.StatusCreated, map[string]string{"status": "vote recorded"})
}

// handleGetVotes retrieves all votes for a given proposal.
func (h *GovernanceHandler) handleGetVotes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		writeError(w, http.StatusBadRequest, "proposal ID is required")
		return
	}

	proposalID, err := uuid.Parse(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid proposal ID format")
		return
	}

	votes, err := h.store.GetVotes(proposalID)
	if err != nil {
		h.logger.Printf("Error getting votes for proposal %s: %v", proposalID, err)
		writeError(w, http.StatusInternalServerError, "failed to retrieve votes")
		return
	}

	writeJSON(w, http.StatusOK, votes)
}

// handleExecuteProposal triggers the execution of a passed proposal.
// This is a highly privileged action.
func (h *GovernanceHandler) handleExecuteProposal(w http.ResponseWriter, r *http.Request) {
	// In a real system, this endpoint would be protected by stringent middleware
	// checking for a specific role, an internal service token, or a multisig confirmation.
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		writeError(w, http.StatusBadRequest, "proposal ID is required")
		return
	}

	proposalID, err := uuid.Parse(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid proposal ID format")
		return
	}

	proposal, err := h.store.GetProposal(proposalID)
	if err != nil {
		writeError(w, http.StatusNotFound, "proposal not found")
		return
	}

	// A background process would typically tally votes and move the proposal
	// from 'Active' to 'Passed' or 'Failed'. This handler assumes that has happened.
	if proposal.Status != StatusPassed {
		writeError(w, http.StatusForbidden, fmt.Sprintf("proposal is not in 'passed' state, current state: %s", proposal.Status))
		return
	}

	// The core execution logic would be in a separate service.
	// This service would be responsible for interpreting the proposal.Payload
	// and applying the change to the system state in a transactional, auditable manner.
	h.logger.Printf("Executing proposal %s of type %s", proposal.ID, proposal.Type)
	// executor.Execute(proposal.ID, proposal.Type, proposal.Payload)
	// For this example, we just log and update the status.

	if err := h.store.UpdateProposalStatus(proposalID, StatusExecuted); err != nil {
		h.logger.Printf("Failed to update proposal %s status to executed: %v", proposalID, err)
		// This is a critical failure. The system might need to halt or alert an operator.
		// The execution logic should be idempotent so it can be retried.
		writeError(w, http.StatusInternalServerError, "failed to finalize proposal execution")
		return
	}

	now := time.Now().UTC()
	proposal.Status = StatusExecuted
	proposal.ExecutedAt = &now

	h.logger.Printf("Successfully executed proposal %s", proposal.ID)
	writeJSON(w, http.StatusOK, proposal)
}

// writeJSON is a helper for writing JSON responses.
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

// writeError is a helper for writing JSON error responses.
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

```