package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/jdbdev/go-cmc/config"
)

// Coinmarketcap recommends using their ID's instead of symbols, slugs, etc.
// Coinmarketcap API endpoint: https://pro-api.coinmarketcap.com/v1/cryptocurrency/map

var client = &http.Client{}

type IDMapInterface interface {
	FetchIDMap() (rBody []byte, err error)
	WriteIDMapToFile() error
}

type IDMapService struct {
	apiKey string
	mapURL string
	client *http.Client
}

// NewIDMapService creates a new instance of IDMapService struct
func NewIDMapService(app *config.AppConfig) *IDMapService {
	return &IDMapService{
		apiKey: app.CMC.APIKey,
		mapURL: app.CMC.IDMapURL,
		client: client,
	}

}

// FetchIDMap fetches CMC ID's for all tokens from Coinmarketcap's dedicated map endpoint.
// Set return limit in queries below ("limit": "5")
func (i *IDMapService) FetchIDMap() (rBody []byte, err error) {
	// Build Request
	client := i.client
	req, err := http.NewRequest("GET", i.mapURL, nil)
	if err != nil {
		return nil, err
	}
	// Set headers & Build queries
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CMC_PRO_API_KEY", i.apiKey)

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5")
	q.Add("sort", "cmc_rank") //must be either "cmc_rank" or "id" for endpoint .../map

	req.URL.RawQuery = q.Encode()

	// Execute request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle Response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("CMC Map Response Body: \n\n%s\n\n", string(respBody))
	return respBody, nil
}

func WriteIDMapToFile(body []byte) error {
	return nil
}
