CREATE TABLE oauth_clients(
    oc_id UUID PRIMARY KEY DEFAULT(uuid_generate_v4()),
    oc_name VARCHAR(255) NOT NULL,
    oc_secret VARCHAR(100) NOT NULL, 
    oc_type VARCHAR(255) NOT NULL,
    oc_revoked BOOLEAN NOT NULL DEFAULT(FALSE),
    oc_created_at TIMESTAMPTZ NOT NULL DEFAULT(NOW())

    -- balance NUMERIC(20,4) NOT NULL,
    -- currency VARCHAR(255) NOT NULL, 
    -- created_at TIMESTAMPTZ NOT NULL DEFAULT(NOW()) 
);

ALTER TABLE oauth_clients ADD CONSTRAINT oc_secret_unique UNIQUE (oc_secret);

-- CREATE INDEX account_index_0 ON accounts (owner);

--     $this->schema->create('oauth_clients', function (Blueprint $table) {
--             $table->bigIncrements('id');
--             $table->unsignedBigInteger('user_id')->nullable()->index();
--             $table->string('name');
--             $table->string('secret', 100)->nullable();
--             $table->string('provider')->nullable();
--             $table->text('redirect');
--             $table->boolean('personal_access_client');
--             $table->boolean('password_client');
--             $table->boolean('revoked');
--             $table->timestamps();
--         });