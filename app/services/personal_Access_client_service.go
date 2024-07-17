package services

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/karkitirtha10/simplebank/app/enums"
	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
	"github.com/karkitirtha10/simplebank/app/repositories"
	"github.com/karkitirtha10/simplebank/app/systemerror"
	"github.com/karkitirtha10/simplebank/config"
)

type PersonalAccessClientService struct {
	OAuthClientRepository      repositories.IOAuthClientRepository
	JWTService                 IJWTService
	OAuthAccessTokenRepository repositories.IOAuthAccessTokenRepository
	OAuthRefreshTokenService   IOAuthRefreshTokenService
	userRepository             repositories.UserRepositoryInterface
	Config                     *config.Config
}

func (yo *PersonalAccessClientService) Generate(user dbmodel.User) (*datamodel.OAuthTokenPair, error) {

	var (
		clientId = yo.Config.OAuthPersonalAccessClientId     //cnf
		secret   = yo.Config.OAuthPersonalAccessClientSecret //cnf
	)

	clientCh := make(chan datamodel.OAuthClientResult)
	go yo.OAuthClientRepository.FindForIdAndSecret(clientCh, clientId, secret, `oc_id`)
	oAuthClientResult := <-clientCh

	if errors.Is(oAuthClientResult.Err, sql.ErrNoRows) {
		return nil, systemerror.NewError(
			"oauth client not found for "+"oc_id = "+clientId+", oc_secret = "+secret,
			"oops! something ent wrong",
			enums.PERSONAL_ACCESS_CLIENT_NOT_FOUND,
			http.StatusInternalServerError,
			oAuthClientResult.Err,
		)
	}
	//todo handle
	if oAuthClientResult.Err != nil {
		return nil, oAuthClientResult.Err
	}

	refreshToken, refreshTokenId, err := yo.OAuthRefreshTokenService.Generate(
		clientId,
		user,
		yo.Config,
	)
	if err != nil {
		return nil, err
	}

	oAuthTokenPair, err := yo.GenerateForUserAndRefreshTokenId(user.Id, refreshTokenId)
	if err != nil {
		return nil, err
	}
	oAuthTokenPair.RefreshToken = refreshToken

	return oAuthTokenPair, err
}

func (yo *PersonalAccessClientService) GenerateForUserAndRefreshTokenId(userId string, refreshTokenId string) (*datamodel.OAuthTokenPair, error) {

	var (
		clientId       = yo.Config.OAuthPersonalAccessClientId                             //cnf
		privateKeyPath = yo.Config.PrivateKeyPath                                          //cnf 		//userResult
		expiresIn      = time.Duration(yo.Config.PersonalAccessTokenExpiresIn) * time.Hour //in hours
	)

	//create claims
	// accessTokenId := uuid.Must(uuid.NewV7()).String()
	uuidObject, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	accessTokenId := uuidObject.String()
	accessClaims := jwt.MapClaims{
		// "exp": nullableExpiresAt, //5 DAYS
		"iat": time.Now().Unix(),
		"aud": yo.Config.AppUrl,
		"iss": yo.Config.AppUrl,
		"sub": userId,
		"jti": accessTokenId,
	}

	expiresAt := time.Now().Add(expiresIn)
	nullableExpiresAt := sql.NullTime{Time: expiresAt, Valid: false}
	if int(expiresIn) > 0 {
		accessClaims["exp"] = expiresAt.Unix()
		nullableExpiresAt.Valid = true
	}

	//Generate access token
	accessToken, err := yo.JWTService.CreateToken(accessClaims, privateKeyPath)
	if err != nil {
		//return "", err
		return nil, err
	}
	var s []string
	//persist access token
	ch := make(chan datamodel.InsertOAuthAccessTokenResult)
	go yo.OAuthAccessTokenRepository.Insert(
		ch,
		accessTokenId,
		refreshTokenId,
		clientId,
		sql.NullString{String: userId, Valid: true},
		"access_token", //to change type to sql.NullString
		s,
		false,
		nullableExpiresAt,
	)
	insertOAuthAccessTokenResult := <-ch

	if insertOAuthAccessTokenResult.Err != nil {
		//return "", err
		return nil, systemerror.NewServerErrorWithPrevious(
			insertOAuthAccessTokenResult.Err.Error(),
			insertOAuthAccessTokenResult.Err,
		)
	}
	//oauthTokenPair.ExpiresAt = nullableExpiresAt
	//oauthTokenPair.TokenType = "bearer"
	return &datamodel.OAuthTokenPair{
		TokenType:   "bearer",
		AccessToken: accessToken,
		ExpiresAt:   nullableExpiresAt,
	}, nil
	//return datamodel.OAuthTokenPair{TokenType: "bearer", AccessToken: token, ExpiresAt: nullableExpiresAt}, nil
}

func (yo *PersonalAccessClientService) GenerateFromRefreshToken(userId string, oAuthRefreshTokenId string) (*datamodel.OAuthTokenPair, error) {

	var (
		clientId = yo.Config.OAuthPersonalAccessClientId     //cnf
		secret   = yo.Config.OAuthPersonalAccessClientSecret //cnf
		clientCh chan datamodel.OAuthClientResult
	)

	//get personal access client
	go yo.OAuthClientRepository.FindForIdAndSecret(clientCh, clientId, secret, `oc_id`)
	oAuthClientResult := <-clientCh

	if oAuthClientResult.Err != nil {
		return nil, oAuthClientResult.Err
	}

	return yo.GenerateForUserAndRefreshTokenId(userId, oAuthRefreshTokenId)
	//refreshTokenId := oAuthRefreshToken.Id
	//oauthTokenPair, err := yo.GenerateForUserAndRefreshTokenId(userId, oAuthRefreshTokenId, cnf)
	//if err != nil {
	//	return "", err
	//}
	//return oauthTokenPair, nil
}

/*
func (yo PersonalAccessClientService) generateBackUp(
	usercontroller dbmodel.User,
	cnf config.Config,
) (string, string, systemerror) {

	var (
		clientId       = yo.Config.OAuthPersonalAccessClientId                             //cnf
		secret         = yo.Config.OAuthPersonalAccessClientSecret                         //cnf
		privateKeyPath = yo.Config.PrivateKeyPath                                          //cnf 		//userResult
		expiresIn      = time.Duration(yo.Config.PersonalAccessTokenExpiresIn) * time.Hour //in hours
		userId         = usercontroller.Id
	)

	var clientCh chan datamodel.OAuthClientResult
	go yo.OAuthClientRepository.FindForIdAndSecret(clientCh, clientId, secret, `id`)
	oAuthClientResult := <-clientCh
	if oAuthClientResult.Err != nil {

		return "", "", oAuthClientResult.Err
	}

	refreshToken, refreshTokenId, err := yo.OAuthRefreshTokenService.Generate(clientId, usercontroller, cnf)
	if err != nil {
		return "", "", err
	}

	//create claims
	accessTokenId := uuid.Must(uuid.NewV7()).String()
	expiresAt := time.Now().Add(expiresIn)
	accessClaims := jwt.MapClaims{
		"exp": expiresAt, //5 DAYS
		"iat": time.Now().Unix(),
		"aud": yo.Config.AppUrl,
		"iss": yo.Config.AppUrl,
		"sub": userId,
		"jti": accessTokenId,
	}

	//Generate access token
	token, err := yo.JWTService.CreateToken(accessClaims, privateKeyPath)
	if err != nil {
		return "", "", err
	}
	var s []string
	//persist access token
	var ch chan datamodel.InsertOAuthAccessTokenResult
	go yo.OAuthAccessTokenRepository.Insert(
		ch,
		accessTokenId,
		refreshTokenId,
		clientId,
		sql.NullString{String: userId, Valid: true},
		"access_token", //to change type to sql.NullString
		s,
		false,
		sql.NullTime{Time: expiresAt, Valid: true},
	)
	insertOAuthAccessTokenResult := <-ch

	if insertOAuthAccessTokenResult.Err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}
*/
//exp,
//nbf,
//iat,
//aud,
//iss,
//sub,

func NewPersonalAccessClientService(
	OAuthClientRepository repositories.IOAuthClientRepository,
	JWTService IJWTService,
	OAuthAccessTokenRepository repositories.IOAuthAccessTokenRepository,
	OAuthRefreshTokenService IOAuthRefreshTokenService,
	userRepository repositories.UserRepositoryInterface,
	Config *config.Config,
) PersonalAccessClientServiceInterface {
	return &PersonalAccessClientService{
		OAuthClientRepository:      OAuthClientRepository,
		JWTService:                 JWTService,
		OAuthAccessTokenRepository: OAuthAccessTokenRepository,
		OAuthRefreshTokenService:   OAuthRefreshTokenService,
		userRepository:             userRepository,
		Config:                     Config,
	}
}
