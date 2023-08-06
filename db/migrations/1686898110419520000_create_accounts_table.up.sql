CREATE TABLE accounts(
    id BIGSERIAL PRIMARY KEY,
    owner VARCHAR(255) NOT NULL, 
    balance NUMERIC(20,4) NOT NULL,
    currency VARCHAR(255) NOT NULL, 
    created_at TIMESTAMPTZ NOT NULL DEFAULT(NOW()) 
);

CREATE INDEX account_index_0 ON accounts (owner);