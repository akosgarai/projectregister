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
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view"})
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	sessionStore := session.NewStore(testConfig)
	sessionStore.Set(testhelper.TestSessionCookieValue, session.New(testUser))
	c := New(
		testhelper.NewRepositoryContainerMock(),
		sessionStore,
		testhelper.CSVStorageMock{},
		render.NewRenderer(testConfig, render.NewTemplates()))
	c.CacheTemplates()

	req, err := testhelper.NewRequestWithSessionCookie("GET", "/dashboard")
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
