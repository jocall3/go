package sustainability

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"time"
)

// Constants for Carbon Calculation (approximate global averages in kg CO2e)
const (
	EmissionFactorEnergyKWh      = 0.233 // kg CO2e per kWh
	EmissionFactorTransportCar   = 0.192 // kg CO2e per mile
	EmissionFactorTransportBus   = 0.105
	EmissionFactorTransportTrain = 0.041
	EmissionFactorTransportPlane = 0.255

	// Daily dietary impact
	EmissionDietVegan      = 2.9
	EmissionDietVegetarian = 3.8
	EmissionDietOmnivore   = 5.6
	EmissionDietHeavyMeat  = 7.2
)

var (
	ErrInvalidInput     = errors.New("invalid input parameters")
	ErrOffsetUnavailable = errors.New("carbon offset provider unavailable")
	ErrInsufficientFunds = errors.New("insufficient funds for offset purchase")
)

// Domain Models for Sustainability Context

type TransportType string

const (
	TransportCar   TransportType = "car"
	TransportBus   TransportType = "bus"
	TransportTrain TransportType = "train"
	TransportPlane TransportType = "plane"
)

type DietType string

const (
	DietVegan      DietType = "vegan"
	DietVegetarian DietType = "vegetarian"
	DietOmnivore   DietType = "omnivore"
	DietHeavyMeat  DietType = "heavy_meat"
)

// FootprintInput represents the user data required to calculate a carbon footprint.
type FootprintInput struct {
	EnergyUsageKWh float64       `json:"energy_usage_kwh"`
	TransportMiles float64       `json:"transport_miles"`
	TransportMode  TransportType `json:"transport_mode"`
	Diet           DietType      `json:"diet"`
	PeriodDays     int           `json:"period_days"` // Number of days this data covers
}

// FootprintResult contains the calculated emissions data.
type FootprintResult struct {
	TotalEmissionsKg float64            `json:"total_emissions_kg"`
	Breakdown        map[string]float64 `json:"breakdown"`
	CalculatedAt     time.Time          `json:"calculated_at"`
	Rating           string             `json:"rating"` // e.g., "Eco-Friendly", "Average", "High"
}

// OffsetPurchaseRequest represents a request to buy carbon credits.
type OffsetPurchaseRequest struct {
	UserID       string  `json:"user_id"`
	AmountKgCO2e float64 `json:"amount_kg_co2e"`
	ProviderID   string  `json:"provider_id"`
	PaymentToken string  `json:"payment_token"` // Tokenized payment info
}

// OffsetPurchaseReceipt represents the confirmation of a purchase.
type OffsetPurchaseReceipt struct {
	TransactionID   string    `json:"transaction_id"`
	CertificateCode string    `json:"certificate_code"`
	AmountOffsetKg  float64   `json:"amount_offset_kg"`
	ProviderName    string    `json:"provider_name"`
	Timestamp       time.Time `json:"timestamp"`
	Status          string    `json:"status"`
}

// Repository defines the data persistence layer requirements for sustainability.
type Repository interface {
	SaveFootprint(ctx context.Context, userID string, result FootprintResult) error
	GetLatestFootprint(ctx context.Context, userID string) (*FootprintResult, error)
	RecordOffsetTransaction(ctx context.Context, req OffsetPurchaseRequest, receipt OffsetPurchaseReceipt) error
	GetTotalOffsetByUser(ctx context.Context, userID string) (float64, error)
}

// Service defines the interface for sustainability business logic.
type Service interface {
	CalculateFootprint(ctx context.Context, userID string, input FootprintInput) (*FootprintResult, error)
	PurchaseOffset(ctx context.Context, req OffsetPurchaseRequest) (*OffsetPurchaseReceipt, error)
	GetSustainabilityProfile(ctx context.Context, userID string) (map[string]interface{}, error)
}

type service struct {
	repo   Repository
	logger *slog.Logger
}

// NewService creates a new instance of the Sustainability Service.
func NewService(repo Repository, logger *slog.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

// CalculateFootprint processes user input to determine carbon emissions and saves the result.
func (s *service) CalculateFootprint(ctx context.Context, userID string, input FootprintInput) (*FootprintResult, error) {
	s.logger.InfoContext(ctx, "calculating footprint", "user_id", userID)

	if input.PeriodDays <= 0 {
		return nil, fmt.Errorf("%w: period days must be positive", ErrInvalidInput)
	}

	breakdown := make(map[string]float64)
	var total float64

	// 1. Energy Calculation
	energyEmission := input.EnergyUsageKWh * EmissionFactorEnergyKWh
	breakdown["energy"] = round(energyEmission)
	total += energyEmission

	// 2. Transport Calculation
	var transportFactor float64
	switch input.TransportMode {
	case TransportCar:
		transportFactor = EmissionFactorTransportCar
	case TransportBus:
		transportFactor = EmissionFactorTransportBus
	case TransportTrain:
		transportFactor = EmissionFactorTransportTrain
	case TransportPlane:
		transportFactor = EmissionFactorTransportPlane
	default:
		transportFactor = EmissionFactorTransportCar // Default fallback
	}
	transportEmission := input.TransportMiles * transportFactor
	breakdown["transport"] = round(transportEmission)
	total += transportEmission

	// 3. Diet Calculation
	var dietFactor float64
	switch input.Diet {
	case DietVegan:
		dietFactor = EmissionDietVegan
	case DietVegetarian:
		dietFactor = EmissionDietVegetarian
	case DietHeavyMeat:
		dietFactor = EmissionDietHeavyMeat
	case DietOmnivore:
		fallthrough
	default:
		dietFactor = EmissionDietOmnivore
	}
	dietEmission := dietFactor * float64(input.PeriodDays)
	breakdown["diet"] = round(dietEmission)
	total += dietEmission

	// Determine Rating
	// Global average is roughly 4000-5000kg per year.
	// Daily average ~12kg.
	dailyAvg := total / float64(input.PeriodDays)
	var rating string
	if dailyAvg < 8.0 {
		rating = "Eco-Warrior"
	} else if dailyAvg < 15.0 {
		rating = "Conscious Citizen"
	} else {
		rating = "High Impact"
	}

	result := FootprintResult{
		TotalEmissionsKg: round(total),
		Breakdown:        breakdown,
		CalculatedAt:     time.Now().UTC(),
		Rating:           rating,
	}

	// Persist result
	if err := s.repo.SaveFootprint(ctx, userID, result); err != nil {
		s.logger.ErrorContext(ctx, "failed to save footprint", "error", err)
		return nil, fmt.Errorf("failed to save footprint history: %w", err)
	}

	return &result, nil
}

// PurchaseOffset handles the logic for buying carbon credits.
func (s *service) PurchaseOffset(ctx context.Context, req OffsetPurchaseRequest) (*OffsetPurchaseReceipt, error) {
	s.logger.InfoContext(ctx, "processing offset purchase", "user_id", req.UserID, "amount", req.AmountKgCO2e)

	if req.AmountKgCO2e <= 0 {
		return nil, fmt.Errorf("%w: amount must be positive", ErrInvalidInput)
	}

	// In a real app, we would call an external Payment Gateway and a Carbon Registry API here.
	// We simulate this process.

	// 1. Simulate Payment Processing
	if req.PaymentToken == "" {
		return nil, fmt.Errorf("%w: missing payment token", ErrInvalidInput)
	}
	// Simulate payment latency
	time.Sleep(100 * time.Millisecond)

	// 2. Generate Receipt
	txID, err := generateID()
	if err != nil {
		return nil, err
	}
	certCode, err := generateID()
	if err != nil {
		return nil, err
	}

	receipt := OffsetPurchaseReceipt{
		TransactionID:   txID,
		CertificateCode: fmt.Sprintf("CERT-%s", certCode[0:8]),
		AmountOffsetKg:  req.AmountKgCO2e,
		ProviderName:    "Global Green Initiative", // Mock provider
		Timestamp:       time.Now().UTC(),
		Status:          "COMPLETED",
	}

	// 3. Record Transaction
	if err := s.repo.RecordOffsetTransaction(ctx, req, receipt); err != nil {
		s.logger.ErrorContext(ctx, "failed to record offset transaction", "error", err)
		// Note: In production, this might require a rollback of the payment or a reconciliation queue
		return nil, fmt.Errorf("transaction recorded failed: %w", err)
	}

	s.logger.InfoContext(ctx, "offset purchase successful", "tx_id", receipt.TransactionID)
	return &receipt, nil
}

// GetSustainabilityProfile aggregates user stats.
func (s *service) GetSustainabilityProfile(ctx context.Context, userID string) (map[string]interface{}, error) {
	latest, err := s.repo.GetLatestFootprint(ctx, userID)
	if err != nil {
		s.logger.WarnContext(ctx, "could not fetch latest footprint", "user_id", userID, "error", err)
		// Continue, as they might be a new user
	}

	totalOffset, err := s.repo.GetTotalOffsetByUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch offset history: %w", err)
	}

	profile := make(map[string]interface{})
	profile["user_id"] = userID
	profile["total_offset_kg"] = totalOffset

	if latest != nil {
		profile["latest_footprint_kg"] = latest.TotalEmissionsKg
		profile["rating"] = latest.Rating
		profile["last_calculated"] = latest.CalculatedAt
		
		// Net Impact = Footprint - Offsets
		// Note: This is a simplification. Usually footprint is per year/period.
		net := latest.TotalEmissionsKg - totalOffset
		if net < 0 {
			net = 0
		}
		profile["net_emissions_kg"] = round(net)
	} else {
		profile["status"] = "No data calculated yet"
	}

	return profile, nil
}

// Helper functions

func round(val float64) float64 {
	return math.Round(val*100) / 100
}

func generateID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}