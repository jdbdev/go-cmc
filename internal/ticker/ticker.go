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
	"net/http"
	"net/url"

	"github.com/jdbdev/go-cmc/config"
)

var client = &http.Client{}

type TickerInterface interface {
	FetchAndDecodeData() error
	UpdateDB() error
}

type TickerService struct {
	apiKey  string
	baseURL string
	client  *http.Client
	// data    []TickerData // Add a field to store the decoded data
}

// NewTickerService creates a new instance of the TickerService
func NewTickerService(app *config.AppConfig) *TickerService {
	return &TickerService{
		apiKey:  app.CMC.APIKey,
		baseURL: app.CMC.BaseURL,
		client:  client,
	}
}

// FetchAndDecodeData gets and decodes data from CMC
func (t *TickerService) FetchAndDecodeData() error {
	client := t.client
	req, err := http.NewRequest("GET", t.baseURL, nil)
	if err != nil {
		return err
	}
	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", t.apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // Now we can safely close here

	// respBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }

	fmt.Println(resp.Status)

	// TODO: Implement decoding logic here
	return nil
}

// UpdateDB updates the database with data from CMC
func (t *TickerService) UpdateDB() error {
	return nil
}
