-- Migration: create_coin_quote_table
-- Description: Creates the coin_quote table to store dynamic price and market data from Coinmarketcap API
-- Maps to: ticker.CoinQuote struct
-- Note: One quote per coin (latest). Remove UNIQUE constraint if you want historical data.

CREATE TABLE IF NOT EXISTS coin_quote (
    id SERIAL PRIMARY KEY,
    coin_id INT NOT NULL REFERENCES coin_info(id) ON DELETE CASCADE,
    price NUMERIC(20, 8) NOT NULL,
    market_cap NUMERIC(20, 2),
    fully_diluted_market_cap NUMERIC(20, 2),
    volume_24h NUMERIC(20, 2),
    percent_change_1h NUMERIC(10, 4),
    percent_change_24h NUMERIC(10, 4),
    percent_change_7d NUMERIC(10, 4),
    last_updated TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- One quote per coin (most recent)
    CONSTRAINT unique_coin_quote UNIQUE(coin_id)
);

-- Indexes for faster lookups and joins
CREATE INDEX IF NOT EXISTS idx_coin_quote_coin_id ON coin_quote(coin_id);
CREATE INDEX IF NOT EXISTS idx_coin_quote_last_updated ON coin_quote(last_updated DESC);

