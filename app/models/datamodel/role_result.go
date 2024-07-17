package datamodel

import (
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

type RoleResult struct {
	Role *dbmodel.Role
	Err  error
}
