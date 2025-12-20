```go
package execution

import (
	"context"
	"errors"

	"foundry/pkg/types/account"
	"foundry/pkg/types/asset"
	"foundry/pkg/types/instrument"
	"foundry/pkg/types/order"
	"foundry/pkg/types/state"
)

// ErrProjectionNotFound is returned when a requested entity (e.g., account, order)
// does not exist in the projection. This is a specific error that callers can
// check for, distinguishing it from transient system or network errors.
var ErrProjectionNotFound = errors.New("projection: entity not found")

// ProjectionReader defines the interface for reading from fast,
// eventually-consistent data projections.
//
// The execution engine relies on these projections for high-speed, non-blocking
// reads of system state, such as account balances, positions, and open orders.
// Implementations of this interface are expected to be highly available and
// performant, but consumers must be aware that the data may be slightly stale
// compared to the canonical, event-sourced state. This trade-off is fundamental
// to achieving exchange-level execution speed while maintaining the integrity
// of the core event log.
//
// The "Halts safely on uncertainty" principle is supported via the
// GetProjectionWatermark method, which allows consumers to verify data freshness
// before acting. This makes the eventual consistency model explicit and manageable.
//
// All methods must be thread-safe.
type ProjectionReader interface {
	// GetAccountBalance retrieves the projected balance for a specific asset in a given account.
	// It returns ErrProjectionNotFound if the account or asset balance does not exist.
	GetAccountBalance(ctx context.Context, accountID account.AccountID, assetID asset.AssetID) (state.Balance, error)

	// GetPosition retrieves the projected position for a specific instrument in a given account.
	// It returns ErrProjectionNotFound if the account has no position in the specified instrument.
	GetPosition(ctx context.Context, accountID account.AccountID, instrumentID instrument.InstrumentID) (state.Position, error)

	// GetOpenOrdersByAccount retrieves all projected open orders for a given account.
	// This is crucial for risk checks and order management. An empty slice is returned
	// if there are no open orders, not an error.
	GetOpenOrdersByAccount(ctx context.Context, accountID account.AccountID) ([]order.Order, error)

	// GetOrderByID retrieves a specific order by its unique ID from the projection.
	// This is useful for checking the status of a specific order without querying the
	// core transactional system.
	// It returns ErrProjectionNotFound if the order does not exist.
	GetOrderByID(ctx context.Context, orderID order.OrderID) (order.Order, error)

	// GetProjectionWatermark returns the sequence number of the last event
	// that has been processed by the projection. This allows consumers to
	// understand the freshness of the data and potentially wait for the
	// projection to catch up to a specific event before proceeding. This is a
	// key mechanism for managing the trade-offs of eventual consistency.
	GetProjectionWatermark(ctx context.Context) (uint64, error)
}

```