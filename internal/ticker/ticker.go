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
)

// TEMP CMCIDMap is a map of CMC ID's
var CMCIDMap = map[string]string{
	"BTC": "1",
	"ETH": "1027",
	"SOL": "5994",
}

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
	// data    []TickerData // Add a field to store the decoded data
}

// NewTickerService creates a new instance of the TickerService
func NewTickerService(app *config.AppConfig) *TickerService {
	return &TickerService{
		apiKey:    app.CMC.APIKey,
		baseURL:   app.CMC.BaseURL,
		quotesURL: app.CMC.QuotesURL,
		client:    client,
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
	for _, id := range CMCIDMap {
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

	fmt.Printf("Making request to: %s\n", req.URL.String()) // Debug URL
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("Response Status: %s\n", resp.Status)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Response Body: %s\n\n\n", string(respBody))

	return nil
}

// UpdateDB updates the database with data from CMC
func (t *TickerService) UpdateDB() error {
	return nil
}
