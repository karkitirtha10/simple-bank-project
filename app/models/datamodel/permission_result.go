package datamodel

import (
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

type PermissionResult struct {
	Permission *dbmodel.Permission
	Err        error
}
