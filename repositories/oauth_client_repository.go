package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/enums"
	datamodel "github.com/karkitirtha10/simplebank/models"
	"github.com/karkitirtha10/simplebank/models/dbmodel"
)

type OAuthClientRepository struct {
	DB *sqlx.DB
}

func (yo OAuthClientRepository) Insert(
	ch chan datamodel.InsertOAuthClientResult,
	name string,
	secret string,
	ocType enums.OAuthClientTypeEnum,
	revoked bool,
) {
	var oAuthClient dbmodel.OAuthClient
	query := "INSERT INTO oauth_clients (oc_name, oc_secret, oc_type, oc_revoked) VALUES ($1,$2,$3,$4) RETURNING oc_id, oc_secret"

	err := yo.DB.QueryRowx(
		query,
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

type IOAuthClientRepository interface {
	Insert(
		ch chan datamodel.InsertOAuthClientResult,
		name string,
		secret string,
		ocType enums.OAuthClientTypeEnum,
		revoked bool,
	)
}

//015970170
