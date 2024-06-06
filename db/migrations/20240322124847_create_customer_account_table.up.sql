 CREATE TABLE customer_account(
    cac_id UUID PRIMARY KEY DEFAULT(uuid_generate_v4()),
    cac_customer_id UUID NOT NULL,
    cac_account_id UUID NOT NULL
 );    

ALTER TABLE customer_account
ADD CONSTRAINT fk_cac_customer_id
FOREIGN KEY (cac_customer_id) REFERENCES customers(cust_id);

ALTER TABLE customer_account
ADD CONSTRAINT fk_cac_account_id
FOREIGN KEY (cac_account_id) REFERENCES accounts(ac_id);

ALTER TABLE customer_account
ADD CONSTRAINT cac_customer_id_cac_account_id_unique 
UNIQUE(cac_customer_id, cac_account_id);


COMMENT ON TABLE customer_account IS 'links customer with accountcontroller with many to many relationship. multiple customer can own a single accountcontroller .short code is cac.';