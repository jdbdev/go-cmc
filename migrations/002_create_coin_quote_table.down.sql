-- Migration: create_coin_quote_table (rollback)
-- Description: Drops the coin_quote table and its indexes

DROP INDEX IF EXISTS idx_coin_quote_last_updated;
DROP INDEX IF EXISTS idx_coin_quote_coin_id;
DROP TABLE IF EXISTS coin_quote;

