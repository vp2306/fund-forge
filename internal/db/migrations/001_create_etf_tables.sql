-- Schema for ETFs and their holdings
-- Create ETFs table
CREATE TABLE IF NOT EXISTS etfs (
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- Create holdings table
CREATE TABLE IF NOT EXISTS etf_holdings (
    id      BIGSERIAL PRIMARY KEY,
    etf_id  BIGINT NOT NULL REFERENCES etfs(id) ON DELETE CASCADE,
    ticker  TEXT NOT NULL,
    weight  DOUBLE PRECISION NOT NULL CHECK (weight >= 0 AND weight <= 1),
    UNIQUE (etf_id, ticker)
);

-- index for joins/filters
CREATE INDEX IF NOT EXISTS idx_etf_holdings_etf_id ON etf_holdings(etf_id);

