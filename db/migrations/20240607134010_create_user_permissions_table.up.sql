CREATE TABLE user_permissions (
    id UUID PRIMARY KEY,
    user_id  UUID NOT NULL, --fk
    permission_id UUID NOT NULL, --fk

    created_by UUID,  --fk
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--foreign keys
ALTER TABLE user_permissions
ADD CONSTRAINT fk_user_permissions_user_id
FOREIGN KEY (user_id) REFERENCES users (u_id);

ALTER TABLE user_permissions
ADD CONSTRAINT fk_user_permissions_permission_id
FOREIGN KEY (permission_id) REFERENCES permissions (id);

ALTER TABLE user_permissions
ADD CONSTRAINT fk_user_permissions_created_by
FOREIGN KEY (created_by) REFERENCES users (u_id);

--unique index
ALTER TABLE user_permissions 
ADD CONSTRAINT user_permissions_user_id_permission_id_unique 
UNIQUE (permission_id, user_id);