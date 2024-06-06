package services

import (
	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

type PersonalAccessClientServiceInterface interface {
	Generate(user dbmodel.User) (*datamodel.OAuthTokenPair, error)
	GenerateForUserAndRefreshTokenId(userId string, refreshTokenId string) (*datamodel.OAuthTokenPair, error)
	GenerateFromRefreshToken(userId string, oAuthRefreshTokenId string) (*datamodel.OAuthTokenPair, error)
}
