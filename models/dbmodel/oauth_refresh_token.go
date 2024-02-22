package dbmodel

import (
	"database/sql"
	"time"
)

type OAuthRefreshToken struct {
	Id       string         `db:"ort_id"`
	ClientId string         `db:"ort_client_id"`
	UserId   sql.NullString `db:"ort_user_id"`
	//Name      sql.NullString `db:"ort_name"`
	//Scopes    []string       `db:"ort_scopes"` //may be null
	Revoked   bool         `db:"ort_revoked"`
	ExpiresAt sql.NullTime `db:"ort_expires_at"`
	//Audience  string       `db:"ort_audience"`
	CreatedAt time.Time `db:"ort_created_at"`
}

//type OAuthClientPart struct {
//	Id     string `db:"oc_id"`
//	Secret string `db:"oc_secret"`
//}
