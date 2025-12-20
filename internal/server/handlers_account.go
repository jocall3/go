```go
package server

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/your-org/your-project/internal/domain/account"
	"github.com/your-org/your-project/internal/domain/ledger"
	"github.com/your-org/your-project/internal/service"
)

// AccountHandler handles HTTP requests for account-related operations.
// It acts as the primary interface for developers to interact with the account
// and ledger primitives of the system.
type AccountHandler struct {
	logger         *slog.Logger
	accountService service.AccountService
	ledgerService  service.LedgerService
}

// NewAccountHandler creates a new AccountHandler with the necessary dependencies.
func NewAccountHandler(
	logger *slog.Logger,
	accountService service.AccountService,
	ledgerService service.LedgerService,
) *AccountHandler {
	return &AccountHandler{
		logger:         logger,
		accountService: accountService,
		ledgerService:  ledgerService,
	}
}

// RegisterRoutes connects the account handler's methods to the HTTP router.
func (h *AccountHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/accounts", h.createAccount).Methods("POST")
	router.HandleFunc("/accounts", h.listAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}", h.getAccount).Methods("GET")
	router.HandleFunc("/accounts/{id}/balance", h.getAccountBalance).Methods("GET")
}

// =============================================================================
// Request & Response Types
// =============================================================================

// CreateAccountRequest defines the structure for a new account creation request.
type CreateAccountRequest struct {
	OwnerID  uuid.UUID              `json:"owner_id"`
	Currency string                 `json:"currency"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// AccountResponse defines the standard structure for a single account returned by the API.
type AccountResponse struct {
	ID        uuid.UUID              `json:"id"`
	OwnerID   uuid.UUID              `json:"owner_id"`
	Currency  string                 `json:"currency"`
	Status    string                 `json:"status"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

// BalanceResponse defines the structure for an account's balance.
// It distinguishes between different types of balances, which is critical for
// financial accuracy and risk management.
type BalanceResponse struct {
	AccountID        uuid.UUID `json:"account_id"`
	Currency         string    `json:"currency"`
	PostedBalance    int64     `json:"posted_balance"`    // Settled funds, part of the official record.
	PendingBalance   int64     `json:"pending_balance"`   // Uncleared funds (e.g., incoming/outgoing transfers).
	AvailableBalance int64     `json:"available_balance"` // Posted - Pending Outgoing - Holds. The spendable amount.
	Timestamp        time.Time `json:"timestamp"`         // The time the balance was calculated.
}

// =============================================================================
// Handlers
// =============================================================================

// createAccount handles the creation of a new financial account.
// POST /accounts
func (h *AccountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("handler", "createAccount")

	var req CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Warn("Failed to decode request body", "error", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Basic validation
	if req.OwnerID == uuid.Nil {
		respondWithError(w, http.StatusBadRequest, "owner_id is required")
		return
	}
	if req.Currency == "" {
		// In a real system, this would be validated against a list of supported currencies.
		respondWithError(w, http.StatusBadRequest, "currency is required")
		return
	}

	params := service.CreateAccountParams{
		OwnerID:  req.OwnerID,
		Currency: req.Currency,
		Metadata: req.Metadata,
	}

	newAccount, err := h.accountService.CreateAccount(ctx, params)
	if err != nil {
		log.Error("Failed to create account", "error", err)
		// TODO: Differentiate between user errors (e.g., duplicate) and server errors.
		respondWithError(w, http.StatusInternalServerError, "Could not create account")
		return
	}

	log.Info("Account created successfully", "account_id", newAccount.ID)
	respondWithJSON(w, http.StatusCreated, toAccountResponse(newAccount))
}

// getAccount retrieves a single account by its ID.
// GET /accounts/{id}
func (h *AccountHandler) getAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("handler", "getAccount")
	vars := mux.Vars(r)

	accountID, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Warn("Invalid account ID format", "id", vars["id"], "error", err)
		respondWithError(w, http.StatusBadRequest, "Invalid account ID format")
		return
	}

	log = log.With("account_id", accountID)

	acc, err := h.accountService.GetAccount(ctx, accountID)
	if err != nil {
		if errors.Is(err, account.ErrAccountNotFound) {
			log.Warn("Account not found")
			respondWithError(w, http.StatusNotFound, "Account not found")
		} else {
			log.Error("Failed to retrieve account", "error", err)
			respondWithError(w, http.StatusInternalServerError, "Could not retrieve account")
		}
		return
	}

	respondWithJSON(w, http.StatusOK, toAccountResponse(acc))
}

// listAccounts retrieves a paginated list of accounts.
// GET /accounts?limit=20&offset=0
func (h *AccountHandler) listAccounts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("handler", "listAccounts")

	limit, offset, err := parsePagination(r)
	if err != nil {
		log.Warn("Invalid pagination parameters", "error", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	params := service.ListAccountsParams{
		Limit:  limit,
		Offset: offset,
	}

	accounts, err := h.accountService.ListAccounts(ctx, params)
	if err != nil {
		log.Error("Failed to list accounts", "error", err)
		respondWithError(w, http.StatusInternalServerError, "Could not list accounts")
		return
	}

	// Even if the list is empty, return a 200 OK with an empty array.
	response := make([]AccountResponse, len(accounts))
	for i, acc := range accounts {
		response[i] = toAccountResponse(acc)
	}

	respondWithJSON(w, http.StatusOK, response)
}

// getAccountBalance retrieves the balance for a specific account.
// GET /accounts/{id}/balance
func (h *AccountHandler) getAccountBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("handler", "getAccountBalance")
	vars := mux.Vars(r)

	accountID, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Warn("Invalid account ID format", "id", vars["id"], "error", err)
		respondWithError(w, http.StatusBadRequest, "Invalid account ID format")
		return
	}

	log = log.With("account_id", accountID)

	// First, verify the account exists. This prevents querying balances for non-existent accounts.
	acc, err := h.accountService.GetAccount(ctx, accountID)
	if err != nil {
		if errors.Is(err, account.ErrAccountNotFound) {
			log.Warn("Account not found")
			respondWithError(w, http.StatusNotFound, "Account not found")
		} else {
			log.Error("Failed to retrieve account before getting balance", "error", err)
			respondWithError(w, http.StatusInternalServerError, "Could not retrieve account")
		}
		return
	}

	balance, err := h.ledgerService.GetAccountBalance(ctx, accountID)
	if err != nil {
		log.Error("Failed to get account balance", "error", err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve account balance")
		return
	}

	respondWithJSON(w, http.StatusOK, toBalanceResponse(acc, balance))
}

// =============================================================================
// Helpers & Mappers
// =============================================================================

// toAccountResponse maps the internal domain account model to the public API response model.
func toAccountResponse(acc account.Account) AccountResponse {
	return AccountResponse{
		ID:        acc.ID,
		OwnerID:   acc.OwnerID,
		Currency:  acc.Currency,
		Status:    acc.Status.String(),
		Metadata:  acc.Metadata,
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
	}
}

// toBalanceResponse maps the internal domain balance model to the public API response model.
func toBalanceResponse(acc account.Account, bal ledger.Balance) BalanceResponse {
	return BalanceResponse{
		AccountID:        bal.AccountID,
		Currency:         acc.Currency, // Currency comes from the account entity for consistency.
		PostedBalance:    bal.PostedBalance,
		PendingBalance:   bal.PendingBalance,
		AvailableBalance: bal.AvailableBalance,
		Timestamp:        bal.Timestamp,
	}
}

// parsePagination extracts limit and offset from query parameters with defaults and validation.
func parsePagination(r *http.Request) (limit, offset int, err error) {
	const defaultLimit = 20
	const maxLimit = 100
	const defaultOffset = 0

	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limit = defaultLimit
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			return 0, 0, errors.New("limit must be a positive integer")
		}
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr == "" {
		offset = defaultOffset
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			return 0, 0, errors.New("offset must be a non-negative integer")
		}
	}

	return limit, offset, nil
}

```