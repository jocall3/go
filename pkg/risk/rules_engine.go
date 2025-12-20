```go
// Copyright (c) 2024. All rights reserved.
// This file is part of the Go-based financial infrastructure project.
//
// The content of this file is proprietary and confidential.
// Unauthorized copying, modification, distribution, or use of this file,
// via any medium, is strictly prohibited.

package risk

import (
	"context"
	"fmt"
	"time"
)

// Decision represents the outcome of a policy evaluation.
type Decision int

const (
	// Fail is the zero value for Decision, ensuring that any uninitialized
	// or default state results in a failed check. This embodies the
	// "fail-closed" principle.
	Fail Decision = iota
	// Pass indicates that a policy check was successful.
	Pass
)

// String provides a human-readable representation of a Decision.
func (d Decision) String() string {
	switch d {
	case Pass:
		return "PASS"
	case Fail:
		return "FAIL"
	default:
		return "UNKNOWN"
	}
}

// EvaluationContext holds the data against which risk policies are evaluated.
// It is designed to be an immutable snapshot of the state at the time of evaluation
// to ensure deterministic outcomes.
type EvaluationContext struct {
	// EventID is a unique identifier for the event being evaluated (e.g., transaction ID).
	EventID string

	// PrincipalID identifies the user, service, or entity initiating the action.
	PrincipalID string

	// ActionType categorizes the event (e.g., "TRANSFER", "WITHDRAWAL", "API_KEY_CREATE").
	ActionType string

	// Amount represents the monetary value in the smallest currency unit (e.g., cents, satoshis).
	Amount int64

	// Asset is the currency or asset identifier (e.g., "USD", "BTC").
	Asset string

	// Source and Destination identify the accounts, wallets, or entities involved.
	Source      string
	Destination string

	// Metadata provides additional context about the request environment.
	Timestamp   time.Time
	IPAddress   string
	UserAgent   string
	GeoLocation string

	// CustomData allows for flexible extension with domain-specific information
	// that may be relevant to certain policies.
	CustomData map[string]interface{}
}

// Policy represents a single, atomic risk rule.
// Each policy is a self-contained unit of logic that evaluates the
// EvaluationContext and returns a decision. This interface-based approach
// allows for easy composition and extension of the risk system.
type Policy interface {
	// Name returns a unique, human-readable identifier for the policy.
	// This name is used for logging, auditing, and configuration.
	Name() string

	// Evaluate executes the risk rule against the provided context.
	// It returns a Decision (Pass/Fail), a string explaining the reason for
	// the decision, and an error.
	//
	// An error should only be returned for system-level failures (e.g.,
	// database connection issue, misconfiguration), not for a standard
	// policy violation. A non-nil error will always result in a Fail decision
	// by the RulesEngine, upholding the fail-closed principle.
	Evaluate(ctx context.Context, evalCtx EvaluationContext) (Decision, string, error)
}

// Result encapsulates the outcome of a full evaluation by the RulesEngine.
type Result struct {
	// Decision is the final outcome (Pass or Fail).
	Decision Decision

	// Reason provides a human-readable explanation for the final decision.
	// If failed, it contains the reason from the failing policy.
	Reason string

	// FailedPolicyName indicates which policy caused the failure, if any.
	// It is empty if the decision is Pass.
	FailedPolicyName string

	// EvaluatedPolicies lists the names of all policies that were executed
	// in order, up to the point of decision (failure or completion).
	EvaluatedPolicies []string

	// Error holds any system-level error that occurred during evaluation.
	// Its presence indicates a system malfunction, not a simple policy violation.
	Error error
}

// IsOK returns true if the decision is Pass and no system error occurred.
func (r Result) IsOK() bool {
	return r.Decision == Pass && r.Error == nil
}

// RulesEngine is the core component that orchestrates the evaluation of a
// set of risk policies. It processes policies in a defined order and implements
// fail-fast logic: the first policy to return a Fail decision halts the evaluation.
type RulesEngine struct {
	policies []Policy
}

// NewRulesEngine creates and configures a new RulesEngine.
// It requires at least one policy to be provided. It also validates that all
// policies have unique names to prevent ambiguity in configuration and results.
func NewRulesEngine(policies ...Policy) (*RulesEngine, error) {
	if len(policies) == 0 {
		return nil, fmt.Errorf("rules engine requires at least one policy")
	}

	names := make(map[string]struct{})
	for _, p := range policies {
		name := p.Name()
		if name == "" {
			return nil, fmt.Errorf("policy name cannot be empty")
		}
		if _, exists := names[name]; exists {
			return nil, fmt.Errorf("duplicate policy name detected: %s", name)
		}
		names[name] = struct{}{}
	}

	return &RulesEngine{
		policies: policies,
	}, nil
}

// Evaluate runs all configured policies in order against the provided context.
//
// The evaluation process is sequential and fail-fast:
// 1. If any policy returns an error, evaluation stops immediately, and the engine
//    returns a Fail decision with the error details.
// 2. If any policy returns a Fail decision, evaluation stops immediately, and the
//    engine returns the failure result.
// 3. If all policies execute successfully and return a Pass decision, the engine
//    returns an overall Pass result.
//
// This design ensures that the system is both efficient and safe, halting on any
// sign of risk or uncertainty.
func (e *RulesEngine) Evaluate(ctx context.Context, evalCtx EvaluationContext) Result {
	evaluated := make([]string, 0, len(e.policies))

	for _, policy := range e.policies {
		policyName := policy.Name()
		evaluated = append(evaluated, policyName)

		decision, reason, err := policy.Evaluate(ctx, evalCtx)
		if err != nil {
			// A system-level error occurred. Fail closed immediately.
			return Result{
				Decision:          Fail,
				Reason:            fmt.Sprintf("System error during evaluation of policy '%s'", policyName),
				FailedPolicyName:  policyName,
				EvaluatedPolicies: evaluated,
				Error:             err,
			}
		}

		if decision == Fail {
			// A policy explicitly failed. Stop processing and return the result.
			return Result{
				Decision:          Fail,
				Reason:            reason,
				FailedPolicyName:  policyName,
				EvaluatedPolicies: evaluated,
				Error:             nil,
			}
		}
	}

	// All policies were evaluated and passed.
	return Result{
		Decision:          Pass,
		Reason:            "All risk policies passed.",
		FailedPolicyName:  "",
		EvaluatedPolicies: evaluated,
		Error:             nil,
	}
}
### END_OF_FILE_COMPLETED ###
```