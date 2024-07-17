package repositories

import (
	"database/sql"

	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

type PermissionRepositoryInterface interface {
	Insert(permission *dbmodel.Permission) (sql.Result, error)
	FindForName(
		ch chan datamodel.PermissionResult,
		name string,
		cols string,
	)
}
