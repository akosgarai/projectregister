package model

// Role type
type Role struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string
}

// RoleRepository interface
type RoleRepository interface {
	CreateRole(name string) (*Role, error)
	GetRoleByName(name string) (*Role, error)
	GetRoleByID(id int64) (*Role, error)
	UpdateRole(role *Role) error
	DeleteRole(id int64) error
	GetRoles() ([]*Role, error)
}
