package enums

type PermissionCategory string

const USER PermissionCategory = "user"
const ROLE PermissionCategory = "role"

var CategoryToPermissions = map[PermissionCategory][]Permission{
	USER: {
		CREATE_NEW_USER,
		VIEW_USER,
		ARCHIVE_USER,
	},
}
