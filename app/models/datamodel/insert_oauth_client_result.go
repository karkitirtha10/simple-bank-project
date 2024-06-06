package datamodel

import (
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

// InsertOAuthClientResult usercontroller result set
type InsertOAuthClientResult struct {
	OAuthClient dbmodel.OAuthClient
	Err         error
}
