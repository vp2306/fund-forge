-- what user wants to do 
CREATE TABLE trades (
    id BIGSERIAL PRIMARY KEY,
    etf_id BIGINT NOT NULL REFERENCES etfs(id),
    trade_type VARCHAR(4) NOT NULL CHECK (trade_type IN ('BUY', 'SELL')),
    amount DECIMAL(12,2) NOT NULL CHECK (amount > 0),
    status VARCHAR(10) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP DEFAULT NOW(),
    executed_at TIMESTAMP
);

-- what got bought/sold
CREATE TABLE trade_executions (
    id BIGSERIAL PRIMARY KEY,
    trade_id BIGINT NOT NULL REFERENCES trades(id) ON DELETE CASCADE,
    ticker VARCHAR(10) NOT NULL,
    shares DECIMAL(10,4) NOT NULL, -- Fractional shares
    price_per_share DECIMAL(8,2) NOT NULL,
    total_cost DECIMAL(12,2) NOT NULL
);

-- users position after trade execution 
CREATE TABLE positions (
    id BIGSERIAL PRIMARY KEY,
    ticker VARCHAR(10) NOT NULL UNIQUE,
    total_shares DECIMAL(10,4) NOT NULL DEFAULT 0,
    average_cost DECIMAL(8,2) NOT NULL DEFAULT 0
);