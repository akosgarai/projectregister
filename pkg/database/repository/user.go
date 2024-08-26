package repository

import (
	"strconv"
	"strings"
	"time"

	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/model"
)

// UserRepository type
type UserRepository struct {
	db *database.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *database.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// CreateUser creates a new user
// it returns the created user and an error
// the input parameters are the username, email and encrypted password
func (u *UserRepository) CreateUser(username, email, password string, roleID int64) (*model.User, error) {
	var user model.User
	query := "INSERT INTO users (name, email, password, role_id) VALUES ($1, $2, $3, $4) RETURNING *"
	err := u.db.QueryRow(query, username, email, password, roleID).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &roleID)
	if err != nil {
		return nil, err
	}
	role := model.Role{}
	role.ID = roleID
	user.Role = &role
	return u.withRole(&user)
}

// GetUserByEmail gets a user by email
func (u *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	query := "SELECT * FROM users WHERE email = $1"
	var roleID int64
	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &roleID)
	if err != nil {
		return nil, err
	}
	role := model.Role{}
	role.ID = roleID
	user.Role = &role
	return u.withRole(&user)
}

// GetUserByID gets a user by id
func (u *UserRepository) GetUserByID(id int64) (*model.User, error) {
	// get user by id
	var user model.User
	query := "SELECT * FROM users WHERE id = $1"
	var roleID int64
	err := u.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &roleID)
	if err != nil {
		return nil, err
	}
	role := model.Role{}
	role.ID = roleID
	user.Role = &role
	return u.withRole(&user)
}

// UpdateUser updates a user
func (u *UserRepository) UpdateUser(user *model.User) error {
	query := "UPDATE users SET name = $1, email = $2, password = $3, updated_at = $4, role_id = $5 WHERE id = $6"
	now := time.Now().Format("2006-01-02 15:04:05.999999-07:00")
	_, err := u.db.Exec(query, user.Name, user.Email, user.Password, now, user.Role.ID, user.ID)
	return err
}

// DeleteUser deletes a user
func (u *UserRepository) DeleteUser(id int64) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := u.db.Exec(query, id)
	return err
}

// GetUsers gets all users
func (u *UserRepository) GetUsers(filters *model.UserFilter) ([]*model.User, error) {
	// get all users
	var users []*model.User
	query := "SELECT users.*, roles.name, roles.created_at, roles.updated_at FROM users JOIN roles ON users.role_id = roles.id"
	params := []interface{}{}
	whereConditions := []string{}
	if filters.Name != "" {
		index := len(params) + 1
		whereConditions = append(whereConditions, "users.name LIKE '%' || $"+strconv.Itoa(index)+" || '%'")
		params = append(params, filters.Name)
	}
	if filters.Email != "" {
		index := len(params) + 1
		whereConditions = append(whereConditions, "users.email LIKE '%' || $"+strconv.Itoa(index)+" || '%'")
		params = append(params, filters.Email)
	}
	if len(filters.RoleIDs) > 0 {
		index := len(params) + 1
		whereConditions = append(whereConditions, "users.role_id = ANY($"+strconv.Itoa(index)+"::bigint[])")
		params = append(params, "{"+strings.Join(filters.RoleIDs, ",")+"}")
	}
	if len(whereConditions) > 0 {
		query += " WHERE " + strings.Join(whereConditions, " AND ")
	}
	rows, err := u.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.User
		var role model.Role
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, err
		}
		user.Role = &role
		users = append(users, &user)
	}
	return users, nil
}

// withRole function gets a user as input and returns a user with the role
func (u *UserRepository) withRole(user *model.User) (*model.User, error) {
	roleRepository := NewRoleRepository(u.db)
	role, err := roleRepository.GetRoleByID(user.Role.ID)
	if err != nil {
		return nil, err
	}
	user.Role = role
	return user, nil
}
