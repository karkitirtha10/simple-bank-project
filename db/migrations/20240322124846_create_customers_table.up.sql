 CREATE TABLE customers(
    cust_id UUID PRIMARY KEY DEFAULT(uuid_generate_v4()),
    cust_first_name VARCHAR(255) NOT NULL,
    cust_last_name VARCHAR(255) NOT NULL,
    cust_address VARCHAR(255) ,
    cust_country VARCHAR(255) ,
    cust_phone_no VARCHAR(255) ,
    cust_email VARCHAR(255) ,
    cust_dob TIMESTAMP ,
    
    cust_created_by UUID,
    cust_created_at TIMESTAMPTZ DEFAULT(NOW()),
    cust_updated_at TIMESTAMPTZ 
 );

ALTER TABLE customers
ADD CONSTRAINT fk_cust_created_by
FOREIGN KEY (cust_created_by) REFERENCES users(u_id);

COMMENT ON TABLE customers IS 'short code is cust.';