package data

// Coin struct
type Coin struct {
	ID                int                   `json:"id"`
	Name              string                `json:"name"`
	Symbol            string                `json:"symbol"`
	Slug              string                `json:"website_slug"`
	Rank              int                   `json:"rank"`
	CirculatingSupply float64               `json:"circulating_supply"`
	TotalSupply       float64               `json:"total_supply"`
	MaxSupply         float64               `json:"max_supply"`
	Quotes            map[string]*CoinQuote `json:"quotes"`
	LastUpdated       int                   `json:"last_updated"`
}

// CoinQuote struct
type CoinQuote struct {
	Price            float64 `json:"price"`
	Volume24H        float64 `json:"volume_24h"`
	MarketCap        float64 `json:"market_cap"`
	PercentChange1H  float64 `json:"percent_change_1h"`
	PercentChange24H float64 `json:"percent_change_24h"`
	PercentChange7D  float64 `json:"percent_change_7d"`
}
