package testhelper

import (
	"net/http"

	"github.com/akosgarai/projectregister/pkg/model"
)

var (
	// TestSessionCookieName is the name of the session cookie that we use for testing.
	TestSessionCookieName = "session"
	// TestSessionCookieValue is the value of the session cookie that we use for testing.
	TestSessionCookieValue = "test"
	// TestConfigData is the test configuration data with the minimal setup to render pages.
	TestConfigData = map[string]string{
		"RENDER_TEMPLATE_DIRECTORY_PATH": "../../web/template",
	}
)

// UserRepositoryMock is a mock for the UserRepository interface.
// It can be used to mock the UserRepository interface.
// Set the LatestUser field to the user you want to return.
// Set the AllUsers field to the list of users you want to return.
// Set the Error field to the error you want to return.
// Set the UpdateUserError field to the error you want to return.
type UserRepositoryMock struct {
	LatestUser *model.User
	AllUsers   []*model.User

	Error           error
	UpdateUserError error
}

// CreateUser mocks the CreateUser method.
func (u *UserRepositoryMock) CreateUser(username, email, password string, roleID int64) (*model.User, error) {
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

// RoleRepositoryMock is a mock for the RoleRepository interface.
// It can be used to mock the RoleRepository interface.
// Set the LatestRole field to the role you want to return.
// Set the AllRoles field to the list of roles you want to return.
// Set the Error field to the error you want to return.
// Set the UpdateRoleError field to the error you want to return.
type RoleRepositoryMock struct {
	LatestRole *model.Role
	AllRoles   []*model.Role

	Error           error
	UpdateRoleError error
}

// CreateRole mocks the CreateRole method.
func (r *RoleRepositoryMock) CreateRole(name string, resourceIDs []int64) (*model.Role, error) {
	return r.LatestRole, r.Error
}

// GetRoleByName mocks the GetRoleByName method.
func (r *RoleRepositoryMock) GetRoleByName(name string) (*model.Role, error) {
	return r.LatestRole, r.Error
}

// GetRoleByID mocks the GetRoleByID method.
func (r *RoleRepositoryMock) GetRoleByID(id int64) (*model.Role, error) {
	return r.LatestRole, r.Error
}

// UpdateRole mocks the UpdateRole method.
func (r *RoleRepositoryMock) UpdateRole(role *model.Role, resourceIDs []int64) error {
	return r.UpdateRoleError
}

// DeleteRole mocks the DeleteRole method.
func (r *RoleRepositoryMock) DeleteRole(id int64) error {
	return r.Error
}

// GetRoles mocks the GetRoles method.
func (r *RoleRepositoryMock) GetRoles() ([]*model.Role, error) {
	return r.AllRoles, r.Error
}

// ResourceRepositoryMock is a mock for the ResourceRepository interface.
// It can be used to mock the ResourceRepository interface.
// Set the LatestResource field to the resource you want to return.
// Set the AllResources field to the list of resources you want to return.
// Set the Error field to the error you want to return.
// Set the UpdateResourceError field to the error you want to return.
type ResourceRepositoryMock struct {
	LatestResource *model.Resource
	AllResources   []*model.Resource

	Error               error
	UpdateResourceError error
}

// CreateResource mocks the CreateResource method.
func (r *ResourceRepositoryMock) CreateResource(name string) (*model.Resource, error) {
	return r.LatestResource, r.Error
}

// GetResourceByName mocks the GetResourceByName method.
func (r *ResourceRepositoryMock) GetResourceByName(name string) (*model.Resource, error) {
	return r.LatestResource, r.Error
}

// GetResourceByID mocks the GetResourceByID method.
func (r *ResourceRepositoryMock) GetResourceByID(id int64) (*model.Resource, error) {
	return r.LatestResource, r.Error
}

// UpdateResource mocks the UpdateResource method.
func (r *ResourceRepositoryMock) UpdateResource(resource *model.Resource) error {
	return r.UpdateResourceError
}

// DeleteResource mocks the DeleteResource method.
func (r *ResourceRepositoryMock) DeleteResource(id int64) error {
	return r.Error
}

// GetResources mocks the GetResources method.
func (r *ResourceRepositoryMock) GetResources() ([]*model.Resource, error) {
	return r.AllResources, r.Error
}

// NewRequestWithSessionCookie creates a new request with the session cookie.
func NewRequestWithSessionCookie(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	// set the session cookie
	req.AddCookie(&http.Cookie{
		Name:  TestSessionCookieName,
		Value: TestSessionCookieValue,
	})
	return req, nil
}
