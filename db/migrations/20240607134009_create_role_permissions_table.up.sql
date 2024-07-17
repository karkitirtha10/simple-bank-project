CREATE TABLE role_permissions (
    id UUID PRIMARY KEY,
    role_id  UUID NOT NULL, --fk
    permission_id UUID NOT NULL, --fk

    created_by UUID,  --fk
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--foreign keys
ALTER TABLE role_permissions
ADD CONSTRAINT fk_role_permissions_role_id
FOREIGN KEY (role_id) REFERENCES roles (id);

ALTER TABLE role_permissions
ADD CONSTRAINT fk_role_permissions_permission_id
FOREIGN KEY (permission_id) REFERENCES permissions (id);

ALTER TABLE role_permissions
ADD CONSTRAINT fk_role_permissions_created_by
FOREIGN KEY (created_by) REFERENCES users (u_id);

--unique index
ALTER TABLE role_permissions 
ADD CONSTRAINT role_permissions_role_id_permission_id_unique 
UNIQUE (role_id, permission_id);