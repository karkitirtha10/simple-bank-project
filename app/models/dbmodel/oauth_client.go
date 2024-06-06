package dbmodel

import (
	"database/sql"
)

type OAuthClient struct {
	Id        string       `db:"oc_id"`
	Name      string       `db:"oc_name"`
	Secret    string       `db:"oc_secret"`
	Type      string       `db:"oc_type"`
	Revoked   bool         `db:"oc_revoked"`
	CreatedAt sql.NullTime `db:"oc_created_at"`
}

type OAuthClientPart struct {
	Id     string `db:"oc_id"`
	Secret string `db:"oc_secret"`
}
