package datamodel

import (
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

// InsertOAuthRefreshTokenResult  is a dto object of insert query
type InsertOAuthRefreshTokenResult struct {
	OAuthRefreshToken dbmodel.OAuthRefreshToken
	Err               error
}
