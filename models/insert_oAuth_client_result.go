package datamodel

import "github.com/karkitirtha10/simplebank/models/dbmodel"

// InsertOAuthClientResult user result set
type InsertOAuthClientResult struct {
	OAuthClient dbmodel.OAuthClient
	Err         error
}
