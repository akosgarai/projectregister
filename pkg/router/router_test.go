package router

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/config"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/session"
)

// UserRepositoryMock is a mock for the UserRepository
/*
	CreateUser(username, email, password string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int64) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int64) error
	GetUsers() ([]*User, error)
*/
type UserRepositoryMock struct{}

func (u *UserRepositoryMock) CreateUser(username, email, password string) (*model.User, error) {
	return &model.User{}, nil
}
func (u *UserRepositoryMock) GetUserByEmail(email string) (*model.User, error) {
	return &model.User{}, nil
}
func (u *UserRepositoryMock) GetUserByID(id int64) (*model.User, error) {
	return &model.User{}, nil
}
func (u *UserRepositoryMock) UpdateUser(user *model.User) error {
	return nil
}
func (u *UserRepositoryMock) DeleteUser(id int64) error {
	return nil
}
func (u *UserRepositoryMock) GetUsers() ([]*model.User, error) {
	return []*model.User{}, nil
}

// TestNew Tests the New function
// It has to return a new router
// The router has to have the given routes
// The router has to have the given controller
func TestNew(t *testing.T) {
	userRepository := &UserRepositoryMock{}
	sessionStore := session.NewStore(config.DefaultEnvironment())
	router := New(userRepository, sessionStore)
	if router == nil {
		t.Error("New router is nil")
	}
	// check the routes
	routesToCheck := []string{
		"/health",
		"/login",
		"/auth/login",

		"/admin/dashboard",
		"/admin/user/view/{userId}",

		"/api/user/create",
		"/api/user/view/{userId}",
		"/api/user/update/{userId}",
		"/api/user/delete/{userId}",
		"/api/user/list",
	}
	for _, route := range routesToCheck {
		if router.Path(route) == nil {
			t.Errorf("Route %s is missing.", route)
		}
	}
}
