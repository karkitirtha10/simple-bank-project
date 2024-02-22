package dbmodel

import (
	"database/sql"
	"time"
)

type OAuthAccessToken struct {
	Id             string         `db:"oat_id"`
	RefreshTokenId string         `db:"oat_refresh_token_id"`
	ClientId       string         `db:"oat_client_id"`
	UserId         sql.NullString `db:"oat_user_id"`
	Name           sql.NullString `db:"oat_name"`
	Scopes         []string       `db:"oat_scopes"` //may be null
	Revoked        bool           `db:"oat_revoked"`
	ExpiresAt      sql.NullTime   `db:"oat_expires_at"`
	Audience       string         `db:"oat_audience"`
	CreatedAt      time.Time      `db:"oat_created_at"`
}

//type OAuthClientPart struct { oat_refresh_token_id
//	Id     string `db:"oc_id"`
//	Secret string `db:"oc_secret"`
//}
