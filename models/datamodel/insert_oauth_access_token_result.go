package datamodel

import "github.com/karkitirtha10/simplebank/models/dbmodel"

// InsertOAuthAccessTokenResult  is a dto object of insert query
type InsertOAuthAccessTokenResult struct {
	OAuthAccessToken dbmodel.OAuthAccessToken
	Err              error
}
