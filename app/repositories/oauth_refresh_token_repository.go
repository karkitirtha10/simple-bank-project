package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

type OAuthRefreshTokenRepository struct {
	DB *sqlx.DB
}

func (yo OAuthRefreshTokenRepository) Insert(
	ch chan datamodel.InsertOAuthRefreshTokenResult,
	oAuthRefreshToken dbmodel.OAuthRefreshToken,
	cols string,
) {
	err := yo.DB.QueryRowx(
		`
		INSERT INTO oauth_refresh_tokens 
		(
			 ort_id,
			 ort_client_id,
			 ort_user_id,
			 ort_revoked,
			 ort_expires_at
		) 
		VALUES ($1,$2,$3,$4,$5) 
		RETURNING `+cols,
		oAuthRefreshToken.Id,
		oAuthRefreshToken.ClientId,
		oAuthRefreshToken.UserId,
		oAuthRefreshToken.Revoked,
		oAuthRefreshToken.ExpiresAt,
	).StructScan(&oAuthRefreshToken)

	_ = oAuthRefreshToken
	ch <- datamodel.InsertOAuthRefreshTokenResult{
		OAuthRefreshToken: oAuthRefreshToken,
		Err:               err,
	}

}

//func (yo OAuthRefreshTokenRepository) FindForIdAndSecret(
//	ch chan datamodel.OAuthClientResult,
//	id string,
//	secret string,
//	cols string,
//) {
//	var oAuthClient dbmodel.OAuthClient
//	err := yo.DB.QueryRowx(
//		`SELECT `+cols+` FROM oauth_clients
//				WHERE oc_id = $1  and oc_secret = $2 LIMIT 1`,
//		id,
//		secret,
//	).StructScan(&oAuthClient)
//
//	ch <- datamodel.OAuthClientResult{
//		OAuthClient: oAuthClient,
//		Err:         err,
//	}
//}

type IOAuthRefreshTokenRepository interface {
	Insert(ch chan datamodel.InsertOAuthRefreshTokenResult, oAuthRefreshToken dbmodel.OAuthRefreshToken, cols string)

	//FindForIdAndSecret(
	//	ch chan datamodel.OAuthClientResult,
	//	id string,
	//	secret string,
	//	cols string,
	//)
}

//015970170

func NewOAuthRefreshTokenRepository(db *sqlx.DB) IOAuthRefreshTokenRepository {
	return &OAuthRefreshTokenRepository{
		DB: db,
	}
}
