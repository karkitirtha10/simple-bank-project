package datamodel

import "github.com/karkitirtha10/simplebank/app/models/dbmodel"

type LoginHandlerResult struct {
	OAuthTokenPair *OAuthTokenPair
	User dbmodel.User
}
