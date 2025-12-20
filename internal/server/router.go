```go
package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/your-org/your-project/internal/handlers"
	"github.com/your-org/your-project/internal/middleware"
)

// NewRouter creates and configures a new mux.Router with all the application routes.
// It establishes a clear, versioned, and secure entry point for all external interactions,
// mapping HTTP requests to their corresponding handler logic.
//
// The routing structure is designed with the following principles:
// - Separation of Concerns: Routes are grouped by domain (e.g., accounts, transfers, risk).
// - Versioning: All API routes are prefixed with a version number (e.g., /api/v1) to allow for future evolution.
// - Security by Default: Middleware for logging, recovery, and authentication is applied broadly,
//   with more granular authorization applied to sensitive route groups.
// - Observability: Includes standard endpoints for health checks (/health) and metrics (/metrics).
// - Operational Control: Provides explicit admin endpoints for system-level actions like halting or resuming operations,
//   embodying the "fail-closed" and "powerful for operators" principles.
func NewRouter() *mux.Router {
	// Create a new main router. StrictSlash(true) ensures that paths like "/api/" and "/api" are treated the same.
	router := mux.NewRouter().StrictSlash(true)

	// --- Global Middleware ---
	// These are applied to every single request that the router receives.
	// The order is important as they are executed in a chain.
	router.Use(middleware.RequestID) // Injects a unique ID into each request context for traceability.
	router.Use(middleware.Logging)   // Logs request details, crucial for audit and debugging.
	router.Use(middleware.Recovery)  // Recovers from panics and returns a 500 error, preventing crashes.

	// --- Public Endpoints ---
	// These endpoints do not require authentication and are used for service health monitoring and observability.
	router.HandleFunc("/health", handlers.HealthCheckHandler).Methods(http.MethodGet)
	router.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet) // Exposes Prometheus metrics.

	// --- API Version 1 Sub-router ---
	// All v1 API routes are grouped under the "/api/v1" prefix.
	// This sub-router has its own middleware chain that applies to all its child routes.
	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(middleware.Authentication) // Enforces that a valid identity is present for all v1 API calls.
	apiV1.Use(middleware.RateLimiter)    // Applies a general rate limit to protect the API from abuse.

	// --- Account Management Routes ---
	// Handles the creation, retrieval, and management of financial accounts.
	accountsRouter := apiV1.PathPrefix("/accounts").Subrouter()
	accountsRouter.HandleFunc("", handlers.CreateAccountHandler).Methods(http.MethodPost)
	accountsRouter.HandleFunc("", handlers.ListAccountsHandler).Methods(http.MethodGet)
	accountsRouter.HandleFunc("/{account_id}", handlers.GetAccountHandler).Methods(http.MethodGet)
	accountsRouter.HandleFunc("/{account_id}/balances", handlers.GetAccountBalancesHandler).Methods(http.MethodGet)
	accountsRouter.HandleFunc("/{account_id}/history", handlers.GetAccountHistoryHandler).Methods(http.MethodGet)

	// --- Transfer & Execution Routes ---
	// Manages the initiation and lifecycle of funds transfers. Idempotency is a key concern,
	// handled within the handler logic using a client-provided key.
	transfersRouter := apiV1.PathPrefix("/transfers").Subrouter()
	transfersRouter.HandleFunc("", handlers.CreateTransferHandler).Methods(http.MethodPost)
	transfersRouter.HandleFunc("/{transfer_id}", handlers.GetTransferHandler).Methods(http.MethodGet)

	// --- Ledger Routes (Auditability) ---
	// Provides read-only access to ledger data, forming the core of the system's auditability.
	// Access is highly restricted to roles like 'auditor' or 'admin'.
	ledgerRouter := apiV1.PathPrefix("/ledger").Subrouter()
	ledgerRouter.Use(middleware.Authorize("auditor", "admin"))
	ledgerRouter.HandleFunc("/entries", handlers.ListLedgerEntriesHandler).Methods(http.MethodGet)
	ledgerRouter.HandleFunc("/entries/{entry_id}", handlers.GetLedgerEntryHandler).Methods(http.MethodGet)

	// --- Risk Management Routes ---
	// Endpoints for configuring and querying the risk management system.
	// These are highly sensitive and require specific 'risk_manager' or 'admin' privileges.
	riskRouter := apiV1.PathPrefix("/risk").Subrouter()
	riskRouter.Use(middleware.Authorize("risk_manager", "admin"))
	riskRouter.HandleFunc("/limits/{account_id}", handlers.GetRiskLimitsHandler).Methods(http.MethodGet)
	riskRouter.HandleFunc("/limits", handlers.SetRiskLimitsHandler).Methods(http.MethodPost)
	riskRouter.HandleFunc("/exposure", handlers.GetSystemExposureHandler).Methods(http.MethodGet)

	// --- Settlement Routes ---
	// Manages the settlement process, which finalizes transfers between accounts.
	// Requires specific operational privileges.
	settlementRouter := apiV1.PathPrefix("/settlements").Subrouter()
	settlementRouter.Use(middleware.Authorize("settlement_operator", "admin"))
	settlementRouter.HandleFunc("", handlers.InitiateSettlementHandler).Methods(http.MethodPost)
	settlementRouter.HandleFunc("/{settlement_id}", handlers.GetSettlementStatusHandler).Methods(http.MethodGet)

	// --- Governance and Admin Routes ---
	// System-level configuration and control. Access is restricted to the highest privilege level ('admin').
	// These endpoints allow operators to safely manage the state of the entire system.
	adminRouter := apiV1.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.Authorize("admin"))
	adminRouter.HandleFunc("/system/status", handlers.GetSystemStatusHandler).Methods(http.MethodGet)
	adminRouter.HandleFunc("/system/halt", handlers.HaltSystemHandler).Methods(http.MethodPost)     // Implements fail-closed principle.
	adminRouter.HandleFunc("/system/resume", handlers.ResumeSystemHandler).Methods(http.MethodPost) // Allows controlled resumption of operations.

	return router
}

```