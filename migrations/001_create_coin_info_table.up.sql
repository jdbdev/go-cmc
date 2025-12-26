-- Migration: create_coin_info_table
-- Description: Creates the coin_info table to store basic coin information from CoinMarketCap API
-- Maps to: ticker.CoinInfo struct

CREATE TABLE IF NOT EXISTS coin_info (
    id SERIAL PRIMARY KEY,
    cmc_id INT NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    circulating_supply NUMERIC(20, 8),
    total_supply NUMERIC(20, 8),
    last_updated TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for faster lookups
CREATE INDEX IF NOT EXISTS idx_coin_info_cmc_id ON coin_info(cmc_id);
CREATE INDEX IF NOT EXISTS idx_coin_info_symbol ON coin_info(symbol);

