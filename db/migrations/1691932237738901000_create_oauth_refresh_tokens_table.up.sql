 CREATE TABLE oauth_refresh_tokens(
     UUID PRIMARY KEY DEFAULT(uuid_generate_v4()),
    ort_client_id UUID NOT NULL,
    ort_user_id UUID, --nullable
--     ort_access_token_id UUID NOT NULL,
    ort_expires_at timestamp without time zone, --nullable
    ort_revoked BOOLEAN NOT NULL DEFAULT(FALSE),
--     ort_created_by UUID NOT NULL,
    ort_created_at TIMESTAMPTZ NOT NULL DEFAULT(NOW())
);    

-- COMMENT ON COLUMN oauth_access_tokens.oat_client_id IS 'client the token is generated for';

ALTER TABLE oauth_refresh_tokens
ADD CONSTRAINT fk_ort_client_id
FOREIGN KEY(ort_client_id) REFERENCES oauth_clients (oc_id);

ALTER TABLE oauth_refresh_tokens
ADD CONSTRAINT fk_ort_user_id
FOREIGN KEY (ort_user_id) REFERENCES users(u_id);

-- ALTER TABLE oauth_refresh_tokens
-- ADD CONSTRAINT fk_ort_access_token_id
-- FOREIGN KEY (ort_access_token_id) REFERENCES oauth_access_tokens(oat_id);

    -- $this->schema->create('oauth_refresh_tokens', function (Blueprint $table) {
    --         $table->string('id', 100)->primary();
    --         $table->string('access_token_id', 100)->index();
    --         $table->boolean('revoked');
    --         $table->dateTime('expires_at')->nullable();