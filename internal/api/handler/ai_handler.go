package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// -----------------------------------------------------------------------------
// Dependencies & Interfaces
// -----------------------------------------------------------------------------

// AIService defines the business logic for AI operations.
// In a full project, this would likely be imported from internal/service or internal/core/ports.
type AIService interface {
	GenerateChatCompletion(ctx context.Context, req ChatRequest) (*ChatResponse, error)
	StreamChatCompletion(ctx context.Context, req ChatRequest) (<-chan ChatStreamChunk, error)
	GenerateImage(ctx context.Context, req ImageRequest) (*ImageResponse, error)
}

// Logger defines a structured logger interface.
type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
	Debug(msg string, args ...any)
}

// -----------------------------------------------------------------------------
// DTOs (Data Transfer Objects)
// -----------------------------------------------------------------------------

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model       string        `json:"model"`
	Messages    []ChatMessage `json:"messages"`
	Temperature float64       `json:"temperature,omitempty"`
	MaxTokens   int           `json:"max_tokens,omitempty"`
}

type ChatResponse struct {
	ID      string      `json:"id"`
	Object  string      `json:"object"`
	Created int64       `json:"created"`
	Message ChatMessage `json:"message"`
	Usage   Usage       `json:"usage"`
}

type ChatStreamChunk struct {
	ID      string `json:"id"`
	Content string `json:"content"` // Delta content
	Done    bool   `json:"done"`
	Error   error  `json:"-"` // Internal error passing
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ImageRequest struct {
	Prompt string `json:"prompt"`
	Size   string `json:"size,omitempty"` // e.g., "1024x1024"
	N      int    `json:"n,omitempty"`
}

type ImageResponse struct {
	Created int64       `json:"created"`
	Data    []ImageItem `json:"data"`
}

type ImageItem struct {
	URL     string `json:"url,omitempty"`
	B64JSON string `json:"b64_json,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

// -----------------------------------------------------------------------------
// Handler Implementation
// -----------------------------------------------------------------------------

// AIHandler manages HTTP requests for AI capabilities.
type AIHandler struct {
	service AIService
	logger  Logger
}

// NewAIHandler initializes a new AIHandler with necessary dependencies.
func NewAIHandler(service AIService, logger Logger) *AIHandler {
	return &AIHandler{
		service: service,
		logger:  logger,
	}
}

// RegisterRoutes attaches the handlers to a standard http.ServeMux.
// This follows Go 1.22+ routing patterns.
func (h *AIHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/ai/chat", h.HandleChatCompletion)
	mux.HandleFunc("POST /api/v1/ai/chat/stream", h.HandleStreamChatCompletion)
	mux.HandleFunc("POST /api/v1/ai/image", h.HandleGenerateImage)
}

// HandleChatCompletion processes a standard request-response chat interaction.
func (h *AIHandler) HandleChatCompletion(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
	defer cancel()

	var req ChatRequest
	if err := h.decodeJSON(w, r, &req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if len(req.Messages) == 0 {
		h.respondError(w, http.StatusBadRequest, "Messages list cannot be empty")
		return
	}

	h.logger.Info("Processing chat completion", "model", req.Model, "msg_count", len(req.Messages))

	resp, err := h.service.GenerateChatCompletion(ctx, req)
	if err != nil {
		h.logger.Error("Failed to generate chat completion", "error", err)
		h.respondError(w, http.StatusInternalServerError, "AI service unavailable")
		return
	}

	h.respondJSON(w, http.StatusOK, resp)
}

// HandleStreamChatCompletion handles Server-Sent Events (SSE) for real-time chat streaming.
func (h *AIHandler) HandleStreamChatCompletion(w http.ResponseWriter, r *http.Request) {
	// Streaming endpoints often require longer timeouts
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Minute)
	defer cancel()

	var req ChatRequest
	if err := h.decodeJSON(w, r, &req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		h.respondError(w, http.StatusInternalServerError, "Streaming not supported")
		return
	}

	h.logger.Info("Starting chat stream", "model", req.Model)

	streamChan, err := h.service.StreamChatCompletion(ctx, req)
	if err != nil {
		h.logger.Error("Failed to initialize stream", "error", err)
		// If headers haven't been flushed, we can send a JSON error.
		// Otherwise, we might need to send a specific SSE error event.
		h.respondError(w, http.StatusInternalServerError, "Failed to start stream")
		return
	}

	for chunk := range streamChan {
		if chunk.Error != nil {
			h.logger.Error("Stream processing error", "error", chunk.Error)
			// Send an error event to the client
			fmt.Fprintf(w, "event: error\ndata: %s\n\n", chunk.Error.Error())
			flusher.Flush()
			return
		}

		// Marshal the chunk data
		data, err := json.Marshal(chunk)
		if err != nil {
			h.logger.Error("Failed to marshal chunk", "error", err)
			continue
		}

		// Write SSE format: data: <json>\n\n
		fmt.Fprintf(w, "data: %s\n\n", data)
		flusher.Flush()

		if chunk.Done {
			fmt.Fprintf(w, "data: [DONE]\n\n")
			flusher.Flush()
			return
		}
	}
}

// HandleGenerateImage processes image generation requests.
func (h *AIHandler) HandleGenerateImage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 90*time.Second)
	defer cancel()

	var req ImageRequest
	if err := h.decodeJSON(w, r, &req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Prompt == "" {
		h.respondError(w, http.StatusBadRequest, "Prompt is required")
		return
	}

	h.logger.Info("Generating image", "prompt_len", len(req.Prompt))

	resp, err := h.service.GenerateImage(ctx, req)
	if err != nil {
		h.logger.Error("Failed to generate image", "error", err)
		h.respondError(w, http.StatusInternalServerError, "Image generation failed")
		return
	}

	h.respondJSON(w, http.StatusOK, resp)
}

// -----------------------------------------------------------------------------
// Helper Methods
// -----------------------------------------------------------------------------

// decodeJSON decodes the request body into the target struct.
// It enforces a maximum body size to prevent DoS attacks.
func (h *AIHandler) decodeJSON(w http.ResponseWriter, r *http.Request, v any) error {
	// Limit request body to 1MB
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(v); err != nil {
		return err
	}

	// Ensure only one JSON object is present
	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("body must only contain a single JSON object")
	}

	return nil
}

// respondJSON writes a JSON response with the given status code.
func (h *AIHandler) respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		h.logger.Error("Failed to encode response", "error", err)
	}
}

// respondError writes a standardized error response.
func (h *AIHandler) respondError(w http.ResponseWriter, code int, message string) {
	h.respondJSON(w, code, ErrorResponse{
		Error: message,
		Code:  code,
	})
}