CREATE TABLE roles (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    display_name VARCHAR(50) NOT NULL,  
    description TEXT,
    is_system BOOLEAN NOT NULL DEFAULT (FALSE),
    tenant_id  UUID, --fk
    created_by UUID,  --fk
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE roles
ADD CONSTRAINT fk_roles_tenant_id
FOREIGN KEY (tenant_id) REFERENCES tenants (tnt_id);


ALTER TABLE roles
ADD CONSTRAINT fk_roles_created_by
FOREIGN KEY (created_by) REFERENCES users (u_id);


ALTER TABLE roles 
ADD CONSTRAINT roles_tenant_id_name_unique 
UNIQUE (tenant_id, name);

-- COMMENT ON TABLE tenant_groups IS 'short code is ro';
COMMENT ON COLUMN roles.is_system IS 'false means custom role created by users. custom roles are editable oppose to system';
--unique : ro_tenant_id, ro_name

-- SUPERADMIN *superadmin
-- TENANT GROUP ADMIN *group admin
-- TENANT ADMIN *admin