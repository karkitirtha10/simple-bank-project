package services

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/karkitirtha10/simplebank/app/enums"
	"github.com/karkitirtha10/simplebank/app/models/datamodel"
	"github.com/karkitirtha10/simplebank/app/models/dbmodel"
	"github.com/karkitirtha10/simplebank/app/repositories"
	"github.com/karkitirtha10/simplebank/app/systemerror"
)

type RolePermissionPersister struct {
	PermissionRepository     repositories.PermissionRepositoryInterface
	RoleRepository           repositories.RoleRepositoryInterface
	RolePermissionRepository repositories.RolePermissionRepositoryInterface
}

func (yo *RolePermissionPersister) FindRoleOrCreate(
	roleName enums.Role,
	isSystem bool,
) (*dbmodel.Role, error) {
	roleCh := make(chan datamodel.RoleResult)
	go yo.RoleRepository.FindForName(roleCh, string(roleName), "id")
	roleResult := <-roleCh

	var err error
	role := roleResult.Role
	if roleResult.Err == sql.ErrNoRows {
		role, err = yo.CreateRole(roleName, isSystem)
		if err != nil {
			return nil, err
		}

		return role, nil
	}

	if roleResult.Err != nil {
		return nil, systemerror.NewServerError(roleResult.Err.Error())
	}
	return role, nil
}

func (yo *RolePermissionPersister) CreateRole(
	roleName enums.Role,
	isSystem bool,
) (*dbmodel.Role, error) {
	roleId, err1 := uuid.NewV7()
	if err1 != nil {
		return nil, systemerror.NewServerError(err1.Error())
	}

	errch := make(chan error)

	role := &dbmodel.Role{
		Id:          roleId.String(),
		Name:        string(roleName),
		DisplayName: string(roleName),
		Description: sql.NullString{},
		IsSystem:    isSystem,
	}
	go func(ch chan error) {
		_, err := yo.RoleRepository.Insert(role)
		ch <- err
	}(errch)

	err := <-errch
	if err != nil {
		return nil, systemerror.NewServerError(err.Error())
	}
	return role, nil
}

func (yo *RolePermissionPersister) FindPermissionOrCreate(
	permissionEnum enums.Permission,
	permissionCategoryEnum enums.PermissionCategory,
) (*dbmodel.Permission, error) {
	ch := make(chan datamodel.PermissionResult)
	go yo.PermissionRepository.FindForName(ch, string(permissionEnum), "id")
	permissionResult := <-ch

	var err error
	permission := permissionResult.Permission
	if permissionResult.Err == sql.ErrNoRows {
		permission, err = yo.CreatePermission(
			permissionEnum,
			permissionCategoryEnum,
		)
		if err != nil {
			return nil, err
		}

		return permission, nil
	}

	if permissionResult.Err != nil {
		return nil, systemerror.NewServerError(
			permissionResult.Err.Error(),
		)
	}
	return permission, nil
}

func (yo *RolePermissionPersister) CreatePermission(
	permissionEnum enums.Permission,
	permissionCategoryEnum enums.PermissionCategory,
) (*dbmodel.Permission, error) {
	permissionId, err1 := uuid.NewV7()
	if err1 != nil {
		return nil, err1
	}

	errch := make(chan error)
	permission := &dbmodel.Permission{
		Id:          permissionId.String(),
		Name:        string(permissionEnum),
		Category:    string(permissionCategoryEnum),
		Description: sql.NullString{},
	}

	go func(ch chan error) {
		_, err := yo.PermissionRepository.Insert(permission)
		ch <- err
	}(errch)

	err := <-errch
	if err != nil {
		// return nil, err
		return nil, systemerror.NewServerError(err.Error())
	}

	return permission, nil
}

func (yo *RolePermissionPersister) SyncPermissionsToRole(
	roleId string,
	permissionIds []string,
) error {
	errch := make(chan error)
	for _, permissionId := range permissionIds {

		rolePermId, err1 := uuid.NewV7()
		if err1 != nil {
			return err1
		}

		rolePermission := &dbmodel.RolePermission{
			Id:           rolePermId.String(),
			RoleId:       roleId,
			PermissionId: permissionId,
		}

		go func(errCh chan error) {
			_, err := yo.RolePermissionRepository.InsertIfNotExists(rolePermission)

			errCh <- err
		}(errch)

		err := <-errch

		if err != nil {
			return err
		}
	}
	return nil
}

func NewRolePermissionPersister(
	permissionRepository repositories.PermissionRepositoryInterface,
	roleRepository repositories.RoleRepositoryInterface,
	rolePermissionRepository repositories.RolePermissionRepositoryInterface,
) RolePermissionPersisterInterface {
	return &RolePermissionPersister{
		PermissionRepository:     permissionRepository,
		RoleRepository:           roleRepository,
		RolePermissionRepository: rolePermissionRepository,
	}
}
