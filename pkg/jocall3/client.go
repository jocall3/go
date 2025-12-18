// Copyright (c) 2023 The JoCall3 Project Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jocall3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"
)

const (
	// defaultBaseURL is the default production API endpoint for JoCall3.
	defaultBaseURL = "https://api.jocall3.com/v1/"
	// defaultTimeout is the default timeout for HTTP requests.
	defaultTimeout = 30 * time.Second
	// libraryVersion is the version of this Go client library.
	libraryVersion = "0.1.0"
)

// Client is the main entry point for interacting with the JoCall3 API.
// It manages authentication, configuration, and provides access to API services.
type Client struct {
	// httpClient is the underlying HTTP client used to make requests.
	// It can be customized using the WithHTTPClient option.
	httpClient *http.Client

	// BaseURL is the base URL for all API requests.
	BaseURL *url.URL

	// UserAgent is the User-Agent header sent with each request.
	UserAgent string

	// apiKey is the API key used for authentication.
	apiKey string

	// common holds a reference to the client for embedding in service structs.
	common service

	// Services for different API resource categories.
	Jobs  *JobService
	Tasks *TaskService
}

// service is a helper struct that holds a reference to the main client.
// It is embedded in API service clients (e.g., JobService) to provide
// them with access to the client's methods.
type service struct {
	client *Client
}

// Option is a functional option for configuring the Client.
type Option func(*Client) error

// NewClient creates a new JoCall3 API client.
// An API key is required for all authenticated endpoints.
//
// Example:
//
//	client, err := jocall3.NewClient("your-api-key")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	jobs, _, err := client.Jobs.List(context.Background(), nil)
func NewClient(apiKey string, opts ...Option) (*Client, error) {
	if strings.TrimSpace(apiKey) == "" {
		return nil, fmt.Errorf("API key cannot be empty")
	}

	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		// This should not happen with a hardcoded URL, but is a safeguard.
		return nil, fmt.Errorf("internal error: failed to parse default base URL: %w", err)
	}

	c := &Client{
		BaseURL:   baseURL,
		apiKey:    apiKey,
		UserAgent: fmt.Sprintf("jocall3-go/%s (%s; %s)", libraryVersion, runtime.GOOS, runtime.GOARCH),
	}

	// Set a default HTTP client if one is not provided via options.
	c.httpClient = &http.Client{
		Timeout: defaultTimeout,
	}

	// Apply functional options to customize the client.
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	// Initialize API services.
	c.common.client = c
	c.Jobs = (*JobService)(&c.common)
	c.Tasks = (*TaskService)(&c.common)

	return c, nil
}

// WithHTTPClient sets a custom HTTP client for the JoCall3 client.
// This is useful for advanced configurations, such as setting custom transports,
// proxies, or timeouts.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) error {
		if httpClient == nil {
			return fmt.Errorf("http client cannot be nil")
		}
		c.httpClient = httpClient
		return nil
	}
}

// WithBaseURL sets a custom base URL for the JoCall3 client.
// This is useful for testing against a mock server or using a different API
// environment (e.g., a staging or self-hosted instance). The URL should
// always end with a trailing slash.
func WithBaseURL(baseURL string) Option {
	return func(c *Client) error {
		u, err := url.Parse(baseURL)
		if err != nil {
			return fmt.Errorf("invalid base URL: %w", err)
		}
		if !strings.HasSuffix(u.Path, "/") {
			u.Path += "/"
		}
		c.BaseURL = u
		return nil
	}
}

// WithUserAgent sets a custom User-Agent string for all requests.
func WithUserAgent(userAgent string) Option {
	return func(c *Client) error {
		if strings.TrimSpace(userAgent) == "" {
			return fmt.Errorf("user agent cannot be empty")
		}
		c.UserAgent = userAgent
		return nil
	}
}

// NewRequest creates an API request. A relative URL path can be provided,
// which will be resolved relative to the BaseURL of the Client.
// If a non-nil body is provided, it will be JSON-encoded and included in the request.
func (c *Client) NewRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse request path: %w", err)
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(body); err != nil {
			return nil, fmt.Errorf("failed to encode request body: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create new HTTP request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// Do sends an API request and returns the API response. The response is JSON-decoded
// and stored in the value pointed to by v, or discarded if v is nil.
// It returns a *Response which wraps the standard *http.Response, and an error.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// If the context was cancelled, we can return that error directly.
		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		default:
		}
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}

	if err := checkResponse(resp); err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
		if err != nil {
			return response, fmt.Errorf("failed to decode response body: %w", err)
		}
	}

	return response, nil
}

// Response is a JoCall3 API response. This wraps the standard http.Response
// and provides access to the raw response.
type Response struct {
	*http.Response
}

// ErrorResponse represents an error returned from the JoCall3 API.
// It includes the HTTP response and a structured error message.
type ErrorResponse struct {
	Response *http.Response `json:"-"`                 // HTTP response that caused this error
	Message  string         `json:"message"`           // Human-readable message
	Code     string         `json:"code,omitempty"`    // API-specific error code
	Details  interface{}    `json:"details,omitempty"` // Additional error details
}

// Error implements the error interface, providing a descriptive error message.
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%s %s: %d %s",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

// checkResponse checks the API response for errors and returns an error if present.
// A response is considered an error if it has a status code outside the 200-299 range.
func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		if jsonErr := json.Unmarshal(data, errorResponse); jsonErr != nil {
			// If we can't unmarshal the error, use the raw body as the message.
			errorResponse.Message = string(data)
		}
	}

	// Re-populate the response body so it can be read again by the caller if needed.
	r.Body = io.NopCloser(bytes.NewBuffer(data))

	return errorResponse
}