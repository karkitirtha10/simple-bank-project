package repositories

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

type PermissionRepository struct {
	DB *sqlx.DB
}

func (yo *PermissionRepository) Insert(
	permission *dbmodel.Permission,
) (sql.Result, error) {
	// var permission dbmodel.permission
	query := `
		INSERT INTO permissions 
		(id, name, category,description) 
		VALUES ($1,$2,$3,$4)
	`
	return yo.DB.Exec(
		query,
		permission.Id,
		permission.Name,
		permission.Category,
		permission.Description,
	)

	// fmt.Println(permission)

}

func (yo *PermissionRepository) FindForName(
	ch chan datamodel.PermissionResult,
	name string,
	cols string,
) {
	var permission dbmodel.Permission
	err := yo.DB.QueryRowx(
		`SELECT `+cols+` FROM permissions 
				WHERE name = $1 LIMIT 1`,
		name,
	).StructScan(&permission)

	// aa := oAuthClient
	ch <- datamodel.PermissionResult{Permission: &permission, Err: err}

	// ch <- datamodel.UserResult{User: user, Err: err}
}

// type PermissionRepositoryInterface interface {
// 	Insert(permission dbmodel.Permission) error
// }

func NewPermissionRepository(db *sqlx.DB) PermissionRepositoryInterface {
	return &PermissionRepository{db}
}
