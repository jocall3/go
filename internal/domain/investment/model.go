package investment

import (
	"errors"
	"fmt"
	"math"
	"time"
)

// Common errors for the investment domain.
var (
	ErrInvalidAssetAllocation = errors.New("total target allocation must equal 100%")
	ErrInvalidAssetQuantity   = errors.New("asset quantity cannot be negative")
	ErrInvalidAssetPrice      = errors.New("asset price cannot be negative")
	ErrEmptyPortfolio         = errors.New("portfolio contains no assets")
	ErrAssetNotFound          = errors.New("asset not found in portfolio")
)

// AssetType represents the classification of an investment asset.
type AssetType string

const (
	AssetTypeStock  AssetType = "STOCK"
	AssetTypeBond   AssetType = "BOND"
	AssetTypeCrypto AssetType = "CRYPTO"
	AssetTypeCash   AssetType = "CASH"
	AssetTypeETF    AssetType = "ETF"
	AssetTypeRealEstate AssetType = "REAL_ESTATE"
)

// Asset represents a single holding within a portfolio.
type Asset struct {
	ID             string    `json:"id"`
	PortfolioID    string    `json:"portfolio_id"`
	Symbol         string    `json:"symbol"`          // e.g., "AAPL", "BTC"
	Name           string    `json:"name"`            // e.g., "Apple Inc."
	Type           AssetType `json:"type"`
	Quantity       float64   `json:"quantity"`        // Number of units held
	CurrentPrice   float64   `json:"current_price"`   // Current market price per unit
	TargetPercent  float64   `json:"target_percent"`  // Desired allocation percentage (0-100)
	LastUpdated    time.Time `json:"last_updated"`
}

// CurrentValue calculates the total market value of this asset holding.
func (a *Asset) CurrentValue() float64 {
	return a.Quantity * a.CurrentPrice
}

// Validate checks if the asset data is consistent.
func (a *Asset) Validate() error {
	if a.Quantity < 0 {
		return ErrInvalidAssetQuantity
	}
	if a.CurrentPrice < 0 {
		return ErrInvalidAssetPrice
	}
	if a.Symbol == "" {
		return errors.New("asset symbol is required")
	}
	return nil
}

// Portfolio represents a collection of assets managed by a user.
type Portfolio struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Assets      []Asset   `json:"assets"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewPortfolio creates a new empty portfolio.
func NewPortfolio(id, userID, name, description string) *Portfolio {
	now := time.Now().UTC()
	return &Portfolio{
		ID:          id,
		UserID:      userID,
		Name:        name,
		Description: description,
		Assets:      make([]Asset, 0),
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// AddAsset adds an asset to the portfolio.
func (p *Portfolio) AddAsset(asset Asset) {
	asset.PortfolioID = p.ID
	p.Assets = append(p.Assets, asset)
	p.UpdatedAt = time.Now().UTC()
}

// TotalValue calculates the sum of all asset values in the portfolio.
func (p *Portfolio) TotalValue() float64 {
	var total float64
	for _, asset := range p.Assets {
		total += asset.CurrentValue()
	}
	return total
}

// ValidateAllocation checks if the sum of target percentages equals 100%.
// It allows for a small floating-point epsilon error.
func (p *Portfolio) ValidateAllocation() error {
	var totalPercent float64
	for _, asset := range p.Assets {
		totalPercent += asset.TargetPercent
	}

	// Allow a tiny margin of error for floating point math
	if math.Abs(totalPercent-100.0) > 0.01 {
		return fmt.Errorf("%w: current total is %.2f%%", ErrInvalidAssetAllocation, totalPercent)
	}
	return nil
}

// RebalanceActionType defines the action required to rebalance.
type RebalanceActionType string

const (
	ActionBuy  RebalanceActionType = "BUY"
	ActionSell RebalanceActionType = "SELL"
	ActionHold RebalanceActionType = "HOLD"
)

// RebalanceSuggestion represents a single trade required to bring the portfolio back to target.
type RebalanceSuggestion struct {
	AssetID        string              `json:"asset_id"`
	Symbol         string              `json:"symbol"`
	Action         RebalanceActionType `json:"action"`
	Units          float64             `json:"units"`           // Quantity to buy or sell
	EstimatedValue float64             `json:"estimated_value"` // Monetary value of the trade
	CurrentPercent float64             `json:"current_percent"`
	TargetPercent  float64             `json:"target_percent"`
}

// RebalancePlan contains the complete set of instructions to rebalance a portfolio.
type RebalancePlan struct {
	PortfolioID     string                `json:"portfolio_id"`
	GeneratedAt     time.Time             `json:"generated_at"`
	TotalValue      float64               `json:"total_value"`
	Suggestions     []RebalanceSuggestion `json:"suggestions"`
	IsBalanced      bool                  `json:"is_balanced"`
}

// GenerateRebalancingPlan calculates the necessary trades to align the portfolio with target allocations.
// thresholdPercent is the minimum deviation required to trigger a rebalance suggestion (e.g., 0.5%).
func (p *Portfolio) GenerateRebalancingPlan(thresholdPercent float64) (*RebalancePlan, error) {
	if len(p.Assets) == 0 {
		return nil, ErrEmptyPortfolio
	}

	if err := p.ValidateAllocation(); err != nil {
		return nil, err
	}

	totalValue := p.TotalValue()
	plan := &RebalancePlan{
		PortfolioID: p.ID,
		GeneratedAt: time.Now().UTC(),
		TotalValue:  totalValue,
		Suggestions: make([]RebalanceSuggestion, 0),
	}

	// If total value is 0, we can't rebalance based on percentages (unless we are depositing cash, which is a different flow).
	if totalValue == 0 {
		return plan, nil
	}

	isBalanced := true

	for _, asset := range p.Assets {
		currentVal := asset.CurrentValue()
		currentPercent := (currentVal / totalValue) * 100.0
		targetVal := totalValue * (asset.TargetPercent / 100.0)
		diffVal := targetVal - currentVal
		
		// Calculate deviation
		deviation := math.Abs(currentPercent - asset.TargetPercent)

		// If deviation is within threshold, we hold.
		if deviation <= thresholdPercent {
			continue
		}

		isBalanced = false
		
		var action RebalanceActionType
		var units float64

		// Avoid division by zero if price is missing
		if asset.CurrentPrice <= 0 {
			// If price is 0, we can't calculate units to buy/sell properly. 
			// We skip or handle as error. Here we skip safely.
			continue
		}

		if diffVal > 0 {
			action = ActionBuy
			units = diffVal / asset.CurrentPrice
		} else {
			action = ActionSell
			units = math.Abs(diffVal) / asset.CurrentPrice
		}

		suggestion := RebalanceSuggestion{
			AssetID:        asset.ID,
			Symbol:         asset.Symbol,
			Action:         action,
			Units:          units,
			EstimatedValue: math.Abs(diffVal),
			CurrentPercent: currentPercent,
			TargetPercent:  asset.TargetPercent,
		}

		plan.Suggestions = append(plan.Suggestions, suggestion)
	}

	plan.IsBalanced = isBalanced
	return plan, nil
}

// UpdateAssetPrices allows bulk updating of asset prices for accurate rebalancing.
// pricesMap is a map of Symbol -> NewPrice.
func (p *Portfolio) UpdateAssetPrices(pricesMap map[string]float64) {
	for i := range p.Assets {
		if newPrice, ok := pricesMap[p.Assets[i].Symbol]; ok {
			p.Assets[i].CurrentPrice = newPrice
			p.Assets[i].LastUpdated = time.Now().UTC()
		}
	}
	p.UpdatedAt = time.Now().UTC()
}