```go
// Copyright (c) 2024. The Bridge Project. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package risk

import (
	"fmt"
)

// Decision represents the outcome of a policy evaluation. It is the fundamental
// unit of judgment produced by the risk engine.
type Decision int

const (
	// Allow indicates the action is permitted by the policy.
	Allow Decision = iota
	// Deny indicates the action is blocked by the policy. This is a terminal state
	// for a risk check.
	Deny
	// Flag indicates the action is permitted but should be flagged for review or
	// further monitoring.
	Flag
)

// String provides a human-readable representation of a Decision.
func (d Decision) String() string {
	switch d {
	case Allow:
		return "Allow"
	case Deny:
		return "Deny"
	case Flag:
		return "Flag"
	default:
		return "Unknown"
	}
}

// Result holds the outcome of a single policy evaluation. It provides a clear,
// auditable record of why a decision was made.
type Result struct {
	PolicyName string
	Decision   Decision
	Reason     string
	Error      error
}

// Context provides the data against which a policy is evaluated.
// It's a flexible map allowing various data points like transaction amount,
// user history, account balances, etc., to be passed to the risk engine
// without tight coupling.
type Context map[string]interface{}

// Policy defines the interface for a single, checkable risk rule.
// Each policy evaluates a given context and returns a result, forming the
// basic building block of the risk management system.
type Policy interface {
	// Name returns the unique, human-readable name of the policy.
	Name() string
	// Evaluate checks the provided context against the policy's rule.
	Evaluate(ctx Context) Result
}

// Operator defines the type of comparison to be performed in a policy.
type Operator int

const (
	LessThan Operator = iota
	LessThanOrEqual
	GreaterThan
	GreaterThanOrEqual
	EqualTo
	NotEqualTo
)

// String provides a human-readable representation of an Operator.
func (o Operator) String() string {
	switch o {
	case LessThan:
		return "<"
	case LessThanOrEqual:
		return "<="
	case GreaterThan:
		return ">"
	case GreaterThanOrEqual:
		return ">="
	case EqualTo:
		return "=="
	case NotEqualTo:
		return "!="
	default:
		return "Unsupported"
	}
}

// ThresholdPolicy is a concrete implementation of a Policy that compares a numeric
// value from the context against a fixed threshold.
// It is designed to handle a wide range of common rules like "daily withdrawal
// limit < $10,000" or "transaction amount > $1,000,000".
// All numeric comparisons are done using int64 to represent the smallest
// currency unit (e.g., cents, satoshis) to avoid floating-point inaccuracies.
type ThresholdPolicy struct {
	PolicyName        string
	PolicyDescription string
	ContextKey        string   // The key to look up in the Context map.
	Op                Operator // The comparison operator to use.
	Threshold         int64    // The value to compare against.
	FailureDecision   Decision // The decision to return if the check fails (e.g., Deny or Flag).
}

// NewThresholdPolicy creates and validates a new ThresholdPolicy.
func NewThresholdPolicy(name, description, contextKey string, op Operator, threshold int64, failureDecision Decision) *ThresholdPolicy {
	return &ThresholdPolicy{
		PolicyName:        name,
		PolicyDescription: description,
		ContextKey:        contextKey,
		Op:                op,
		Threshold:         threshold,
		FailureDecision:   failureDecision,
	}
}

// Name returns the name of the policy.
func (p *ThresholdPolicy) Name() string {
	return p.PolicyName
}

// Evaluate checks if the value associated with ContextKey in the context
// satisfies the comparison with the Threshold. It adheres to fail-closed semantics:
// any missing data or type mismatch results in a denial.
func (p *ThresholdPolicy) Evaluate(ctx Context) Result {
	// 1. Retrieve the value from the context.
	rawValue, ok := ctx[p.ContextKey]
	if !ok {
		err := fmt.Errorf("context key '%s' not found", p.ContextKey)
		return Result{
			PolicyName: p.Name(),
			Decision:   Deny, // Fail-closed: missing data is a denial.
			Reason:     "Required risk data is missing from the context.",
			Error:      err,
		}
	}

	// 2. Assert the value to int64.
	// This enforces that the system consistently uses int64 for comparable monetary values.
	value, ok := rawValue.(int64)
	if !ok {
		err := fmt.Errorf("value for context key '%s' is not of type int64, got %T", p.ContextKey, rawValue)
		return Result{
			PolicyName: p.Name(),
			Decision:   Deny, // Fail-closed: incorrect data type is a denial.
			Reason:     "Risk data has an unexpected type.",
			Error:      err,
		}
	}

	// 3. Perform the comparison.
	passed := compare(value, p.Threshold, p.Op)

	// 4. Construct the result based on the comparison outcome.
	if passed {
		return Result{
			PolicyName: p.Name(),
			Decision:   Allow,
			Reason:     fmt.Sprintf("Check passed: value %d %s threshold %d.", value, p.Op.String(), p.Threshold),
		}
	}

	// The check failed.
	return Result{
		PolicyName: p.Name(),
		Decision:   p.FailureDecision,
		Reason:     fmt.Sprintf("Check failed: value %d is not %s threshold %d. %s", value, p.Op.String(), p.Threshold, p.PolicyDescription),
	}
}

// compare performs a numeric comparison between two int64 values based on the operator.
func compare(val1, val2 int64, op Operator) bool {
	switch op {
	case LessThan:
		return val1 < val2
	case LessThanOrEqual:
		return val1 <= val2
	case GreaterThan:
		return val1 > val2
	case GreaterThanOrEqual:
		return val1 >= val2
	case EqualTo:
		return val1 == val2
	case NotEqualTo:
		return val1 != val2
	default:
		// This should not be reached if policies are configured correctly.
		// We default to false (fail-closed) for any unknown or unsupported operators.
		return false
	}
}

```