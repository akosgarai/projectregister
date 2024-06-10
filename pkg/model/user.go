package model

// User type
type User struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt string
	UpdatedAt string
	Role      *Role

	Password string
}

// HasPrivilege checks if the user has the privilege
func (u *User) HasPrivilege(resource string) bool {
	return u.Role.HasResource(resource)
}

// UserRepository interface
type UserRepository interface {
	CreateUser(username, email, password string, roleID int64) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int64) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int64) error
	GetUsers() ([]*User, error)
}
