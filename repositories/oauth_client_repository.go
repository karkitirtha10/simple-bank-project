package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/enums"
	"github.com/karkitirtha10/simplebank/models/datamodel"
	"github.com/karkitirtha10/simplebank/models/dbmodel"
)

type OAuthClientRepository struct {
	DB *sqlx.DB
}

func (yo OAuthClientRepository) Insert(ch chan datamodel.InsertOAuthClientResult, id string, name string, secret string, ocType enums.OAuthClientTypeEnum, revoked bool) {
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

func (yo OAuthClientRepository) FindForIdAndSecret(
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

	ch <- datamodel.OAuthClientResult{
		OAuthClient: oAuthClient,
		Err:         err,
	}
}

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

//015970170
