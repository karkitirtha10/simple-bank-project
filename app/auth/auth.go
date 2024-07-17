package auth

type Authorization struct {
}

func (yo *Authorization) hasPermission() bool {
	return true
}

func (yo *Authorization) hasAnyPermission() bool {
	return true
}

func (yo *Authorization) hasAllPermissions() bool {
	return true
}

func (yo *Authorization) hasRole() bool {
	return true
}

func (yo *Authorization) hasAnyRole() bool {
	return true
}

func (yo *Authorization) hasAllRoles() bool {
	return true
}
