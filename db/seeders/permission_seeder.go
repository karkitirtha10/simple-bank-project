package seeders

import (
	"github.com/karkitirtha10/simplebank/app/enums"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
	"github.com/karkitirtha10/simplebank/app/services"
)

//Restore may not be needed
// 1.	Create New Post: 'create_new_post'
// 2.	View Post: 'view_post'
// 3.	Edit Post: 'edit_post'
// 4.	Delete Post: 'delete_post'
// 5.	Archive Post: 'archive_post'
// 6.	Restore Archived Post: 'restore_archived_post'

type PermissionSeeder struct {
	RolePermissionPersister services.RolePermissionPersisterInterface
}

func (yo *PermissionSeeder) Seed() error {

	var err error
	permissionIds, err := yo.getIdsAfterSeedingPermissions()
	if err != nil {
		// fmt.Println(" Failed to seed permissions. " + err.Error())
		return err
	}

	err = yo.SeedingRoles(permissionIds)
	if err != nil {
		// fmt.Println(" Failed to seed roles. " + err.Error())
		return err
	}

	return nil

}

func (yo *PermissionSeeder) getIdsAfterSeedingPermissions() ([]string, error) {
	var existingPermissionIds []string

	for category, permissions := range enums.CategoryToPermissions {
		for _, permission := range permissions {
			permissionModel, err := yo.RolePermissionPersister.FindPermissionOrCreate(
				permission,
				category,
			) ////TODO: PROBLEM INSIDE HERE

			if err != nil {
				// fmt.Println("Failed to seed UserSeeder. " + err.Error())
				return nil, err
			}
			existingPermissionIds = append(existingPermissionIds, permissionModel.Id)
		}
	}
	return existingPermissionIds, nil
}

func (yo *PermissionSeeder) SeedingRoles(permissionIds []string) error {
	var (
		role *dbmodel.Role
		err  error
	)
	// for _, roleEnum = range []enums.Role{
	// 	enums.SUPERADMIN,
	// 	enums.TENANT_GROUP_ADMIN,
	// 	enums.ADMIN,
	// } {
	// 	role, err = yo.findRoleOrCreate(roleEnum, true)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	err = yo.syncPermissionsToRole(role.Id, permissionIds)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	/*
	 *  superadmin
	 */
	yo.RolePermissionPersister.FindRoleOrCreate(enums.SUPERADMIN, true)

	/*
	 *  TENANT_GROUP_ADMIN
	 */
	role, err = yo.RolePermissionPersister.FindRoleOrCreate(enums.TENANT_GROUP_ADMIN, true)
	if err != nil {
		return err
	}
	err = yo.RolePermissionPersister.SyncPermissionsToRole(role.Id, permissionIds)
	if err != nil {
		return err
	}

	/*
	 *  ADMIN
	 */
	role, err = yo.RolePermissionPersister.FindRoleOrCreate(enums.ADMIN, true)
	if err != nil {
		return err
	}
	err = yo.RolePermissionPersister.SyncPermissionsToRole(role.Id, permissionIds)
	if err != nil {
		// fmt.Println("failed SyncPermissionsToRole 2" + err.Error())
		return err
	}

	return nil
}

func NewPermissionSeeder(rolePermissionPersister services.RolePermissionPersisterInterface) SeederInterface {
	return &PermissionSeeder{
		RolePermissionPersister: rolePermissionPersister,
	}
}
