CREATE TABLE permissions (
    id UUID PRIMARY KEY,
    name VARCHAR(400) NOT NULL,
    category VARCHAR(255) NOT NULL,  -- users, product model names
    description TEXT,
    -- is_system BOOLEAN NOT NULL DEFAULT (FALSE),
    -- tenant_id  UUID, --fk
    -- created_by UUID,  --fk
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE permissions
ADD CONSTRAINT permissions_name_unique 
UNIQUE (name);