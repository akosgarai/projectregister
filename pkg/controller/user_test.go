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

// TestUserViewControllerUserFound tests the UserViewController function.
// It creates a new controller, and calls the UserViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserViewController function with the recorder and the request.
func TestUserViewControllerUserFound(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("GET", "/admin/user/view/1", nil)
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
		"<p>View User</p>",
		"<p>ID: " + strconv.Itoa((int)(testUser.ID)) + "</p>",
		"<p>Email: " + testUser.Email + "</p>",
		"<p>Name: " + testUser.Name + "</p>",
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
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("GET", "/admin/user/view/Wrong", nil)
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
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("GET", "/admin/user/view/", nil)
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
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	missingDataError := "Missing data error"
	userRepository.Error = errors.New(missingDataError)
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("GET", "/admin/user/view/2", nil)
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
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	missingDataError := "Missing data error"
	userRepository.Error = errors.New(missingDataError)
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("GET", "/api/user/view/2", nil)
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
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("GET", "/api/user/view/1", nil)
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

// TestUserCreateViewControllerRendersTemplate tests the UserCreateViewController function.
// It creates a new controller, and calls the UserCreateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateViewController function with the recorder and the request.
func TestUserCreateViewControllerRendersTemplate(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("GET", "/admin/user/create", nil)
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
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "email@address.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/admin/user/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{""},
		"password": []string{"password"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/create", c.UserCreateViewController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestUserCreateViewControllerLongPasswd tests the UserCreateViewController function.
// It creates a new controller, and calls the UserCreateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateViewController function with the recorder and the request.
func TestUserCreateViewControllerLongPasswd(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "email@address.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/admin/user/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User update"},
		"password": []string{"passwordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpassword"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/create", c.UserCreateViewController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserCreateViewControllerSave tests the UserCreateViewController function.
// It creates a new controller, and calls the UserCreateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateViewController function with the recorder and the request.
func TestUserCreateViewControllerSave(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "email@address.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/admin/user/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User update"},
		"password": []string{"password"},
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
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "email@address.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	userRepository.Error = errors.New("Create error")
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/admin/user/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User update"},
		"password": []string{"password"},
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/admin/user/create", c.UserCreateViewController)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// TestUserCreateAPIController tests the UserCreateAPIController function.
// It creates a new controller, and calls the UserCreateAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserCreateAPIController function with the recorder and the request.
func TestUserCreateAPIController(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "email@address.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/api/user/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User"},
		"password": []string{"password"},
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
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "email@address.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	userRepository.Error = errors.New("Create error")
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/api/user/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User"},
		"password": []string{"password"},
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
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "email@address.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	userRepository.Error = errors.New("Create error")
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/api/user/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User"},
		"password": []string{"passwordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpassword"},
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

// TestUserUpdateViewControllerInvalidUserId tests the UserUpdateViewController function.
// It creates a new controller, and calls the UserUpdateViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserUpdateViewController function with the recorder and the request.
func TestUserUpdateViewControllerInvalidUserId(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("GET", "/admin/user/update/Wrong", nil)
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
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("GET", "/admin/user/update/2", nil)
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
	userRepository := &testhelper.UserRepositoryMock{}
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("GET", "/admin/user/update/1", nil)
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
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "email@address.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/admin/user/update/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email": []string{"email@address.com"},
		"name":  []string{""},
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
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "email@address.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/admin/user/update/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User update"},
		"password": []string{"passwordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpassword"},
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
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "email@address.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/admin/user/update/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User update"},
		"password": []string{"password"},
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
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "email@address.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	userRepository.UpdateUserError = errors.New("Update error")
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/admin/user/update/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User update"},
		"password": []string{"password"},
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
// It calls the UserViewController function with the recorder and the request.
func TestUserUpdateAPIControllerBadUserId(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/admin/user/update/Wrong", nil)
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
// It calls the UserViewController function with the recorder and the request.
func TestUserUpdateAPIControllerMissingUserId(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	userRepository.Error = errors.New("Missing data error")
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/admin/user/update/2", nil)
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
// It calls the UserViewController function with the recorder and the request.
func TestUserUpdateAPIControllerWrongNewPassword(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/admin/user/update/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User"},
		"password": []string{"passwordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpassword"},
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
// It calls the UserViewController function with the recorder and the request.
func TestUserUpdateAPIControllerSave(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/admin/user/update/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User"},
		"password": []string{"password"},
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
// It calls the UserViewController function with the recorder and the request.
func TestUserUpdateAPIControllerSaveError(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	userRepository.UpdateUserError = errors.New("Save error")
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("POST", "/admin/user/update/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"email":    []string{"email@address.com"},
		"name":     []string{"Test User"},
		"password": []string{"password"},
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

// TestUserDeleteAPIControllerBadUserId tests the UserDeleteAPIController function.
// It creates a new controller, and calls the UserDeleteAPIController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserViewController function with the recorder and the request.
func TestUserDeleteAPIControllerBadUserId(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	userRepository.Error = errors.New("Missing data error")
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("DELETE", "/admin/user/delete/Wrong", nil)
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
// It calls the UserViewController function with the recorder and the request.
func TestUserDeleteAPIControllerMissingUserId(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	userRepository.Error = errors.New("Missing data error")
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("DELETE", "/admin/user/delete/2", nil)
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
// It calls the UserViewController function with the recorder and the request.
func TestUserDeleteAPIControllerDelete(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	// Set the user data for the mock.
	testUser := &model.User{
		ID:    1,
		Email: "test@email.com",
		Name:  "Test User",
	}
	userRepository.LatestUser = testUser
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)

	req, err := http.NewRequest("DELETE", "/admin/user/delete/1", nil)
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

// TestUserListViewControllerError tests the UserListViewController function.
// It creates a new controller, and calls the UserListViewController function.
// The test checks the status code of the response.
// The test creates a new request with a new response recorder.
// It calls the UserListViewController function with the recorder and the request.
func TestUserListViewControllerError(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.Error = errors.New("List error")
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)
	req, err := http.NewRequest("GET", "/admin/user/list", nil)
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
	userRepository := &testhelper.UserRepositoryMock{}
	testUser := &model.User{
		ID:    1,
		Email: "one@email.com",
		Name:  "Test User 01",
	}
	testUser2 := &model.User{
		ID:    2,
		Email: "two@email.com",
		Name:  "Test User 02",
	}
	userRepository.AllUsers = []*model.User{testUser, testUser2}
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)
	req, err := http.NewRequest("GET", "/admin/user/list", nil)
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
	needles := []string{
		"<title>User List</title>",
		"<h1>User List</h1>",
		"<td>" + strconv.Itoa((int)(testUser.ID)) + "</td>",
		"<td>" + testUser.Email + "</td>",
		"<td>" + testUser.Name + "</td>",
		"<td>" + strconv.Itoa((int)(testUser2.ID)) + "</td>",
		"<td>" + testUser2.Email + "</td>",
		"<td>" + testUser2.Name + "</td>",
		"<a href=\"/admin/user/update/" + strconv.Itoa((int)(testUser.ID)) + "\">Edit</a>",
		"<a href=\"/admin/user/update/" + strconv.Itoa((int)(testUser2.ID)) + "\">Edit</a>",
		"<a href=\"/admin/user/view/" + strconv.Itoa((int)(testUser.ID)) + "\">View</a>",
		"<a href=\"/admin/user/view/" + strconv.Itoa((int)(testUser2.ID)) + "\">View</a>",
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
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)
	req, err := http.NewRequest("GET", "/api/user/list", nil)
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
	userRepository := &testhelper.UserRepositoryMock{}
	testUser := &model.User{
		ID:    1,
		Email: "01@email.com",
		Name:  "Test User 01",
	}
	testUser2 := &model.User{
		ID:    2,
		Email: "02@email.com",
		Name:  "Test User 02",
	}
	userRepository.AllUsers = []*model.User{testUser, testUser2}
	testConfig := config.NewEnvironment(testConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(userRepository, sessionStore, renderer)
	req, err := http.NewRequest("GET", "/api/user/list", nil)
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
