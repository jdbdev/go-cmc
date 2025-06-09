package ticker

// Coinmarketcap (CMC) API Documentation: https://coinmarketcap.com/api/documentation/v1/
// CMC recommends using CoinMarketCap ID's instead of ID or other identifiers
// Common endpoints:
// https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest
// https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest

// Sample CMD ID's:
// Bitcoin CMC ID: 1
// Ethereum CMC ID: 1027
// Solana CMC ID: 5994

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/jdbdev/go-cmc/config"
	"github.com/jdbdev/go-cmc/internal/mapper"
	"github.com/jdbdev/go-cmc/utils"
)

var client = &http.Client{}

type TickerInterface interface {
	FetchAndDecodeData() error
	UpdateDB() error
}

type TickerService struct {
	apiKey    string
	baseURL   string
	quotesURL string
	client    *http.Client
	mapper    mapper.IDMapInterface // example: mapper.GetIDMap()
	// data    []TickerData // Add a field to store the decoded data
}

// NewTickerService creates a new instance of the TickerService struct
func NewTickerService(app *config.AppConfig, mapService mapper.IDMapInterface) *TickerService {
	return &TickerService{
		apiKey:    app.CMC.APIKey,
		baseURL:   app.CMC.BaseURL,
		quotesURL: app.CMC.QuotesURL,
		client:    client,
		mapper:    mapService,
	}
}

// FetchAndDecodeData gets and decodes data from CMC
func (t *TickerService) FetchAndDecodeData() error {
	client := t.client
	req, err := http.NewRequest("GET", t.quotesURL, nil)
	if err != nil {
		return err
	}

	// Build query parameters
	q := url.Values{}

	// Collect all IDs from the map
	var ids []string
	for _, id := range t.mapper.GetIDMap() {
		ids = append(ids, id)
	}
	// Join IDs with commas and add to query
	q.Add("id", strings.Join(ids, ","))
	q.Add("convert", "USD")

	// Only get requested fields (automatically get price, market_cap, volume_24h, etc. in "quotes"):
	// Available aux fields: num_market_pairs, cmc_rank, date_added, tags, platform, max_supply,
	// circulating_supply, total_supply, market_cap_by_total_supply, volume_24h_reported,
	// volume_7d, volume_7d_reported, volume_30d, volume_30d_reported, is_active, is_fiat
	q.Add("aux", "circulating_supply,total_supply,volume_24h_reported")

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CMC_PRO_API_KEY", t.apiKey)

	// Add query parameters to URL
	req.URL.RawQuery = q.Encode()

	// Debug URL
	fmt.Printf("Making request to: %s\n", req.URL.String())

	// Execute request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Print Response Status
	fmt.Printf("Response Status: %s\n", resp.Status)

	// Read and debug response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Response Body: %s\n\n\n", string(respBody))

	// Write response body to file (provide file name)
	err = utils.WriteJSONToFile(respBody, "sample_response")
	if err != nil {
		return err
	}

	return nil
}

// UpdateDB updates the database with data from CMC
func (t *TickerService) UpdateDB() error {
	return nil
}
