 CREATE TABLE system_error_logs(
    sel_id BIGSERIAL PRIMARY KEY,
    sel_log JSONB NOT NULL,
    sel_created_at TIMESTAMPTZ DEFAULT(NOW())
 );

COMMENT ON TABLE customers IS 'short code is sel.';