package repositories

import (
	"database/sql"

	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

type RolePermissionRepositoryInterface interface {
	InsertIfNotExists(
		user *dbmodel.RolePermission,
	) (sql.Result, error)
}
