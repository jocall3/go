```go
// Copyright (c) 2024. The financial_bridge project authors. All rights reserved.
// Use of this source code is governed by a license that can be found in the LICENSE file.

package observability

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.comcom/prometheus/client_golang/prometheus/promauto"
)

const (
	namespace = "financial_bridge"
)

// Metrics are registered with Prometheus automatically on package initialization
// thanks to the use of the promauto package. This ensures that metrics are
// available as soon as the application starts, and it panics on any registration
// errors (e.g., duplicate metrics), which is desirable behavior at startup.

// --- Execution Metrics ---

var (
	// TransactionsProcessed is a counter for transactions processed by the system.
	// It provides insights into the system's throughput and success/failure rates.
	// Labels:
	// - type: The type of transaction (e.g., "deposit", "withdrawal", "internal_transfer").
	// - status: The final status of the transaction (e.g., "success", "failed_risk", "failed_execution").
	TransactionsProcessed = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "execution",
			Name:      "transactions_processed_total",
			Help:      "Total number of transactions processed, labeled by type and status.",
		},
		[]string{"type", "status"},
	)

	// OperationLatency measures the duration of key system operations.
	// This is critical for identifying performance bottlenecks and ensuring the system
	// meets its execution speed goals.
	// Labels:
	// - operation: The name of the operation being measured (e.g., "execute_transaction", "commit_ledger_update").
	OperationLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: "system",
			Name:      "operation_latency_seconds",
			Help:      "Latency of key system operations in seconds.",
			// Buckets are tailored for financial operations, ranging from sub-millisecond to a few seconds.
			Buckets: []float64{0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5},
		},
		[]string{"operation"},
	)
)

// --- Risk Metrics ---

var (
	// RiskChecks is a counter for risk checks performed by the risk engine.
	// This metric is essential for monitoring the effectiveness and activity of risk controls.
	// Labels:
	// - check_name: The specific risk check being performed (e.g., "aml_screen", "velocity_limit", "balance_check").
	// - result: The outcome of the check (e.g., "pass", "fail", "error").
	RiskChecks = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "risk",
			Name:      "checks_total",
			Help:      "Total number of risk checks performed, labeled by check name and result.",
		},
		[]string{"check_name", "result"},
	)

	// SystemHalts is a counter for system halts triggered by the risk engine or other critical components.
	// This is a key indicator of system stability and its fail-safe mechanisms.
	// Labels:
	// - reason: The reason for the system halt (e.g., "market_volatility", "data_feed_failure", "capital_inadequacy").
	SystemHalts = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "risk",
			Name:      "system_halts_total",
			Help:      "Total number of times the system has been halted due to risk or uncertainty.",
		},
		[]string{"reason"},
	)
)

// --- Capital & Settlement Metrics ---

var (
	// SystemCapital tracks the total value of assets held by the system.
	// This is a fundamental metric for ensuring capital safety and solvency.
	// Labels:
	// - asset: The currency or asset being tracked (e.g., "USD", "BTC", "ETH").
	SystemCapital = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "capital",
			Name:      "system_assets_total",
			Help:      "Total amount of assets managed by the system, labeled by asset.",
		},
		[]string{"asset"},
	)

	// SettlementEvents tracks settlement-related events, providing visibility into the finality of transactions.
	// Labels:
	// - type: The type of settlement event (e.g., "batch_commit", "network_confirmation").
	// - status: The status of the event (e.g., "success", "pending", "failed").
	SettlementEvents = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "settlement",
			Name:      "events_total",
			Help:      "Total number of settlement events, labeled by type and status.",
		},
		[]string{"type", "status"},
	)
)

// --- System Health Metrics ---

var (
	// ComponentStatus indicates the operational status of a system component.
	// This allows operators to quickly assess the health of the entire system.
	// Values: 0=OK, 1=Degraded, 2=Halted/Failed
	// Labels:
	// - component: The name of the system component (e.g., "execution_engine", "risk_engine", "ledger").
	ComponentStatus = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "system",
			Name:      "component_status",
			Help:      "Current operational status of a system component (0=OK, 1=Degraded, 2=Halted).",
		},
		[]string{"component"},
	)

	// EventQueueDepth tracks the number of events in various processing queues.
	// This is useful for detecting back-pressure and potential processing delays.
	// Labels:
	// - queue_name: The name of the queue (e.g., "inbound_transactions", "settlement_batch").
	EventQueueDepth = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "system",
			Name:      "event_queue_depth",
			Help:      "Current number of events in a processing queue.",
		},
		[]string{"queue_name"},
	)
)
### END_OF_FILE_COMPLETED ###
```