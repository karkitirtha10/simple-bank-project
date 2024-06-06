package datamodel

import (
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

type OAuthClientResult struct {
	OAuthClient dbmodel.OAuthClient
	Err         error
}
