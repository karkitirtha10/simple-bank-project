 CREATE TABLE oauth_access_tokens(
    oat_id UUID PRIMARY KEY,
    oat_refresh_token_id UUID NOT NULL,
    oat_client_id UUID NOT NULL,
    oat_user_id UUID, --nullable
    oat_name VARCHAR(255),
    oat_scopes TEXT[],
    oat_revoked BOOLEAN NOT NULL DEFAULT(FALSE),
    oat_expires_at timestamp without time zone, --nullable
--     oat_audience VARCHAR(255) NOT NULL,
--     oat_created_by UUID, //is  needed ?
    oat_created_at TIMESTAMPTZ NOT NULL DEFAULT(NOW())
 );  

-- COMMENT ON COLUMN oauth_access_tokens.oat_client_id IS 'client the token is generated for';

ALTER TABLE oauth_access_tokens
ADD CONSTRAINT fk_oat_client_id
FOREIGN KEY(oat_client_id) REFERENCES oauth_clients (oc_id);

ALTER TABLE  oauth_access_tokens
ADD CONSTRAINT fk_oat_refresh_token_id
FOREIGN KEY (oat_refresh_token_id) REFERENCES oauth_refresh_tokens(ort_id);

 ALTER TABLE oauth_access_tokens
 ADD CONSTRAINT fk_oat_user_id
 FOREIGN KEY (oat_user_id) REFERENCES users(u_id);

-- ALTER TABLE oauth_access_tokens
-- ADD CONSTRAINT fk_oat_created_by
-- FOREIGN KEY (oat_created_by) REFERENCES users(u_id);

-- CREATE INDEX entries_index_1 ON entries (account_id);

      
    --    $this->schema->create('oauth_access_tokens', function (Blueprint $table) {
    --         $table->string('id', 100)->primary();
    --         $table->unsignedBigInteger('user_id')->nullable()->index();
    --         $table->unsignedBigInteger('client_id');
    --         $table->string('name')->nullable();
    --         $table->text('scopes')->nullable();
    --         $table->boolean('revoked');
    --         $table->timestamps();
    --         $table->dateTime('expires_at')->nullable();
    --     });