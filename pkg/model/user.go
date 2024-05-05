package model

// User type
type User struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt string
	UpdatedAt string

	Password string
}

// UserRepository interface
type UserRepository interface {
	CreateUser(username, email, password string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int64) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int64) error
	GetUsers() ([]*User, error)
}
