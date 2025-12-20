```go
package ledger

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
)

// ReconcilableEntry represents a single, comparable financial record from any source.
// It serves as a standardized format for reconciliation, abstracting away the
// specific details of different source systems.
type ReconcilableEntry struct {
	// ID is the unique identifier used to match entries across different systems.
	// This could be a transaction ID, a reference number, or a composite key.
	// It is the primary key for matching.
	ID string

	// Amount is the value of the entry in the smallest currency unit (e.g., cents, satoshis).
	// Using integers helps avoid floating-point precision issues, which is a critical
	// invariant for financial systems.
	Amount int64

	// Currency is the ISO 4217 currency code (e.g., "USD", "BTC").
	Currency string

	// Timestamp is the time the entry was recorded or settled. Consistency in timezone
	// (e.g., always UTC) across sources is crucial for accurate comparison.
	Timestamp time.Time

	// Description provides a human-readable summary of the entry.
	Description string

	// Metadata holds additional, source-specific key-value pairs for deeper comparison.
	// This allows for flexible and detailed reconciliation beyond the core fields.
	Metadata map[string]string
}

// DataSource defines the interface for fetching reconcilable entries from a source system.
// This abstraction allows the reconciliation engine to be agnostic about whether it's
// reading from an internal ledger database, an external bank's API, a CSV file, etc.
type DataSource interface {
	// GetName returns a unique, human-readable name for the data source (e.g., "InternalLedgerDB", "BankOfAmerica-API").
	GetName() string
	// GetEntries fetches all reconcilable entries within a specified time window [startTime, endTime).
	GetEntries(ctx context.Context, startTime, endTime time.Time) ([]ReconcilableEntry, error)
}

// DiscrepancyType enumerates the kinds of differences found during reconciliation.
type DiscrepancyType string

const (
	// Mismatch indicates an entry exists in both sources but has differing data.
	Mismatch DiscrepancyType = "MISMATCH"
	// InternalOnly indicates an entry exists only in the internal (source A) system.
	InternalOnly DiscrepancyType = "INTERNAL_ONLY"
	// ExternalOnly indicates an entry exists only in the external (source B) system.
	ExternalOnly DiscrepancyType = "EXTERNAL_ONLY"
)

// Discrepancy captures a single identified difference between two data sources.
// Each discrepancy is an atomic, actionable item for investigation.
type Discrepancy struct {
	Type         DiscrepancyType
	EntryID      string
	InternalData *ReconcilableEntry // Populated for Mismatch and InternalOnly
	ExternalData *ReconcilableEntry // Populated for Mismatch and ExternalOnly
	// Differences provides a list of human-readable descriptions of what differs for Mismatch types.
	// e.g., ["Amount mismatch: internal=100, external=101", "Metadata['key'] mismatch"]
	Differences []string
}

// ReconciliationStatus represents the state of a reconciliation process.
type ReconciliationStatus string

const (
	StatusPending   ReconciliationStatus = "PENDING"
	StatusRunning   ReconciliationStatus = "RUNNING"
	StatusCompleted ReconciliationStatus = "COMPLETED"
	StatusFailed    ReconciliationStatus = "FAILED"
)

// ReconciliationSummary provides high-level statistics about the reconciliation result.
// This is useful for dashboards and quick operational health checks.
type ReconciliationSummary struct {
	InternalTotalCount  int
	ExternalTotalCount  int
	MatchedCount        int
	MismatchedCount     int
	InternalOnlyCount   int
	ExternalOnlyCount   int
	InternalTotalAmount int64
	ExternalTotalAmount int64
	// Currency assumes a single currency for the total amount summary.
	// A multi-currency system would require a map[string]int64 here.
	Currency string
}

// ReconciliationReport is the final, auditable output of a reconciliation process.
// It is an immutable record of the state of two systems at a point in time.
type ReconciliationReport struct {
	ID               uuid.UUID
	StartTime        time.Time
	EndTime          time.Time
	InternalSource   string
	ExternalSource   string
	Status           ReconciliationStatus
	Summary          ReconciliationSummary
	Discrepancies    []Discrepancy
	GeneratedAt      time.Time
	ProcessingErrors []string
}

// Reconciler orchestrates the reconciliation process.
type Reconciler struct {
	// Configuration options can be added here, e.g., comparison tolerance for amounts or timestamps.
}

// NewReconciler creates a new Reconciler instance.
func NewReconciler() *Reconciler {
	return &Reconciler{}
}

// Reconcile performs a full reconciliation between two data sources for a given time period.
// It is designed to be deterministic and fail-closed. If data cannot be reliably fetched
// from either source, the process will halt and return a FAILED report, preventing
// decisions based on incomplete or incorrect data.
func (r *Reconciler) Reconcile(ctx context.Context, internal, external DataSource, startTime, endTime time.Time) *ReconciliationReport {
	report := &ReconciliationReport{
		ID:             uuid.New(),
		StartTime:      startTime,
		EndTime:        endTime,
		InternalSource: internal.GetName(),
		ExternalSource: external.GetName(),
		Status:         StatusRunning,
		GeneratedAt:    time.Now().UTC(),
	}

	var internalEntries, externalEntries []ReconcilableEntry
	var internalErr, externalErr error
	var wg sync.WaitGroup

	wg.Add(2)

	// Fetch data from both sources concurrently to minimize I/O wait time.
	go func() {
		defer wg.Done()
		internalEntries, internalErr = internal.GetEntries(ctx, startTime, endTime)
	}()

	go func() {
		defer wg.Done()
		externalEntries, externalErr = external.GetEntries(ctx, startTime, endTime)
	}()

	wg.Wait()

	// Fail-closed semantics: if any source fails, abort the reconciliation immediately.
	if internalErr != nil {
		report.ProcessingErrors = append(report.ProcessingErrors, fmt.Sprintf("failed to get internal entries: %v", internalErr))
	}
	if externalErr != nil {
		report.ProcessingErrors = append(report.ProcessingErrors, fmt.Sprintf("failed to get external entries: %v", externalErr))
	}
	if len(report.ProcessingErrors) > 0 {
		report.Status = StatusFailed
		report.GeneratedAt = time.Now().UTC()
		return report
	}

	// Core reconciliation logic is performed after all data is successfully fetched.
	discrepancies, summary := r.compareEntrySets(internalEntries, externalEntries)

	report.Discrepancies = discrepancies
	report.Summary = summary
	report.Status = StatusCompleted
	report.GeneratedAt = time.Now().UTC()

	return report
}

// compareEntrySets contains the core logic for comparing two slices of ReconcilableEntry.
// It uses a map-based approach for efficient O(N+M) complexity.
func (r *Reconciler) compareEntrySets(internal, external []ReconcilableEntry) ([]Discrepancy, ReconciliationSummary) {
	summary := ReconciliationSummary{}
	summary.InternalTotalCount = len(internal)
	summary.ExternalTotalCount = len(external)

	discrepancies := make([]Discrepancy, 0)

	externalMap := make(map[string]ReconcilableEntry, len(external))
	for _, entry := range external {
		externalMap[entry.ID] = entry
		summary.ExternalTotalAmount += entry.Amount
		if summary.Currency == "" {
			summary.Currency = entry.Currency
		}
	}

	// 1. Iterate through internal entries and match against the external map.
	for _, internalEntry := range internal {
		summary.InternalTotalAmount += internalEntry.Amount
		if summary.Currency == "" {
			summary.Currency = internalEntry.Currency
		}

		externalEntry, found := externalMap[internalEntry.ID]
		if found {
			// Entry found in both systems, now compare contents.
			diffs := compareEntries(internalEntry, externalEntry)
			if len(diffs) == 0 {
				summary.MatchedCount++
			} else {
				summary.MismatchedCount++
				// Create a copy of the entries to avoid retaining the entire slice.
				iEntry := internalEntry
				eEntry := externalEntry
				discrepancies = append(discrepancies, Discrepancy{
					Type:         Mismatch,
					EntryID:      internalEntry.ID,
					InternalData: &iEntry,
					ExternalData: &eEntry,
					Differences:  diffs,
				})
			}
			// Remove from map to track which external entries have been processed.
			delete(externalMap, internalEntry.ID)
		} else {
			// Entry only exists in the internal system.
			summary.InternalOnlyCount++
			iEntry := internalEntry
			discrepancies = append(discrepancies, Discrepancy{
				Type:         InternalOnly,
				EntryID:      internalEntry.ID,
				InternalData: &iEntry,
			})
		}
	}

	// 2. Any remaining entries in the external map exist only in the external system.
	for _, externalEntry := range externalMap {
		summary.ExternalOnlyCount++
		eEntry := externalEntry
		discrepancies = append(discrepancies, Discrepancy{
			Type:         ExternalOnly,
			EntryID:      externalEntry.ID,
			ExternalData: &eEntry,
		})
	}

	// For deterministic output and easier review, sort discrepancies by EntryID.
	sort.Slice(discrepancies, func(i, j int) bool {
		return discrepancies[i].EntryID < discrepancies[j].EntryID
	})

	return discrepancies, summary
}

// compareEntries checks for differences between two ReconcilableEntry objects.
// It returns a slice of strings describing any differences found.
// An empty slice indicates the entries are identical.
func compareEntries(a, b ReconcilableEntry) []string {
	var diffs []string

	if a.Amount != b.Amount {
		diffs = append(diffs, fmt.Sprintf("Amount mismatch: internal=%d, external=%d", a.Amount, b.Amount))
	}
	if a.Currency != b.Currency {
		diffs = append(diffs, fmt.Sprintf("Currency mismatch: internal=%s, external=%s", a.Currency, b.Currency))
	}
	// Note: Timestamp comparison can be tricky due to timezone and precision.
	// Using a tolerance or standardizing to UTC is recommended. Here we do an exact match for simplicity.
	if !a.Timestamp.Equal(b.Timestamp) {
		diffs = append(diffs, fmt.Sprintf("Timestamp mismatch: internal=%s, external=%s", a.Timestamp.UTC(), b.Timestamp.UTC()))
	}
	if a.Description != b.Description {
		diffs = append(diffs, fmt.Sprintf("Description mismatch: internal='%s', external='%s'", a.Description, b.Description))
	}

	// Deep comparison of metadata maps.
	if !reflect.DeepEqual(a.Metadata, b.Metadata) {
		// A more granular comparison could be implemented here to show which keys/values differ.
		diffs = append(diffs, "Metadata mismatch")
	}

	return diffs
}

```