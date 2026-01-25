package ticker

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/jdbdev/go-cmc/config"
// )

// // TestNewTickerService tests the creation of a new ticker service
// func TestNewTickerService(t *testing.T) {
// 	// 1. Setup test data
// 	cfg := &config.AppConfig{
// 		CMC: config.CMCSettings{
// 			APIKey:    "test-key",             // Test value
// 			QuotesURL: "https://test-url.com", // Test value
// 		},
// 	}

// 	// 2. Execute the code we want to test
// 	service := NewTickerService(cfg)

// 	// 3. Assert our expectations
// 	if service.apiKey != "test-key" {
// 		// t.Errorf reports a test failure with a formatted message
// 		t.Errorf("Expected apiKey to be 'test-key', got %s", service.apiKey)
// 	}
// }

// // Version 2: Using subtests for better organization
// func TestNewTickerService_WithSubtests(t *testing.T) {
// 	// Setup once for all subtests
// 	cfg := &config.AppConfig{
// 		CMC: config.CMCSettings{
// 			APIKey:    "test-key",
// 			QuotesURL: "https://test-url.com",
// 		},
// 	}
// 	service := NewTickerService(cfg)

// 	// Subtest for API Key
// 	t.Run("API Key", func(t *testing.T) {
// 		if service.apiKey != "test-key" {
// 			t.Errorf("Expected apiKey to be 'test-key', got %s", service.apiKey)
// 		}
// 	})

// 	// Subtest for Quotes URL
// 	t.Run("Quotes URL", func(t *testing.T) {
// 		if service.quotesURL != "https://test-url.com" {
// 			t.Errorf("Expected quotesURL to be 'https://test-url.com', got %s", service.quotesURL)
// 		}
// 	})
// }

// // Version 3: Using table tests for multiple cases
// func TestNewTickerService_TableDriven(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		config  *config.AppConfig
// 		wantKey string
// 		wantURL string
// 		wantErr bool
// 	}{
// 		{
// 			name: "valid config",
// 			config: &config.AppConfig{
// 				CMC: config.CMCSettings{
// 					APIKey:    "test-key",
// 					QuotesURL: "https://test-url.com",
// 				},
// 			},
// 			wantKey: "test-key",
// 			wantURL: "https://test-url.com",
// 			wantErr: false,
// 		},
// 		{
// 			name: "empty config",
// 			config: &config.AppConfig{
// 				CMC: config.CMCSettings{},
// 			},
// 			wantKey: "",
// 			wantURL: "",
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			service := NewTickerService(tt.config)

// 			if service.apiKey != tt.wantKey {
// 				t.Errorf("apiKey = %v, want %v", service.apiKey, tt.wantKey)
// 			}

// 			if service.quotesURL != tt.wantURL {
// 				t.Errorf("quotesURL = %v, want %v", service.quotesURL, tt.wantURL)
// 			}
// 		})
// 	}
// }

// // TestFetchAndDecodeData tests the API call functionality
// func TestFetchAndDecodeData(t *testing.T) {
// 	// Create a mock server
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Test request headers
// 		if r.Header.Get("X-CMC_PRO_API_KEY") != "test-key" {
// 			t.Errorf("Expected API key header 'test-key', got %s", r.Header.Get("X-CMC_PRO_API_KEY"))
// 		}

// 		// Test query parameters
// 		query := r.URL.Query()
// 		if !query.Has("convert") {
// 			t.Error("Missing 'convert' parameter")
// 		}
// 		if query.Get("convert") != "USD" {
// 			t.Errorf("Expected convert=USD, got %s", query.Get("convert"))
// 		}

// 		// Return a mock response
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{
// 			"status": {
// 				"timestamp": "2024-01-01T00:00:00.000Z",
// 				"error_code": 0,
// 				"error_message": null
// 			},
// 			"data": {
// 				"1": {
// 					"symbol": "BTC",
// 					"quote": {
// 						"USD": {
// 							"price": 50000.00
// 						}
// 					}
// 				}
// 			}
// 		}`))
// 	}))
// 	defer server.Close()

// 	// Create service with mock server URL
// 	cfg := &config.AppConfig{
// 		CMC: config.CMCSettings{
// 			APIKey:    "test-key",
// 			QuotesURL: server.URL,
// 		},
// 	}
// 	service := NewTickerService(cfg)

// 	// Test the API call
// 	err := service.FetchAndDecodeData()
// 	if err != nil {
// 		t.Errorf("Expected no error, got %v", err)
// 	}
// }

// // Example of table-driven test
// func TestCMCIDMap(t *testing.T) {
// 	tests := []struct {
// 		symbol string
// 		id     string
// 	}{
// 		{"BTC", "1"},
// 		{"ETH", "1027"},
// 		{"SOL", "5994"},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.symbol, func(t *testing.T) {
// 			if id := CMCIDMap[tt.symbol]; id != tt.id {
// 				t.Errorf("Expected %s ID to be %s, got %s", tt.symbol, tt.id, id)
// 			}
// 		})
// 	}
// }
