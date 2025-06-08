package utils

import (
	"net/http"

	"github.com/jdbdev/go-cmc/config"
)

// Coinmarketcap recommends using their ID's instead of symbols, slugs, etc.
// Coinmarketcap API endpoint: https://pro-api.coinmarketcap.com/v1/cryptocurrency/map

var client = &http.Client{}

type IDMapInterface interface {
	GetCMCIDMap() error
	WriteCMCIDMapToFile() error
}

type IDMapService struct {
	apiKey string
	mapURL string
	client *http.Client
}

func NewIDMapService(app *config.AppConfig) *IDMapService {
	return &IDMapService{
		apiKey: app.CMC.APIKey,
		mapURL: app.CMC.IDMapURL,
		client: client,
	}

}

// GetIDMap fetches CMC ID's for all tokens
func GetIDMap(i *IDMapService) error {
	return nil
}

func WriteIDMapToFile() error {
	return nil
}
