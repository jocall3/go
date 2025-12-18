// Package jocall3 provides the client SDK for interacting with the JoCall v3 API.
// It offers a simple, idiomatic Go interface for creating, managing, and retrieving
// results from remote job executions.
package jocall3

import (
	"encoding/json"
	"fmt"
	"time"
)

// JobStatus represents the lifecycle state of a job.
type JobStatus string

// Defines the possible statuses for a job.
const (
	StatusPending   JobStatus = "pending"
	StatusRunning   JobStatus = "running"
	StatusCompleted JobStatus = "completed"
	StatusFailed    JobStatus = "failed"
	StatusCancelled JobStatus = "cancelled"
)

// IsTerminal returns true if the job status is a final state (Completed, Failed, or Cancelled).
// A job in a terminal state will not change its status again.
func (s JobStatus) IsTerminal() bool {
	switch s {
	case StatusCompleted, StatusFailed, StatusCancelled:
		return true
	default:
		return false
	}
}

// Error represents a detailed error returned by the JoCall API.
// It provides more context than a standard Go error, including a machine-readable
// code and optional structured details.
type Error struct {
	// Code is a machine-readable string identifying the error type.
	// e.g., "INVALID_ARGUMENT", "NOT_FOUND", "INTERNAL_ERROR".
	Code string `json:"code"`

	// Message is a human-readable description of the error.
	Message string `json:"message"`

	// Details provides additional structured information about the error,
	// which can be useful for programmatic error handling.
	Details map[string]any `json:"details,omitempty"`
}

// Error implements the standard Go error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("jocall error: code=%s, message=%s", e.Code, e.Message)
}

// JobResult holds the outcome of a completed or failed job.
// Exactly one of Output or Error will be non-nil, depending on the job's success.
type JobResult struct {
	// Output contains the successful result of the job execution as raw JSON.
	// It is nil if the job failed. You can unmarshal this into a specific Go type.
	Output json.RawMessage `json:"output,omitempty"`

	// Error contains detailed information about why the job failed.
	// It is nil if the job completed successfully.
	Error *Error `json:"error,omitempty"`
}

// Job represents a single unit of work executed by the JoCall service.
// It contains all metadata, input, and results associated with a specific execution.
type Job struct {
	// ID is the unique identifier for the job.
	ID string `json:"id"`

	// Function is the name of the registered function that was executed.
	Function string `json:"function"`

	// Status indicates the current state of the job in its lifecycle.
	Status JobStatus `json:"status"`

	// Input is the raw JSON payload provided when the job was created.
	Input json.RawMessage `json:"input"`

	// Result holds the final outcome of the job. It is nil until the job
	// reaches a terminal state (Completed or Failed).
	Result *JobResult `json:"result,omitempty"`

	// CreatedAt is the timestamp when the job was accepted by the API.
	CreatedAt time.Time `json:"createdAt"`

	// UpdatedAt is the timestamp of the last status change for the job.
	UpdatedAt time.Time `json:"updatedAt"`

	// StartedAt is the timestamp when the job execution began.
	// It is nil if the job has not yet started.
	StartedAt *time.Time `json:"startedAt,omitempty"`

	// CompletedAt is the timestamp when the job finished execution,
	// either successfully or with an error. It is nil if the job is not finished.
	CompletedAt *time.Time `json:"completedAt,omitempty"`
}

// JobRequest is used to create a new job. It specifies the function to run
// and the input data to provide.
type JobRequest struct {
	// Function is the name of the function to be executed. This must match
	// a function registered with the JoCall service.
	Function string `json:"function"`

	// Input is the data to be passed to the function. It can be any value
	// that is serializable to JSON.
	Input any `json:"input"`

	// CallbackURL is an optional URL where the final Job object will be sent
	// via a POST request upon completion. If not provided, the client must
	// poll for the result.
	CallbackURL string `json:"callbackUrl,omitempty"`

	// CorrelationID is an optional client-provided identifier that is passed
	// through and included in the final Job object. Useful for tracking.
	CorrelationID string `json:"correlationId,omitempty"`
}

// FunctionDefinition describes a function that is available for execution
// via the JoCall API. This is useful for service discovery and client generation.
type FunctionDefinition struct {
	// Name is the unique identifier for the function.
	Name string `json:"name"`

	// Description provides a human-readable explanation of what the function does.
	Description string `json:"description,omitempty"`

	// InputSchema is a JSON Schema object describing the expected structure and
	// constraints of the input payload.
	InputSchema json.RawMessage `json:"inputSchema,omitempty"`

	// OutputSchema is a JSON Schema object describing the structure of the
	// successful output payload.
	OutputSchema json.RawMessage `json:"outputSchema,omitempty"`
}