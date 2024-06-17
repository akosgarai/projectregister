package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/config"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/render"
	"github.com/akosgarai/projectregister/pkg/session"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

func getStaticViewUser() *model.User {
	return &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
		Role: &model.Role{
			ID:   1,
			Name: "Test Role",
			Resources: []model.Resource{
				model.Resource{
					ID:   1,
					Name: "users.view",
				},
			},
		},
	}
}

// getViewController
func getViewController() *Controller {
	testUser := getStaticViewUser()
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

// TestUserViewControllerUserFound tests the UserViewController function.
// It creates a new controller, and calls the UserViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserViewController function with the recorder and the request.
func TestUserViewControllerUserFound(t *testing.T) {
	testUser := getStaticViewUser()
	c := getViewController()

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

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	needles := []string{
		"<title>User View</title>",
		"<h1>User View</h1>",
		"<p>ID: " + strconv.Itoa((int)(testUser.ID)) + "</p>",
		"<p>Email: " + testUser.Email + "</p>",
		"<p>Name: " + testUser.Name + "</p>",
		"<p>Role: " + testUser.Role.Name + "</p>",
	}
	body := rr.Body.String()
	for _, needle := range needles {
		if !strings.Contains(body, needle) {
			t.Errorf("handler returned unexpected body: got %v want %v", body, needle)
		}
	}
}

// TestUserViewControllerBadUserId tests the UserViewController function.
// It creates a new controller, and calls the UserViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserViewController function with the recorder and the request.
func TestUserViewControllerBadUserId(t *testing.T) {
	c := getViewController()

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

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestUserViewControllerMissingUserId tests the UserViewController function.
// It creates a new controller, and calls the UserViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserViewController function with the recorder and the request.
func TestUserViewControllerMissingUserId(t *testing.T) {
	c := getViewController()

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

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

// TestUserViewControllerRepositoryError tests the UserViewController function.
// It creates a new controller, and calls the UserViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserViewController function with the recorder and the request.
func TestUserViewControllerRepositoryError(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.LatestUser = getStaticViewUser()
	missingDataError := "Missing data error"
	userRepository.Error = errors.New(missingDataError)
	c := getViewController()
	c.userRepository = userRepository

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

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserViewAPIControllerError tests the UserViewAPIController function.
// It creates a new controller, and calls the UserViewAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserViewAPIController function with the recorder and the request.
func TestUserViewAPIControllerError(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	missingDataError := "Missing data error"
	userRepository.Error = errors.New(missingDataError)
	userRepository.LatestUser = getStaticViewUser()
	c := getViewController()
	c.userRepository = userRepository

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

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserViewAPIController tests the UserViewAPIController function.
// It creates a new controller, and calls the UserViewAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserViewAPIController function with the recorder and the request.
func TestUserViewAPIController(t *testing.T) {
	c := getViewController()

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

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// getCreateController
func getCreateController() *Controller {
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
		Role: &model.Role{
			ID:   1,
			Name: "Test Role",
			Resources: []model.Resource{
				model.Resource{
					ID:   1,
					Name: "users.view",
				},
				model.Resource{
					ID:   2,
					Name: "users.create",
				},
			},
		},
	}
	userRepository.LatestUser = testUser
	resourceRepository := &testhelper.ResourceRepositoryMock{}
	roleRepository := &testhelper.RoleRepositoryMock{}
	roleRepository.LatestRole = testUser.Role
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	// Add the user to the session store.
	sessionStore.Set(testhelper.TestSessionCookieValue, session.New(testUser))
	renderer := render.NewRenderer(testConfig)
	return New(
		userRepository,
		roleRepository,
		resourceRepository,
		sessionStore,
		renderer,
	)
}

// TestUserCreateViewControllerRendersTemplate tests the UserCreateViewController function.
// It creates a new controller, and calls the UserCreateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateViewController function with the recorder and the request.
func TestUserCreateViewControllerRendersTemplate(t *testing.T) {
	c := getCreateController()

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/create")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/create", c.UserCreateViewController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	needles := []string{
		"<title>User Create</title>",
		"<h1>User Create</h1>",
		"<label for=\"email\">Email</label>",
		"<input type=\"email\" class=\"form-control\" id=\"email\" name=\"email\" value=\"\" required>",
		"<label for=\"name\">Name</label>",
		"<input type=\"text\" class=\"form-control\" id=\"name\" name=\"name\" value=\"\" required>",
		"<label for=\"password\">Password</label>",
		"<input type=\"password\" class=\"form-control\" id=\"password\" name=\"password\">",
	}
	body := rr.Body.String()
	for _, needle := range needles {
		if !strings.Contains(body, needle) {
			t.Errorf("handler returned unexpected body: got %v want %v", body, needle)
		}
	}
}

// TestUserCreateViewControllerEmptyNameError tests the UserCreateViewController function.
// It creates a new controller, and calls the UserCreateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateViewController function with the recorder and the request.
func TestUserCreateViewControllerEmptyNameError(t *testing.T) {
	c := getCreateController()

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

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
	body := rr.Body.String()
	if !strings.Contains(body, UserCreateRequiredFieldMissing) {
		t.Errorf("handler returned unexpected body: got %v want %v", body, UserCreateRequiredFieldMissing)
	}
}

// TestUserCreateViewControllerLongPasswd tests the UserCreateViewController function.
// It creates a new controller, and calls the UserCreateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateViewController function with the recorder and the request.
func TestUserCreateViewControllerLongPasswd(t *testing.T) {
	c := getCreateController()

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

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
	body := rr.Body.String()
	if !strings.Contains(body, UserPasswordEncriptionFailedErrorMessage) {
		t.Errorf("handler returned unexpected body: got %v want %v", body, UserPasswordEncriptionFailedErrorMessage)
	}
}

// TestUserCreateViewControllerSave tests the UserCreateViewController function.
// It creates a new controller, and calls the UserCreateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateViewController function with the recorder and the request.
func TestUserCreateViewControllerSave(t *testing.T) {
	c := getCreateController()

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

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}
}

// TestUserCreateViewControllerCreateError tests the UserCreateViewController function.
// It creates a new controller, and calls the UserCreateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateViewController function with the recorder and the request.
func TestUserCreateViewControllerCreateError(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.LatestUser = getStaticUpdateUser()
	userRepository.Error = errors.New("Create error")
	c := getCreateController()
	c.userRepository = userRepository

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

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
	body := rr.Body.String()
	if !strings.Contains(body, UserCreateCreateUserErrorMessagePrefix) {
		t.Errorf("handler returned unexpected body: got %v want %v", body, UserCreateCreateUserErrorMessagePrefix)
	}
}

// TestUserCreateAPIController tests the UserCreateAPIController function.
// It creates a new controller, and calls the UserCreateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateAPIController function with the recorder and the request.
func TestUserCreateAPIController(t *testing.T) {
	c := getCreateController()

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

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v. Response: %v",
			status, http.StatusOK, rr.Body.String())
	}
}

// TestUserCreateAPIControllerCreateError tests the UserCreateAPIController function.
// It creates a new controller, and calls the UserCreateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateAPIController function with the recorder and the request.
func TestUserCreateAPIControllerCreateError(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.LatestUser = getStaticUpdateUser()
	userRepository.Error = errors.New("Create error")
	c := getCreateController()
	c.userRepository = userRepository

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

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserCreateAPIControllerCreateErrorLongPwd tests the UserCreateAPIController function.
// It creates a new controller, and calls the UserCreateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateAPIController function with the recorder and the request.
func TestUserCreateAPIControllerCreateErrorLongPwd(t *testing.T) {
	c := getCreateController()

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

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func getStaticUpdateUser() *model.User {
	return &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
		Role: &model.Role{
			ID:   1,
			Name: "Test Role",
			Resources: []model.Resource{
				model.Resource{
					ID:   1,
					Name: "users.view",
				},
				model.Resource{
					ID:   2,
					Name: "users.create",
				},
				model.Resource{
					ID:   3,
					Name: "users.update",
				},
			},
		},
	}
}

// getUpdateController
func getUpdateController() *Controller {
	testUpdateUser := getStaticUpdateUser()
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.LatestUser = testUpdateUser
	resourceRepository := &testhelper.ResourceRepositoryMock{}
	roleRepository := &testhelper.RoleRepositoryMock{}
	roleRepository.LatestRole = testUpdateUser.Role
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	// Add the user to the session store.
	sessionStore.Set("test", session.New(testUpdateUser))
	renderer := render.NewRenderer(testConfig)
	return New(
		userRepository,
		roleRepository,
		resourceRepository,
		sessionStore,
		renderer,
	)
}

// TestUserUpdateViewControllerInvalidUserId tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerInvalidUserId(t *testing.T) {
	c := getUpdateController()

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/update/Wrong")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateViewController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestUserUpdateViewControllerMissingUserId tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerMissingUserId(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.Error = errors.New("Missing data error")
	c := getUpdateController()
	c.userRepository = userRepository

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/update/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateViewController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserUpdateViewControllerRendersTemplate tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerRendersTemplate(t *testing.T) {
	c := getUpdateController()

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/update/1")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/update/{userId}", c.UserUpdateViewController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	testUser := getStaticUpdateUser()
	needles := []string{
		"<title>User Update</title>",
		"<h1>User Update</h1>",
		"<label for=\"email\">Email</label>",
		"<input type=\"email\" class=\"form-control\" id=\"email\" name=\"email\" value=\"" + testUser.Email + "\" required>",
		"<label for=\"name\">Name</label>",
		"<input type=\"text\" class=\"form-control\" id=\"name\" name=\"name\" value=\"" + testUser.Name + "\" required>",
		"<label for=\"password\">Password</label>",
		"<input type=\"password\" class=\"form-control\" id=\"password\" name=\"password\">",
	}
	body := rr.Body.String()
	for _, needle := range needles {
		if !strings.Contains(body, needle) {
			t.Errorf("handler returned unexpected body: got %v want %v", body, needle)
		}
	}
}

// TestUserUpdateViewControllerEmptyNameError tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerEmptyNameError(t *testing.T) {
	c := getUpdateController()

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

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestUserUpdateViewControllerLongPasswd tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerLongPasswd(t *testing.T) {
	c := getUpdateController()

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

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserUpdateViewControllerSave tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerSave(t *testing.T) {
	c := getUpdateController()

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

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}
}

// TestUserUpdateViewControllerUpdateError tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerUpdateError(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.LatestUser = getStaticUpdateUser()
	userRepository.UpdateUserError = errors.New("Update error")
	c := getUpdateController()
	c.userRepository = userRepository

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

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserUpdateAPIControllerBadUserId tests the UserUpdateAPIController function.
// It creates a new controller, and calls the UserUpdateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateAPIController function with the recorder and the request.
func TestUserUpdateAPIControllerBadUserId(t *testing.T) {
	c := getUpdateController()

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

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestUserUpdateAPIControllerMissingUserId tests the UserUpdateAPIController function.
// It creates a new controller, and calls the UserUpdateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateAPIController function with the recorder and the request.
func TestUserUpdateAPIControllerMissingUserId(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.LatestUser = getStaticUpdateUser()
	userRepository.Error = errors.New("Missing data error")
	c := getUpdateController()
	c.userRepository = userRepository

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

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserUpdateAPIControllerWrongNewPassword tests the UserUpdateAPIController function.
// It creates a new controller, and calls the UserUpdateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateAPIController function with the recorder and the request.
func TestUserUpdateAPIControllerWrongNewPassword(t *testing.T) {
	c := getUpdateController()

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

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserUpdateAPIControllerSave tests the UserUpdateAPIController function.
// It creates a new controller, and calls the UserUpdateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateAPIController function with the recorder and the request.
func TestUserUpdateAPIControllerSave(t *testing.T) {
	c := getUpdateController()

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

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestUserUpdateAPIControllerSaveError tests the UserUpdateAPIController function.
// It creates a new controller, and calls the UserUpdateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateAPIController function with the recorder and the request.
func TestUserUpdateAPIControllerSaveError(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.LatestUser = getStaticUpdateUser()
	userRepository.UpdateUserError = errors.New("Save error")
	c := getUpdateController()
	c.userRepository = userRepository

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

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
func getStaticDeleteUser() *model.User {
	return &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
		Role: &model.Role{
			ID:   1,
			Name: "Test Role",
			Resources: []model.Resource{
				model.Resource{
					ID:   1,
					Name: "users.view",
				},
				model.Resource{
					ID:   2,
					Name: "users.create",
				},
				model.Resource{
					ID:   3,
					Name: "users.update",
				},
				model.Resource{
					ID:   4,
					Name: "users.delete",
				},
			},
		},
	}
}

// getDeleteController
func getDeleteController() *Controller {
	testDeleteUser := getStaticDeleteUser()
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.LatestUser = testDeleteUser
	resourceRepository := &testhelper.ResourceRepositoryMock{}
	roleRepository := &testhelper.RoleRepositoryMock{}
	roleRepository.LatestRole = testDeleteUser.Role
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	// Add the user to the session store.
	sessionStore.Set(testhelper.TestSessionCookieValue, session.New(testDeleteUser))
	renderer := render.NewRenderer(testConfig)
	return New(
		userRepository,
		roleRepository,
		resourceRepository,
		sessionStore,
		renderer,
	)
}

// TestUserDeleteViewControllerWrongUserId tests the UserDeleteViewController function.
// It creates a new controller, and calls the UserDeleteViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserDeleteViewController function with the recorder and the request.
func TestUserDeleteViewControllerWrongUserId(t *testing.T) {
	c := getDeleteController()
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/delete/Wrong")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/delete/{userId}", c.UserDeleteViewController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestUserDeleteViewControllerMissingUserId tests the UserDeleteViewController function.
// It creates a new controller, and calls the UserDeleteViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserDeleteViewController function with the recorder and the request.
func TestUserDeleteViewControllerMissingUserId(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.Error = errors.New("Missing data error")
	c := getDeleteController()
	c.userRepository = userRepository
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/delete/2")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/delete/{userId}", c.UserDeleteViewController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserDeleteViewControllerRedirects tests the UserDeleteViewController function.
// It creates a new controller, and calls the UserDeleteViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserDeleteViewController function with the recorder and the request.
func TestUserDeleteViewControllerRedirects(t *testing.T) {
	c := getDeleteController()
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/delete/1")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/delete/{userId}", c.UserDeleteViewController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}
}

// TestUserDeleteViewControllerDeleteError tests the UserDeleteViewController function.
// It creates a new controller, and calls the UserDeleteViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserDeleteViewController function with the recorder and the request.
func TestUserDeleteViewControllerDeleteError(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.LatestUser = getStaticDeleteUser()
	userRepository.Error = errors.New("Delete error")
	c := getDeleteController()
	c.userRepository = userRepository
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/delete/1")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/delete/{userId}", c.UserDeleteViewController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserDeleteAPIControllerBadUserId tests the UserDeleteAPIController function.
// It creates a new controller, and calls the UserDeleteAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserDeleteAPIController function with the recorder and the request.
func TestUserDeleteAPIControllerBadUserId(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.LatestUser = getStaticDeleteUser()
	userRepository.Error = errors.New("Missing data error")
	c := getDeleteController()
	c.userRepository = userRepository

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

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestUserDeleteAPIControllerMissingUserId tests the UserDeleteAPIController function.
// It creates a new controller, and calls the UserDeleteAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserDeleteAPIController function with the recorder and the request.
func TestUserDeleteAPIControllerMissingUserId(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.LatestUser = getStaticDeleteUser()
	userRepository.Error = errors.New("Missing data error")
	c := getDeleteController()
	c.userRepository = userRepository

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

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserDeleteAPIControllerDelete tests the UserDeleteAPIController function.
// It creates a new controller, and calls the UserDeleteAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserDeleteAPIController function with the recorder and the request.
func TestUserDeleteAPIControllerDelete(t *testing.T) {
	c := getDeleteController()

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

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func getStaticListUsers() []*model.User {
	return []*model.User{
		&model.User{
			ID:    1,
			Email: "one@email.com",
			Name:  "Test User 01",
			Role: &model.Role{
				ID:   1,
				Name: "Test Role",
				Resources: []model.Resource{
					model.Resource{
						ID:   1,
						Name: "users.view",
					},
				},
			},
		},
		&model.User{
			ID:    2,
			Email: "two@email.com",
			Name:  "Test User 02",
			Role: &model.Role{
				ID:   1,
				Name: "Test Role",
				Resources: []model.Resource{
					model.Resource{
						ID:   1,
						Name: "users.view",
					},
				},
			},
		},
	}
}

// getListController
func getListController() *Controller {
	testUser := getStaticUpdateUser()
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.AllUsers = getStaticListUsers()
	resourceRepository := &testhelper.ResourceRepositoryMock{}
	roleRepository := &testhelper.RoleRepositoryMock{}
	roleRepository.LatestRole = testUser.Role
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	// Add the user to the session store.
	sessionStore.Set(testhelper.TestSessionCookieValue, session.New(testUser))
	renderer := render.NewRenderer(testConfig)
	return New(
		userRepository,
		roleRepository,
		resourceRepository,
		sessionStore,
		renderer,
	)
}

// TestUserListViewControllerError tests the UserListViewController function.
// It creates a new controller, and calls the UserListViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserListViewController function with the recorder and the request.
func TestUserListViewControllerError(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.Error = errors.New("List error")
	c := getListController()
	c.userRepository = userRepository
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/list")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/list", c.UserListViewController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserListViewController tests the UserListViewController function.
// It creates a new controller, and calls the UserListViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserListViewController function with the recorder and the request.
func TestUserListViewController(t *testing.T) {
	c := getListController()
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/admin/user/list")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/list", c.UserListViewController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	userList := getStaticListUsers()
	needles := []string{
		"<title>User List</title>",
		"<h1>User List</h1>",
	}
	for _, user := range userList {
		needles = append(needles, "<td>"+strconv.Itoa((int)(user.ID))+"</td>")
		needles = append(needles, "<td>"+user.Email+"</td>")
		needles = append(needles, "<td>"+user.Name+"</td>")
		needles = append(needles, "<a href=\"/admin/user/update/"+strconv.Itoa((int)(user.ID))+"\">Edit</a>")
		needles = append(needles, "<a href=\"/admin/user/view/"+strconv.Itoa((int)(user.ID))+"\">View</a>")
	}
	body := rr.Body.String()
	for _, needle := range needles {
		if !strings.Contains(body, needle) {
			t.Errorf("handler returned unexpected body: got %v want %v", body, needle)
		}
	}
}

// TestUserListAPIControllerError tests the UserListAPIController function.
// It creates a new controller, and calls the UserListAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserListAPIController function with the recorder and the request.
func TestUserListAPIControllerError(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.Error = errors.New("List error")
	c := getListController()
	c.userRepository = userRepository
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/api/user/list")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/user/list", c.UserListAPIController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserListAPIController tests the UserListAPIController function.
// It creates a new controller, and calls the UserListAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserListAPIController function with the recorder and the request.
func TestUserListAPIController(t *testing.T) {
	c := getListController()
	req, err := testhelper.NewRequestWithSessionCookie("GET", "/api/user/list")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/user/list", c.UserListAPIController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
