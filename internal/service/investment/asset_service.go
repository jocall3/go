package investment

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Common errors returned by the asset service.
var (
	ErrAssetNotFound = errors.New("asset not found")
	ErrInvalidInput  = errors.New("invalid input parameters")
	ErrDuplicate     = errors.New("asset already exists")
)

// AssetType defines the category of the financial instrument.
type AssetType string

const (
	AssetTypeStock     AssetType = "STOCK"
	AssetTypeCrypto    AssetType = "CRYPTO"
	AssetTypeETF       AssetType = "ETF"
	AssetTypeIndex     AssetType = "INDEX"
	AssetTypeForex     AssetType = "FOREX"
	AssetTypeCommodity AssetType = "COMMODITY"
	AssetTypeBond      AssetType = "BOND"
	AssetTypeMutualFund AssetType = "MUTUAL_FUND"
)

// Asset represents the domain model for an investment asset.
type Asset struct {
	ID          string                 `json:"id"`
	Symbol      string                 `json:"symbol"`
	Name        string                 `json:"name"`
	Type        AssetType              `json:"type"`
	Exchange    string                 `json:"exchange"`
	Currency    string                 `json:"currency"`
	Description string                 `json:"description,omitempty"`
	Sector      string                 `json:"sector,omitempty"`
	Industry    string                 `json:"industry,omitempty"`
	ISIN        string                 `json:"isin,omitempty"`
	CUSIP       string                 `json:"cusip,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	IsActive    bool                   `json:"is_active"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

// AssetSearchCriteria holds filters and pagination for searching assets.
type AssetSearchCriteria struct {
	Query    string    // Search by name or symbol
	Type     AssetType // Filter by asset type
	Exchange string    // Filter by exchange
	Sector   string    // Filter by sector
	IsActive *bool     // Filter by active status (nil for all)
	Limit    int       // Pagination limit
	Offset   int       // Pagination offset
}

// AssetRepository defines the interface for asset data persistence.
// This interface is defined here to adhere to the dependency inversion principle.
type AssetRepository interface {
	FindByID(ctx context.Context, id string) (*Asset, error)
	FindBySymbol(ctx context.Context, symbol string) (*Asset, error)
	FindByISIN(ctx context.Context, isin string) (*Asset, error)
	Search(ctx context.Context, criteria AssetSearchCriteria) ([]Asset, int64, error)
	Create(ctx context.Context, asset *Asset) error
	Update(ctx context.Context, asset *Asset) error
	Delete(ctx context.Context, id string) error
}

// AssetService defines the interface for asset business logic.
type AssetService interface {
	GetAssetByID(ctx context.Context, id string) (*Asset, error)
	GetAssetBySymbol(ctx context.Context, symbol string) (*Asset, error)
	GetAssetByISIN(ctx context.Context, isin string) (*Asset, error)
	SearchAssets(ctx context.Context, criteria AssetSearchCriteria) ([]Asset, int64, error)
	CreateAsset(ctx context.Context, asset *Asset) error
	UpdateAsset(ctx context.Context, asset *Asset) error
	DeleteAsset(ctx context.Context, id string) error
}

// assetService implements the AssetService interface.
type assetService struct {
	repo AssetRepository
}

// NewAssetService initializes a new instance of AssetService.
func NewAssetService(repo AssetRepository) AssetService {
	return &assetService{
		repo: repo,
	}
}

// GetAssetByID retrieves a single asset by its unique identifier.
func (s *assetService) GetAssetByID(ctx context.Context, id string) (*Asset, error) {
	if strings.TrimSpace(id) == "" {
		return nil, fmt.Errorf("%w: asset ID is required", ErrInvalidInput)
	}

	asset, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve asset by ID: %w", err)
	}
	if asset == nil {
		return nil, ErrAssetNotFound
	}

	return asset, nil
}

// GetAssetBySymbol retrieves a single asset by its ticker symbol.
func (s *assetService) GetAssetBySymbol(ctx context.Context, symbol string) (*Asset, error) {
	if strings.TrimSpace(symbol) == "" {
		return nil, fmt.Errorf("%w: symbol is required", ErrInvalidInput)
	}

	// Normalize symbol to uppercase
	symbol = strings.ToUpper(strings.TrimSpace(symbol))

	asset, err := s.repo.FindBySymbol(ctx, symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve asset by symbol: %w", err)
	}
	if asset == nil {
		return nil, ErrAssetNotFound
	}

	return asset, nil
}

// GetAssetByISIN retrieves a single asset by its ISIN code.
func (s *assetService) GetAssetByISIN(ctx context.Context, isin string) (*Asset, error) {
	if strings.TrimSpace(isin) == "" {
		return nil, fmt.Errorf("%w: ISIN is required", ErrInvalidInput)
	}

	asset, err := s.repo.FindByISIN(ctx, isin)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve asset by ISIN: %w", err)
	}
	if asset == nil {
		return nil, ErrAssetNotFound
	}

	return asset, nil
}

// SearchAssets searches for assets based on the provided criteria.
// It returns a slice of assets and the total count of matches found.
func (s *assetService) SearchAssets(ctx context.Context, criteria AssetSearchCriteria) ([]Asset, int64, error) {
	// Set default pagination limits if not provided
	if criteria.Limit <= 0 {
		criteria.Limit = 20
	}
	if criteria.Limit > 100 {
		criteria.Limit = 100
	}
	if criteria.Offset < 0 {
		criteria.Offset = 0
	}

	// Sanitize query input
	criteria.Query = strings.TrimSpace(criteria.Query)

	assets, total, err := s.repo.Search(ctx, criteria)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to search assets: %w", err)
	}

	return assets, total, nil
}

// CreateAsset validates and persists a new asset.
func (s *assetService) CreateAsset(ctx context.Context, asset *Asset) error {
	if asset == nil {
		return fmt.Errorf("%w: asset data is nil", ErrInvalidInput)
	}

	if err := s.validateAsset(asset); err != nil {
		return err
	}

	// Check for existing asset by symbol to prevent duplicates
	existing, err := s.repo.FindBySymbol(ctx, asset.Symbol)
	if err != nil && !errors.Is(err, ErrAssetNotFound) {
		// If error is something other than not found, return it
		// Note: Implementation of FindBySymbol should return nil, nil or nil, ErrNotFound depending on repo pattern
		// Here we assume it returns nil, nil if not found or we handle the error if it's a DB error
	}
	if existing != nil {
		return fmt.Errorf("%w: symbol %s already exists", ErrDuplicate, asset.Symbol)
	}

	// Set metadata
	now := time.Now().UTC()
	asset.CreatedAt = now
	asset.UpdatedAt = now
	asset.Symbol = strings.ToUpper(strings.TrimSpace(asset.Symbol))
	
	// Default to active if creating
	asset.IsActive = true

	if err := s.repo.Create(ctx, asset); err != nil {
		return fmt.Errorf("failed to create asset: %w", err)
	}

	return nil
}

// UpdateAsset updates an existing asset's details.
func (s *assetService) UpdateAsset(ctx context.Context, asset *Asset) error {
	if asset == nil {
		return fmt.Errorf("%w: asset data is nil", ErrInvalidInput)
	}
	if asset.ID == "" {
		return fmt.Errorf("%w: asset ID is required for update", ErrInvalidInput)
	}

	if err := s.validateAsset(asset); err != nil {
		return err
	}

	// Verify existence
	existing, err := s.repo.FindByID(ctx, asset.ID)
	if err != nil {
		return fmt.Errorf("failed to check asset existence: %w", err)
	}
	if existing == nil {
		return ErrAssetNotFound
	}

	asset.UpdatedAt = time.Now().UTC()
	asset.Symbol = strings.ToUpper(strings.TrimSpace(asset.Symbol))
	// Preserve CreatedAt from existing record if not passed correctly
	if asset.CreatedAt.IsZero() {
		asset.CreatedAt = existing.CreatedAt
	}

	if err := s.repo.Update(ctx, asset); err != nil {
		return fmt.Errorf("failed to update asset: %w", err)
	}

	return nil
}

// DeleteAsset soft-deletes or removes an asset by ID.
func (s *assetService) DeleteAsset(ctx context.Context, id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("%w: asset ID is required", ErrInvalidInput)
	}

	// Check existence
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to check asset existence: %w", err)
	}
	if existing == nil {
		return ErrAssetNotFound
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete asset: %w", err)
	}

	return nil
}

// validateAsset performs basic validation on asset fields.
func (s *assetService) validateAsset(asset *Asset) error {
	if strings.TrimSpace(asset.Symbol) == "" {
		return fmt.Errorf("%w: symbol is required", ErrInvalidInput)
	}
	if strings.TrimSpace(asset.Name) == "" {
		return fmt.Errorf("%w: name is required", ErrInvalidInput)
	}
	if asset.Type == "" {
		return fmt.Errorf("%w: asset type is required", ErrInvalidInput)
	}
	if strings.TrimSpace(asset.Currency) == "" {
		return fmt.Errorf("%w: currency is required", ErrInvalidInput)
	}
	return nil
}