package mapper

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/jdbdev/go-cmc/config"
)

// IDMapInterface defines the contract for CMC ID mapping operations
type IDMapInterface interface {
	GetCMCID(symbol string) ([]byte, error)
	UnmarshalCMCID(body []byte, client *http.Client)
	// LookupCMCTop10() (string, error)
	// LookupCMCTop100() (string, error)
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

	// Set headers & Build queries
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

// UnmarshallCMCID unmarshalls the response body into CmcIdResponse struct (symbol -> CMCID)
func (i *IDMapService) UnmarshalCMCID(body []byte, client *http.Client) {}
