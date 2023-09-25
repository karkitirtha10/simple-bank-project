CREATE TABLE users(
    u_id UUID PRIMARY KEY DEFAULT(uuid_generate_v4()),
    u_name VARCHAR(255) NOT NULL,
    u_email VARCHAR(255) NOT NULL,
    u_password VARCHAR(255) NOT NULL,
    u_email_verified_at TIMESTAMP,
    u_active SMALLINT NOT NULL DEFAULT(1),
    u_created_at TIMESTAMPTZ DEFAULT(NOW()),
    u_updated_at TIMESTAMPTZ 
);   

ALTER TABLE users 
ADD CONSTRAINT u_email_unique UNIQUE (u_email);

COMMENT ON TABLE users IS 'short code is u.';

