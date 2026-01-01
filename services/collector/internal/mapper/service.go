package mapper

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jdbdev/go-cmc/config"
)

// Mapper service provides utilities to get CMC ID's for coins and unmarshal the response for use in other services.
// Mapper only gets data from Coinmarketcap API. DB Updates are handled by internal/coins and internal/ticker services.

// IDMapInterface defines the contract for CMC ID mapping operations
type IDMapInterface interface {
	GetCMCID(symbol string) ([]byte, error)
	GetCMCTopCoins(limit int) ([]byte, error)
	UnmarshalCMCID(body []byte, client *http.Client)
}

// IDMapService implements the IDMapInterface
type IDMapService struct {
	apiKey string
	mapURL string
	client *http.Client
	logger *slog.Logger
}

// NewIDMapService creates a new instance of IDMapService struct
func NewIDMapService(app *config.AppConfig, logger *slog.Logger, client *http.Client) *IDMapService {
	return &IDMapService{
		apiKey: app.CMC.APIKey,
		mapURL: app.CMC.IDMapURL,
		client: client,
		logger: logger,
	}
}

// LookupCMCID builds the request to look up the corresponding Coinmarketcap ID for a given symbol (ex. ETH -> 1027)
func (i *IDMapService) GetCMCID(symbol string) ([]byte, error) {
	i.logger.Info("Looking up Coinmarketcap ID for:", "symbol", symbol)

	// Build request
	req, err := http.NewRequest("GET", i.mapURL, nil)
	if err != nil {
		return nil, err
	}

	// Set headers & Build query parameters
	q := url.Values{}
	q.Add("symbol", symbol)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CMC_PRO_API_KEY", i.apiKey)

	// Execute the request
	resp, err := i.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Just for testing purposes
	fmt.Println(string(body))

	return body, nil
}

// GetCMCTopCoins gets a set of top coins based on limit parameter (top 10, top 50, etc.)
func (i *IDMapService) GetCMCTopCoins(limit int) ([]byte, error) {
	i.logger.Info("getting top coins for:", "limit", limit)
	// Build request
	req, err := http.NewRequest("GET", i.mapURL, nil)
	if err != nil {
		return nil, err
	}

	// Set headers & Build query parameters
	q := url.Values{}
	q.Add("limit", strconv.Itoa(limit))
	q.Add("sort", "cmc_rank")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CMC_PRO_API_KEY", i.apiKey)

	// Execute the request
	resp, err := i.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// FOR TESTING ONLY. REMOVE WHEN DONE.
	fmt.Println(string(body))

	return body, nil
}

// UnmarshallCMCID unmarshalls the response body into CmcIdResponse struct (symbol -> CMCID)
func (i *IDMapService) UnmarshalCMCID(body []byte, client *http.Client) {}
