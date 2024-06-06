package datamodel

import (
	"database/sql"
)

type OAuthTokenPair struct {
	TokenType    string
	AccessToken  string
	RefreshToken string
	ExpiresAt    sql.NullTime
}
