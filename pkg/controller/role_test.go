package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/config"
	"github.com/akosgarai/projectregister/pkg/render"
	"github.com/akosgarai/projectregister/pkg/session"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestRoleViewController tests the RoleViewController function.
func TestRoleViewControllerWithoutPrivilege(t *testing.T) {
	// user without read access to roles
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view"})
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	// Add the user to the session store.
	sessionStore.Set(testhelper.TestSessionCookieValue, session.New(testUser))
	renderer := render.NewRenderer(testConfig)
	c := New(
		userRepository,
		&testhelper.RoleRepositoryMock{},
		&testhelper.ResourceRepositoryMock{},
		sessionStore,
		renderer,
	)
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
		// To add the vars to the context,
		// we need to create a router through which we can pass the request.
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
func getRoleViewController() *Controller {
	// user with read access to roles
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"roles.view"})
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	// Add the user to the session store.
	sessionStore.Set(testhelper.TestSessionCookieValue, session.New(testUser))
	renderer := render.NewRenderer(testConfig)
	return New(
		userRepository,
		&testhelper.RoleRepositoryMock{},
		&testhelper.ResourceRepositoryMock{},
		sessionStore,
		renderer,
	)
}

// The user has the required privilege to view the role.
// The roleID get parameter is invalid.
func TestRoleViewControllerInvalidRoleId(t *testing.T) {
	c := getRoleViewController()
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/view/invalid")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/role/view/{userId}", c.RoleViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusBadRequest, []string{RoleFailedToGetRoleErrorMessage})
}

// The user has the required privilege to view the role.
// The roleID get parameter is missing.
func TestRoleViewControllerMissingRoleId(t *testing.T) {
	c := getRoleViewController()
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/view/")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/role/view/{userId}", c.RoleViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusNotFound)
}

// The user has the required privilege to view the role.
// The roleID get parameter is valid.
// The role is missing from the db, so that the role repository returns error.
func TestRoleViewControllerRepositoryError(t *testing.T) {
	roleRepository := &testhelper.RoleRepositoryMock{}
	missingDataError := "Missing data error"
	roleRepository.Error = errors.New(missingDataError)
	c := getRoleViewController()
	c.roleRepository = roleRepository
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/view/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/role/view/{roleId}", c.RoleViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{RoleFailedToGetRoleErrorMessage, missingDataError})
}
func TestRoleViewControllerRoleViewSuccess(t *testing.T) {
	roleRepository := &testhelper.RoleRepositoryMock{}
	roleRepository.LatestRole = testhelper.GetRoleWithAccessToResources(1, []string{"roles.view"})
	c := getRoleViewController()
	c.roleRepository = roleRepository
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/role/view/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/role/view/{roleId}", c.RoleViewController)
	router.ServeHTTP(rr, req)

	needles := []string{
		"<title>Role View</title>",
		"<h1>Role View</h1>",
		"<p>ID: " + strconv.Itoa((int)(roleRepository.LatestRole.ID)) + "</p>",
		"<p>Name: " + roleRepository.LatestRole.Name + "</p>",
		"<p>Resources:</p>",
		"<li>roles.view</li>",
	}
	testhelper.CheckResponse(t, rr, http.StatusOK, needles)
}
