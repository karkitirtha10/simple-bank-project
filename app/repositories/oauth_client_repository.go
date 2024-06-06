package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/app/enums"
	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

type OAuthClientRepository struct {
	DB *sqlx.DB
}

func (yo *OAuthClientRepository) Insert(
	ch chan datamodel.InsertOAuthClientResult,
	id string,
	name string,
	secret string,
	ocType enums.OAuthClientTypeEnum,
	revoked bool,
) {
	var oAuthClient dbmodel.OAuthClient
	query := "INSERT INTO oauth_clients (oc_id, oc_name, oc_secret, oc_type, oc_revoked) VALUES ($1,$2,$3,$4,$5) RETURNING oc_id, oc_secret"

	err := yo.DB.QueryRowx(
		query,
		id,
		name,
		secret,
		ocType,
		revoked,
	).StructScan(&oAuthClient)

	// fmt.Println(oAuthClient)

	ch <- datamodel.InsertOAuthClientResult{
		OAuthClient: oAuthClient,
		Err:         err,
	}

}

func (yo *OAuthClientRepository) FindForIdAndSecret(
	ch chan datamodel.OAuthClientResult,
	id string,
	secret string,
	cols string,
) {
	var oAuthClient dbmodel.OAuthClient
	err := yo.DB.QueryRowx(
		`SELECT `+cols+` FROM oauth_clients 
				WHERE oc_id = $1  and oc_secret = $2 LIMIT 1`,
		id,
		secret,
	).StructScan(&oAuthClient)

	// aa := oAuthClient
	ch <- datamodel.OAuthClientResult{OAuthClient: oAuthClient, Err: err}

	// ch <- datamodel.UserResult{User: user, Err: err}
}

// FindUnRevokedAndUnExpiredForTokenIdAndUserId implements IOAuthAccessTokenRepository.
// func (yo OAuthClientRepository) FindUnRevokedAndUnExpiredForTokenIdAndUserId(ch chan datamodel.OAuthAccessTokenResult, tokenId string, userId string, cols string) {
// 	panic("unimplemented")
// }

// Insert implements IOAuthAccessTokenRepository.
// func (yo OAuthClientRepository) Insert(ch chan datamodel.InsertOAuthAccessTokenResult, id string, refreshTokenId string, clientId string, userId sql.NullString, name string, scopes []string, revoked bool, expiresAt sql.NullTime) {
// 	panic("unimplemented")
// }

type IOAuthClientRepository interface {
	Insert(
		ch chan datamodel.InsertOAuthClientResult,
		id string,
		name string,
		secret string,
		ocType enums.OAuthClientTypeEnum,
		revoked bool,
	)

	FindForIdAndSecret(
		ch chan datamodel.OAuthClientResult,
		id string,
		secret string,
		cols string,
	)
}

func NewOAuthClientRepository(db *sqlx.DB) IOAuthClientRepository {
	return &OAuthClientRepository{
		DB: db,
	}
}

//015970170
