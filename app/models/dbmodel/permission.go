package dbmodel

import (
	"database/sql"
	"time"
)

// Account db
type Permission struct {
	Id          string         `db:"id"`
	Name        string         `db:"name"`
	Category    string         `db:"category"`
	Description sql.NullString `db:"description"`
	CreatedAt   time.Time      `db:"created_at"`
}

// Balance  float32 `json:"balance"`
