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

func TestHealthCheckHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	c := New(
		testhelper.NewRepositoryContainerMock(),
		session.NewStore(testConfig),
		testhelper.CSVStorageMock{},
		render.NewRenderer(testConfig, render.NewTemplates()))

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.HealthController)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	testhelper.CheckResponseCode(t, rr, http.StatusOK)

	// Check the response body is what we expect.
	expected := ""
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
