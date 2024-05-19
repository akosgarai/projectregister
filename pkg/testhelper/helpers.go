package testhelper

import (
	"github.com/akosgarai/projectregister/pkg/model"
)

// UserRepositoryMock is a mock for the UserRepository interface.
// It can be used to mock the UserRepository interface.
// Set the LatestUser field to the user you want to return.
// Set the AllUsers field to the list of users you want to return.
// Set the Error field to the error you want to return.
type UserRepositoryMock struct {
	LatestUser *model.User
	AllUsers   []*model.User

	Error           error
	UpdateUserError error
}

// CreateUser mocks the CreateUser method.
func (u *UserRepositoryMock) CreateUser(username, email, password string) (*model.User, error) {
	return u.LatestUser, u.Error
}

// GetUserByEmail mocks the GetUserByEmail method.
func (u *UserRepositoryMock) GetUserByEmail(email string) (*model.User, error) {
	return u.LatestUser, u.Error
}

// GetUserByID mocks the GetUserByID method.
func (u *UserRepositoryMock) GetUserByID(id int64) (*model.User, error) {
	return u.LatestUser, u.Error
}

// UpdateUser mocks the UpdateUser method.
func (u *UserRepositoryMock) UpdateUser(user *model.User) error {
	return u.UpdateUserError
}

// DeleteUser mocks the DeleteUser method.
func (u *UserRepositoryMock) DeleteUser(id int64) error {
	return u.Error
}

// GetUsers mocks the GetUsers method.
func (u *UserRepositoryMock) GetUsers() ([]*model.User, error) {
	return u.AllUsers, u.Error
}
