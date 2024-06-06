package repositories

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
	"github.com/lib/pq"
)

type OAuthAccessTokenRepository struct {
	DB *sqlx.DB
}

func (yo *OAuthAccessTokenRepository) Insert(
	ch chan datamodel.InsertOAuthAccessTokenResult,
	id string,
	refreshTokenId string,
	clientId string,
	userId sql.NullString,
	name string,
	scopes []string,
	revoked bool,
	expiresAt sql.NullTime,
	// audience string,
	// createdBy string,
) {
	var oAuthAccessToken dbmodel.OAuthAccessToken
	query := `
		INSERT INTO oauth_access_tokens 
		(
			 oat_id,
		 	 oat_refresh_token_id,
			 oat_client_id,
			 oat_user_id,
			 oat_name,
			 oat_scopes,
			 oat_revoked,
			 oat_expires_at
		) 
		VALUES ($1,$2,$3,$4,$5, $6, $7, $8) 
		RETURNING oat_id 
	`

	err := yo.DB.QueryRowx(
		query,
		id,
		refreshTokenId,
		clientId,
		userId,
		name,
		pq.Array(scopes),
		revoked,
		expiresAt,
		//createdBy,
	).StructScan(&oAuthAccessToken)

	ch <- datamodel.InsertOAuthAccessTokenResult{
		OAuthAccessToken: oAuthAccessToken,
		Err:              err,
	}

}

// func (yo *OAuthAccessTokenRepository) FindForIdAndSecret(
// 	ch chan datamodel.OAuthClientResult,
// 	id string,
// 	secret string,
// 	cols string,
// ) {
// 	var oAuthClient dbmodel.OAuthClient
// 	err := yo.DB.QueryRowx(
// 		`SELECT `+cols+` FROM oauth_clients
// 				WHERE oc_id = $1  and oc_secret = $2 LIMIT 1`,
// 		id,
// 		secret,
// 	).StructScan(&oAuthClient)

// 	ch <- datamodel.OAuthClientResult{
// 		OAuthClient: oAuthClient,
// 		Err:         err,
// 	}
// }

// FindUnRevokedAndUnExpiredForTokenIdAndUserId
func (yo *OAuthAccessTokenRepository) FindUnRevokedAndUnExpiredForTokenIdAndUserId(
	ch chan datamodel.OAuthAccessTokenResult,
	tokenId string,
	userId string,
	cols string,
) {
	var oAuthToken dbmodel.OAuthAccessToken
	err := yo.DB.QueryRowx(
		`SELECT `+cols+` 
		FROM oauth_access_tokens
		WHERE oat_id = $1  
		and oat_user_id = $2 
		and oat_revoked = false
		and oat_expires_at = $3
		LIMIT 1`,
		tokenId,
		userId,
		time.Now(),
	).StructScan(&oAuthToken)

	ch <- datamodel.OAuthAccessTokenResult{
		OAuthAccessToken: oAuthToken,
		Error:            err,
	}
}

type IOAuthAccessTokenRepository interface {
	Insert(
		ch chan datamodel.InsertOAuthAccessTokenResult,
		id string,
		refreshTokenId string,
		clientId string,
		userId sql.NullString,
		name string,
		scopes []string,
		revoked bool,
		expiresAt sql.NullTime,
		// audience string,
		// createdBy string,
	)

	FindUnRevokedAndUnExpiredForTokenIdAndUserId(
		ch chan datamodel.OAuthAccessTokenResult,
		tokenId string,
		userId string,
		cols string,
	)
	

	//FindForIdAndSecret(
	//	ch chan datamodel.OAuthClientResult,
	//	id string,
	//	secret string,
	//	cols string,
	//)
}

func NewOAuthAccessTokenRepository(db *sqlx.DB) IOAuthAccessTokenRepository {
	return &OAuthAccessTokenRepository{DB: db}
}

//015970170
