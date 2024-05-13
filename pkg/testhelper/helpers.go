package testhelper

import (
	"github.com/akosgarai/projectregister/pkg/model"
)

// UserRepositoryMock is a mock for the UserRepository interface.
type UserRepositoryMock struct {
}

// CreateUser mocks the CreateUser method.
func (u *UserRepositoryMock) CreateUser(username, email, password string) (*model.User, error) {
	return nil, nil
}

// GetUserByEmail mocks the GetUserByEmail method.
func (u *UserRepositoryMock) GetUserByEmail(email string) (*model.User, error) {
	return nil, nil
}

// GetUserByID mocks the GetUserByID method.
func (u *UserRepositoryMock) GetUserByID(id int64) (*model.User, error) {
	return nil, nil
}

// UpdateUser mocks the UpdateUser method.
func (u *UserRepositoryMock) UpdateUser(user *model.User) error {
	return nil
}

// DeleteUser mocks the DeleteUser method.
func (u *UserRepositoryMock) DeleteUser(id int64) error {
	return nil
}

// GetUsers mocks the GetUsers method.
func (u *UserRepositoryMock) GetUsers() ([]*model.User, error) {
	return nil, nil
}
