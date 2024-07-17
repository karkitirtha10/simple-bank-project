package repositories

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

type RolePermissionRepository struct {
	DB *sqlx.DB
}

func (yo *RolePermissionRepository) InsertIfNotExists(
	user *dbmodel.RolePermission,
) (sql.Result, error) {
	return yo.DB.Exec(
		`
		INSERT INTO role_permissions(
			id,
			role_id, 
			permission_id,
			created_by
		) 
		values ($1, $2, $3, $4)
		ON CONFLICT (role_id, permission_id) DO NOTHING
		`,
		user.Id,
		user.RoleId,
		user.PermissionId,
		user.CreatedBy,
	)
	// _ = user
	// ch <- datamodel.InsertOAuthRefreshTokenResult{
	// 	OAuthRefreshToken: user,
	// 	Err:               err,
	// }

}

func NewRolePermissionRepository(db *sqlx.DB) RolePermissionRepositoryInterface {
	return &RolePermissionRepository{
		DB: db,
	}
}
