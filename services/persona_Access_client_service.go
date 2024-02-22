package services

import (
	"database/sql"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/karkitirtha10/simplebank/config"
	"github.com/karkitirtha10/simplebank/models/datamodel"
	"github.com/karkitirtha10/simplebank/models/dbmodel"
	services "github.com/karkitirtha10/simplebank/pkg/jwttoken"
	"github.com/karkitirtha10/simplebank/repositories"
	"time"
)

type PersonalAccessClientService struct {
	OAuthClientRepository repositories.IOAuthClientRepository
	AuthService           IOAuthService
}

func (yo PersonalAccessClientService) generateToken(
// id string,
// secret string,
// privateKeyPath string,
// user dbmodel.User,
) {
	/*
		start creating token\
	*/
	var cnf config.Config //service
	var (
		id             string        = cnf.OAuthPersonalAccessClientId                             //cnf
		secret         string        = cnf.OAuthPersonalAccessClientSecret                         //cnf
		privateKeyPath string        = cnf.PrivateKeyPath                                          //cnf
		user           dbmodel.User                                                                //userResult
		expiresIn      time.Duration = time.Duration(cnf.PersonalAccessTokenExpiresIn) * time.Hour //in hours//cnf
		//clientCredentialsExpiresIn time.Duration
		audience string = cnf.AppUrl //base url of resource server the token is issued for.
		issuer   string = cnf.AppUrl //base url of auth server //cnf

		/*
			baseURL := "http://" + context.Request.Host (in gin)
		*/
	)

	var AuthService services.IOAuthService                                 // service
	var OAuthClientRepository repositories.IOAuthClientRepository          // service
	var oAuthAccessTokenRepository repositories.OAuthAccessTokenRepository // service
	var oAuthRefreshTokenRepository repositories.OAuthRefreshTokenRepository

	user = userResult.User

	var ch2 chan datamodel.OAuthClientResult
	go OAuthClientRepository.FindForIdAndSecret(ch2, id, secret, `id`)
	oAuthClientResult := <-ch2
	if oAuthClientResult.Err != nil {
		return
	}

	//var num int
	//fmt.Println(num)
	/*
	* start refresh token
	 */
	refreshTokenId := uuid.Must(uuid.NewV7()).String()
	refreshExpiresAt := time.Now().Add(expiresIn)
	refreshClaims := jwt.MapClaims{
		"exp": refreshExpiresAt, //5 DAYS
		"iat": time.Now().Unix(),
		"aud": audience,
		"iss": issuer,
		"sub": user.Id,
		"jti": refreshTokenId,
	}

	_, err := AuthService.CreateToken(refreshClaims, privateKeyPath)
	if err != nil {
		return
	}

	var chrefresh chan datamodel.InsertOAuthRefreshTokenResult
	go oAuthRefreshTokenRepository.Insert(
		chrefresh,
		refreshTokenId,
		oAuthClientResult.OAuthClient.Id,
		sql.NullString{String: userResult.User.Id, Valid: true}, //to change type to sql.NullString
		false,
		sql.NullTime{Time: refreshExpiresAt, Valid: true},
	)
	insertOAuthRefreshTokenResult := <-chrefresh

	/*
	* start access token
	 */

	accessTokenId := uuid.Must(uuid.NewV7()).String()
	s := []string{"read", "write"}
	expiresAt := time.Now().Add(expiresIn)
	claims := jwt.MapClaims{
		"exp":   expiresAt, //5 DAYS
		"iat":   time.Now().Unix(),
		"aud":   audience,
		"iss":   issuer,
		"sub":   user.Id,
		"jti":   accessTokenId,
		"scope": s,
	}

	//ACCESSTOKEN SERVICE
	//config bata aune. and pathaunu naparne :
	//iat
	//aud
	//iss
	//jti,
	//PATHAUNIU PARNE : exp, sub(USER ID), jti, scope

	// var AuthService services.IOAuthService // service
	accessToken, err := AuthService.CreateToken(claims, privateKeyPath)
	if err != nil {
		return
	}

	var ch3 chan datamodel.InsertOAuthAccessTokenResult
	go oAuthAccessTokenRepository.Insert(
		ch3,
		accessTokenId,
		insertOAuthRefreshTokenResult.OAuthRefreshToken.Id,
		oAuthClientResult.OAuthClient.Id,
		userResult.User.Id,
		"access_token",
		[]string{},
		false,
		expiresAt,
		audience,
	)
	_ = <-ch3
}

//exp,
//nbf,
//iat,
//aud,
//iss,
//sub,
