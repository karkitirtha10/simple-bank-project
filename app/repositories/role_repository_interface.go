package repositories

import (
	"database/sql"

	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

type RoleRepositoryInterface interface {
	Insert(role *dbmodel.Role) (sql.Result, error)

	FindForName(
		ch chan datamodel.RoleResult,
		name string,
		cols string,
	)
}
