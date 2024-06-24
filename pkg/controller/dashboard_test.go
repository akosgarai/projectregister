package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akosgarai/projectregister/pkg/config"
	"github.com/akosgarai/projectregister/pkg/render"
	"github.com/akosgarai/projectregister/pkg/session"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestDashboardController tests the DashboardController function.
// In has to render the dashboard page.
// The test checks the status code of the response.
// The test creates a new controller, a new request with a new response recorder.
// It calls the DashboardController function with the recorder and the request.
func TestDashboardController(t *testing.T) {
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	renderer := render.NewRenderer(testConfig)
	c := New(
		&testhelper.UserRepositoryMock{},
		&testhelper.RoleRepositoryMock{},
		&testhelper.ResourceRepositoryMock{},
		sessionStore,
		renderer,
	)

	req, err := http.NewRequest("GET", "/dashboard", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.DashboardController)

	handler.ServeHTTP(rr, req)

	needles := []string{
		"<title>Dashboard</title>",
	}
	testhelper.CheckResponse(t, rr, http.StatusOK, needles)
}
