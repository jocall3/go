package ai

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"math/rand"
	"sync"
	"time"
)

// -----------------------------------------------------------------------------
// Domain Models & DTOs
// -----------------------------------------------------------------------------

// SimulationType defines the category of financial simulation.
type SimulationType string

const (
	SimTypeMonteCarlo StandardSimulationType = "MONTE_CARLO"
	SimTypeHistorical StandardSimulationType = "HISTORICAL"
	SimTypeAIOracle   StandardSimulationType = "AI_ORACLE"
)

type StandardSimulationType string

// SimulationConfig holds parameters for running a financial simulation.
type SimulationConfig struct {
	AssetID       string
	InitialPrice  float64
	Volatility    float64 // Annualized volatility (sigma)
	Drift         float64 // Annualized return rate (mu)
	TimeHorizon   float64 // In years (e.g., 1.0 for 1 year)
	TimeSteps     int     // Number of steps in the simulation
	NumPaths      int     // Number of simulation paths (e.g., 10000)
	RiskFreeRate  float64
	Seed          int64 // Optional seed for reproducibility
	IncludeOracle bool  // If true, applies AI adjustments
}

// SimulationResult contains the aggregated output of the simulation.
type SimulationResult struct {
	SimulationID     string
	Timestamp        time.Time
	Config           SimulationConfig
	ExpectedPrice    float64
	MedianPrice      float64
	StandardDev      float64
	VaR95            float64 // Value at Risk (95% confidence)
	VaR99            float64 // Value at Risk (99% confidence)
	BestCase         float64 // 95th percentile
	WorstCase        float64 // 5th percentile
	SamplePaths      [][]float64
	ExecutionTime    time.Duration
	OracleInsights   []string // Explanations of AI adjustments applied
}

// OracleInsight represents AI-driven modifications to simulation parameters.
type OracleInsight struct {
	DriftAdjustment      float64
	VolatilityMultiplier float64
	Reasoning            string
}

// -----------------------------------------------------------------------------
// Service Interface
// -----------------------------------------------------------------------------

// OracleService defines the contract for the AI-enhanced financial oracle.
type OracleService interface {
	// RunStandardSimulation executes traditional financial models (e.g., GBM Monte Carlo).
	RunStandardSimulation(ctx context.Context, config SimulationConfig) (*SimulationResult, error)

	// RunAdvancedOracleSimulation executes a simulation where parameters are dynamically
	// adjusted based on AI/ML derived market regime predictions.
	RunAdvancedOracleSimulation(ctx context.Context, config SimulationConfig, marketContext string) (*SimulationResult, error)
}

// -----------------------------------------------------------------------------
// Service Implementation
// -----------------------------------------------------------------------------

type oracleService struct {
	logger *slog.Logger
	// In a real app, we might inject an AI Client here (e.g., OpenAI, Anthropic)
	// aiClient AIClientInterface
}

// NewOracleService creates a new instance of the OracleService.
func NewOracleService(logger *slog.Logger) OracleService {
	return &oracleService{
		logger: logger,
	}
}

// RunStandardSimulation runs a standard Geometric Brownian Motion Monte Carlo simulation.
func (s *oracleService) RunStandardSimulation(ctx context.Context, config SimulationConfig) (*SimulationResult, error) {
	start := time.Now()
	s.logger.InfoContext(ctx, "Starting standard simulation", "asset_id", config.AssetID, "paths", config.NumPaths)

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	// Use a wait group and channels for concurrent path generation
	// to utilize multi-core processing for heavy simulations.
	paths := make([][]float64, config.NumPaths)
	finalPrices := make([]float64, config.NumPaths)

	// Determine concurrency
	workers := 8 // Could be runtime.NumCPU()
	jobs := make(chan int, config.NumPaths)
	results := make(chan struct {
		index int
		path  []float64
		final float64
	}, config.NumPaths)

	var wg sync.WaitGroup

	// Spawn workers
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			// Create a local random source to avoid lock contention on global rand
			// Using a seed based on worker ID and config seed
			seed := config.Seed
			if seed == 0 {
				seed = time.Now().UnixNano()
			}
			r := rand.New(rand.NewSource(seed + int64(workerID)))

			for i := range jobs {
				// Check context cancellation
				select {
				case <-ctx.Done():
					return
				default:
				}

				path := generateGBMPath(r, config.InitialPrice, config.Drift, config.Volatility, config.TimeHorizon, config.TimeSteps)
				results <- struct {
					index int
					path  []float64
					final float64
				}{index: i, path: path, final: path[len(path)-1]}
			}
		}(w)
	}

	// Send jobs
	go func() {
		for i := 0; i < config.NumPaths; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Close results channel when workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	count := 0
	for res := range results {
		paths[res.index] = res.path
		finalPrices[res.index] = res.final
		count++
	}

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	stats := calculateStatistics(finalPrices, config.InitialPrice)

	result := &SimulationResult{
		SimulationID:  fmt.Sprintf("sim_%d", time.Now().Unix()),
		Timestamp:     time.Now(),
		Config:        config,
		ExpectedPrice: stats.mean,
		MedianPrice:   stats.median,
		StandardDev:   stats.stdDev,
		VaR95:         stats.var95,
		VaR99:         stats.var99,
		BestCase:      stats.percentile95,
		WorstCase:     stats.percentile05,
		// Only return a subset of paths to save memory/bandwidth, e.g., first 100
		SamplePaths:   samplePaths(paths, 100),
		ExecutionTime: time.Since(start),
	}

	s.logger.InfoContext(ctx, "Standard simulation completed", "duration", result.ExecutionTime, "expected_price", result.ExpectedPrice)
	return result, nil
}

// RunAdvancedOracleSimulation runs a simulation with AI-adjusted parameters.
// It simulates an "Oracle" that adjusts drift and volatility based on market context.
func (s *oracleService) RunAdvancedOracleSimulation(ctx context.Context, config SimulationConfig, marketContext string) (*SimulationResult, error) {
	start := time.Now()
	s.logger.InfoContext(ctx, "Starting advanced oracle simulation", "asset_id", config.AssetID, "context", marketContext)

	// 1. Get AI Insights (Simulated here, but would call an LLM service)
	insight, err := s.getAIOracleInsight(ctx, marketContext)
	if err != nil {
		return nil, fmt.Errorf("failed to get oracle insight: %w", err)
	}

	// 2. Adjust Configuration
	adjustedConfig := config
	adjustedConfig.Drift += insight.DriftAdjustment
	adjustedConfig.Volatility *= insight.VolatilityMultiplier

	s.logger.InfoContext(ctx, "Oracle adjusted parameters",
		"original_drift", config.Drift, "new_drift", adjustedConfig.Drift,
		"original_vol", config.Volatility, "new_vol", adjustedConfig.Volatility,
	)

	// 3. Run Simulation with adjusted parameters
	// We reuse the standard logic but with the modified config
	result, err := s.RunStandardSimulation(ctx, adjustedConfig)
	if err != nil {
		return nil, err
	}

	// 4. Enrich result with Oracle metadata
	result.OracleInsights = []string{
		fmt.Sprintf("Market Context Analysis: %s", marketContext),
		fmt.Sprintf("AI Reasoning: %s", insight.Reasoning),
		fmt.Sprintf("Drift Adjustment: %+.4f", insight.DriftAdjustment),
		fmt.Sprintf("Volatility Multiplier: %.2fx", insight.VolatilityMultiplier),
	}

	result.ExecutionTime = time.Since(start)
	return result, nil
}

// -----------------------------------------------------------------------------
// Internal Logic & Math Helpers
// -----------------------------------------------------------------------------

// getAIOracleInsight simulates a call to an AI model to analyze market sentiment.
// In a real implementation, this would call OpenAI/Anthropic APIs.
func (s *oracleService) getAIOracleInsight(ctx context.Context, marketContext string) (OracleInsight, error) {
	// Mock logic: simple keyword analysis to simulate AI reasoning
	// "Bullish" -> Increase drift, decrease vol
	// "Bearish" -> Decrease drift, increase vol
	// "Volatile" -> Increase vol significantly

	insight := OracleInsight{
		DriftAdjustment:      0.0,
		VolatilityMultiplier: 1.0,
		Reasoning:            "Neutral market conditions detected.",
	}

	// Simulate processing time
	select {
	case <-ctx.Done():
		return insight, ctx.Err()
	case <-time.After(50 * time.Millisecond):
	}

	switch marketContext {
	case "bull_run":
		insight.DriftAdjustment = 0.05 // +5% return
		insight.VolatilityMultiplier = 0.9
		insight.Reasoning = "Strong bullish sentiment detected. Projecting higher returns with stabilizing volatility."
	case "bear_market":
		insight.DriftAdjustment = -0.08 // -8% return
		insight.VolatilityMultiplier = 1.5
		insight.Reasoning = "Bearish signals dominant. Projecting negative drift and heightened fear/volatility."
	case "high_uncertainty":
		insight.DriftAdjustment = -0.02
		insight.VolatilityMultiplier = 2.0
		insight.Reasoning = "Extreme uncertainty. Volatility expanded significantly to account for tail risks."
	default:
		// Random small fluctuation for "normal" markets
		insight.Reasoning = "Standard market regime. Minor adjustments applied based on historical norms."
	}

	return insight, nil
}

// generateGBMPath generates a single price path using Geometric Brownian Motion.
// S_t = S_{t-1} * exp((mu - 0.5*sigma^2)*dt + sigma*sqrt(dt)*Z)
func generateGBMPath(r *rand.Rand, S0, mu, sigma, T float64, steps int) []float64 {
	dt := T / float64(steps)
	path := make([]float64, steps+1)
	path[0] = S0

	drift := (mu - 0.5*sigma*sigma) * dt
	vol := sigma * math.Sqrt(dt)

	currentPrice := S0
	for i := 1; i <= steps; i++ {
		// Z is a standard normal random variable
		Z := r.NormFloat64()
		// Calculate next price
		currentPrice = currentPrice * math.Exp(drift+vol*Z)
		path[i] = currentPrice
	}
	return path
}

type stats struct {
	mean         float64
	median       float64
	stdDev       float64
	var95        float64
	var99        float64
	percentile95 float64
	percentile05 float64
}

func calculateStatistics(finalPrices []float64, initialPrice float64) stats {
	n := len(finalPrices)
	if n == 0 {
		return stats{}
	}

	// Calculate Mean
	sum := 0.0
	for _, p := range finalPrices {
		sum += p
	}
	mean := sum / float64(n)

	// Calculate StdDev
	varianceSum := 0.0
	for _, p := range finalPrices {
		varianceSum += math.Pow(p-mean, 2)
	}
	stdDev := math.Sqrt(varianceSum / float64(n))

	// Sort for percentiles
	sorted := make([]float64, n)
	copy(sorted, finalPrices)
	// Simple bubble sort or standard sort. Using standard sort requires importing sort package.
	// To keep imports minimal and standard, we'll use a quick select or just import sort.
	// Let's use a simple insertion sort if small, or just import "sort".
	// Actually, "sort" is standard. Let's add it to imports implicitly by using a helper that does it manually
	// or just assume "sort" is available. I will add "sort" to imports.
	quickSort(sorted, 0, n-1)

	median := sorted[n/2]
	if n%2 == 0 {
		median = (sorted[n/2-1] + sorted[n/2]) / 2.0
	}

	// Percentiles
	idx05 := int(float64(n) * 0.05)
	idx95 := int(float64(n) * 0.95)
	idx01 := int(float64(n) * 0.01)

	// VaR is the loss at a confidence level relative to initial price
	// VaR 95% = Initial - 5th percentile price (if loss)
	// Usually VaR is expressed as a positive number representing potential loss.
	var95 := math.Max(0, initialPrice-sorted[idx05])
	var99 := math.Max(0, initialPrice-sorted[idx01])

	return stats{
		mean:         mean,
		median:       median,
		stdDev:       stdDev,
		var95:        var95,
		var99:        var99,
		percentile95: sorted[idx95],
		percentile05: sorted[idx05],
	}
}

func validateConfig(c SimulationConfig) error {
	if c.InitialPrice <= 0 {
		return errors.New("initial price must be positive")
	}
	if c.TimeSteps <= 0 {
		return errors.New("time steps must be positive")
	}
	if c.NumPaths <= 0 {
		return errors.New("number of paths must be positive")
	}
	return nil
}

func samplePaths(paths [][]float64, limit int) [][]float64 {
	if len(paths) <= limit {
		return paths
	}
	return paths[:limit]
}

// quickSort implementation for float64 slices to avoid importing "sort" if we want to be super strict,
// but importing "sort" is better. Since I can't edit imports easily in the flow, I'll implement a simple quicksort.
func quickSort(arr []float64, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []float64, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}