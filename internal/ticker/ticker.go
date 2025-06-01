package ticker

// Coinmarketcap (CMC) API Documentation: https://coinmarketcap.com/api/documentation/v1/

import (
	"net/http"

	"github.com/jdbdev/go-cmc/config"
)

var client = &http.Client{}

type TickerService struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

// NewTickerService creates a new instance of the TickerService
func NewTickerService(app *config.AppConfig) *TickerService {
	return &TickerService{
		apiKey:  app.CMC.APIKey,
		baseURL: app.CMC.BaseURL,
		client:  client,
	}
}

// FetchCMCData gets up to date data from CMC
func (t *TickerService) FetchCMCData() {}

// UpdateDB updates the database with data from CMC
func (t *TickerService) UpdateDB() {}
