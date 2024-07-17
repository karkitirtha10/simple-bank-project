package repositories

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

type RoleRepository struct {
	DB *sqlx.DB
}

func (yo *RoleRepository) Insert(
	role *dbmodel.Role,
) (sql.Result, error) {
	// var Role dbmodel.Role
	query := `INSERT INTO roles (id, name, display_name, description,is_system,tenant_id,created_by) 
	VALUES ($1,$2,$3,$4,$5,$6,$7)`

	return yo.DB.Exec(
		query,
		role.Id,
		role.Name,
		role.DisplayName,
		role.Description,
		role.IsSystem,
		role.TenantId,
		role.CreatedBy,
	)

}

func (yo *RoleRepository) FindForName(
	ch chan datamodel.RoleResult,
	name string,
	cols string,
) {
	var role dbmodel.Role
	err := yo.DB.QueryRowx(
		`SELECT `+cols+` FROM Roles 
				WHERE name = $1 LIMIT 1`,
		name,
	).StructScan(&role)

	// aa := oAuthClient
	ch <- datamodel.RoleResult{Role: &role, Err: err}

	// ch <- datamodel.UserResult{User: user, Err: err}
}

// func (yo *RoleRepository) FindForIdAndSecret(
// 	ch chan datamodel.RoleResult,
// 	id string,
// 	secret string,
// 	cols string,
// ) {
// 	var Role dbmodel.Role
// 	err := yo.DB.QueryRowx(
// 		`SELECT `+cols+` FROM oauth_clients
// 				WHERE oc_id = $1  and oc_secret = $2 LIMIT 1`,
// 		id,
// 		secret,
// 	).StructScan(&Role)

// 	// aa := Role
// 	ch <- datamodel.RoleResult{Role: Role, Err: err}

// 	// ch <- datamodel.UserResult{User: user, Err: err}
// }

// FindUnRevokedAndUnExpiredForTokenIdAndUserId implements IOAuthAccessTokenRepository.
// func (yo RoleRepository) FindUnRevokedAndUnExpiredForTokenIdAndUserId(ch chan datamodel.OAuthAccessTokenResult, tokenId string, userId string, cols string) {
// 	panic("unimplemented")
// }

// Insert implements IOAuthAccessTokenRepository.
// func (yo RoleRepository) Insert(ch chan datamodel.InsertOAuthAccessTokenResult, id string, refreshTokenId string, clientId string, userId sql.NullString, name string, scopes []string, revoked bool, expiresAt sql.NullTime) {
// 	panic("unimplemented")
// }

func NewRoleRepository(db *sqlx.DB) RoleRepositoryInterface {
	return &RoleRepository{
		DB: db,
	}
}

//015970170
