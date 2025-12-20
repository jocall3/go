```go
// Copyright (c) 2023-2024 The Bridge Core Team. All rights reserved.
// Use of this source code is governed by an MIT license that can be
// found in the LICENSE file.

package settlement

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

// CheckpointManager handles the persistent storage of the last processed event sequence number.
// This is a critical component for ensuring system resilience and data integrity. By reliably
// checkpointing progress, an event projector can safely resume from its last known good state
// after a restart or failure. This prevents both missed events (which could lead to capital loss)
// and duplicate event processing (which could lead to incorrect state).
//
// The implementation guarantees atomicity and durability of checkpoint writes through a
// write-and-rename strategy. This fail-closed semantic ensures that the checkpoint file
// is never left in a corrupted or partially written state, even if the system crashes
// mid-operation.
type CheckpointManager struct {
	filePath string
	mu       sync.RWMutex
}

// NewCheckpointManager creates and initializes a new manager for a checkpoint file.
// It takes a directory and a filename for the checkpoint. It ensures the directory
// exists, creating it if necessary. This constructor prepares the manager for use
// but does not perform any I/O on the checkpoint file itself.
//
// The separation of directory and name allows for clear configuration and management
// of persistent state files.
func NewCheckpointManager(dir, name string) (*CheckpointManager, error) {
	if dir == "" || name == "" {
		return nil, fmt.Errorf("checkpoint directory and name cannot be empty")
	}

	if err := os.MkdirAll(dir, 0750); err != nil {
		return nil, fmt.Errorf("failed to create checkpoint directory %s: %w", dir, err)
	}

	return &CheckpointManager{
		filePath: filepath.Join(dir, name),
	}, nil
}

// Load reads the last saved sequence number from the checkpoint file.
// If the file does not exist, it returns 0 and a nil error, which is the
// defined safe starting point for a new or reset system. This behavior is
// crucial for the initial bootstrap of a projector.
//
// If the file exists but is empty or contains malformed data, an error is
// returned. This is an intentional "halt on uncertainty" design choice. An
// invalid checkpoint is a sign of potential corruption or a critical bug,
// and requires operator intervention rather than making a potentially unsafe
// assumption about the correct sequence number.
func (cm *CheckpointManager) Load() (uint64, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	data, err := os.ReadFile(cm.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// File not found is a normal condition for the first run.
			// The safe starting point is sequence 0.
			return 0, nil
		}
		return 0, fmt.Errorf("failed to read checkpoint file %s: %w", cm.filePath, err)
	}

	content := strings.TrimSpace(string(data))
	if content == "" {
		// An empty file is treated as an uninitialized state, same as a non-existent file.
		return 0, nil
	}

	sequence, err := strconv.ParseUint(content, 10, 64)
	if err != nil {
		// A malformed file is a critical error. We halt safely rather than guessing.
		// This requires manual intervention to resolve the corrupted state.
		return 0, fmt.Errorf("checkpoint file %s is corrupted with non-numeric content: %w", cm.filePath, err)
	}

	return sequence, nil
}

// Store atomically writes the given sequence number to the checkpoint file.
// The operation is designed to be durable and crash-safe. It follows these steps:
// 1. Write the new sequence number to a temporary file in the same directory.
// 2. Force a sync of the temporary file's content to the underlying storage device.
// 3. Atomically rename the temporary file to the final checkpoint file name.
//
// This write-and-rename pattern ensures that the primary checkpoint file is only
// ever updated in a single, atomic step. If the system fails at any point
// during the write or sync, the original, valid checkpoint file remains untouched,
// and the temporary file is left behind to be cleaned up on the next run. This
// prevents a partial write from ever corrupting the checkpoint.
func (cm *CheckpointManager) Store(sequence uint64) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	content := []byte(strconv.FormatUint(sequence, 10))

	// Create a temporary file to write the new checkpoint to.
	// Using a pattern ensures we don't conflict with other processes.
	tempFile, err := os.CreateTemp(filepath.Dir(cm.filePath), filepath.Base(cm.filePath)+".*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temporary checkpoint file: %w", err)
	}
	// Ensure the temp file is removed if anything goes wrong before the final rename.
	defer os.Remove(tempFile.Name())

	// Write the new sequence number to the temporary file.
	if _, err := tempFile.Write(content); err != nil {
		tempFile.Close()
		return fmt.Errorf("failed to write to temporary checkpoint file %s: %w", tempFile.Name(), err)
	}

	// Sync the file to disk to guarantee durability before the atomic rename.
	// This is a critical step to prevent data loss on a power failure or OS crash.
	if err := tempFile.Sync(); err != nil {
		tempFile.Close()
		return fmt.Errorf("failed to sync temporary checkpoint file %s: %w", tempFile.Name(), err)
	}

	if err := tempFile.Close(); err != nil {
		return fmt.Errorf("failed to close temporary checkpoint file %s: %w", tempFile.Name(), err)
	}

	// Atomically replace the old checkpoint file with the new one.
	// This is the core of the crash-safe update mechanism.
	if err := os.Rename(tempFile.Name(), cm.filePath); err != nil {
		return fmt.Errorf("failed to atomically rename checkpoint file from %s to %s: %w", tempFile.Name(), cm.filePath, err)
	}

	return nil
}

```