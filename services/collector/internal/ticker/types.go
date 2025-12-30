package ticker

// All JSON fields that can be null in CMC API response are pointers allowing null values to avoid
// unmarshalling errors or setting zero values instead of nil.
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
	CmcID             int                  `json:"id"` // CMC ID is recommended by CMC API documentation
	Name              string               `json:"name"`
	Symbol            string               `json:"symbol"`
	Slug              string               `json:"slug"`
	CirculatingSupply float64              `json:"circulating_supply"`
	TotalSupply       float64              `json:"total_supply"`
	LastUpdated       string               `json:"last_updated"`
	Quote             map[string]CoinQuote `json:"quote"` // map key is "USD" from CMC API Response
}

// CoinQuote holds the quote data for a coin from CMC API. Only uses USD for map key in CoinInfo.
type CoinQuote struct {
	Price                 float64 `json:"price"`
	MarketCap             float64 `json:"market_cap"`
	FullyDilutedMarketCap float64 `json:"fully_diluted_market_cap"`
	Volume24H             float64 `json:"volume_24h"`
	PercentChange1H       float64 `json:"percent_change_1h"`
	PercentChange24h      float64 `json:"percent_change_24h"`
	PercentChange7d       float64 `json:"percent_change_7d"`
	LastUpdated           string  `json:"last_updated"`
}
