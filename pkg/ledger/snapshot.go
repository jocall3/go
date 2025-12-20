```go
package ledger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// Snapshot represents the state of the ledger at a specific event sequence.
// It contains all the necessary information to restore the ledger to that point in time,
// allowing for faster startup by avoiding a full replay of the event log.
type Snapshot struct {
	// LastEventSequence is the sequence number of the last event included in this snapshot.
	// When restoring, event replay should start from LastEventSequence + 1.
	LastEventSequence uint64 `json:"last_event_sequence"`

	// Timestamp is the UTC time when the snapshot was created.
	Timestamp time.Time `json:"timestamp"`

	// Accounts holds a deep copy of the state of all accounts at the time of the snapshot.
	// The key is the account ID.
	Accounts map[string]Account `json:"accounts"`

	// Version indicates the version of the application or data schema. This is crucial
	// for handling migrations and ensuring compatibility when restoring from a snapshot.
	Version string `json:"version"`
}

// SnapshotStore defines the interface for persisting and retrieving ledger snapshots.
// This abstraction allows for different storage backends (e.g., local disk, cloud storage)
// without changing the core snapshotting logic.
type SnapshotStore interface {
	// Save persists a snapshot. Implementations should ensure this operation is atomic
	// to prevent corrupted snapshot files.
	Save(snapshot *Snapshot) error

	// LoadLatest retrieves the most recent valid snapshot.
	// It should return (nil, nil) if no snapshots are found, which is a valid state
	// for a new ledger.
	LoadLatest() (*Snapshot, error)
}

// FileSnapshotStore is an implementation of SnapshotStore that uses the local filesystem.
// Snapshots are stored as versioned, timestamped JSON files in a specified directory.
type FileSnapshotStore struct {
	mu          sync.RWMutex
	dir         string
	filePattern string
}

// NewFileSnapshotStore creates a new FileSnapshotStore.
// It ensures the provided directory exists, creating it if necessary.
func NewFileSnapshotStore(dir string) (*FileSnapshotStore, error) {
	if dir == "" {
		return nil, fmt.Errorf("snapshot directory cannot be empty")
	}
	if err := os.MkdirAll(dir, 0750); err != nil {
		return nil, fmt.Errorf("failed to create snapshot directory %s: %w", dir, err)
	}
	return &FileSnapshotStore{
		dir:         dir,
		// Filename format is designed for easy, correct lexicographical sorting.
		// The 20-digit padding for the sequence number handles up to 10^20 events.
		filePattern: "snapshot-%020d.json",
	}, nil
}

// Save atomically saves a snapshot to a file.
// It first writes to a temporary file and then renames it to the final destination.
// This ensures that a partial write due to a crash doesn't corrupt the snapshot store.
func (s *FileSnapshotStore) Save(snapshot *Snapshot) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.MarshalIndent(snapshot, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal snapshot: %w", err)
	}

	filename := fmt.Sprintf(s.filePattern, snapshot.LastEventSequence)
	filepath := filepath.Join(s.dir, filename)
	tempFilepath := filepath + ".tmp"

	// Write to a temporary file first to ensure atomicity.
	if err := ioutil.WriteFile(tempFilepath, data, 0640); err != nil {
		return fmt.Errorf("failed to write temporary snapshot file: %w", err)
	}

	// Atomically rename the temporary file to the final filename.
	if err := os.Rename(tempFilepath, filepath); err != nil {
		return fmt.Errorf("failed to rename snapshot file: %w", err)
	}

	return nil
}

// LoadLatest finds and loads the most recent snapshot from the directory.
// It identifies the latest snapshot by finding the file with the highest sequence number
// in its name. The zero-padded filename format makes lexicographical sorting reliable.
func (s *FileSnapshotStore) LoadLatest() (*Snapshot, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	files, err := ioutil.ReadDir(s.dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read snapshot directory: %w", err)
	}

	var snapshotFiles []string
	prefix := "snapshot-"
	suffix := ".json"
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), prefix) && strings.HasSuffix(file.Name(), suffix) {
			snapshotFiles = append(snapshotFiles, file.Name())
		}
	}

	if len(snapshotFiles) == 0 {
		return nil, nil // No snapshots found, not an error.
	}

	// Sort files in descending order to find the latest one.
	sort.Sort(sort.Reverse(sort.StringSlice(snapshotFiles)))

	latestFilename := snapshotFiles[0]
	filepath := filepath.Join(s.dir, latestFilename)

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read latest snapshot file %s: %w", latestFilename, err)
	}

	var snapshot Snapshot
	if err := json.Unmarshal(data, &snapshot); err != nil {
		return nil, fmt.Errorf("failed to unmarshal snapshot from file %s: %w", latestFilename, err)
	}

	return &snapshot, nil
}

// Snapshotter is responsible for orchestrating the creation of ledger snapshots.
// It determines when a snapshot should be taken based on a configured frequency.
type Snapshotter struct {
	store     SnapshotStore
	frequency uint64 // Take a snapshot every 'frequency' events.
}

// NewSnapshotter creates a new Snapshotter.
// The frequency must be greater than zero.
func NewSnapshotter(store SnapshotStore, frequency uint64) (*Snapshotter, error) {
	if frequency == 0 {
		return nil, fmt.Errorf("snapshot frequency must be greater than zero")
	}
	if store == nil {
		return nil, fmt.Errorf("snapshot store cannot be nil")
	}
	return &Snapshotter{
		store:     store,
		frequency: frequency,
	}, nil
}

// ShouldTakeSnapshot determines if a snapshot should be taken for a given event sequence number.
// This is typically called by the Ledger after successfully processing and persisting an event.
func (s *Snapshotter) ShouldTakeSnapshot(eventSequence uint64) bool {
	return eventSequence > 0 && eventSequence%s.frequency == 0
}

// TakeSnapshot creates a snapshot of the current ledger state and saves it using the SnapshotStore.
// This method requires a consistent, read-locked view of the ledger's state.
func (s *Snapshotter) TakeSnapshot(ledgerState *Ledger) error {
	if ledgerState == nil {
		return fmt.Errorf("cannot take snapshot of a nil ledger")
	}

	// It is the responsibility of the caller (the Ledger) to ensure that its state
	// is locked for reading during this operation to guarantee a consistent snapshot.
	// We create a deep copy of the accounts to release the lock as quickly as possible.
	accountsCopy := make(map[string]Account, len(ledgerState.accounts))
	for id, acc := range ledgerState.accounts {
		// Account is a struct, so this assignment creates a copy.
		// If Account contained pointers or slices, a more careful deep copy would be needed.
		// Our design mandates Account to be a simple, copyable struct.
		accountsCopy[id] = acc
	}

	snapshot := &Snapshot{
		LastEventSequence: ledgerState.lastEventSequence,
		Timestamp:         time.Now().UTC(),
		Accounts:          accountsCopy,
		Version:           "1.0.0", // This should be tied to the application version for production systems.
	}

	if err := s.store.Save(snapshot); err != nil {
		// A failure to snapshot is a critical operational risk. The system should
		// halt or enter a safe mode rather than continue processing without snapshots.
		return fmt.Errorf("CRITICAL: failed to save snapshot at sequence %d: %w", snapshot.LastEventSequence, err)
	}

	return nil
}
### END_OF_FILE_COMPLETED ###
```