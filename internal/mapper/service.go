package mapper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jdbdev/go-cmc/config"
)

var client = &http.Client{}

const limit = 10 // Set limit of tokens to fetch. 10 will fetch top ten tokens by CMC rank

// Fallback ID map for top tokens used if DB or Coinmarketcap map API fails
var fallbackIDMap = map[string]string{
	"BTC": "1",
	"ETH": "1027",
	"SOL": "5994",
}

type IDMapService struct {
	apiKey string
	mapURL string
	client *http.Client
	idMap  map[string]string
}

// NewIDMapService creates a new instance of IDMapService struct
func NewIDMapService(app *config.AppConfig) *IDMapService {
	return &IDMapService{
		apiKey: app.CMC.APIKey,
		mapURL: app.CMC.IDMapURL,
		client: client,
		idMap:  make(map[string]string),
	}
}

// FetchIDMap fetches CMC ID's for all tokens from Coinmarketcap's map endpoint
func (i *IDMapService) FetchIDMap() (*CmcIdMapResponse, error) {
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
	q.Add("limit", strconv.Itoa(limit))
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

	// Create CmcIdMapResponse instance in stack memory
	response := CmcIdMapResponse{}

	// Unmarshal response into CmcIdMapResponse struct instance
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}

	fmt.Printf("CMC Map Response Body: \n\n%s\n\n", string(respBody))
	return &response, nil
}

// Initialize loads the IDMapService.idMap with fallbacks following this order;
// 1. Load from DB
// 2. Fetch from API
// 3. Fallback to hard coded map
func (i *IDMapService) Initialize() error {
	// Load from DB
	// If DB fails, load from CMC API
	if resp, err := i.FetchIDMap(); err == nil {
		for _, coin := range resp.Data {
			i.idMap[coin.Symbol] = strconv.Itoa(coin.ID)
		}
		return nil
	}
	// If DB and API call fail, load fallback ID map
	i.idMap = fallbackIDMap
	return nil
}

func (i *IDMapService) GetIDMap() map[string]string {
	return i.idMap
}

// UpdateDB updates the database with the latest CMC ID mappings
func (i *IDMapService) UpdateDB() error {
	return nil
}
