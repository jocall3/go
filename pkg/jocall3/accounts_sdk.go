// Copyright (c) 2024, The Jocall3 Project Authors.
// All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jocall3

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Account represents a user account in the system.
// It holds information about the account's identity, status, and metadata.
type Account struct {
	ID        string                 `json:"id"`
	OwnerID   string                 `json:"owner_id"` // The ID of the user who owns this account.
	Name      string                 `json:"name"`
	Email     string                 `json:"email"`
	Status    string                 `json:"status"` // e.g., "active", "suspended", "pending_verification".
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// CreateAccountParams defines the parameters for creating a new account.
type CreateAccountParams struct {
	OwnerID  string                 `json:"owner_id"`           // Required: The ID of the user who will own this account.
	Name     string                 `json:"name"`               // Required: The name of the account.
	Email    string                 `json:"email"`              // Required: The primary contact email for the account.
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Optional: A set of key-value pairs to store additional information.
}

// UpdateAccountParams defines the parameters for updating an existing account.
// Fields are pointers to allow for partial updates, distinguishing between a zero value and a field not being set.
type UpdateAccountParams struct {
	Name     *string                `json:"name,omitempty"`     // Optional: A new name for the account.
	Email    *string                `json:"email,omitempty"`    // Optional: A new primary contact email.
	Status   *string                `json:"status,omitempty"`   // Optional: A new status for the account.
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Optional: New metadata. Note: This will replace the entire existing metadata object.
}

// ListAccountsParams defines the parameters for listing accounts, including pagination and filtering.
type ListAccountsParams struct {
	Limit   int    `url:"limit,omitempty"`   // The maximum number of accounts to return. Defaults to 20, max 100.
	Offset  int    `url:"offset,omitempty"`  // The number of accounts to skip before starting to collect the result set.
	OwnerID string `url:"owner_id,omitempty"`// Filter accounts by the owner's user ID.
	Status  string `url:"status,omitempty"`  // Filter accounts by status.
}

// AccountList represents a paginated list of accounts.
type AccountList struct {
	Data       []Account `json:"data"`
	TotalCount int       `json:"total_count"`
	HasMore    bool      `json:"has_more"`
}

// CreateAccount creates a new account.
// It takes a context and parameters for the new account.
// On success, it returns the newly created Account object.
func (c *Client) CreateAccount(ctx context.Context, params CreateAccountParams) (*Account, error) {
	if params.OwnerID == "" {
		return nil, fmt.Errorf("%w: OwnerID is a required parameter", ErrInvalidParameters)
	}
	if params.Name == "" {
		return nil, fmt.Errorf("%w: Name is a required parameter", ErrInvalidParameters)
	}
	if params.Email == "" {
		return nil, fmt.Errorf("%w: Email is a required parameter", ErrInvalidParameters)
	}

	req, err := c.newRequest(ctx, http.MethodPost, "/v1/accounts", params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var account Account
	if err := c.do(req, &account); err != nil {
		return nil, err // The error from do should be descriptive enough
	}

	return &account, nil
}

// GetAccount retrieves the details of a specific account by its ID.
// It takes a context and the account ID.
// On success, it returns the requested Account object.
func (c *Client) GetAccount(ctx context.Context, accountID string) (*Account, error) {
	if accountID == "" {
		return nil, fmt.Errorf("%w: accountID cannot be empty", ErrInvalidParameters)
	}

	path := fmt.Sprintf("/v1/accounts/%s", accountID)
	req, err := c.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var account Account
	if err := c.do(req, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

// UpdateAccount updates the details of a specific account.
// It takes a context, the account ID, and the parameters to update.
// Only the non-nil fields in params will be updated.
// On success, it returns the updated Account object.
func (c *Client) UpdateAccount(ctx context.Context, accountID string, params UpdateAccountParams) (*Account, error) {
	if accountID == "" {
		return nil, fmt.Errorf("%w: accountID cannot be empty", ErrInvalidParameters)
	}

	path := fmt.Sprintf("/v1/accounts/%s", accountID)
	req, err := c.newRequest(ctx, http.MethodPatch, path, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var account Account
	if err := c.do(req, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

// DeleteAccount permanently deletes an account.
// It takes a context and the ID of the account to delete.
// This operation is irreversible.
func (c *Client) DeleteAccount(ctx context.Context, accountID string) error {
	if accountID == "" {
		return fmt.Errorf("%w: accountID cannot be empty", ErrInvalidParameters)
	}

	path := fmt.Sprintf("/v1/accounts/%s", accountID)
	req, err := c.newRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// For DELETE, we often don't expect a body, so we pass nil for the result.
	// The `do` method should handle this, typically by checking for a 204 No Content status.
	return c.do(req, nil)
}

// ListAccounts retrieves a paginated list of accounts, with optional filtering.
// It takes a context and parameters for pagination and filtering.
// On success, it returns an AccountList object.
func (c *Client) ListAccounts(ctx context.Context, params ListAccountsParams) (*AccountList, error) {
	// The `newRequest` method in the client should handle encoding params into a query string.
	req, err := c.newRequest(ctx, http.MethodGet, "/v1/accounts", params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var accountList AccountList
	if err := c.do(req, &accountList); err != nil {
		return nil, err
	}

	return &accountList, nil
}