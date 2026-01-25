package ticker

// REFERENCE FILE - NOT USED IN CODE
// This file contains a "medium" level of struct definitions with most fields.
// It is kept as reference only for later inclusion when needed.
// The actual structs used in the codebase are in types.go (minimal version).

// NOTE: Structs use "Ref" suffix to avoid conflicts. Change this in types.go if needed.
// This file is for reference only and won't be used in compilation.

// CMCResponseRef holds the response from the CMC API (reference version with all fields).
type CMCResponseRef struct {
	Status StatusRef              `json:"status"`
	Data   map[string]CoinInfoRef `json:"data"`
}

// StatusRef holds the response status from CMC API (reference version).
type StatusRef struct {
	Timestamp    string  `json:"timestamp"`
	ErrorCode    int     `json:"error_code"`
	ErrorMessage *string `json:"error_message"`
	Elapsed      int     `json:"elapsed"`
	CreditCount  int     `json:"credit_count"`
	Notice       *string `json:"notice"`
}

// CoinInfoRef holds the coin related information from CMC API, including CoinQuote data (reference version).
type CoinInfoRef struct {
	ID                            int                     `json:"id"`
	Name                          string                  `json:"name"`
	Symbol                        string                  `json:"symbol"`
	Slug                          string                  `json:"slug"`
	CirculatingSupply             float64                 `json:"circulating_supply"`
	TotalSupply                   float64                 `json:"total_supply"`
	InfiniteSupply                bool                    `json:"infinite_supply"`
	SelfReportedCirculatingSupply *float64                `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         *float64                `json:"self_reported_market_cap"`
	TvlRatio                      *float64                `json:"tvl_ratio"`
	LastUpdated                   string                  `json:"last_updated"`
	Quote                         map[string]CoinQuoteRef `json:"quote"` // map key is "USD" from CMC API Response
}

// CoinQuoteRef holds the quote data for a coin from CMC API (reference version with all fields).
type CoinQuoteRef struct {
	Price                 float64  `json:"price"`
	MarketCap             float64  `json:"market_cap"`
	FullyDilutedMarketCap float64  `json:"fully_diluted_market_cap"`
	Volume24H             float64  `json:"volume_24h"`
	Volume24HReported     float64  `json:"volume_24h_reported"`
	VolumeChange24H       float64  `json:"volume_change_24h"`
	PercentChange1H       float64  `json:"percent_change_1h"`
	PercentChange24h      float64  `json:"percent_change_24h"`
	PercentChange7d       float64  `json:"percent_change_7d"`
	PercentChange30d      float64  `json:"percent_change_30d"`
	PercentChange60d      float64  `json:"percent_change_60d"`
	PercentChange90d      float64  `json:"percent_change_90d"`
	Tvl                   *float64 `json:"tvl"`
	LastUpdated           string   `json:"last_updated"`
}
