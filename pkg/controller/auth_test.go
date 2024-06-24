package controller

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akosgarai/projectregister/pkg/config"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/passwd"
	"github.com/akosgarai/projectregister/pkg/render"
	"github.com/akosgarai/projectregister/pkg/session"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

func getNewAuthController() *Controller {
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	return New(
		&testhelper.UserRepositoryMock{},
		&testhelper.RoleRepositoryMock{},
		&testhelper.ResourceRepositoryMock{},
		sessionStore,
		renderer)
}

// TestLoginPageControllerWithoutSession tests the LoginPageController function without session.
// It creates a new controller, and a new request with a new response recorder.
// It calls the LoginPageController function with the recorder and the request.
// It checks the status code of the response.
func TestLoginPageControllerWithoutSession(t *testing.T) {
	c := getNewAuthController()

	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.LoginPageController)

	handler.ServeHTTP(rr, req)

	// Check the response body is what we expect.
	// The body has to contain the <title>Login</title> tag.
	// The body has to contain a text input with the name 'username'.
	// The body has to contain a password input with the name 'password'.
	// The body has to contain a submit input.
	needles := []string{
		"<title>Login</title>",
		"<input type=\"text\" name=\"username\"",
		"<input type=\"password\" name=\"password\"",
		"<input type=\"submit\"",
	}
	testhelper.CheckResponse(t, rr, http.StatusOK, needles)
}

// TestLoginPageControllerWithSession tests the LoginPageController function with session.
// It creates a new controller, and a new request with a new response recorder.
// It calls the LoginPageController function with the recorder and the request.
// It checks the status code of the response.
func TestLoginPageControllerWithSession(t *testing.T) {
	c := getNewAuthController()
	sessionKey := "my-session-key"

	// Set the session key in the session store.
	c.sessionStore.Set(sessionKey, session.New(&model.User{}))

	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Set the session cookie for the request.
	sessionCookie := &http.Cookie{
		Name:  "session",
		Value: sessionKey,
		Path:  "/",
	}
	req.AddCookie(sessionCookie)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.LoginPageController)

	handler.ServeHTTP(rr, req)

	// On case of the session is set, the handler has to redirect to the /admin/dashboard.
	// On this case the status code is 303.
	testhelper.CheckResponseCode(t, rr, http.StatusSeeOther)
}

// TestLoginActionControllerNoInput tests the LoginActionController function. With missing input.
// It creates a new controller, and a new request with a new response recorder.
// It calls the LoginActionController function with the recorder and the request.
// It checks the status code of the response.
func TestLoginActionControllerNoInput(t *testing.T) {
	c := getNewAuthController()

	req, err := http.NewRequest("POST", "/auth/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.LoginActionController)

	handler.ServeHTTP(rr, req)

	// On case of missing input, the handler has to redirect to the /login.
	// On this case the status code is 303.
	testhelper.CheckResponseCode(t, rr, http.StatusSeeOther)
}

// TestLoginActionControllerNoUser tests the LoginActionController function. With missing username.
// It creates a new controller, and a new request with a new response recorder.
// It calls the LoginActionController function with the recorder and the request.
// It checks the status code of the response.
func TestLoginActionControllerNoUser(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.Error = sql.ErrNoRows
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(
		userRepository,
		&testhelper.RoleRepositoryMock{},
		&testhelper.ResourceRepositoryMock{},
		sessionStore,
		renderer,
	)

	// Send request with the username and password.
	// The user db is empty, so that the user is not found.
	req, err := http.NewRequest("POST", "/auth/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"username": []string{"missing"},
		"password": []string{"missing"},
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.LoginActionController)

	handler.ServeHTTP(rr, req)

	// On case of missing user, the handler returns error with status code 500.
	testhelper.CheckResponse(t, rr, http.StatusInternalServerError, []string{UserFailedToGetUserErrorMessage})
}

// TestLoginActionControllerWrongPassword tests the LoginActionController function. With wrong password.
// It creates a new controller, and a new request with a new response recorder.
// It calls the LoginActionController function with the recorder and the request.
// It checks the status code of the response.
func TestLoginActionControllerWrongPassword(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.Error = nil
	passwordHash, _ := passwd.HashPassword("test-password")
	userRepository.LatestUser = &model.User{
		Email:    "test-email@address.com",
		Password: passwordHash,
	}
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(
		userRepository,
		&testhelper.RoleRepositoryMock{},
		&testhelper.ResourceRepositoryMock{},
		sessionStore,
		renderer)

	// Send request with the username and password.
	// The user db is not empty, but the password is wrong.
	req, err := http.NewRequest("POST", "/auth/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"username": []string{"test-email@address.com"},
		"password": []string{"wrong-password"},
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.LoginActionController)

	handler.ServeHTTP(rr, req)

	// On case of wrong password, the handler has to redirect to the /login.
	// On this case the status code is 303.
	testhelper.CheckResponseCode(t, rr, http.StatusSeeOther)
}

// TestLoginActionController tests the LoginActionController function. With correct input.
// It creates a new controller, and a new request with a new response recorder.
// It calls the LoginActionController function with the recorder and the request.
// It checks the status code of the response.
func TestLoginActionController(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	userRepository.Error = nil
	passwordHash, _ := passwd.HashPassword("test-password")
	email := "test-email@address.com"
	userRepository.LatestUser = &model.User{
		Email:    email,
		Password: passwordHash,
	}
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(
		userRepository,
		&testhelper.RoleRepositoryMock{},
		&testhelper.ResourceRepositoryMock{},
		sessionStore,
		renderer)

	// Send request with the username and password.
	// The user db is not empty, and the password is correct.
	req, err := http.NewRequest("POST", "/auth/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Form = map[string][]string{
		"username": []string{email},
		"password": []string{"test-password"},
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.LoginActionController)

	handler.ServeHTTP(rr, req)

	// On case of correct input, the handler has to redirect to the /admin/dashboard.
	// On this case the status code is 303.
	testhelper.CheckResponseCode(t, rr, http.StatusSeeOther)
}
