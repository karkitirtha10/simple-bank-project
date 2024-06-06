CREATE TABLE tenant_groups (
    tntg_id UUID PRIMARY KEY,
    tntg_name VARCHAR(255) NOT NULL,
    tntg_type VARCHAR(50) NOT NULL,  --'mart', 'store' enum in golang
    tntg_description TEXT,
    tntg_contact_person_name VARCHAR(255),
    tntg_contact_email VARCHAR(255),
    tntg_contact_phone VARCHAR(20),
    tntg_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tntg_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE tenant_groups IS 'short code is tntg';