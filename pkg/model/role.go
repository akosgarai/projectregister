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

// Roles type is a slice of Role
type Roles []*Role

// ToMap converts the Roles to a map
// The key is the string role ID
// The value is the role Name
func (r Roles) ToMap() map[int64]string {
	result := make(map[int64]string)
	for _, role := range r {
		result[role.ID] = role.Name
	}
	return result
}

// RoleRepository interface
type RoleRepository interface {
	CreateRole(name string, resourceIDs []int64) (*Role, error)
	GetRoleByName(name string) (*Role, error)
	GetRoleByID(id int64) (*Role, error)
	UpdateRole(role *Role, resourceIDs []int64) error
	DeleteRole(id int64) error
	GetRoles() (*Roles, error)
}
