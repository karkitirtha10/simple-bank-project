package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/models/datamodel"
	"github.com/karkitirtha10/simplebank/models/dbmodel"
)

type OAuthAccessTokenRepository struct {
	DB *sqlx.DB
}

func (yo OAuthAccessTokenRepository) Insert(
	ch chan datamodel.InsertOAuthAccessTokenResult,
	id string,
	RefreshTokenId string,
	clientId string,
	userId string,
	name string,
	scopes []string,
	revoked bool,
	expiresAt time.Time,
	audience string,
	// createdBy string,
) {
	var oAuthAccessToken dbmodel.OAuthAccessToken
	query := `
		INSERT INTO oauth_access_tokens 
		(
			 oat_id,
			 oat_client_id,
			 oat_user_id,
			 oat_name,
			 oat_scopes,
			 oat_revoked,
			 oat_expires_at,
			oat_audience
		) 
		VALUES ($1,$2,$3,$4,$5, $6, $7, $8) 
		RETURNING oc_id 
	`

	err := yo.DB.QueryRowx(
		query,
		id,
		clientId,
		userId,
		name,
		scopes,
		revoked,
		expiresAt,
		audience,
		//createdBy,
	).StructScan(&oAuthAccessToken)

	ch <- datamodel.InsertOAuthAccessTokenResult{
		OAuthAccessToken: oAuthAccessToken,
		Err:              err,
	}

}

//func (yo OAuthAccessTokenRepository) FindForIdAndSecret(
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

type IOAuthAccessTokenRepository interface {
	Insert(
		ch chan datamodel.InsertOAuthAccessTokenResult,
		id string,
		RefreshTokenId string,
		clientId string,
		userId string,
		name string,
		scopes []string,
		revoked bool,
		expiresAt time.Time,
		audience string,
		// createdBy string,
	)

	//FindForIdAndSecret(
	//	ch chan datamodel.OAuthClientResult,
	//	id string,
	//	secret string,
	//	cols string,
	//)
}

//015970170
