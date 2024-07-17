package services

import (
	"github.com/karkitirtha10/simplebank/app/enums"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
)

type RolePermissionPersisterInterface interface {
	FindRoleOrCreate(
		roleName enums.Role,
		isSystem bool,
	) (*dbmodel.Role, error)

	FindPermissionOrCreate(
		permissionEnum enums.Permission,
		permissionCategoryEnum enums.PermissionCategory,
	) (*dbmodel.Permission, error)

	SyncPermissionsToRole(
		roleId string,
		permissionIds []string,
	) error
}
