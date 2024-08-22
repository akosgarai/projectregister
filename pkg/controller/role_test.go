package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/config"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/render"
	"github.com/akosgarai/projectregister/pkg/session"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

var (
	viewRoleResources   = []string{"roles.view"}
	createRoleResources = []string{"roles.view", "roles.create"}
	updateRoleResources = []string{"roles.view", "roles.create", "roles.update"}
	deleteRoleResources = []string{"roles.view", "roles.create", "roles.update", "roles.delete"}
)

// TestRoleViewController tests the RoleViewController function.
func TestRoleViewControllerWithoutPrivilege(t *testing.T) {
	// user without read access to roles
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view"})
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.LatestUser = testUser
	sessionStore := session.NewStore(testConfig)
	// Add the user to the session store.
	sessionStore.Set(testhelper.TestSessionCookieValue, session.New(testUser))
	c := New(
		testhelper.NewRepositoryContainerMock(),
		sessionStore,
		testhelper.CSVStorageMock{},
		render.NewRenderer(testConfig, render.NewTemplates()))

	testData := []struct {
		Method       string
		Route        string
		RoutePattern string
		Handler      func(http.ResponseWriter, *http.Request)
	}{
		{"GET", "/admin/role/view/1", "/admin/role/view/{userId}", c.RoleViewController},
		{"GET", "/admin/role/create", "/admin/role/create", c.RoleCreateViewController},
		{"POST", "/admin/role/create", "/admin/role/create", c.RoleCreateViewController},
		{"GET", "/admin/role/update/1", "/admin/role/update/{userId}", c.RoleUpdateViewController},
		{"POST", "/admin/role/update/1", "/admin/role/update/{userId}", c.RoleUpdateViewController},
		{"POST", "/admin/role/delete/1", "/admin/role/delete/{userId}", c.RoleDeleteViewController},
		{"GET", "/admin/role/list", "/admin/role/list", c.RoleListViewController},
	}
	for _, d := range testData {
		req, err := testhelper.NewRequestWithSessionCookie(d.Method, d.Route)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc(d.RoutePattern, d.Handler)
		router.ServeHTTP(rr, req)

		needles := []string{
			"Forbidden",
		}
		testhelper.CheckResponse(t, rr, http.StatusForbidden, needles)
	}
}

// getRoleViewController
// It returns a controller with a user who has the required privilege to view the role.
func getRoleViewController(resources []string, repositoryMock *testhelper.RepositoryContainerMock) *Controller {
	testUser := testhelper.GetUserWithAccessToResources(1, resources)
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	// Add the user to the session store.
	sessionStore.Set(testhelper.TestSessionCookieValue, session.New(testUser))
	c := New(
		repositoryMock,
		sessionStore,
		testhelper.CSVStorageMock{},
		render.NewRenderer(testConfig, render.NewTemplates()))
	c.CacheTemplates()
	return c
}

// The user has the required privilege to view the role.
// The roleID get parameter is invalid.
func TestRoleViewControllerInvalidRoleId(t *testing.T) {
	c := getRoleViewController(viewRoleResources, testhelper.NewRepositoryContainerMock())
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/view/invalid")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/view/{userId}", c.RoleViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusBadRequest, []string{RoleFailedToGetRoleErrorMessage})
}

// The user has the required privilege to view the role.
// The roleID get parameter is missing.
func TestRoleViewControllerMissingRoleId(t *testing.T) {
	c := getRoleViewController(viewRoleResources, testhelper.NewRepositoryContainerMock())
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/view/")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/view/{userId}", c.RoleViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusNotFound)
}

// The user has the required privilege to view the role.
// The roleID get parameter is valid.
// The role is missing from the db, so that the role repository returns error.
func TestRoleViewControllerRepositoryError(t *testing.T) {
	missingDataError := "Missing data error"
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.Error = errors.New(missingDataError)

	c := getRoleViewController(viewRoleResources, repositoryMock)
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/view/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/view/{roleId}", c.RoleViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{RoleFailedToGetRoleErrorMessage, missingDataError})
}

// The user has the required privilege to view the role.
// The roleID get parameter is valid.
// The role has been found, the view template has to be returned.
func TestRoleViewControllerRoleViewSuccess(t *testing.T) {
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.LatestRole = testhelper.GetRoleWithAccessToResources(1, []string{"roles.view"})
	c := getRoleViewController(viewRoleResources, repositoryMock)
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/view/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/view/{roleId}", c.RoleViewController)
	router.ServeHTTP(rr, req)

	needles := []string{
		"<title>Role Detail</title>",
		"<h1>Role Detail</h1>",
		"<div class=\"label\">ID<\\/div>\\s+<div class=\"value\">\\s+<p>\\s+" + strconv.Itoa((int)(repositoryMock.Roles.LatestRole.ID)) + "\\s+<\\/p>\\s+<\\/div>\\s+<\\/div>",
		"<div class=\"label\">Name<\\/div>\\s+<div class=\"value\">\\s+<p>\\s+" + repositoryMock.Roles.LatestRole.Name + "\\s+<\\/p>\\s+<\\/div>\\s+<\\/div>",
		"<div class=\"label\">Resources<\\/div>\\s+<div class=\"value\">\\s+<p>\\s+roles.view\\s+<\\/p>\\s+<\\/div>\\s+<\\/div>",
	}
	testhelper.CheckResponse(t, rr, http.StatusOK, needles)
}

// <div class=\"label\">ID<\/div>\\n\\s+<div class=\"value\">\\s+<p>\\s+" + 1 + "\\s+<\/p>\\s+<\/div>\\s+<\/div>

// The user has the required privilege to create roles.
// The resource repository returns error.
func TestRoleCreateViewControllerResourceRepositoryError(t *testing.T) {
	errorMessage := "Resource repository error"
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Resources.Error = errors.New(errorMessage)
	c := getRoleViewController(createRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/create")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/create", c.RoleCreateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{RoleFailedToGetResourcesErrorMessage, errorMessage})
}

// The user has the required privilege to create roles.
// The resource repository returns the resources.
func TestRoleCreateViewControllerRendersTemplate(t *testing.T) {
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.LatestRole = testhelper.GetRoleWithAccessToResources(1, []string{"roles.view"})
	repositoryMock.Resources.AllResources = &repositoryMock.Roles.LatestRole.Resources
	c := getRoleViewController(createRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/create")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/create", c.RoleCreateViewController)
	router.ServeHTTP(rr, req)

	needles := []string{
		"<title>Create Role</title>",
		"<h1>Create Role</h1>",
		"<div class=\"form-group\">\\s+<label for=\"name\">Name<\\/label>\\s+<input type=\"text\" class=\"form-control\" id=\"name\" name=\"name\" placeholder=\"Name\" value=\"\" required >\\s+<\\/div>",
		"<div class=\"form-group\">\\s+<label for=\"resources\">Resources<\\/label>\\s+<div class=\"checkbox\">\\s+<input type=\"checkbox\" id=\"r_1\" name=\"resources\" value=\"1\" >\\s+<label for=\"r_1\">roles.view<\\/label>\\s+<\\/div>\\s+<\\/div>",
	}
	testhelper.CheckResponse(t, rr, http.StatusOK, needles)
}

// The user has the required privilege to create roles.
// The resource repository returns the resources.
// The name parameter is missing from the request.
func TestRoleCreateViewControllerMissingRequiredParameter(t *testing.T) {
	c := getRoleViewController(createRoleResources, testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/role/create")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"name":      []string{""},
		"resources": []string{"1"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/role/create", c.RoleCreateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusBadRequest, []string{RoleCreateRequiredFieldMissing})
}

// The user has the required privilege to create roles.
// The resource repository returns the resources.
// The name parameter is valid, the resource is invalid.
func TestRoleCreateViewControllerWrongResourceID(t *testing.T) {
	c := getRoleViewController(createRoleResources, testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/role/create")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"name":      []string{"New Role"},
		"resources": []string{"wrong"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/role/create", c.RoleCreateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusBadRequest, []string{RoleResourceIDInvalidErrorMessage})
}

// The user has the required privilege to create roles.
// The resource repository returns the resources.
// The name and resource parameters are valid.
// The role repository returns error.
func TestRoleCreateViewControllerCreateRoleRepositoryError(t *testing.T) {
	errorMessage := "Create role repository error"
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.Error = errors.New(errorMessage)
	c := getRoleViewController(createRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/role/create")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"name":      []string{"New Role"},
		"resources": []string{"1"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/role/create", c.RoleCreateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{RoleCreateCreateRoleErrorMessage, errorMessage})
}

// The user has the required privilege to create roles.
// The resource repository returns the resources.
// The name and resource parameters are valid.
// The role repository does not return error.
// redirect to the list page.
func TestRoleCreateViewControllerSuccess(t *testing.T) {
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.LatestRole = testhelper.GetRoleWithAccessToResources(2, []string{"roles.view"})
	c := getRoleViewController(createRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/role/create")
	if err != nil {
		t.Fatal(err)
	}
	// cast the resouce id to string
	req.Form = map[string][]string{
		"name":      []string{repositoryMock.Roles.LatestRole.Name},
		"resources": []string{strconv.Itoa((int)(repositoryMock.Roles.LatestRole.Resources[0].ID))},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/role/create", c.RoleCreateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusSeeOther)
}

// The roleID get parameter is invalid.
func TestRoleUpdateViewControllerWrongRoleID(t *testing.T) {
	c := getRoleViewController(updateRoleResources, testhelper.NewRepositoryContainerMock())
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/update/invalid")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/update/{userId}", c.RoleUpdateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusBadRequest, []string{RoleRoleIDInvalidErrorMessage})
}

// The user has the required privilege to view the role.
// The roleID get parameter is missing.
func TestRoleUpdateViewControllerMissingRoleId(t *testing.T) {
	c := getRoleViewController(updateRoleResources, testhelper.NewRepositoryContainerMock())
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/update/")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/update/{userId}", c.RoleUpdateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusNotFound)
}

// The user has the required privilege to update the role.
// The roleID get parameter is valid.
// The role is missing from the db, so that the role repository returns error.
func TestRoleUpdateViewControllerGetRoleRepositoryError(t *testing.T) {
	missingDataError := "Missing data error"
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.Error = errors.New(missingDataError)
	c := getRoleViewController(updateRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/update/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/update/{roleId}", c.RoleUpdateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{RoleFailedToGetRoleErrorMessage, missingDataError})
}

// The user has the required privilege to update the role.
// The roleID get parameter is valid.
// The role is found, but the resource repository returns error.
func TestRoleUpdateViewControllerResourceRepositoryError(t *testing.T) {
	missingDataError := "Missing data error"
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.LatestRole = testhelper.GetRoleWithAccessToResources(2, []string{"roles.view"})
	repositoryMock.Resources.Error = errors.New(missingDataError)

	c := getRoleViewController(updateRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/update/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/update/{roleId}", c.RoleUpdateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{RoleFailedToGetResourcesErrorMessage, missingDataError})
}

// The user has the required privilege to update the role.
// The roleID get parameter is valid.
// The role repository returns the role.
// The resource repository returns the resources.
func TestRoleUpdateViewControllerRendersTemplate(t *testing.T) {
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.LatestRole = testhelper.GetRoleWithAccessToResources(1, []string{"roles.view"})
	repositoryMock.Resources.AllResources = &repositoryMock.Roles.LatestRole.Resources

	c := getRoleViewController(updateRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/update/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/update/{roleId}", c.RoleUpdateViewController)
	router.ServeHTTP(rr, req)

	needles := []string{
		"<title>Update Role</title>",
		"<h1>Update Role</h1>",
		"<div class=\"form-group\">\\s+<label for=\"name\">Name<\\/label>\\s+<input type=\"text\" class=\"form-control\" id=\"name\" name=\"name\" placeholder=\"Name\" value=\"Test Role\" required >\\s+<\\/div>",
		"<div class=\"form-group\">\\s+<label for=\"resources\">Resources<\\/label>\\s+<div class=\"checkbox\">\\s+<input type=\"checkbox\" id=\"r_1\" name=\"resources\" value=\"1\" checked>\\s+<label for=\"r_1\">roles.view<\\/label>\\s+<\\/div>\\s+<\\/div>",
	}
	testhelper.CheckResponse(t, rr, http.StatusOK, needles)
}

// The user has the required privilege to update the role.
// The roleID get parameter is valid.
// The role repository returns the role.
// The name request parameter is missing.
func TestRoleUpdateViewControllerMissingRequiredParameter(t *testing.T) {
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.LatestRole = testhelper.GetRoleWithAccessToResources(1, []string{"roles.view"})
	c := getRoleViewController(updateRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/role/update/1")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"name":      []string{""},
		"resources": []string{"1"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/role/update/{roleId}", c.RoleUpdateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusBadRequest, []string{RoleUpdateRequiredFieldMissing})
}

// The user has the required privilege to update the role.
// The roleID get parameter is valid.
// The role repository returns the role.
// The name request parameter is valid, but the resource is invalid.
func TestRoleUpdateViewControllerWrongResourceID(t *testing.T) {
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.LatestRole = testhelper.GetRoleWithAccessToResources(1, []string{"roles.view"})
	c := getRoleViewController(updateRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/role/update/1")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"name":      []string{"New Role"},
		"resources": []string{"invalid"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/role/update/{roleId}", c.RoleUpdateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusBadRequest, []string{RoleResourceIDInvalidErrorMessage})
}

// The user has the required privilege to update the role.
// The roleID get parameter is valid.
// The role repository returns the role.
// The name and resource request parameters are valid.
// The role repository returns error.
func TestRoleUpdateViewControllerUpdateRoleRepositoryError(t *testing.T) {
	errorMessage := "Update role repository error"
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.LatestRole = testhelper.GetRoleWithAccessToResources(1, []string{"roles.view"})
	repositoryMock.Roles.UpdateRoleError = errors.New(errorMessage)
	c := getRoleViewController(updateRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/role/update/1")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"name":      []string{"New Role"},
		"resources": []string{"1"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/role/update/{roleId}", c.RoleUpdateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{RoleUpdateUpdateRoleErrorMessage, errorMessage})
}

// The user has the required privilege to update the role.
// The roleID get parameter is valid.
// The role repository returns the role.
// The name and resource request parameters are valid.
// The role repository does not return error.
// It redirects to the list page.
func TestRoleUpdateViewControllerSuccess(t *testing.T) {
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.LatestRole = testhelper.GetRoleWithAccessToResources(1, []string{"roles.view"})
	c := getRoleViewController(updateRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/role/update/1")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"name":      []string{"New Role"},
		"resources": []string{"1"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/role/update/{roleId}", c.RoleUpdateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusSeeOther)
}

// The roleID get parameter is invalid.
func TestRoleDeleteViewControllerWrongRoleID(t *testing.T) {
	c := getRoleViewController(deleteRoleResources, testhelper.NewRepositoryContainerMock())
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/delete/invalid")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/delete/{userId}", c.RoleDeleteViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusBadRequest, []string{RoleRoleIDInvalidErrorMessage})
}

// The user has the required privilege to delete the role.
// The roleID get parameter is missing.
func TestRoleDeleteViewControllerMissingRoleId(t *testing.T) {
	c := getRoleViewController(deleteRoleResources, testhelper.NewRepositoryContainerMock())
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/delete/")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/delete/{userId}", c.RoleDeleteViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusNotFound)
}

// The user has the required privilege to delete the role.
// The roleID get parameter is valid.
// The role repository returns error.
func TestRoleDeleteViewControllerDeleteRoleRepositoryError(t *testing.T) {
	errorMessage := "Update role repository error"
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.Error = errors.New(errorMessage)
	c := getRoleViewController(deleteRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/role/delete/1")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/delete/{roleId}", c.RoleDeleteViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{RoleDeleteFailedToDeleteErrorMessage})
}

// The user has the required privilege to delete the role.
// The roleID get parameter is valid.
// The role repository does not return error.
// It redirects to the list page
func TestRoleDeleteViewControllerSuccess(t *testing.T) {
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.LatestRole = testhelper.GetRoleWithAccessToResources(1, []string{"roles.view"})
	c := getRoleViewController(deleteRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/role/delete/1")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/delete/{roleId}", c.RoleDeleteViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusSeeOther)
}

// The user has the required privilege to view the role.
// The role repository returns error.
func TestRoleListViewControllerGetRolesRepositoryError(t *testing.T) {
	errorMessage := "Update role repository error"
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.Error = errors.New(errorMessage)
	c := getRoleViewController(viewRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/list")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/list", c.RoleListViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{RoleListFailedToGetRolesErrorMessage})
}

// The user has the required privilege to view the role.
// The role repository returns the list.
func TestRoleListViewControllerRendersTemplate(t *testing.T) {
	repositoryMock := testhelper.NewRepositoryContainerMock()
	repositoryMock.Roles.AllRoles = &model.Roles{
		testhelper.GetRoleWithAccessToResources(1, []string{"roles.view"}),
		testhelper.GetRoleWithAccessToResources(2, []string{"roles.view"}),
	}
	c := getRoleViewController(deleteRoleResources, repositoryMock)

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/list")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/admin/role/list", c.RoleListViewController)
	router.ServeHTTP(rr, req)

	needles := []string{
		"<title>Role List</title>",
		"<h1>Role List</h1>",
	}
	for _, role := range *repositoryMock.Roles.AllRoles {
		needles = append(needles, "<td>\\s+"+strconv.Itoa((int)(role.ID))+"\\s+<\\/td>")
		needles = append(needles, "<td>\\s+"+role.Name+"\\s+<\\/td>")
		needles = append(needles, "<a href=\"/admin/role/update/"+strconv.Itoa((int)(role.ID))+"\">Update</a>")
		needles = append(needles, "<a href=\"/admin/role/view/"+strconv.Itoa((int)(role.ID))+"\">View</a>")
		needles = append(needles, "<form action=\"/admin/role/delete/"+strconv.Itoa((int)(role.ID))+"\" method=\"post\" class=\"form-link\">\\s+<input type=\"submit\" value=\"Delete\">\\s+<\\/form>")
	}
	testhelper.CheckResponse(t, rr, http.StatusOK, needles)
}
