```go
// Copyright (c) 2024. The Bridge Project. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package server provides the HTTP server implementation, including routing,
// middleware, and request/response handling for the financial infrastructure API.
package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// envelope is a generic wrapper for JSON responses, providing a consistent
// top-level structure for both successful and error responses.
type envelope map[string]interface{}

// writeJSON is an internal helper for sending JSON responses. It sets the
// appropriate headers, marshals the data, and handles potential encoding errors,
// ensuring a fail-closed behavior.
func writeJSON(w http.ResponseWriter, r *http.Request, status int, data envelope, headers http.Header) {
	// Marshal the data to JSON. Using MarshalIndent for development readability.
	// In a high-performance production environment, json.Marshal would be used.
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		// This is a server-side problem. Log the detailed error and send a
		// generic 500 response to the client to avoid leaking implementation details.
		log.Printf("ERROR: failed to marshal JSON response for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	// Append any additional headers provided in the call.
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Set the standard JSON content type header and write the status code.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	// Write the JSON body to the response.
	if _, err := w.Write(js); err != nil {
		// This could happen if the client closes the connection prematurely.
		// Log it for visibility, as it's not a server error but is good to know.
		log.Printf("NOTICE: failed to write JSON response for %s %s: %v", r.Method, r.URL.Path, err)
	}
}

// Respond sends a standard success JSON response.
// It wraps the provided data payload within a "data" key.
func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	response := envelope{"data": data}
	writeJSON(w, r, status, response, nil)
}

// RespondError sends a standard error JSON response.
// It wraps the error details within an "error" key.
func RespondError(w http.ResponseWriter, r *http.Request, status int, code, message string) {
	response := envelope{
		"error": envelope{
			"code":    code,
			"message": message,
		},
	}
	writeJSON(w, r, status, response, nil)
}

// --- Standard Error Response Helpers ---

// ServerError responds with a 500 Internal Server Error.
// It logs the underlying error but sends a generic message to the client,
// adhering to the principle of not leaking internal implementation details.
func ServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("SERVER ERROR: %s %s: %v", r.Method, r.URL.Path, err)
	message := "The server encountered a problem and could not process your request."
	RespondError(w, r, http.StatusInternalServerError, "internal_server_error", message)
}

// NotFound responds with a 404 Not Found error.
// This is used when the requested resource does not exist.
func NotFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found."
	RespondError(w, r, http.StatusNotFound, "not_found", message)
}

// MethodNotAllowed responds with a 405 Method Not Allowed error.
// This is used when the client uses an HTTP method that is not supported by the resource.
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource.", r.Method)
	RespondError(w, r, http.StatusMethodNotAllowed, "method_not_allowed", message)
}

// BadRequest responds with a 400 Bad Request error.
// The details parameter can be used to provide specific information about the error,
// such as a malformed request body.
func BadRequest(w http.ResponseWriter, r *http.Request, details interface{}) {
	response := envelope{
		"error": envelope{
			"code":    "bad_request",
			"message": "The request is invalid or malformed.",
			"details": details,
		},
	}
	writeJSON(w, r, http.StatusBadRequest, response, nil)
}

// FailedValidation responds with a 422 Unprocessable Entity error.
// This is the standard response for validation failures, providing clear,
// actionable feedback to the client. The details map should contain the
// fields and their corresponding error messages.
func FailedValidation(w http.ResponseWriter, r *http.Request, details map[string]string) {
	response := envelope{
		"error": envelope{
			"code":    "validation_failed",
			"message": "The request failed validation checks.",
			"details": details,
		},
	}
	writeJSON(w, r, http.StatusUnprocessableEntity, response, nil)
}

// Unauthorized responds with a 401 Unauthorized error.
// This is used when authentication is required but has failed or has not been provided.
// It includes the WWW-Authenticate header, which is standard for 401 responses.
func Unauthorized(w http.ResponseWriter, r *http.Request) {
	message := "Authentication is required and has failed or has not yet been provided."
	header := http.Header{}
	header.Set("WWW-Authenticate", `Bearer realm="restricted"`)

	response := envelope{
		"error": envelope{
			"code":    "unauthorized",
			"message": message,
		},
	}
	writeJSON(w, r, http.StatusUnauthorized, response, header)
}

// Forbidden responds with a 403 Forbidden error.
// This is used when the client is authenticated but lacks the necessary
// permissions to access or modify the resource.
func Forbidden(w http.ResponseWriter, r *http.Request) {
	message := "You do not have permission to perform this action."
	RespondError(w, r, http.StatusForbidden, "forbidden", message)
}

// RateLimitExceeded responds with a 429 Too Many Requests error.
// This indicates that the user has sent too many requests in a given amount of time.
func RateLimitExceeded(w http.ResponseWriter, r *http.Request) {
	message := "You have exceeded the API rate limit."
	RespondError(w, r, http.StatusTooManyRequests, "rate_limit_exceeded", message)
}

```