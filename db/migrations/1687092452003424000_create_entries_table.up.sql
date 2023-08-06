CREATE TABLE entries(
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT,/* NOT NULL*/
    amount NUMERIC(20,4) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT(NOW())     
);
COMMENT ON COLUMN entries.amount IS 'can be negative or positive';

ALTER TABLE entries
ADD FOREIGN KEY(account_id) REFERENCES accounts (id);

CREATE INDEX entries_index_1 ON entries (account_id);