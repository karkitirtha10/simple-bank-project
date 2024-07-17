CREATE TABLE role_users  (
    id UUID PRIMARY KEY,
    role_id  UUID NOT NULL, --fk
    user_id UUID NOT NULL, --fk

    created_by UUID,  --fk
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--foreign keys
ALTER TABLE role_users
ADD CONSTRAINT fk_role_users_role_id
FOREIGN KEY (role_id) REFERENCES roles (id);

ALTER TABLE role_users
ADD CONSTRAINT fk_role_users_user_id
FOREIGN KEY (user_id) REFERENCES users (u_id);

ALTER TABLE role_users
ADD CONSTRAINT fk_role_users_created_by
FOREIGN KEY (created_by) REFERENCES users (u_id);

--unique index
ALTER TABLE role_users 
ADD CONSTRAINT role_users_role_id_user_id_unique 
UNIQUE (role_id, user_id);

