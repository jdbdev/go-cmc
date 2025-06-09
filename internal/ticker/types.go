package ticker

// CoinInfo holds static data and time of last CMC update
type CoinInfo struct {
	ID          int                   `json:"id"`           // static
	CMCID       int                   `json:"cmc_rank"`     // static
	Name        string                `json:"name"`         // static
	Symbol      string                `json:"symbol"`       // static
	Slug        string                `json:"website_slug"` // static
	Quotes      map[string]*CoinQuote `json:"quotes"`       // dynamic
	LastUpdated int                   `json:"last_updated"` // dynamic
}

// CoinSupply holds data related to supply
type CoinSupply struct {
	CirculatingSupply float64 `json:"circulating_supply"` // dynamic
	TotalSupply       float64 `json:"total_supply"`       // dynamic
	MaxSupply         float64 `json:"max_supply"`         // dynamic
}

// CoinQuote holds dynamic price and volume data
type CoinQuote struct {
	Price            float64 `json:"price"`
	Volume24H        float64 `json:"volume_24h"`
	MarketCap        float64 `json:"market_cap"`
	PercentChange1H  float64 `json:"percent_change_1h"`
	PercentChange24H float64 `json:"percent_change_24h"`
	PercentChange7D  float64 `json:"percent_change_7d"`
}
