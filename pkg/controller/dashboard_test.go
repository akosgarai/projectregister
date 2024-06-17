package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
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

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	needles := []string{
		"<title>Dashboard</title>",
	}
	body := rr.Body.String()
	for _, needle := range needles {
		if !strings.Contains(body, needle) {
			t.Errorf("handler returned unexpected body: got %v want %v", body, needle)
		}
	}
}
