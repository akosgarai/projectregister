package repository

import (
	"time"

	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/model"
)

// RoleRepository type
type RoleRepository struct {
	db *database.DB
}

// NewRoleRepository creates a new role repository
func NewRoleRepository(db *database.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

// CreateRole creates a new role
// it returns the created role and an error
// the input parameter is the name and the resource ids
func (r *RoleRepository) CreateRole(name string, resourceIDs []int64) (*model.Role, error) {
	var role model.Role
	query := "INSERT INTO roles (name) VALUES ($1) RETURNING *"
	err := r.db.QueryRow(query, name).Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}
	// insert the role_to_resource records
	for _, resourceID := range resourceIDs {
		query = "INSERT INTO role_to_resource (role_id, resource_id) VALUES ($1, $2)"
		_, err = r.db.Exec(query, role.ID, resourceID)
		if err != nil {
			return nil, err
		}
	}
	return &role, err
}

// GetRoleByName gets a role by name
// the input parameter is the rolename
// it returns the role and an error
func (r *RoleRepository) GetRoleByName(name string) (*model.Role, error) {
	var role model.Role
	query := "SELECT * FROM roles WHERE name = $1"
	err := r.db.QueryRow(query, name).Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}
	// get the resources
	query = "SELECT r.id, r.name FROM resources r JOIN role_to_resources rr ON r.id = rr.resource_id WHERE rr.role_id = $1"
	rows, err := r.db.Query(query, role.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var resource model.Resource
		err = rows.Scan(&resource.ID, &resource.Name)
		if err != nil {
			return nil, err
		}
		role.Resources = append(role.Resources, resource)
	}
	return &role, err
}

// GetRoleByID gets a role by id
// the input parameter is the role id
// it returns the role and an error
func (r *RoleRepository) GetRoleByID(id int64) (*model.Role, error) {
	// get role by id
	var role model.Role
	query := "SELECT * FROM roles WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}
	// get the resources
	query = "SELECT r.id, r.name FROM resources r JOIN role_to_resources rr ON r.id = rr.resource_id WHERE rr.role_id = $1"
	rows, err := r.db.Query(query, role.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var resource model.Resource
		err = rows.Scan(&resource.ID, &resource.Name)
		if err != nil {
			return nil, err
		}
		role.Resources = append(role.Resources, resource)
	}
	return &role, err
}

// UpdateRole updates a role
// the input parameter is the role
// it returns an error
func (r *RoleRepository) UpdateRole(role *model.Role) error {
	query := "UPDATE roles SET name = $1, updated_at = $2 WHERE id = $3"
	now := time.Now().Format("2006-01-02 15:04:05.999999-07:00")
	_, err := r.db.Exec(query, role.Name, now, role.ID)
	return err
}

// DeleteRole deletes a role
// the input parameter is the role id
// it returns an error
func (r *RoleRepository) DeleteRole(id int64) error {
	query := "DELETE FROM roles WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

// GetRoles gets all roles
// it returns the roles and an error
func (r *RoleRepository) GetRoles() ([]*model.Role, error) {
	// get all roles
	var roles []*model.Role
	query := "SELECT * FROM roles"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var role model.Role
		err = rows.Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}
	return roles, nil
}
