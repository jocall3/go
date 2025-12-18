// Package payment provides services related to payment processing,
// including foreign exchange, transaction handling, and account management.
package payment

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/google/uuid"

	"github.com/your-org/your-repo/internal/domain/payment"
	"github.com/your-org/your-repo/internal/repository"
	"github.com/your-org/your-repo/pkg/errors"
	"github.com/your-org/your-repo/pkg/log"
)

const (
	// DefaultBaseCurrency is the currency used for triangulation if a direct rate is not found.
	// Using a major currency like USD is a common practice.
	DefaultBaseCurrency = payment.USD

	// rateCacheTTL is the time-to-live for FX rates in the cache.
	// FX rates don't change dramatically second-by-second for most applications,
	// so caching for a few minutes to an hour is reasonable.
	rateCacheTTL = 10 * time.Minute
)

// FXService defines the interface for foreign exchange operations.
// It abstracts the complexities of rate retrieval, including caching,
// triangulation, and fetching from external providers.
type FXService interface {
	// GetRate retrieves the exchange rate between two currencies.
	GetRate(ctx context.Context, from, to payment.Currency) (*payment.FXRate, error)

	// Convert transforms a monetary amount from one currency to another.
	Convert(ctx context.Context, amount payment.Money, toCurrency payment.Currency) (*payment.Money, error)
}

// fxService is the concrete implementation of the FXService.
type fxService struct {
	fxRepo repository.FXRateRepository
	logger log.Logger
	// A cache client would be injected here for production use.
	// For example: cache cache.Client
}

// NewFXService creates a new instance of the FX service.
// It requires a repository for data access and a logger for observability.
func NewFXService(fxRepo repository.FXRateRepository, logger log.Logger) FXService {
	return &fxService{
		fxRepo: fxRepo,
		logger: logger.With(log.String("service", "fx")),
	}
}

// GetRate retrieves the exchange rate between two currencies.
// The retrieval strategy is as follows:
// 1. Handle the trivial case where 'from' and 'to' currencies are the same.
// 2. Attempt to fetch the direct rate (e.g., USD -> EUR).
// 3. If not found, attempt to fetch the inverse rate (e.g., EUR -> USD) and calculate its reciprocal.
// 4. If still not found, attempt to triangulate the rate via a standard base currency (e.g., USD).
// 5. If all attempts fail, return a 'not found' error.
// A caching layer would typically wrap these steps to improve performance.
func (s *fxService) GetRate(ctx context.Context, from, to payment.Currency) (*payment.FXRate, error) {
	const op = "payment.fxService.GetRate"

	if err := from.Validate(); err != nil {
		return nil, errors.New(op, errors.KindBadRequest, fmt.Errorf("invalid 'from' currency: %w", err))
	}
	if err := to.Validate(); err != nil {
		return nil, errors.New(op, errors.KindBadRequest, fmt.Errorf("invalid 'to' currency: %w", err))
	}

	// 1. Trivial case: same currency.
	if from == to {
		return payment.NewFXRate(uuid.New(), from, to, big.NewRat(1, 1), time.Now().UTC()), nil
	}

	// TODO: Implement a caching layer here.
	// e.g., check cache for "fx:rate:USD:EUR"
	// if found, return cached rate

	// 2. Try to find the direct rate (e.g., USD -> EUR).
	rate, err := s.fxRepo.FindLatestRate(ctx, from, to)
	if err == nil && rate != nil {
		// TODO: Cache the result here.
		return rate, nil
	}
	if err != nil && !errors.Is(err, errors.KindNotFound) {
		return nil, errors.Wrap(op, err, "failed to find direct rate")
	}

	// 3. Try to find the inverse rate (e.g., EUR -> USD) and calculate its inverse.
	inverseRate, err := s.fxRepo.FindLatestRate(ctx, to, from)
	if err == nil && inverseRate != nil {
		if inverseRate.Rate().Sign() == 0 {
			return nil, errors.New(op, errors.KindInternal, fmt.Errorf("inverse rate for %s to %s is zero, cannot invert", to, from))
		}
		// Invert the rate: 1 / inverseRate
		inverted := new(big.Rat).Inv(inverseRate.Rate())
		calculatedRate := payment.NewFXRate(uuid.New(), from, to, inverted, inverseRate.Timestamp())

		// TODO: Cache and potentially store the calculated rate here for future direct lookups.
		return calculatedRate, nil
	}
	if err != nil && !errors.Is(err, errors.KindNotFound) {
		return nil, errors.Wrap(op, err, "failed to find inverse rate")
	}

	// 4. Try triangulation via a base currency (e.g., USD).
	if from != DefaultBaseCurrency && to != DefaultBaseCurrency {
		s.logger.Info(ctx, "Direct and inverse rates not found, attempting triangulation",
			log.String("from", string(from)),
			log.String("to", string(to)),
			log.String("base", string(DefaultBaseCurrency)),
		)
		rate, err := s.getTriangulatedRate(ctx, from, to, DefaultBaseCurrency)
		if err == nil && rate != nil {
			// TODO: Cache and potentially store the calculated rate here.
			return rate, nil
		}
		if err != nil {
			s.logger.Warn(ctx, "Triangulation failed", log.Error(err))
		}
	}

	// 5. If all attempts fail, return a not found error.
	return nil, errors.New(op, errors.KindNotFound, fmt.Errorf("exchange rate not found for pair %s/%s", from, to))
}

// getTriangulatedRate calculates a rate by converting through a base currency.
// For example, to get JPY -> EUR, it can use (JPY -> USD) and (EUR -> USD).
// The formula is: (from -> to) = (from -> base) / (to -> base).
func (s *fxService) getTriangulatedRate(ctx context.Context, from, to, base payment.Currency) (*payment.FXRate, error) {
	const op = "payment.fxService.getTriangulatedRate"

	// Get rate from the 'from' currency to the base currency.
	fromToBaseRate, err := s.GetRate(ctx, from, base)
	if err != nil {
		return nil, errors.Wrap(op, err, fmt.Sprintf("could not get rate from %s to base %s for triangulation", from, base))
	}

	// Get rate from the 'to' currency to the base currency.
	toToBaseRate, err := s.GetRate(ctx, to, base)
	if err != nil {
		return nil, errors.Wrap(op, err, fmt.Sprintf("could not get rate from %s to base %s for triangulation", to, base))
	}

	// Ensure rates are not zero to avoid division by zero.
	if fromToBaseRate.Rate().Sign() == 0 || toToBaseRate.Rate().Sign() == 0 {
		return nil, errors.New(op, errors.KindInternal, "one of the triangulation rates is zero")
	}

	// Calculate the final rate: (from -> base) / (to -> base) = from -> to
	// Example: (JPY -> USD) / (EUR -> USD) = (JPY/USD) / (EUR/USD) = JPY/EUR
	finalRateValue := new(big.Rat).Quo(fromToBaseRate.Rate(), toToBaseRate.Rate())

	// Use the most recent timestamp of the two component rates.
	timestamp := fromToBaseRate.Timestamp()
	if toToBaseRate.Timestamp().After(timestamp) {
		timestamp = toToBaseRate.Timestamp()
	}

	finalRate := payment.NewFXRate(uuid.New(), from, to, finalRateValue, timestamp)
	return finalRate, nil
}

// Convert transforms a monetary amount from one currency to another using the latest exchange rate.
func (s *fxService) Convert(ctx context.Context, amount payment.Money, toCurrency payment.Currency) (*payment.Money, error) {
	const op = "payment.fxService.Convert"

	fromCurrency := amount.Currency()

	if err := toCurrency.Validate(); err != nil {
		return nil, errors.New(op, errors.KindBadRequest, fmt.Errorf("invalid 'to' currency: %w", err))
	}

	// If currencies are the same, no conversion is needed.
	if fromCurrency == toCurrency {
		return &amount, nil
	}

	s.logger.Info(ctx, "Performing currency conversion",
		log.String("from_amount", amount.String()),
		log.String("to_currency", string(toCurrency)),
	)

	// Get the exchange rate.
	rate, err := s.GetRate(ctx, fromCurrency, toCurrency)
	if err != nil {
		return nil, errors.Wrap(op, err, "failed to get exchange rate for conversion")
	}

	// Perform the conversion. The domain.Money object handles the precision
	// and rounding logic, keeping the service layer clean.
	convertedAmount, err := amount.ConvertTo(toCurrency, rate.Rate())
	if err != nil {
		return nil, errors.New(op, errors.KindInternal, fmt.Errorf("failed to perform conversion calculation: %w", err))
	}

	s.logger.Info(ctx, "Currency conversion successful",
		log.String("from_amount", amount.String()),
		log.String("to_amount", convertedAmount.String()),
		log.String("rate", rate.Rate().FloatString(8)),
	)

	return convertedAmount, nil
}