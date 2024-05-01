package repository

import (
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
func (u *UserRepository) CreateUser(username, email, password string) (model.User, error) {
	var user model.User
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *"
	err := u.db.QueryRow(query, username, email, password).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

// GetUserByEmail gets a user by email
func (u *UserRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	query := "SELECT * FROM users WHERE email = $1"
	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

// GetUserByID gets a user by id
func (u *UserRepository) GetUserByID(id int64) (model.User, error) {
	// get user by id
	var user model.User
	query := "SELECT * FROM users WHERE id = $1"
	err := u.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

// UpdateUser updates a user
func (u *UserRepository) UpdateUser(user model.User) error {
	query := "UPDATE users SET name = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5"
	now := time.Now().Format("2006-01-02 15:04:05")
	_, err := u.db.Exec(query, user.Name, user.Email, user.Password, now, user.ID)
	return err
}

// DeleteUser deletes a user
func (u *UserRepository) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := u.db.Exec(query, id)
	return err
}

// GetUsers gets all users
func (u *UserRepository) GetUsers() ([]model.User, error) {
	// get all users
	var users []model.User
	query := "SELECT * FROM users"
	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
