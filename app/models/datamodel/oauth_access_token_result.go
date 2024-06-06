package datamodel

import "github.com/karkitirtha10/simplebank/app/models/dbmodel"

type OAuthAccessTokenResult struct {
	OAuthAccessToken dbmodel.OAuthAccessToken
	Error            error
}
