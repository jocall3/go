```go
// Copyright (c) 2024. The Bridge Project Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package settlement

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

// LedgerView defines the interface for accessing the source-of-truth ledger state.
// This represents the canonical state of all accounts, derived directly from the
// immutable log of transactions.
type LedgerView interface {
	// GetAllAccountBalances retrieves the balances of all accounts directly from the ledger.
	// This is expected to be the definitive source of truth.
	GetAllAccountBalances(ctx context.Context) (map[string]decimal.Decimal, error)
}

// ProjectionView defines the interface for accessing a read-side projection of the ledger.
// Projections are optimized for fast queries but might lag or, in case of bugs,
// become inconsistent with the source-of-truth ledger.
type ProjectionView interface {
	// GetAllProjectedAccountBalances retrieves the balances of all accounts from a read-side projection.
	GetAllProjectedAccountBalances(ctx context.Context) (map[string]decimal.Decimal, error)
}

// Alerter defines an interface for sending critical alerts to monitoring systems.
type Alerter interface {
	// Alert sends a high-priority notification about a critical system event.
	Alert(ctx context.Context, level string, message string, details map[string]interface{})
}

// SystemStatusController defines an interface to control the system's operational state.
// This is a critical component for implementing fail-closed semantics.
type SystemStatusController interface {
	// Halt gracefully stops the system from processing new transactions due to a critical error,
	// preserving the last known consistent state.
	Halt(ctx context.Context, reason string)
}

// Reconciler is a background process that periodically compares the ledger's state
// with read-side projections to detect and report inconsistencies. Its primary function
// is to act as a safety net, ensuring the system's view of the financial state
// remains consistent. If an inconsistency is found, it triggers a system halt
// to prevent any potentially incorrect operations.
type Reconciler struct {
	ledger               LedgerView
	projection           ProjectionView
	alerter              Alerter
	systemStatus         SystemStatusController
	logger               *slog.Logger
	reconciliationInterval time.Duration

	ticker *time.Ticker
	done   chan struct{}
	wg     sync.WaitGroup
}

// NewReconciler creates a new Reconciler instance.
// It requires implementations for ledger and projection views, an alerter, a system status controller,
// a logger, and the interval at which to run the reconciliation checks.
func NewReconciler(
	ledger LedgerView,
	projection ProjectionView,
	alerter Alerter,
	systemStatus SystemStatusController,
	logger *slog.Logger,
	reconciliationInterval time.Duration,
) (*Reconciler, error) {
	if ledger == nil {
		return nil, fmt.Errorf("ledger view cannot be nil")
	}
	if projection == nil {
		return nil, fmt.Errorf("projection view cannot be nil")
	}
	if alerter == nil {
		return nil, fmt.Errorf("alerter cannot be nil")
	}
	if systemStatus == nil {
		return nil, fmt.Errorf("system status controller cannot be nil")
	}
	if logger == nil {
		return nil, fmt.Errorf("logger cannot be nil")
	}
	if reconciliationInterval <= 0 {
		return nil, fmt.Errorf("reconciliation interval must be positive")
	}

	return &Reconciler{
		ledger:               ledger,
		projection:           projection,
		alerter:              alerter,
		systemStatus:         systemStatus,
		logger:               logger.With(slog.String("component", "reconciler")),
		reconciliationInterval: reconciliationInterval,
		done:                 make(chan struct{}),
	}, nil
}

// Start begins the periodic reconciliation process in a background goroutine.
func (r *Reconciler) Start(ctx context.Context) {
	r.logger.Info("Starting settlement reconciler", "interval", r.reconciliationInterval.String())
	r.ticker = time.NewTicker(r.reconciliationInterval)
	r.wg.Add(1)
	go r.run(ctx)
}

// Stop gracefully shuts down the reconciler.
func (r *Reconciler) Stop() {
	r.logger.Info("Stopping settlement reconciler")
	close(r.done)
	r.ticker.Stop()
	r.wg.Wait()
	r.logger.Info("Settlement reconciler stopped")
}

// run is the main loop for the reconciler. It listens for ticker events or a shutdown signal.
func (r *Reconciler) run(ctx context.Context) {
	defer r.wg.Done()

	// Perform an initial reconciliation on startup to ensure the system starts in a consistent state.
	r.logger.Info("Performing initial reconciliation check")
	if err := r.reconcile(ctx); err != nil {
		r.handleInconsistency(ctx, err)
		// If the initial check fails, the system is halted. The reconciler's job is done.
		return
	}

	for {
		select {
		case <-r.ticker.C:
			if err := r.reconcile(ctx); err != nil {
				r.handleInconsistency(ctx, err)
				// After a critical inconsistency, the system will be halted.
				// The reconciler's job is done, so we exit the loop.
				return
			}
		case <-r.done:
			return
		case <-ctx.Done():
			r.logger.Info("Context cancelled, shutting down reconciler")
			return
		}
	}
}

// reconcile performs a single comparison between the ledger and the projection.
// It returns an error if any inconsistency is found.
func (r *Reconciler) reconcile(ctx context.Context) error {
	r.logger.Debug("Running reconciliation check")

	ledgerBalances, err := r.ledger.GetAllAccountBalances(ctx)
	if err != nil {
		// This is a failure to fetch, not an inconsistency. We log and continue.
		r.logger.Error("Failed to get ledger balances for reconciliation", "error", err)
		return nil // Or return a specific error type if we want to retry differently.
	}

	projectionBalances, err := r.projection.GetAllProjectedAccountBalances(ctx)
	if err != nil {
		// This is a failure to fetch, not an inconsistency. We log and continue.
		r.logger.Error("Failed to get projection balances for reconciliation", "error", err)
		return nil
	}

	r.logger.Debug("Comparing balances", "ledger_accounts", len(ledgerBalances), "projection_accounts", len(projectionBalances))

	// Compare the two maps for any discrepancies.
	return r.compareBalances(ledgerBalances, projectionBalances)
}

// compareBalances checks for discrepancies between two maps of account balances.
// It is a pure function that returns an error describing the first inconsistency found.
func (r *Reconciler) compareBalances(ledger, projection map[string]decimal.Decimal) error {
	// Check for accounts in ledger but not in projection, or with mismatched balances.
	for accountID, ledgerBalance := range ledger {
		projectionBalance, ok := projection[accountID]
		if !ok {
			return fmt.Errorf("account %q exists in ledger but not in projection", accountID)
		}
		if !ledgerBalance.Equal(projectionBalance) {
			return fmt.Errorf("balance mismatch for account %q: ledger=%s, projection=%s", accountID, ledgerBalance.String(), projectionBalance.String())
		}
	}

	// Check for accounts in projection but not in ledger.
	// This check is efficient because we already know all ledger accounts exist in the projection.
	// If the lengths differ, there must be extra accounts in the projection.
	if len(ledger) != len(projection) {
		for accountID := range projection {
			if _, ok := ledger[accountID]; !ok {
				return fmt.Errorf("account %q exists in projection but not in ledger", accountID)
			}
		}
	}

	r.logger.Info("Reconciliation check passed successfully")
	return nil
}

// handleInconsistency logs the error, sends an alert, and triggers a system halt.
// This is the "fail-closed" response to a detected inconsistency.
func (r *Reconciler) handleInconsistency(ctx context.Context, err error) {
	r.logger.Error("CRITICAL: Reconciliation failed, data inconsistency detected", "error", err)

	alertDetails := map[string]interface{}{
		"error":        err.Error(),
		"component":    "settlement_reconciler",
		"action_taken": "system_halt_initiated",
	}
	r.alerter.Alert(ctx, "CRITICAL", "Ledger and projection data inconsistency detected", alertDetails)

	haltReason := fmt.Sprintf("Reconciliation failure: %v. System halted to prevent further divergence and ensure data integrity.", err)
	r.systemStatus.Halt(ctx, haltReason)
}

```