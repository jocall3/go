```go
package risk

import (
	"context"
	"errors"
	"time"

	"financial-infra/pkg/asset"

	"github.com/shopspring/decimal"
)

// Sentinel errors for the Oracle interface.
// These allow consumers to programmatically handle different failure modes,
// adhering to the fail-closed principle of the system.
var (
	// ErrPriceNotFound indicates that a price for the requested asset pair
	// could not be found from any available source. This is a definitive
	// "no data" response.
	ErrPriceNotFound = errors.New("price not found for asset pair")

	// ErrStalePrice indicates that the most recent price available is older
	// than the acceptable threshold. The system must halt rather than operate
	// on potentially inaccurate, outdated data. This prevents risk calculations
	// based on a non-representative market state.
	ErrStalePrice = errors.New("price is stale")

	// ErrOracleUnavailable indicates a failure in the underlying data source,
	// such as a network error, an API outage, or a timeout. This signals that
	// the oracle's state is uncertain and cannot be trusted.
	ErrOracleUnavailable = errors.New("oracle data source is unavailable")

	// ErrLowConfidence indicates that while a price may be available, it is not
	// considered reliable. This could be due to high volatility, wide bid-ask
	// spreads from underlying sources, or disagreement between multiple feeds.
	// Acting on low-confidence data is a risk the system is designed to avoid.
	ErrLowConfidence = errors.New("price confidence is below threshold")
)

// Price represents a snapshot of an asset's price at a specific time.
// It is an immutable data structure designed for deterministic risk calculations.
type Price struct {
	// Pair is the asset pair for which this price is quoted (e.g., BTC/USD).
	Pair asset.Pair

	// Rate is the exchange rate, representing how much of the Quote asset
	// is needed to purchase one unit of the Base asset.
	// Using a high-precision decimal type is non-negotiable to prevent
	// floating-point errors in financial calculations.
	Rate decimal.Decimal

	// Timestamp is the exact time at which the price was observed.
	// This is critical for auditability and deterministic replay of events.
	Timestamp time.Time

	// Source identifies the origin of the price data (e.g., "ExchangeA", "CompositeFeed").
	// This provides traceability for risk, compliance, and debugging purposes.
	Source string
}

// Oracle defines the interface for a price data provider.
// It abstracts the source of external data, such as asset prices or FX rates,
// allowing the risk engine to operate on a consistent data model regardless of
// the underlying data feed (e.g., live exchange APIs, historical data, mock data).
//
// An Oracle implementation must be fail-closed. If it cannot provide a timely,
// high-confidence price, it MUST return an error. This forces the system to
// halt operations for the affected assets rather than proceeding with
// uncertain data, a core principle of the architecture. This interface makes
// risk assessment deterministic given a set of oracle inputs.
type Oracle interface {
	// GetPrice fetches the most recent, valid price for a given asset pair.
	//
	// Implementations must ensure that the returned price is not stale according
	// to their internal configuration. If a fresh, high-confidence price cannot
	// be obtained, it must return an appropriate error (e.g., ErrStalePrice,
	// ErrOracleUnavailable, ErrLowConfidence). The context can be used to enforce
	// deadlines on the fetch operation, preventing indefinite hangs.
	GetPrice(ctx context.Context, pair asset.Pair) (Price, error)

	// GetPriceAtTime fetches the price for a given asset pair as it was at a
	// specific point in time.
	//
	// This method is essential for:
	//  - Auditing: Verifying historical calculations and positions.
	//  - Replayability: Re-running event sequences for debugging or analysis.
	//  - Back-testing: Evaluating risk models against historical data.
	//
	// If a price for the exact timestamp is not available, an implementation may
	// choose to return a price from the nearest preceding point in time, but it
	// must not return a price from the future relative to the requested time.
	// If no suitable historical price can be found, it should return ErrPriceNotFound.
	GetPriceAtTime(ctx context.Context, pair asset.Pair, t time.Time) (Price, error)
}
### END_OF_FILE_COMPLETED ###
```