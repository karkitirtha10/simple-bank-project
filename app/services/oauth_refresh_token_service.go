package services

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
	"github.com/karkitirtha10/simplebank/app/repositories"
	"github.com/karkitirtha10/simplebank/config"
)

type OAuthRefreshTokenService struct {
	OAuthClientRepository       repositories.IOAuthClientRepository
	JWTService                  IJWTService
	OAuthRefreshTokenRepository repositories.IOAuthRefreshTokenRepository
}

func (yo *OAuthRefreshTokenService) Generate(
	clientId string,
	user dbmodel.User,
	cnf *config.Config,
) (string, string, error) {

	//create claims
	expiresIn := time.Duration(cnf.RefreshTokenExpiresIn) * time.Hour
	refreshTokenId := uuid.Must(uuid.NewV7()).String()
	refreshExpiresAt := time.Now().Add(expiresIn)
	refreshClaims := jwt.MapClaims{
		"exp": refreshExpiresAt, //5 DAYS
		"iat": time.Now().Unix(),
		"aud": cnf.AppUrl,
		"iss": cnf.AppUrl,
		"sub": user.Id,
		"jti": refreshTokenId,
	}

	//Generate refresh token
	token, err := yo.JWTService.CreateToken(refreshClaims, cnf.PrivateKeyPath)
	if err != nil {
		return "", "", err
	}

	//persist refresh token
	ch := make(chan datamodel.InsertOAuthRefreshTokenResult)
	oAuthRefreshToken := dbmodel.OAuthRefreshToken{
		Id:        refreshTokenId,
		ClientId:  clientId,
		UserId:    sql.NullString{String: user.Id, Valid: true},
		Revoked:   false,
		ExpiresAt: sql.NullTime{Time: refreshExpiresAt, Valid: true},
	}
	go yo.OAuthRefreshTokenRepository.Insert(ch, oAuthRefreshToken, "ort_id, ort_created_at")
	insertOAuthRefreshTokenResult := <-ch

	if insertOAuthRefreshTokenResult.Err != nil {
		return "", "", err
	}

	return token, insertOAuthRefreshTokenResult.OAuthRefreshToken.Id, nil
}

type IOAuthRefreshTokenService interface {
	Generate(
		clientId string,
		user dbmodel.User,
		cnf *config.Config,
	) (string, string, error)
}

func NewOAuthRefreshTokenService(
	OAuthClientRepository repositories.IOAuthClientRepository,
	JWTService IJWTService,
	OAuthRefreshTokenRepository repositories.IOAuthRefreshTokenRepository,
) IOAuthRefreshTokenService {
	return &OAuthRefreshTokenService{
		OAuthClientRepository:       OAuthClientRepository,
		JWTService:                  JWTService,
		OAuthRefreshTokenRepository: OAuthRefreshTokenRepository,
	}
}
