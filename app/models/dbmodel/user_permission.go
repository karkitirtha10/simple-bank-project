package dbmodel

import (
	"database/sql"
	"time"
)

// Account db
type UserPermission struct {
	Id           string `db:"id"`
	UserId       string `db:"user_id"`
	PermissionId string `db:"permission_id"`

	CreatedBy sql.NullString `db:"created_by"`
	CreatedAt time.Time      `db:"created_at"`
}

// Balance  float32 `json:"balance"`
