package dbmodel

import (
	"database/sql"
	"time"
)

// Account db
type Role struct {
	Id          string         `db:"id"`
	Name        string         `db:"name"`
	DisplayName string         `db:"display_name"`
	Description sql.NullString `db:"description"`
	IsSystem    bool           `db:"is_system"`
	TenantId    sql.NullString `db:"tenant_id"`
	CreatedBy   sql.NullString `db:"created_by"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
}

// Balance  float32 `json:"balance"`
