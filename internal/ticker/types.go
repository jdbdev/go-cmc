package ticker

// CoinInfo holds static data and time of last CMC update
// DB table: coin_info
type CoinInfo struct {
	ID      int                   `json:"id"`         // static
	CmcID   int                   `json:"cmc_id"`     // static
	Name    string                `json:"name"`       // static
	Symbol  string                `json:"symbol"`     // static
	Slug    string                `json:"slug"`       // static
	Created int                   `json:"created_at"` // dynamic
	Updated int                   `json:"updated_at"` // dynamic
	Quotes  map[string]*CoinQuote `json:"quotes"`     // dynamic
}

// CoinSupply holds data related to supply
// DB table: coin_supply
type CoinSupply struct {
	CirculatingSupply float64 `json:"circulating_supply"` // dynamic
	TotalSupply       float64 `json:"total_supply"`       // dynamic
	MaxSupply         float64 `json:"max_supply"`         // dynamic
}

// CoinQuote holds dynamic price and volume data
// DB table: coin_quote
type CoinQuote struct {
	Price            float64 `json:"price"`              // dynamic
	Volume24H        float64 `json:"volume_24h"`         // dynamic
	MarketCap        float64 `json:"market_cap"`         // dynamic
	PercentChange1H  float64 `json:"percent_change_1h"`  // dynamic
	PercentChange24H float64 `json:"percent_change_24h"` // dynamic
	PercentChange7D  float64 `json:"percent_change_7d"`  // dynamic
}
