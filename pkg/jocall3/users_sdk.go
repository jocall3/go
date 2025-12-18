// Copyright (c) 2024 The Jocall3 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jocall3

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// User represents a user account in the system.
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest defines the parameters for the CreateUser method.
type CreateUserRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// UpdateUserRequest defines the parameters for the UpdateUser method.
// Fields are pointers to distinguish between a zero value (e.g., false for a bool)
// and a field that should not be updated.
type UpdateUserRequest struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	IsActive  *bool   `json:"is_active,omitempty"`
}

// ListUsersParams defines the query parameters for listing users.
// These parameters are encoded into the URL query string.
type ListUsersParams struct {
	Page    int    `url:"page,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
	SortBy  string `url:"sort_by,omitempty"` // e.g., "created_at_desc"
	Email   string `url:"email,omitempty"`   // Filter by email
}

// ListUsersResponse represents the response from the ListUsers method,
// containing a slice of users and pagination details.
type ListUsersResponse struct {
	Users      []User     `json:"users"`
	Pagination Pagination `json:"pagination"`
}

// CreateUser creates a new user.
// It takes a context and a CreateUserRequest object containing the new user's details.
// It returns the newly created User object or an error on failure.
func (c *Client) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	if req == nil {
		return nil, ErrRequestBodyRequired
	}

	var user User
	err := c.do(ctx, http.MethodPost, "/users", req, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser retrieves a single user by their unique ID.
// It takes a context and the user's ID.
// It returns the User object or an error if the user is not found or another error occurs.
func (c *Client) GetUser(ctx context.Context, userID string) (*User, error) {
	if userID == "" {
		return nil, ErrMissingUserID
	}

	path := fmt.Sprintf("/users/%s", userID)

	var user User
	err := c.do(ctx, http.MethodGet, path, nil, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetCurrentUser retrieves the details of the currently authenticated user.
// This endpoint typically uses the authentication token provided to the client to identify the user.
func (c *Client) GetCurrentUser(ctx context.Context) (*User, error) {
	path := "/users/me"

	var user User
	err := c.do(ctx, http.MethodGet, path, nil, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser modifies an existing user's details.
// It takes a context, the user's ID, and an UpdateUserRequest object with the fields to update.
// Only non-nil fields in the request will be updated.
// It returns the updated User object or an error.
func (c *Client) UpdateUser(ctx context.Context, userID string, req *UpdateUserRequest) (*User, error) {
	if userID == "" {
		return nil, ErrMissingUserID
	}
	if req == nil {
		return nil, ErrRequestBodyRequired
	}

	path := fmt.Sprintf("/users/%s", userID)

	var user User
	// Using PATCH is conventional for partial updates.
	err := c.do(ctx, http.MethodPatch, path, req, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser permanently removes a user by their ID.
// It takes a context and the user's ID.
// It returns an error if the operation fails. A nil error indicates success.
func (c *Client) DeleteUser(ctx context.Context, userID string) error {
	if userID == "" {
		return ErrMissingUserID
	}

	path := fmt.Sprintf("/users/%s", userID)

	// The response body is typically empty on a successful DELETE, so we pass nil for the result.
	return c.do(ctx, http.MethodDelete, path, nil, nil)
}

// ListUsers retrieves a paginated list of users, with optional filtering and sorting.
// It takes a context and a ListUsersParams object. If params is nil, default values are used.
// It returns a ListUsersResponse object containing the users and pagination info, or an error.
func (c *Client) ListUsers(ctx context.Context, params *ListUsersParams) (*ListUsersResponse, error) {
	path := "/users"

	// The `addOptions` helper function (assumed to exist in client.go or a utility file)
	// converts the params struct into a URL query string.
	pathWithOptions, err := addOptions(path, params)
	if err != nil {
		return nil, fmt.Errorf("failed to encode list users params: %w", err)
	}

	var resp ListUsersResponse
	err = c.do(ctx, http.MethodGet, pathWithOptions, nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}