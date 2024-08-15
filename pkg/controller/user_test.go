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

// getViewController
func getViewController(repositoryMock *testhelper.RepositoryContainerMock) *Controller {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view"})
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

// TestUserViewControllerUserFound tests the UserViewController function.
// It creates a new controller, and calls the UserViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserViewController function with the recorder and the request.
func TestUserViewControllerUserFound(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view"})
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.LatestUser = testUser
	c := getViewController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/view/1")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/view/{userId}", c.UserViewController)
	router.ServeHTTP(rr, req)

	needles := []string{
		"<title>User Detail</title>",
		"<h1>User Detail</h1>",
		"<div class=\"label\">ID<\\/div>\\s+<div class=\"value\">\\s+<p>\\s+" + strconv.Itoa((int)(testUser.ID)) + "\\s+<\\/p>\\s+<\\/div>\\s+<\\/div>",
		"<div class=\"label\">Email<\\/div>\\s+<div class=\"value\">\\s+<p>\\s+" + testUser.Email + "\\s+<\\/p>\\s+<\\/div>\\s+<\\/div>",
		"<div class=\"label\">Name<\\/div>\\s+<div class=\"value\">\\s+<p>\\s+" + testUser.Name + "\\s+<\\/p>\\s+<\\/div>\\s+<\\/div>",
		"<div class=\"label\">Role<\\/div>\\s+<div class=\"value\">\\s+<p>\\s+" + testUser.Role.Name + "\\s+<\\/p>\\s+<\\/div>\\s+<\\/div>",
	}
	testhelper.CheckResponse(t, rr, http.StatusOK, needles)
}

// TestUserViewControllerBadUserId tests the UserViewController function.
// It creates a new controller, and calls the UserViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserViewController function with the recorder and the request.
func TestUserViewControllerBadUserId(t *testing.T) {
	c := getViewController(testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/view/Wrong")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/view/{userId}", c.UserViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusBadRequest, []string{UserFailedToGetUserErrorMessage})
}

// TestUserViewControllerMissingUserId tests the UserViewController function.
// It creates a new controller, and calls the UserViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserViewController function with the recorder and the request.
func TestUserViewControllerMissingUserId(t *testing.T) {
	c := getViewController(testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/view/")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/view/{userId}", c.UserViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusNotFound)
}

// TestUserViewControllerRepositoryError tests the UserViewController function.
// It creates a new controller, and calls the UserViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserViewController function with the recorder and the request.
func TestUserViewControllerRepositoryError(t *testing.T) {
	missingDataError := "Missing data error"
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.Error = errors.New(missingDataError)
	repositoryContainer.Users.LatestUser = testhelper.GetUserWithAccessToResources(1, []string{"users.view"})
	c := getViewController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/view/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/view/{userId}", c.UserViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{UserFailedToGetUserErrorMessage})
}

// TestUserViewAPIControllerError tests the UserViewAPIController function.
// It creates a new controller, and calls the UserViewAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserViewAPIController function with the recorder and the request.
func TestUserViewAPIControllerError(t *testing.T) {
	missingDataError := "Missing data error"
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.Error = errors.New(missingDataError)
	repositoryContainer.Users.LatestUser = testhelper.GetUserWithAccessToResources(1, []string{"users.view"})
	c := getViewController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/api/user/view/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/api/user/view/{userId}", c.UserViewAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{UserFailedToGetUserErrorMessage})
}

// TestUserViewAPIController tests the UserViewAPIController function.
// It creates a new controller, and calls the UserViewAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserViewAPIController function with the recorder and the request.
func TestUserViewAPIController(t *testing.T) {
	c := getViewController(testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/api/user/view/1")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/api/user/view/{userId}", c.UserViewAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusOK)
}

// getCreateController
func getCreateController(repositoryMock *testhelper.RepositoryContainerMock) *Controller {
	// Set the user data for the mock.
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create"})
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

// TestUserCreateViewControllerRendersTemplate tests the UserCreateViewController function.
// It creates a new controller, and calls the UserCreateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateViewController function with the recorder and the request.
func TestUserCreateViewControllerRendersTemplate(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Roles.AllRoles = &model.Roles{
		testhelper.GetRoleWithAccessToResources(1, []string{"roles.view"}),
		testhelper.GetRoleWithAccessToResources(2, []string{"roles.view"}),
	}
	c := getCreateController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/create")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/create", c.UserCreateViewController)
	router.ServeHTTP(rr, req)

	needles := []string{
		"<title>Create User</title>",
		"<h1>Create User</h1>",
		"<div class=\"form-group\">\\s+<label for=\"email\">Email<\\/label>\\s+<input type=\"email\" class=\"form-control\" id=\"email\" name=\"email\" placeholder=\"Email\" value=\"\" required >\\s+<\\/div>",
		"<div class=\"form-group\">\\s+<label for=\"name\">Name<\\/label>\\s+<input type=\"text\" class=\"form-control\" id=\"name\" name=\"name\" placeholder=\"Name\" value=\"\" required >\\s+<\\/div>",
		"<div class=\"form-group\">\\s+<label for=\"password\">Password<\\/label>\\s+<input type=\"password\" class=\"form-control\" id=\"password\" name=\"password\"  >\\s+<\\/div>",
		"<div class=\"form-group\">\\s+<label for=\"role\">Role<\\/label>\\s+<select class=\"form-control\" id=\"role\" name=\"role\" required >\\s+<option value=\"\">--Pick One--<\\/option>\\s+<option value=\"1\" >Test Role<\\/option>\\s+<option value=\"2\" >Test Role<\\/option>\\s+<\\/select>\\s+<\\/div>",
	}
	testhelper.CheckResponse(t, rr, http.StatusOK, needles)
}

// TestUserCreateViewControllerEmptyNameError tests the UserCreateViewController function.
// It creates a new controller, and calls the UserCreateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateViewController function with the recorder and the request.
func TestUserCreateViewControllerEmptyNameError(t *testing.T) {
	c := getCreateController(testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/user/create")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{""},
		"password": []string{"password"},
		"role":     []string{"1"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/create", c.UserCreateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusBadRequest, []string{UserCreateRequiredFieldMissing})
}

// TestUserCreateViewControllerLongPasswd tests the UserCreateViewController function.
// It creates a new controller, and calls the UserCreateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateViewController function with the recorder and the request.
func TestUserCreateViewControllerLongPasswd(t *testing.T) {
	c := getCreateController(testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/user/create")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User update"},
		"password": []string{"passwordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpassword"},
		"role":     []string{"1"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/create", c.UserCreateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{UserPasswordEncriptionFailedErrorMessage})
}

// TestUserCreateViewControllerSave tests the UserCreateViewController function.
// It creates a new controller, and calls the UserCreateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateViewController function with the recorder and the request.
func TestUserCreateViewControllerSave(t *testing.T) {
	c := getCreateController(testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/user/create")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User update"},
		"password": []string{"password"},
		"role":     []string{"1"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/create", c.UserCreateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusSeeOther)
}

// TestUserCreateViewControllerCreateError tests the UserCreateViewController function.
// It creates a new controller, and calls the UserCreateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateViewController function with the recorder and the request.
func TestUserCreateViewControllerCreateError(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.LatestUser = testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create", "users.update"})
	repositoryContainer.Users.Error = errors.New("Create error")
	c := getCreateController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/user/create")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User update"},
		"password": []string{"password"},
		"role":     []string{"1"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/create", c.UserCreateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{UserCreateCreateUserErrorMessagePrefix})
}

// TestUserCreateAPIController tests the UserCreateAPIController function.
// It creates a new controller, and calls the UserCreateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateAPIController function with the recorder and the request.
func TestUserCreateAPIController(t *testing.T) {
	c := getCreateController(testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/api/user/create")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User"},
		"password": []string{"password"},
		"role":     []string{"1"},
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/api/user/create", c.UserCreateAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusOK)
}

// TestUserCreateAPIControllerCreateError tests the UserCreateAPIController function.
// It creates a new controller, and calls the UserCreateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateAPIController function with the recorder and the request.
func TestUserCreateAPIControllerCreateError(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.LatestUser = testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create"})
	repositoryContainer.Users.Error = errors.New("Create error")
	c := getCreateController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/api/user/create")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User"},
		"password": []string{"password"},
		"role":     []string{"1"},
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/api/user/create", c.UserCreateAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusInternalServerError)
}

// TestUserCreateAPIControllerCreateErrorLongPwd tests the UserCreateAPIController function.
// It creates a new controller, and calls the UserCreateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateAPIController function with the recorder and the request.
func TestUserCreateAPIControllerCreateErrorLongPwd(t *testing.T) {
	c := getCreateController(testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/api/user/create")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User"},
		"password": []string{"passwordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpassword"},
		"role":     []string{"1"},
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/api/user/create", c.UserCreateAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusInternalServerError)
}

// getUpdateController
func getUpdateController(repositoryMock *testhelper.RepositoryContainerMock) *Controller {
	testUpdateUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create", "users.update"})
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	// Add the user to the session store.
	sessionStore.Set("test", session.New(testUpdateUser))
	c := New(
		repositoryMock,
		sessionStore,
		testhelper.CSVStorageMock{},
		render.NewRenderer(testConfig, render.NewTemplates()))
	c.CacheTemplates()
	return c
}

// TestUserUpdateViewControllerInvalidUserId tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerInvalidUserId(t *testing.T) {
	c := getUpdateController(testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/update/Wrong")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusBadRequest, []string{UserUserIDInvalidErrorMessagePrefix})
}

// TestUserUpdateViewControllerMissingUserId tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerMissingUserId(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.Error = errors.New("Missing data error")
	c := getUpdateController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/update/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{UserUpdateFailedToGetUserErrorMessage})
}

// TestUserUpdateViewControllerRendersTemplate tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerRendersTemplate(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Roles.AllRoles = &model.Roles{
		testhelper.GetRoleWithAccessToResources(1, []string{"roles.view"}),
		testhelper.GetRoleWithAccessToResources(2, []string{"roles.view"}),
	}
	c := getUpdateController(repositoryContainer)
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create", "users.update"})
	repositoryContainer.Users.LatestUser = testUser
	repositoryContainer.Roles.LatestRole = testUser.Role

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/update/1")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateViewController)
	router.ServeHTTP(rr, req)

	needles := []string{
		"<title>Update User</title>",
		"<h1>Update User</h1>",
		"<div class=\"form-group\">\\s+<label for=\"email\">Email<\\/label>\\s+<input type=\"email\" class=\"form-control\" id=\"email\" name=\"email\" placeholder=\"Email\" value=\"" + testUser.Email + "\" required >\\s+<\\/div>",
		"<div class=\"form-group\">\\s+<label for=\"name\">Name<\\/label>\\s+<input type=\"text\" class=\"form-control\" id=\"name\" name=\"name\" placeholder=\"Name\" value=\"" + testUser.Name + "\" required >\\s+<\\/div>",
		"<div class=\"form-group\">\\s+<label for=\"password\">Password<\\/label>\\s+<input type=\"password\" class=\"form-control\" id=\"password\" name=\"password\"  >\\s+<\\/div>",
	}
	testhelper.CheckResponse(t, rr, http.StatusOK, needles)
}

// TestUserUpdateViewControllerEmptyNameError tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerEmptyNameError(t *testing.T) {
	c := getUpdateController(testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/user/update/1")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{""},
		"password": []string{"password"},
		"role":     []string{"1"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusBadRequest, []string{UserUpdateRequiredFieldMissing})
}

// TestUserUpdateViewControllerLongPasswd tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerLongPasswd(t *testing.T) {
	c := getUpdateController(testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/user/update/1")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User update"},
		"password": []string{"passwordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpassword"},
		"role":     []string{"1"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{UserPasswordEncriptionFailedErrorMessage})
}

// TestUserUpdateViewControllerSave tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerSave(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create", "users.update"})
	repositoryContainer.Users.LatestUser = testUser
	repositoryContainer.Roles.LatestRole = testUser.Role
	c := getUpdateController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/user/update/1")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User update"},
		"password": []string{"password"},
		"role":     []string{"1"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusSeeOther)
}

// TestUserUpdateViewControllerUpdateError tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerUpdateError(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.LatestUser = testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create", "users.update"})
	repositoryContainer.Users.UpdateUserError = errors.New("Update error")
	c := getUpdateController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/user/update/1")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User update"},
		"password": []string{"password"},
		"role":     []string{"1"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{UserUpdateFailedToUpdateUserErrorMessage})
}

// TestUserUpdateAPIControllerBadUserId tests the UserUpdateAPIController function.
// It creates a new controller, and calls the UserUpdateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateAPIController function with the recorder and the request.
func TestUserUpdateAPIControllerBadUserId(t *testing.T) {
	c := getUpdateController(testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/user/update/Wrong")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusBadRequest)
}

// TestUserUpdateAPIControllerMissingUserId tests the UserUpdateAPIController function.
// It creates a new controller, and calls the UserUpdateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateAPIController function with the recorder and the request.
func TestUserUpdateAPIControllerMissingUserId(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.LatestUser = testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create", "users.update"})
	repositoryContainer.Users.Error = errors.New("Missing data error")
	c := getUpdateController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/user/update/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusInternalServerError)
}

// TestUserUpdateAPIControllerWrongNewPassword tests the UserUpdateAPIController function.
// It creates a new controller, and calls the UserUpdateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateAPIController function with the recorder and the request.
func TestUserUpdateAPIControllerWrongNewPassword(t *testing.T) {
	c := getUpdateController(testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/user/update/1")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User"},
		"password": []string{"passwordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpassword"},
		"role":     []string{"1"},
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusInternalServerError)
}

// TestUserUpdateAPIControllerSave tests the UserUpdateAPIController function.
// It creates a new controller, and calls the UserUpdateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateAPIController function with the recorder and the request.
func TestUserUpdateAPIControllerSave(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create", "users.update"})
	repositoryContainer.Users.LatestUser = testUser
	repositoryContainer.Roles.LatestRole = testUser.Role
	c := getUpdateController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/user/update/1")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User"},
		"password": []string{"password"},
		"role":     []string{"1"},
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusOK)
}

// TestUserUpdateAPIControllerSaveError tests the UserUpdateAPIController function.
// It creates a new controller, and calls the UserUpdateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateAPIController function with the recorder and the request.
func TestUserUpdateAPIControllerSaveError(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.LatestUser = testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create", "users.update"})
	repositoryContainer.Users.UpdateUserError = errors.New("Save error")
	c := getUpdateController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("POST", "/admin/user/update/1")
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User"},
		"password": []string{"password"},
		"role":     []string{"1"},
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusInternalServerError)
}

// getDeleteController
func getDeleteController(repositoryMock *testhelper.RepositoryContainerMock) *Controller {
	testDeleteUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create", "users.update", "users.delete"})
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	// Add the user to the session store.
	sessionStore.Set(testhelper.TestSessionCookieValue, session.New(testDeleteUser))
	c := New(
		repositoryMock,
		sessionStore,
		testhelper.CSVStorageMock{},
		render.NewRenderer(testConfig, render.NewTemplates()))
	c.CacheTemplates()
	return c
}

// TestUserDeleteViewControllerWrongUserId tests the UserDeleteViewController function.
// It creates a new controller, and calls the UserDeleteViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserDeleteViewController function with the recorder and the request.
func TestUserDeleteViewControllerWrongUserId(t *testing.T) {
	c := getDeleteController(testhelper.NewRepositoryContainerMock())
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/delete/Wrong")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/delete/{userId}", c.UserDeleteViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusBadRequest)
	testhelper.CheckResponse(t, rr, http.StatusBadRequest, []string{UserUserIDInvalidErrorMessagePrefix})
}

// TestUserDeleteViewControllerMissingUserId tests the UserDeleteViewController function.
// It creates a new controller, and calls the UserDeleteViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserDeleteViewController function with the recorder and the request.
func TestUserDeleteViewControllerMissingUserId(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.Error = errors.New("Missing data error")
	c := getDeleteController(repositoryContainer)
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/delete/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/delete/{userId}", c.UserDeleteViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{UserDeleteFailedToDeleteErrorMessage})
}

// TestUserDeleteViewControllerRedirects tests the UserDeleteViewController function.
// It creates a new controller, and calls the UserDeleteViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserDeleteViewController function with the recorder and the request.
func TestUserDeleteViewControllerRedirects(t *testing.T) {
	c := getDeleteController(testhelper.NewRepositoryContainerMock())
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/delete/1")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/delete/{userId}", c.UserDeleteViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusSeeOther)
}

// TestUserDeleteViewControllerDeleteError tests the UserDeleteViewController function.
// It creates a new controller, and calls the UserDeleteViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserDeleteViewController function with the recorder and the request.
func TestUserDeleteViewControllerDeleteError(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.Error = errors.New("Delete error")
	repositoryContainer.Users.LatestUser = testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create", "users.update", "users.delete"})
	c := getDeleteController(repositoryContainer)
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/delete/1")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/delete/{userId}", c.UserDeleteViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{UserDeleteFailedToDeleteErrorMessage})
}

// TestUserDeleteAPIControllerBadUserId tests the UserDeleteAPIController function.
// It creates a new controller, and calls the UserDeleteAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserDeleteAPIController function with the recorder and the request.
func TestUserDeleteAPIControllerBadUserId(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.Error = errors.New("Missing data error")
	repositoryContainer.Users.LatestUser = testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create", "users.update", "users.delete"})
	c := getDeleteController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("DELETE", "/admin/user/delete/Wrong")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/delete/{userId}", c.UserDeleteAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusBadRequest)
}

// TestUserDeleteAPIControllerMissingUserId tests the UserDeleteAPIController function.
// It creates a new controller, and calls the UserDeleteAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserDeleteAPIController function with the recorder and the request.
func TestUserDeleteAPIControllerMissingUserId(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.Error = errors.New("Missing data error")
	repositoryContainer.Users.LatestUser = testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create", "users.update", "users.delete"})
	c := getDeleteController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("DELETE", "/admin/user/delete/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/delete/{userId}", c.UserDeleteAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusInternalServerError)
}

// TestUserDeleteAPIControllerDelete tests the UserDeleteAPIController function.
// It creates a new controller, and calls the UserDeleteAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserDeleteAPIController function with the recorder and the request.
func TestUserDeleteAPIControllerDelete(t *testing.T) {
	c := getDeleteController(testhelper.NewRepositoryContainerMock())

	req, err := testhelper.NewRequestWithSessionCookie("DELETE", "/admin/user/delete/1")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// To add the vars to the context,
	// we need to create a router through which we can pass the request.
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/delete/{userId}", c.UserDeleteAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusOK)
}

func getStaticListUsers() []*model.User {
	return []*model.User{
		testhelper.GetUserWithAccessToResources(2, []string{"users.view"}),
		testhelper.GetUserWithAccessToResources(3, []string{"users.view"}),
	}
}

// getListController
func getListController(repositoryMock *testhelper.RepositoryContainerMock) *Controller {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create", "users.update"})
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

// TestUserListViewControllerError tests the UserListViewController function.
// It creates a new controller, and calls the UserListViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserListViewController function with the recorder and the request.
func TestUserListViewControllerError(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.Error = errors.New("List error")
	c := getListController(repositoryContainer)
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/list")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/list", c.UserListViewController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{UserFailedToGetUserErrorMessage})
}

// TestUserListViewController tests the UserListViewController function.
// It creates a new controller, and calls the UserListViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserListViewController function with the recorder and the request.
func TestUserListViewController(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	userList := getStaticListUsers()
	repositoryContainer.Users.AllUsers = userList
	c := getListController(repositoryContainer)

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/list")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/list", c.UserListViewController)
	router.ServeHTTP(rr, req)

	needles := []string{
		"<title>User List</title>",
		"<h1>User List</h1>",
	}
	for _, user := range userList {
		needles = append(needles, "<td>\\s+"+strconv.Itoa((int)(user.ID))+"\\s+<\\/td>\\s+<td>\\s+"+user.Name+"\\s+<\\/td>\\s+<td>\\s+"+user.Email+"\\s+<\\/td>")
		needles = append(needles, "<a href=\"/admin/user/update/"+strconv.Itoa((int)(user.ID))+"\">Update</a>")
		needles = append(needles, "<a href=\"/admin/user/view/"+strconv.Itoa((int)(user.ID))+"\">View</a>")
	}
	testhelper.CheckResponse(t, rr, http.StatusOK, needles)
}

// TestUserListAPIControllerError tests the UserListAPIController function.
// It creates a new controller, and calls the UserListAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserListAPIController function with the recorder and the request.
func TestUserListAPIControllerError(t *testing.T) {
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	repositoryContainer.Users.Error = errors.New("List error")
	c := getListController(repositoryContainer)
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/api/user/list")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/user/list", c.UserListAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusInternalServerError)
}

// TestUserListAPIController tests the UserListAPIController function.
// It creates a new controller, and calls the UserListAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserListAPIController function with the recorder and the request.
func TestUserListAPIController(t *testing.T) {
	c := getListController(testhelper.NewRepositoryContainerMock())
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/api/user/list")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/user/list", c.UserListAPIController)
	router.ServeHTTP(rr, req)

	testhelper.CheckResponseCode(t, rr, http.StatusOK)
}
