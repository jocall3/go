```go
package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/build-trust/ockam-go/processor/key"
	"github.com/fi-bridge/bridge/internal/command"
	"github.com/fi-bridge/bridge/internal/domain/payment"
	"github.com/fi-bridge/bridge/internal/engine"
	"github.com/fi-bridge/bridge/internal/idempotency"
	"github.com/fi-bridge/bridge/internal/query"
	"github.com/fi-bridge/bridge/internal/server/api"
)

// PaymentHandler translates HTTP requests into payment commands and queries.
// It acts as the primary anti-corruption layer between the external world and the core domain.
type PaymentHandler struct {
	engine           *engine.ExecutionEngine
	queryService     query.PaymentReader
	idempotencyStore idempotency.Store
	logger           *log.Logger
}

// NewPaymentHandler creates a new PaymentHandler with necessary dependencies.
func NewPaymentHandler(
	engine *engine.ExecutionEngine,
	queryService query.PaymentReader,
	idempotencyStore idempotency.Store,
	logger *log.Logger,
) *PaymentHandler {
	return &PaymentHandler{
		engine:           engine,
		queryService:     queryService,
		idempotencyStore: idempotencyStore,
		logger:           logger,
	}
}

// RegisterPaymentRoutes registers the payment API routes with the provided router.
// Using the standard library's ServeMux with method-based routing (Go 1.22+)
// keeps dependencies minimal while providing sufficient functionality.
func (h *PaymentHandler) RegisterPaymentRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /v1/payments", h.handleCreatePayment)
	router.HandleFunc("GET /v1/payments/{id}", h.handleGetPayment)
}

// handleCreatePayment handles the creation of a new payment.
// This endpoint is designed to be idempotent and asynchronous, which is critical
// for financial operations to prevent duplicate transactions in the face of network retries.
func (h *PaymentHandler) handleCreatePayment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idempotencyKey := r.Header.Get("Idempotency-Key")
	if idempotencyKey == "" {
		writeErrorResponse(w, http.StatusBadRequest, "Idempotency-Key header is required")
		return
	}

	// 1. Check for a cached response for this idempotency key. This prevents re-processing.
	cachedResponse, found, err := h.idempotencyStore.Get(ctx, idempotencyKey)
	if err != nil {
		h.logger.Printf("ERROR: Failed to check idempotency key %s: %v", idempotencyKey, err)
		writeErrorResponse(w, http.StatusInternalServerError, "Internal server error while checking idempotency")
		return
	}
	if found {
		// Request has been processed before. Return the stored response to ensure
		// the client receives a consistent result on retries.
		w.WriteHeader(cachedResponse.StatusCode)
		w.Header().Set("Content-Type", "application/json")
		w.Write(cachedResponse.Body)
		return
	}

	// 2. Decode and validate the request body.
	var req api.CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleRequestBodyError(w, err)
		return
	}

	if err := validateCreatePaymentRequest(&req); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// 3. Construct the command for the execution engine.
	// This translates the API model into the internal command model.
	amount, _ := decimal.NewFromString(req.Amount) // Validation ensures this won't fail.
	paymentID := uuid.New()

	cmd := &command.CreatePaymentCommand{
		CommandID:            uuid.New(),
		PaymentID:            paymentID,
		SourceAccountID:      req.SourceAccountID,
		DestinationAccountID: req.DestinationAccountID,
		Amount:               amount,
		Currency:             req.Currency,
		IdempotencyKey:       idempotencyKey,
	}

	// 4. Submit the command to the engine. The engine is responsible for all subsequent
	// state transitions. The handler's job is done once the command is accepted.
	if err := h.engine.Submit(ctx, cmd); err != nil {
		h.logger.Printf("ERROR: Failed to submit CreatePaymentCommand %s: %v", cmd.CommandID, err)
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to process payment request")
		return
	}

	// 5. Prepare and cache the successful response before sending it to the client.
	resp := api.CreatePaymentResponse{
		PaymentID: paymentID.String(),
		Status:    string(payment.StatusPending),
	}
	respBody, err := json.Marshal(resp)
	if err != nil {
		// This is a critical server-side failure. The command is submitted, but we can't
		// form a response. This must be logged with high severity.
		h.logger.Printf("FATAL: Failed to marshal successful response for idempotency key %s: %v", idempotencyKey, err)
		writeErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	// The system is asynchronous. We return 202 Accepted to signal that the request
	// has been accepted for processing, but is not yet complete.
	responseToCache := idempotency.Response{
		StatusCode: http.StatusAccepted,
		Body:       respBody,
	}
	// Cache the response for a reasonable duration (e.g., 24 hours).
	if err := h.idempotencyStore.Set(ctx, idempotencyKey, responseToCache, 24*time.Hour); err != nil {
		h.logger.Printf("ERROR: Failed to store idempotency response for key %s: %v", idempotencyKey, err)
		// CRITICAL: The command was submitted, but caching failed. A retry of this request
		// could lead to a duplicate payment. This failure mode requires robust monitoring
		// and potentially a manual reconciliation process. We still return success to the client
		// as the payment *will* be processed.
	}

	// 6. Send the response to the client.
	writeJSONResponse(w, http.StatusAccepted, resp)
}

// handleGetPayment handles retrieving the status and details of a payment.
// This is a read-only operation that queries a separate read-model (CQRS pattern),
// ensuring that read operations do not impact the performance of the write path.
func (h *PaymentHandler) handleGetPayment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	paymentIDStr := r.PathValue("id")

	paymentID, err := uuid.Parse(paymentIDStr)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid payment ID format")
		return
	}

	// Query the read model for the payment's current state.
	paymentView, err := h.queryService.GetPaymentByID(ctx, paymentID)
	if err != nil {
		if errors.Is(err, query.ErrNotFound) {
			writeErrorResponse(w, http.StatusNotFound, "Payment not found")
			return
		}
		h.logger.Printf("ERROR: Failed to query for payment ID %s: %v", paymentID, err)
		writeErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	// Map the internal query view to the public API response.
	resp := api.PaymentResponse{
		PaymentID:            paymentView.PaymentID.String(),
		SourceAccountID:      paymentView.SourceAccountID,
		DestinationAccountID: paymentView.DestinationAccountID,
		Amount:               paymentView.Amount.String(),
		Currency:             paymentView.Currency,
		Status:               string(paymentView.Status),
		CreatedAt:            paymentView.CreatedAt,
		UpdatedAt:            paymentView.UpdatedAt,
	}

	writeJSONResponse(w, http.StatusOK, resp)
}

// validateCreatePaymentRequest performs business logic validation on the request.
// This ensures that only well-formed and logical requests proceed to the command stage.
func validateCreatePaymentRequest(req *api.CreatePaymentRequest) error {
	if req.SourceAccountID == "" {
		return errors.New("source_account_id is required")
	}
	if req.DestinationAccountID == "" {
		return errors.New("destination_account_id is required")
	}
	if req.SourceAccountID == req.DestinationAccountID {
		return errors.New("source and destination accounts cannot be the same")
	}

	if req.Amount == "" {
		return errors.New("amount is required")
	}
	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		return errors.New("amount must be a valid number")
	}
	if !amount.IsPositive() {
		return errors.New("amount must be positive")
	}

	// A robust implementation would use a library or a predefined list of valid currencies.
	if len(req.Currency) != 3 {
		return errors.New("currency must be a 3-letter ISO 4217 code")
	}

	return nil
}

// handleRequestBodyError provides more specific error messages for JSON decoding failures.
func handleRequestBodyError(w http.ResponseWriter, err error) {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	switch {
	case errors.As(err, &syntaxError):
		msg := fmt.Sprintf("Request body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		http.Error(w, msg, http.StatusBadRequest)
	case errors.Is(err, io.ErrUnexpectedEOF):
		msg := "Request body contains badly-formed JSON"
		http.Error(w, msg, http.StatusBadRequest)
	case errors.As(err, &unmarshalTypeError):
		msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at character %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
		http.Error(w, msg, http.StatusBadRequest)
	case errors.Is(err, io.EOF):
		msg := "Request body must not be empty"
		http.Error(w, msg, http.StatusBadRequest)
	default:
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// writeJSONResponse is a helper to write a JSON response with a given status code.
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			// Headers are already written, so we can't send a new error.
			// Log the failure to aid in debugging.
			log.Printf("ERROR: Could not write JSON response: %v", err)
		}
	}
}

// writeErrorResponse is a helper to write a standardized JSON error response.
func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	writeJSONResponse(w, statusCode, api.ErrorResponse{Error: message})
}
### END_OF_FILE_COMPLETED ###
```