package mapper

// IDMapInterface defines the contract for CMC ID mapping operations
type IDMapInterface interface {
	FetchIDMap() (*CmcIdMapResponse, error)
	UpdateDB() error
	Initialize() error
	GetIDMap() map[string]string
}

// CmcIdMapResponse is the struct to store the ID map from Coinmarketcap.
// The CMC endpoint /map returns multiple tokens under the key "data"
type CmcIdMapResponse struct {
	Data []CmcCoinID `json:"data"`
}

// CmcCoinID stores only the required fields for the app
type CmcCoinID struct {
	ID     int    `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
}
