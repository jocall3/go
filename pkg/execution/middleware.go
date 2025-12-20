```go
// Copyright (c) 2024 Your Org. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package execution

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/your-org/your-project/pkg/command"
	"github.com/your-org/your-project/pkg/observability/metrics"
)

// Middleware is a function that wraps a Handler to provide additional,
// cross-cutting functionality like logging, metrics, or validation.
// This pattern allows for a clean separation of concerns, where the core
// command handler focuses solely on business logic.
type Middleware func(Handler) Handler

// Chain applies a series of middlewares to a Handler. The middlewares are
// applied in the order they are passed, meaning the first middleware in the
// slice becomes the outermost layer, and the last becomes the innermost.
// For example, Chain(h, m1, m2, m3) results in m1(m2(m3(h))).
func Chain(h Handler, middlewares ...Middleware) Handler {
	// Apply middlewares in reverse order to achieve the desired nesting.
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// LoggingMiddleware creates a middleware that provides structured logging for
// command execution. It logs the beginning and end of command processing,
// including execution duration and success or failure status.
// This is crucial for auditability and operational insight, making the system's
// behavior "boring to regulators" and "powerful to operators".
func LoggingMiddleware(logger logrus.FieldLogger) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, cmd command.Command) (err error) {
			startTime := time.Now()

			// Create a logger with context-specific fields for this command.
			cmdLogger := logger.WithFields(logrus.Fields{
				"command_type": string(cmd.Type()),
				"command_id":   cmd.ID(),
			})

			cmdLogger.Info("Executing command")

			// Defer the final log message to ensure it's always logged,
			// regardless of panics or errors.
			defer func() {
				duration := time.Since(startTime)
				logFields := logrus.Fields{
					"duration_ms": duration.Milliseconds(),
				}
				if err != nil {
					cmdLogger.WithFields(logFields).WithError(err).Error("Command execution failed")
				} else {
					cmdLogger.WithFields(logFields).Info("Command execution succeeded")
				}
			}()

			return next(ctx, cmd)
		}
	}
}

// MetricsMiddleware creates a middleware that records key performance indicators
// about command execution. It tracks the total number of commands processed
// and their execution latency.
// This data is essential for monitoring system health, identifying performance
// bottlenecks, and ensuring the system feels "fast to builders".
func MetricsMiddleware(recorder metrics.CommandRecorder) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, cmd command.Command) (err error) {
			startTime := time.Now()
			cmdType := string(cmd.Type())

			defer func() {
				duration := time.Since(startTime)
				status := metrics.StatusSuccess
				if err != nil {
					status = metrics.StatusFailure
				}
				recorder.IncCommandTotal(cmdType, status)
				recorder.ObserveCommandDuration(cmdType, duration)
			}()

			return next(ctx, cmd)
		}
	}
}

// ValidationMiddleware creates a middleware that performs validation on commands
// before they are executed. It checks if a command implements the `command.Validatable`
// interface and, if so, calls its `Validate` method.
// This is a critical safety mechanism that enforces invariants at the system's
// entry point, making it "slow to attackers" by rejecting malformed or
// unauthorized requests early.
func ValidationMiddleware() Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, cmd command.Command) error {
			if validatable, ok := cmd.(command.Validatable); ok {
				if err := validatable.Validate(ctx); err != nil {
					// Return the validation error immediately, halting the execution chain.
					// This prevents invalid state from ever reaching the core logic.
					return err
				}
			}
			return next(ctx, cmd)
		}
	}
}

```