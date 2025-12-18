package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10" // Standard for validation
)

// -------------------------------------------------------------------------
// Domain Models & Service Interface
// -------------------------------------------------------------------------
// In a real repo, these would be imported from "internal/domain" and "internal/service".
// Defined here to ensure this file is self-contained and compilable as requested,
// while maintaining the architecture of the "best app".

// User represents the domain model for a user.
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"-"` // Never return password
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserService defines the business logic for user operations.
// This interface allows for dependency injection and easy mocking.
type UserService interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, page, limit int) ([]User, int64, error)
}

// -------------------------------------------------------------------------
// DTOs (Data Transfer Objects)
// -------------------------------------------------------------------------

// CreateUserRequest defines the payload for creating a user.
type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=8"`
}

// UpdateUserRequest defines the payload for updating a user.
type UpdateUserRequest struct {
	Email string `json:"email" validate:"omitempty,email"`
	Name  string `json:"name" validate:"omitempty,min=2,max=100"`
}

// UserResponse defines the public view of a user.
type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ListUsersResponse defines the paginated response.
type ListUsersResponse struct {
	Data  []UserResponse `json:"data"`
	Meta  PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
	TotalItems  int64 `json:"total_items"`
	TotalPages  int   `json:"total_pages"`
}

// -------------------------------------------------------------------------
// Handler Implementation
// -------------------------------------------------------------------------

// UserHandler handles HTTP requests for user resources.
type UserHandler struct {
	service   UserService
	logger    *slog.Logger
	validator *validator.Validate
}

// NewUserHandler creates a new instance of UserHandler.
func NewUserHandler(service UserService, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		service:   service,
		logger:    logger,
		validator: validator.New(),
	}
}

// RegisterRoutes registers the user endpoints on the provided mux.
// Uses Go 1.22+ routing patterns.
func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /users", h.Create)
	mux.HandleFunc("GET /users/{id}", h.Get)
	mux.HandleFunc("PUT /users/{id}", h.Update)
	mux.HandleFunc("DELETE /users/{id}", h.Delete)
	mux.HandleFunc("GET /users", h.List)
}

// Create handles the creation of a new user.
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := h.decodeJSON(r, &req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		h.respondValidationError(w, err)
		return
	}

	user := &User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	if err := h.service.Create(r.Context(), user); err != nil {
		if errors.Is(err, ErrUserAlreadyExists) {
			h.respondError(w, http.StatusConflict, "User already exists")
			return
		}
		h.logger.Error("failed to create user", "error", err)
		h.respondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	h.respondJSON(w, http.StatusCreated, toUserResponse(user))
}

// Get retrieves a user by ID.
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.respondError(w, http.StatusBadRequest, "Missing user ID")
		return
	}

	user, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			h.respondError(w, http.StatusNotFound, "User not found")
			return
		}
		h.logger.Error("failed to get user", "id", id, "error", err)
		h.respondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	h.respondJSON(w, http.StatusOK, toUserResponse(user))
}

// Update modifies an existing user.
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.respondError(w, http.StatusBadRequest, "Missing user ID")
		return
	}

	var req UpdateUserRequest
	if err := h.decodeJSON(r, &req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		h.respondValidationError(w, err)
		return
	}

	// Fetch existing to ensure existence and merge updates
	user, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			h.respondError(w, http.StatusNotFound, "User not found")
			return
		}
		h.logger.Error("failed to fetch user for update", "id", id, "error", err)
		h.respondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	// Apply updates
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Name != "" {
		user.Name = req.Name
	}

	if err := h.service.Update(r.Context(), user); err != nil {
		h.logger.Error("failed to update user", "id", id, "error", err)
		h.respondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	h.respondJSON(w, http.StatusOK, toUserResponse(user))
}

// Delete removes a user.
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.respondError(w, http.StatusBadRequest, "Missing user ID")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		if errors.Is(err, ErrUserNotFound) {
			h.respondError(w, http.StatusNotFound, "User not found")
			return
		}
		h.logger.Error("failed to delete user", "id", id, "error", err)
		h.respondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// List retrieves a paginated list of users.
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	page := h.getQueryInt(r, "page", 1)
	limit := h.getQueryInt(r, "limit", 10)

	// Enforce reasonable limits
	if limit > 100 {
		limit = 100
	}

	users, total, err := h.service.List(r.Context(), page, limit)
	if err != nil {
		h.logger.Error("failed to list users", "error", err)
		h.respondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	userResponses := make([]UserResponse, len(users))
	for i, u := range users {
		userResponses[i] = toUserResponse(&u)
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	resp := ListUsersResponse{
		Data: userResponses,
		Meta: PaginationMeta{
			CurrentPage: page,
			PageSize:    limit,
			TotalItems:  total,
			TotalPages:  totalPages,
		},
	}

	h.respondJSON(w, http.StatusOK, resp)
}

// -------------------------------------------------------------------------
// Helpers
// -------------------------------------------------------------------------

func (h *UserHandler) decodeJSON(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return errors.New("body is nil")
	}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Strict JSON decoding
	return decoder.Decode(v)
}

func (h *UserHandler) respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			h.logger.Error("failed to encode response", "error", err)
		}
	}
}

func (h *UserHandler) respondError(w http.ResponseWriter, status int, message string) {
	h.respondJSON(w, status, map[string]string{"error": message})
}

func (h *UserHandler) respondValidationError(w http.ResponseWriter, err error) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		errs := make(map[string]string)
		for _, e := range validationErrors {
			errs[e.Field()] = fmt.Sprintf("failed on tag '%s'", e.Tag())
		}
		h.respondJSON(w, http.StatusBadRequest, map[string]interface{}{
			"error":   "Validation failed",
			"details": errs,
		})
		return
	}
	h.respondError(w, http.StatusBadRequest, "Invalid input")
}

func (h *UserHandler) getQueryInt(r *http.Request, key string, defaultVal int) int {
	valStr := r.URL.Query().Get(key)
	if valStr == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(valStr)
	if err != nil || val < 1 {
		return defaultVal
	}
	return val
}

func toUserResponse(u *User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// -------------------------------------------------------------------------
// Sentinel Errors (Mocking domain errors for self-containment)
// -------------------------------------------------------------------------

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)