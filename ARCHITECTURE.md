# Collector Service Architecture

## Data Flow Diagram

### Initialization Flow
```
Mapper Service
    ↓ GetCMCTopCoins(100)
CoinMarketCap API
    ↓ Returns JSON
Unmarshal → CmcIdMapResponse
    ↓ Loop through coins
Coins Service
    ↓ AddTrackedCoin()
tracked_coins table
```

### Runtime Loop (Every X seconds)
```
tracked_coins table
    ↓ GetTrackedCoinIDs()
Ticker Service
    ↓ Build query with CMC IDs
CoinMarketCap API
    ↓ Returns price data
Unmarshal → CMCResponse
    ↓ Save data
coin_info table (using cmc_id from tracked_coins)
coin_quote table (using coin_info.id)
```

## Table Relationships

```
tracked_coins (Source of Truth)
├── id (PK)
├── cmc_id (CMC ID - links to coin_info)
├── symbol
├── name
├── enabled
└── created_at

coin_info (Coin Data)
├── id (PK)
├── cmc_id (UNIQUE - matches tracked_coins.cmc_id)
├── name
├── symbol
├── slug
├── circulating_supply
└── total_supply

coin_quote (Price Data)
├── id (PK)
├── coin_id (FK → coin_info.id)
├── price
├── market_cap
├── volume_24h
└── percent_change_*
```

## Service Responsibilities

### Mapper Service
- **Purpose:** Helper tool for looking up CMC IDs
- **Methods:**
  - `GetCMCID(symbol)` - Look up single coin ID
  - `GetCMCTopCoins(limit)` - Get top N coins
  - `UnmarshalCMCID(body)` - Parse API response
- **No database operations**

### Coins Service
- **Purpose:** Manages tracked_coins table
- **Methods:**
  - `AddTrackedCoin(cmcID, symbol)` - Add coin to tracking
  - `GetTrackedCoinIDs()` - Get list of CMC IDs to fetch
  - `InitializeCoinTable()` - Setup table
- **Source of truth for which coins to track**

### Ticker Service
- **Purpose:** Fetches and stores coin data
- **Methods:**
  - `FetchAndDecodeData()` - Get data from API
  - `UpdateDB()` - Save to database
- **Flow:**
  1. Read CMC IDs from tracked_coins (via coins service)
  2. Fetch data from CoinMarketCap API
  3. Save to coin_info and coin_quote tables

## Data Flow Summary

1. **Initialization:**
   - Mapper gets top 100 coins from API
   - Coins service adds them to tracked_coins table

2. **Runtime (every X seconds):**
   - Ticker reads CMC IDs from tracked_coins
   - Ticker fetches data for those IDs
   - Ticker saves to coin_info and coin_quote

3. **Website Service:**
   - Users add coins → Coins service → tracked_coins
   - Website reads from coin_info and coin_quote (no API calls)

## Key Points

- **tracked_coins** is the source of truth for which coins to track
- **coin_info** and **coin_quote** are linked by `coin_info.id` (FK)
- **tracked_coins** and **coin_info** are linked by `cmc_id` (no FK, just join)
- Ticker always uses CMC IDs from tracked_coins to ensure consistency

