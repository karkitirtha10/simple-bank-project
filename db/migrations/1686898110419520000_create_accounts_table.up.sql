CREATE TABLE accounts(
    ac_id UUID PRIMARY KEY DEFAULT(uuid_generate_v4()),
    -- ac_owner VARCHAR(255) NOT NULL, 
    ac_balance NUMERIC(20,4) NOT NULL,
    ac_currency VARCHAR(255) NOT NULL, 
    ac_created_by UUID,
    ac_status SMALLINT DEFAULT(1), --opened closed
    ac_type VARCHAR(255) NOT NULL,  
    ac_created_at TIMESTAMPTZ DEFAULT(NOW()) 
);

COMMENT ON TABLE accounts IS 'short code is ac';
COMMENT ON COLUMN accounts.ac_type IS 'ac_type values = saving,current. actually is a enum';

ALTER TABLE accounts
ADD CONSTRAINT fk_ac_created_by
FOREIGN KEY (ac_created_by) REFERENCES users(u_id);

--CREATE INDEX account_index_0 ON accounts (ac_owner);