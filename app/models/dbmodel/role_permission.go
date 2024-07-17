package dbmodel

import (
	"database/sql"
	"time"
)

// Account db
type RolePermission struct {
	Id           string `db:"id"`
	RoleId       string `db:"role_id"`
	PermissionId string `db:"permission_id"`

	CreatedBy sql.NullString `db:"created_by"`
	CreatedAt time.Time      `db:"created_at"`
}

// Balance  float32 `json:"balance"`
