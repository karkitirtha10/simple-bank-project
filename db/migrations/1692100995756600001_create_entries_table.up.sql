CREATE TABLE entries(
    en_id UUID PRIMARY KEY DEFAULT(uuid_generate_v4()),
    en_account_id BIGINT,/* NOT NULL*/
    en_amount NUMERIC(20,4) NOT NULL,
    en_e_created_by UUID ,
    en_created_at TIMESTAMPTZ NOT NULL DEFAULT(NOW())
);
COMMENT ON COLUMN entries.amount IS 'can be negative or positive';

ALTER TABLE entries
ADD CONSTRAINT fk_en_account_id_created_at
FOREIGN KEY(en_account_id) REFERENCES accounts (ac_id);

ALTER TABLE entries
ADD CONSTRAINT fk_en_created_at
FOREIGN KEY (en_created_at) REFERENCES users(u_id);

CREATE INDEX entries_index_1 ON entries (en_account_id);