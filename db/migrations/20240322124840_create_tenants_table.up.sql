CREATE TABLE tenants (
    tnt_id UUID PRIMARY KEY,
    tnt_name VARCHAR(255) NOT NULL,
    tnt_type VARCHAR(50) NOT NULL ,  --'mart', 'store' enum in golang
    tnt_description TEXT,
    tnt_contact_person_name VARCHAR(255),
    tnt_contact_email VARCHAR(255),
    tnt_contact_phone VARCHAR(20)[],
    tnt_address VARCHAR(255),
    tnt_registration_no VARCHAR(100),
    tnt_city VARCHAR(100),
    tnt_state VARCHAR(100),
    tnt_country VARCHAR(100),
    tnt_path_to_logo VARCHAR(20),
    tnt_url VARCHAR(20),
    tnt_established_year VARCHAR(20),
    tnt_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tnt_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE tenants IS 'short code is tnt';