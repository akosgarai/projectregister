package model

// Role type
type Role struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string

	Resources []*Resource
}

// HasResource checks if the role has the resource
func (r *Role) HasResource(resource string) bool {
	for _, res := range r.Resources {
		if res.Name == resource {
			return true
		}
	}
	return false
}

// RoleRepository interface
type RoleRepository interface {
	CreateRole(name string, resourceIDs []int64) (*Role, error)
	GetRoleByName(name string) (*Role, error)
	GetRoleByID(id int64) (*Role, error)
	UpdateRole(role *Role, resourceIDs []int64) error
	DeleteRole(id int64) error
	GetRoles() ([]*Role, error)
}
