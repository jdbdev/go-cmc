-- Migration: create_coin_info_table (rollback)
-- Description: Drops the coin_info table and its indexes

DROP INDEX IF EXISTS idx_coin_info_symbol;
DROP INDEX IF EXISTS idx_coin_info_cmc_id;
DROP TABLE IF EXISTS coin_info;

