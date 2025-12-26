package ticker

// All JSON fields that can be null in CMC response are pointers allowing null values to avoid
//unmarshalling errors or setting zero values instead of nil.
// Always check documentation when adding new fields.
// Always check for nil if trying to dereference a pointer to avoid runtime errors (panic).

// CMCResponse holds the response from the CMC API.
type CMCResponse struct {
	Status Status              `json:"status"`
	Data   map[string]CoinInfo `json:"data"`
}

// Status holds the response status from CMC API.
type Status struct {
	Timestamp    string  `json:"timestamp"`
	ErrorCode    int     `json:"error_code"`
	ErrorMessage *string `json:"error_message"`
	Elapsed      int     `json:"elapsed"`
	CreditCount  int     `json:"credit_count"`
	Notice       *string `json:"notice"`
}

// CoinInfo holds the coin related information from CMC API, including CoinQuote data.
type CoinInfo struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Symbol            string  `json:"symbol"`
	Slug              string  `json:"slug"`
	CirculatingSupply float64 `json:"circulating_supply"`
	TotalSupply       float64 `json:"total_supply"`
	InfiniteSupply    bool    `json:"infinite_supply"`
	MarketCap         float64 `json:"market_cap"`
	// SelfReportedCirculatingSupply *float64 `json:"self_reported_circulating_supply"`
	// SelfReportedMarketCap *float64 `json:"self_reported_market_cap"`
	// TvlRatio *float64 `json:"tvl_ratio"`
	// Tvl *float64 `json:"tvl"`
	PercentChange24h float64              `json:"percent_change_24h"`
	PercentChange7d  float64              `json:"percent_change_7d"`
	PercentChange30d float64              `json:"percent_change_30d"`
	PercentChange60d float64              `json:"percent_change_60d"`
	PercentChange90d float64              `json:"percent_change_90d"`
	LastUpdated      string               `json:"last_updated"`
	Quote            map[string]CoinQuote `json:"quote"`
}

// CoinQuote holds the quote data for a coin from CMC API. Only uses USD for map key in CoinInfo.
type CoinQuote struct {
	Price             float64 `json:"price"`
	Volume24H         float64 `json:"volume_24h"`
	Volume24HReported float64 `json:"volume_24h_reported"`
	VolumeChange24H   float64 `json:"volume_change_24h"`
	PercentChange1H   float64 `json:"percent_change_1h"`
}
