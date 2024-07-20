package testhelper

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

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

// ClientRepositoryMock is a mock for the ClientRepository interface.
// It can be used to mock the ClientRepository interface.
// Set the LatestClient field to the client you want to return.
// Set the AllClients field to the list of clients you want to return.
// Set the Error field to the error you want to return.
// Set the UpdateClientError field to the error you want to return.
type ClientRepositoryMock struct {
	LatestClient *model.Client
	AllClients   []*model.Client

	Error             error
	UpdateClientError error
}

// CreateClient mocks the CreateClient method.
func (r *ClientRepositoryMock) CreateClient(name string) (*model.Client, error) {
	return r.LatestClient, r.Error
}

// GetClientByName mocks the GetClientByName method.
func (r *ClientRepositoryMock) GetClientByName(name string) (*model.Client, error) {
	return r.LatestClient, r.Error
}

// GetClientByID mocks the GetClientByID method.
func (r *ClientRepositoryMock) GetClientByID(id int64) (*model.Client, error) {
	return r.LatestClient, r.Error
}

// UpdateClient mocks the UpdateClient method.
func (r *ClientRepositoryMock) UpdateClient(client *model.Client) error {
	return r.UpdateClientError
}

// DeleteClient mocks the DeleteClient method.
func (r *ClientRepositoryMock) DeleteClient(id int64) error {
	return r.Error
}

// GetClients mocks the GetClients method.
func (r *ClientRepositoryMock) GetClients() ([]*model.Client, error) {
	return r.AllClients, r.Error
}

// ProjectRepositoryMock is a mock for the ProjectRepository interface.
// It can be used to mock the ProjectRepository interface.
// Set the LatestProject field to the project you want to return.
// Set the AllProjects field to the list of projects you want to return.
// Set the Error field to the error you want to return.
// Set the UpdateProjectError field to the error you want to return.
type ProjectRepositoryMock struct {
	LatestProject *model.Project
	AllProjects   []*model.Project

	Error              error
	UpdateProjectError error
}

// CreateProject mocks the CreateProject method.
func (r *ProjectRepositoryMock) CreateProject(name string) (*model.Project, error) {
	return r.LatestProject, r.Error
}

// GetProjectByName mocks the GetProjectByName method.
func (r *ProjectRepositoryMock) GetProjectByName(name string) (*model.Project, error) {
	return r.LatestProject, r.Error
}

// GetProjectByID mocks the GetProjectByID method.
func (r *ProjectRepositoryMock) GetProjectByID(id int64) (*model.Project, error) {
	return r.LatestProject, r.Error
}

// UpdateProject mocks the UpdateProject method.
func (r *ProjectRepositoryMock) UpdateProject(project *model.Project) error {
	return r.UpdateProjectError
}

// DeleteProject mocks the DeleteProject method.
func (r *ProjectRepositoryMock) DeleteProject(id int64) error {
	return r.Error
}

// GetProjects mocks the GetProjects method.
func (r *ProjectRepositoryMock) GetProjects() ([]*model.Project, error) {
	return r.AllProjects, r.Error
}

// DomainRepositoryMock is a mock for the DomainRepository interface.
// It can be used to mock the DomainRepository interface.
// Set the LatestDomain field to the domain you want to return.
// Set the AllDomains field to the list of domains you want to return.
// Set the Error field to the error you want to return.
// Set the UpdateDomainError field to the error you want to return.
type DomainRepositoryMock struct {
	LatestDomain *model.Domain
	AllDomains   []*model.Domain

	Error             error
	UpdateDomainError error
}

// CreateDomain mocks the CreateDomain method.
func (r *DomainRepositoryMock) CreateDomain(name string) (*model.Domain, error) {
	return r.LatestDomain, r.Error
}

// GetDomainByName mocks the GetDomainByName method.
func (r *DomainRepositoryMock) GetDomainByName(name string) (*model.Domain, error) {
	return r.LatestDomain, r.Error
}

// GetDomainByID mocks the GetDomainByID method.
func (r *DomainRepositoryMock) GetDomainByID(id int64) (*model.Domain, error) {
	return r.LatestDomain, r.Error
}

// UpdateDomain mocks the UpdateDomain method.
func (r *DomainRepositoryMock) UpdateDomain(domain *model.Domain) error {
	return r.UpdateDomainError
}

// DeleteDomain mocks the DeleteDomain method.
func (r *DomainRepositoryMock) DeleteDomain(id int64) error {
	return r.Error
}

// GetDomains mocks the GetDomains method.
func (r *DomainRepositoryMock) GetDomains() ([]*model.Domain, error) {
	return r.AllDomains, r.Error
}

// EnvironmentRepositoryMock is a mock for the EnvironmentRepository interface.
// It can be used to mock the EnvironmentRepository interface.
// Set the LatestEnvironment field to the environment you want to return.
// Set the AllEnvironments field to the list of environments you want to return.
// Set the Error field to the error you want to return.
// Set the UpdateEnvironmentError field to the error you want to return.
type EnvironmentRepositoryMock struct {
	LatestEnvironment *model.Environment
	AllEnvironments   []*model.Environment

	Error                  error
	UpdateEnvironmentError error
}

// CreateEnvironment mocks the CreateEnvironment method.
func (r *EnvironmentRepositoryMock) CreateEnvironment(name string) (*model.Environment, error) {
	return r.LatestEnvironment, r.Error
}

// GetEnvironmentByName mocks the GetEnvironmentByName method.
func (r *EnvironmentRepositoryMock) GetEnvironmentByName(name string) (*model.Environment, error) {
	return r.LatestEnvironment, r.Error
}

// GetEnvironmentByID mocks the GetEnvironmentByID method.
func (r *EnvironmentRepositoryMock) GetEnvironmentByID(id int64) (*model.Environment, error) {
	return r.LatestEnvironment, r.Error
}

// UpdateEnvironment mocks the UpdateEnvironment method.
func (r *EnvironmentRepositoryMock) UpdateEnvironment(environment *model.Environment) error {
	return r.UpdateEnvironmentError
}

// DeleteEnvironment mocks the DeleteEnvironment method.
func (r *EnvironmentRepositoryMock) DeleteEnvironment(id int64) error {
	return r.Error
}

// GetEnvironments mocks the GetEnvironments method.
func (r *EnvironmentRepositoryMock) GetEnvironments() ([]*model.Environment, error) {
	return r.AllEnvironments, r.Error
}

// RuntimeRepositoryMock is a mock for the RuntimeRepository interface.
// It can be used to mock the RuntimeRepository interface.
// Set the LatestRuntime field to the runtime you want to return.
// Set the AllRuntimes field to the list of runtimes you want to return.
// Set the Error field to the error you want to return.
// Set the UpdateRuntimeError field to the error you want to return.
type RuntimeRepositoryMock struct {
	LatestRuntime *model.Runtime
	AllRuntimes   []*model.Runtime

	Error              error
	UpdateRuntimeError error
}

// CreateRuntime mocks the CreateRuntime method.
func (r *RuntimeRepositoryMock) CreateRuntime(name string) (*model.Runtime, error) {
	return r.LatestRuntime, r.Error
}

// GetRuntimeByName mocks the GetRuntimeByName method.
func (r *RuntimeRepositoryMock) GetRuntimeByName(name string) (*model.Runtime, error) {
	return r.LatestRuntime, r.Error
}

// GetRuntimeByID mocks the GetRuntimeByID method.
func (r *RuntimeRepositoryMock) GetRuntimeByID(id int64) (*model.Runtime, error) {
	return r.LatestRuntime, r.Error
}

// UpdateRuntime mocks the UpdateRuntime method.
func (r *RuntimeRepositoryMock) UpdateRuntime(runtime *model.Runtime) error {
	return r.UpdateRuntimeError
}

// DeleteRuntime mocks the DeleteRuntime method.
func (r *RuntimeRepositoryMock) DeleteRuntime(id int64) error {
	return r.Error
}

// GetRuntimes mocks the GetRuntimes method.
func (r *RuntimeRepositoryMock) GetRuntimes() ([]*model.Runtime, error) {
	return r.AllRuntimes, r.Error
}

// PoolRepositoryMock is a mock for the PoolRepository interface.
// It can be used to mock the PoolRepository interface.
// Set the LatestPool field to the pool you want to return.
// Set the AllPools field to the list of pools you want to return.
// Set the Error field to the error you want to return.
// Set the UpdatePoolError field to the error you want to return.
type PoolRepositoryMock struct {
	LatestPool *model.Pool
	AllPools   []*model.Pool

	Error           error
	UpdatePoolError error
}

// CreatePool mocks the CreatePool method.
func (r *PoolRepositoryMock) CreatePool(name string) (*model.Pool, error) {
	return r.LatestPool, r.Error
}

// GetPoolByName mocks the GetPoolByName method.
func (r *PoolRepositoryMock) GetPoolByName(name string) (*model.Pool, error) {
	return r.LatestPool, r.Error
}

// GetPoolByID mocks the GetPoolByID method.
func (r *PoolRepositoryMock) GetPoolByID(id int64) (*model.Pool, error) {
	return r.LatestPool, r.Error
}

// UpdatePool mocks the UpdatePool method.
func (r *PoolRepositoryMock) UpdatePool(pool *model.Pool) error {
	return r.UpdatePoolError
}

// DeletePool mocks the DeletePool method.
func (r *PoolRepositoryMock) DeletePool(id int64) error {
	return r.Error
}

// GetPools mocks the GetPools method.
func (r *PoolRepositoryMock) GetPools() ([]*model.Pool, error) {
	return r.AllPools, r.Error
}

// DatabaseRepositoryMock is a mock for the DatabaseRepository interface.
// It can be used to mock the DatabaseRepository interface.
// Set the LatestDatabase field to the database you want to return.
// Set the AllDatabases field to the list of databases you want to return.
// Set the Error field to the error you want to return.
// Set the UpdateDatabaseError field to the error you want to return.
type DatabaseRepositoryMock struct {
	LatestDatabase *model.Database
	AllDatabases   []*model.Database

	Error               error
	UpdateDatabaseError error
}

// CreateDatabase mocks the CreateDatabase method.
func (r *DatabaseRepositoryMock) CreateDatabase(name string) (*model.Database, error) {
	return r.LatestDatabase, r.Error
}

// GetDatabaseByName mocks the GetDatabaseByName method.
func (r *DatabaseRepositoryMock) GetDatabaseByName(name string) (*model.Database, error) {
	return r.LatestDatabase, r.Error
}

// GetDatabaseByID mocks the GetDatabaseByID method.
func (r *DatabaseRepositoryMock) GetDatabaseByID(id int64) (*model.Database, error) {
	return r.LatestDatabase, r.Error
}

// UpdateDatabase mocks the UpdateDatabase method.
func (r *DatabaseRepositoryMock) UpdateDatabase(database *model.Database) error {
	return r.UpdateDatabaseError
}

// DeleteDatabase mocks the DeleteDatabase method.
func (r *DatabaseRepositoryMock) DeleteDatabase(id int64) error {
	return r.Error
}

// GetDatabases mocks the GetDatabases method.
func (r *DatabaseRepositoryMock) GetDatabases() ([]*model.Database, error) {
	return r.AllDatabases, r.Error
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

// CheckBodyContains checks if the body contains the needles.
func CheckBodyContains(t *testing.T, body string, needles []string) {
	for _, needle := range needles {
		if !strings.Contains(body, needle) {
			t.Errorf("Missing needle in the body: %s / %s", needle, body)
		}
	}
}

// CheckResponseCode checks if the response code is the expected.
func CheckResponseCode(t *testing.T, rr *httptest.ResponseRecorder, expectedCode int) {
	if rr.Code != expectedCode {
		t.Errorf("The response code is not correct. Expected: %d, got: %d", expectedCode, rr.Code)
	}
}

// CheckResponse check if the response code is the expected and the body contains the needles.
func CheckResponse(t *testing.T, rr *httptest.ResponseRecorder, expectedCode int, needles []string) {
	CheckResponseCode(t, rr, expectedCode)
	CheckBodyContains(t, rr.Body.String(), needles)
}

// GetUserWithAccessToResources returns a user with access to the given resources.
// The user has the given ID, and the given resource names.
// The email of the user has the format of test+ID@email.com.
// The name of the user is Test User ID.
// The role of the user is Test Role with 1 as role ID.
func GetUserWithAccessToResources(userID int, resourceNames []string) *model.User {
	email := "test" + strconv.Itoa(userID) + "@email.com"
	user := &model.User{
		ID:    int64(userID),
		Email: email,
		Name:  "Test User " + strconv.Itoa(userID),
		Role:  GetRoleWithAccessToResources(1, resourceNames),
	}
	return user
}

// GetRoleWithAccessToResources returns a role with access to the given resources.
// The role has the given ID, and the given resource names.
func GetRoleWithAccessToResources(roleID int, resourceNames []string) *model.Role {
	role := &model.Role{
		ID:        int64(roleID),
		Name:      "Test Role",
		Resources: []*model.Resource{},
	}
	for _, resourceName := range resourceNames {
		id := int64(len(role.Resources) + 1)
		resource := &model.Resource{Name: resourceName, ID: id}
		role.Resources = append(role.Resources, resource)
	}
	return role
}
