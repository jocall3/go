// Copyright (c) 2024 The JoCall Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jocall3

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Constants for the AI service.
const (
	DefaultAPIBaseURL = "https://api.jocall.ai/v3"
	DefaultModel      = "jocall-3-pro"
	APIVersion        = "v3.0.0"
)

// Role constants for chat messages.
const (
	RoleSystem    = "system"
	RoleUser      = "user"
	RoleAssistant = "assistant"
	RoleTool      = "tool"
)

// FinishReason constants for chat completion choices.
const (
	FinishReasonStop          = "stop"
	FinishReasonLength        = "length"
	FinishReasonToolCalls     = "tool_calls"
	FinishReasonContentFilter = "content_filter"
	FinishReasonFunctionCall  = "function_call" // For legacy compatibility
)

// --- Client and Configuration ---

// Client is the main entry point for interacting with the JoCall AI API.
// It manages API authentication, endpoint configuration, and HTTP requests.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// ClientOption is a functional option for configuring the Client.
type ClientOption func(*Client)

// NewClient creates a new JoCall AI client.
// It requires an API key, which can be provided directly or will be
// read from the JOCALL_API_KEY environment variable if the provided key is empty.
func NewClient(apiKey string, opts ...ClientOption) (*Client, error) {
	if apiKey == "" {
		apiKey = os.Getenv("JOCALL_API_KEY")
	}
	if apiKey == "" {
		return nil, errors.New("API key is required; please provide it directly or set the JOCALL_API_KEY environment variable")
	}

	c := &Client{
		apiKey:  apiKey,
		baseURL: DefaultAPIBaseURL,
		httpClient: &http.Client{
			Timeout: 90 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

// WithBaseURL sets a custom base URL for the API, allowing for use with
// different environments or a self-hosted instance.
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = strings.TrimRight(url, "/")
	}
}

// WithHTTPClient sets a custom http.Client, enabling users to configure
// custom transports, timeouts, or middleware.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		if client != nil {
			c.httpClient = client
		}
	}
}

// --- API Error Handling ---

// APIError represents a structured error returned by the JoCall API.
type APIError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	Type       string `json:"type"`
	Code       any    `json:"code,omitempty"` // Can be string or int
}

// Error implements the error interface for APIError.
func (e *APIError) Error() string {
	return fmt.Sprintf("jocall api error (status %d): %s (type: %s, code: %v)", e.StatusCode, e.Message, e.Type, e.Code)
}

// --- Chat Completion ---

// ChatCompletionRequest represents a request to the chat completion endpoint.
type ChatCompletionRequest struct {
	Model            string        `json:"model"`
	Messages         []ChatMessage `json:"messages"`
	Temperature      float32       `json:"temperature,omitempty"`
	TopP             float32       `json:"top_p,omitempty"`
	N                int           `json:"n,omitempty"`
	Stream           bool          `json:"stream,omitempty"`
	Stop             []string      `json:"stop,omitempty"`
	MaxTokens        int           `json:"max_tokens,omitempty"`
	PresencePenalty  float32       `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32       `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	User             string        `json:"user,omitempty"`
	Tools            []Tool        `json:"tools,omitempty"`
	ToolChoice       any           `json:"tool_choice,omitempty"` // Can be "none", "auto", or a specific tool object
}

// ChatMessage represents a single message in a conversation.
type ChatMessage struct {
	Role       string     `json:"role"`
	Content    string     `json:"content"`
	Name       string     `json:"name,omitempty"`       // For tool/function calls
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`  // Assistant message with tool calls
	ToolCallID string     `json:"tool_call_id,omitempty"` // Tool message with result
}

// Tool represents a tool the model can call, currently supporting functions.
type Tool struct {
	Type     string   `json:"type"` // e.g., "function"
	Function Function `json:"function"`
}

// Function represents the definition of a function tool.
type Function struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Parameters  any    `json:"parameters"` // Should be a JSON Schema object
}

// ToolCall represents a call to a tool made by the model.
type ToolCall struct {
	ID       string       `json:"id"`
	Type     string       `json:"type"` // e.g., "function"
	Function FunctionCall `json:"function"`
}

// FunctionCall represents the name and arguments for a function call.
type FunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"` // A JSON string of arguments
}

// ChatCompletionResponse represents a response from the chat completion endpoint.
type ChatCompletionResponse struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	SystemFingerprint string   `json:"system_fingerprint"`
}

// Choice represents a single completion choice.
type Choice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	LogProbs     any         `json:"logprobs,omitempty"`
	FinishReason string      `json:"finish_reason"`
}

// Usage represents token usage statistics for a request.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatCompletion creates a model response for the given chat conversation.
// For streaming responses, use ChatCompletionStream instead.
func (c *Client) ChatCompletion(ctx context.Context, req ChatCompletionRequest) (*ChatCompletionResponse, error) {
	if req.Stream {
		return nil, errors.New("use ChatCompletionStream for streaming requests")
	}
	if req.Model == "" {
		req.Model = DefaultModel
	}

	var respData ChatCompletionResponse
	err := c.do(ctx, http.MethodPost, "/chat/completions", req, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

// --- Streaming Chat Completion ---

// ChatCompletionStreamResponse is a chunk of a streaming chat completion response.
type ChatCompletionStreamResponse struct {
	ID                string         `json:"id"`
	Object            string         `json:"object"`
	Created           int64          `json:"created"`
	Model             string         `json:"model"`
	Choices           []StreamChoice `json:"choices"`
	SystemFingerprint string         `json:"system_fingerprint"`
}

// StreamChoice represents a single choice in a streaming response.
type StreamChoice struct {
	Index        int         `json:"index"`
	Delta        ChatMessage `json:"delta"`
	LogProbs     any         `json:"logprobs,omitempty"`
	FinishReason string      `json:"finish_reason"`
}

// ChatCompletionStream handles a streaming response from the chat completion endpoint.
type ChatCompletionStream struct {
	resp   *http.Response
	reader *bufio.Reader
}

// Recv receives the next chunk from the stream.
// It returns io.EOF when the stream is finished successfully.
func (s *ChatCompletionStream) Recv() (*ChatCompletionStreamResponse, error) {
	for {
		line, err := s.reader.ReadBytes('\n')
		if err != nil {
			// This includes io.EOF at the end of the stream
			return nil, err
		}

		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue // Skip empty lines between events
		}

		if !bytes.HasPrefix(line, []byte("data: ")) {
			continue // Skip non-data lines (e.g., comments)
		}
		data := bytes.TrimPrefix(line, []byte("data: "))

		if string(data) == "[DONE]" {
			return nil, io.EOF
		}

		var chunk ChatCompletionStreamResponse
		if err := json.Unmarshal(data, &chunk); err != nil {
			return nil, fmt.Errorf("failed to unmarshal stream chunk: %w", err)
		}
		return &chunk, nil
	}
}

// Close closes the underlying stream. It is the caller's responsibility to close the stream.
func (s *ChatCompletionStream) Close() error {
	return s.resp.Body.Close()
}

// ChatCompletionStream creates a streaming model response for the given chat conversation.
// The caller is responsible for closing the returned stream via the Close() method.
func (c *Client) ChatCompletionStream(ctx context.Context, req ChatCompletionRequest) (*ChatCompletionStream, error) {
	req.Stream = true
	if req.Model == "" {
		req.Model = DefaultModel
	}

	httpReq, err := c.newRequest(ctx, http.MethodPost, "/chat/completions", req)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "text/event-stream")
	httpReq.Header.Set("Cache-Control", "no-cache")
	httpReq.Header.Set("Connection", "keep-alive")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		// Ensure body is closed to prevent resource leaks
		defer resp.Body.Close()
		return nil, c.handleAPIError(resp)
	}

	stream := &ChatCompletionStream{
		resp:   resp,
		reader: bufio.NewReader(resp.Body),
	}

	return stream, nil
}

// --- Internal Request Handling ---

// newRequest creates a new HTTP request with appropriate headers and body.
func (c *Client) newRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	fullURL := c.baseURL + path

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "jocall-go/"+APIVersion)

	return req, nil
}

// do performs a non-streaming HTTP request and decodes the response body into v.
func (c *Client) do(ctx context.Context, method, path string, body, v interface{}) error {
	req, err := c.newRequest(ctx, method, path, body)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return c.handleAPIError(resp)
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return fmt.Errorf("failed to decode response body: %w", err)
		}
	}

	return nil
}

// handleAPIError reads an error response body and returns a structured APIError.
func (c *Client) handleAPIError(resp *http.Response) error {
	apiErr := &APIError{StatusCode: resp.StatusCode}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		apiErr.Message = fmt.Sprintf("failed to read error response body: %v", readErr)
		return apiErr
	}

	// Attempt to unmarshal the standard error structure.
	if err := json.Unmarshal(body, &apiErr); err != nil {
		// If unmarshaling fails, the body is likely not JSON or has a different structure.
		// Use the raw body as the error message for diagnostics.
		apiErr.Message = string(body)
	}

	return apiErr
}