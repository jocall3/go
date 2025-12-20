```go
// Copyright (c) 2024. The Bridge Project. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package governance

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// PolicyLoader defines the interface for loading a governance Ruleset.
// This abstraction allows for different storage backends like local files,
// a database, or a remote configuration service, promoting testability and
// flexibility in deployment.
type PolicyLoader interface {
	// Load retrieves and parses the governance Ruleset from its persistent source.
	// It returns a validated Ruleset or an error if loading, parsing,
	// or validation fails. A non-nil error from Load should be treated as a
	// fatal condition on system startup, enforcing fail-closed semantics.
	Load() (*Ruleset, error)
}

// FilePolicyLoader implements the PolicyLoader interface for loading rules
// from a local JSON file. This is suitable for deployments where governance
// rules are managed as version-controlled configuration files, providing a
// clear audit trail for policy changes.
type FilePolicyLoader struct {
	filePath string
}

// NewFilePolicyLoader creates a new loader for a specific file path.
// It returns an error if the provided path is empty, ensuring that the loader
// is always configured with a valid target.
func NewFilePolicyLoader(path string) (*FilePolicyLoader, error) {
	if path == "" {
		return nil, fmt.Errorf("policy file path cannot be empty")
	}
	return &FilePolicyLoader{filePath: path}, nil
}

// Load reads the policy file from the configured path, unmarshals it into a
// Ruleset, and performs validation.
// This implementation adheres to fail-closed semantics: if the file cannot be
// read, parsed, or validated, the system will not start with an incomplete
// or invalid policy set. This prevents the system from operating in an
// undefined or unsafe state.
func (l *FilePolicyLoader) Load() (*Ruleset, error) {
	// Ensure the file path is clean and absolute for clarity in logs and errors.
	absPath, err := filepath.Abs(l.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for policy file '%s': %w", l.filePath, err)
	}

	// Read the entire file into memory.
	// For very large policy files, a streaming parser might be considered,
	// but governance rulesets are typically small enough for this approach,
	// and it simplifies the loading logic.
	data, err := os.ReadFile(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("governance policy file not found at '%s': %w", absPath, err)
		}
		return nil, fmt.Errorf("failed to read governance policy file '%s': %w", absPath, err)
	}

	// Unmarshal the JSON data into the Ruleset struct.
	var ruleset Ruleset
	if err := json.Unmarshal(data, &ruleset); err != nil {
		return nil, fmt.Errorf("failed to parse governance policy JSON from '%s': %w", absPath, err)
	}

	// After successful parsing, perform validation on the loaded ruleset.
	// This is a critical step to ensure that the rules are internally consistent
	// and adhere to system invariants before they are put into effect.
	if err := ruleset.Validate(); err != nil {
		return nil, fmt.Errorf("governance policy validation failed for '%s': %w", absPath, err)
	}

	// The ruleset is valid and ready to be used by the governance engine.
	return &ruleset, nil
}

// MockPolicyLoader is a test implementation of PolicyLoader.
// It allows injecting a specific Ruleset or an error for testing purposes,
// decoupling tests from the filesystem.
type MockPolicyLoader struct {
	RulesetToLoad *Ruleset
	ErrorToReturn error
}

// Load returns the pre-configured Ruleset or error.
// This is useful for unit testing components that depend on a PolicyLoader
// without needing to interact with the filesystem, enabling fast and reliable
// tests.
func (m *MockPolicyLoader) Load() (*Ruleset, error) {
	if m.ErrorToReturn != nil {
		return nil, m.ErrorToReturn
	}
	if m.RulesetToLoad != nil {
		// In a real test, we might want to return a deep copy to avoid mutation
		// across different parts of a test. For simplicity here, we return the pointer.
		return m.RulesetToLoad, nil
	}
	// Default behavior if neither is set: return an empty, valid ruleset.
	// This is a safe default for tests that don't care about the specific
	// policies but require a valid, non-nil Ruleset.
	rs := NewRuleset()
	return rs, nil
}

```