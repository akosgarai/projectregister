package render

import (
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/akosgarai/projectregister/pkg/config"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

var (
	httpStatusCodes = []int{200, 201, 202, 203, 204, 205, 206, 207, 208, 226, 300, 301, 302, 303, 304, 305, 306, 307, 308}
)

// TestNewRenderer is a test function for the NewRenderer function.
func TestNewRenderer(t *testing.T) {
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	renderer := NewRenderer(testConfig, NewTemplates())
	if renderer == nil {
		t.Error("The renderer is nil.")
	}
	if renderer.Template == nil {
		t.Errorf("The template is nil.")
	}
}

// TestGetTemplateDirectoryPath is a test function for the GetTemplateDirectoryPath function.
func TestGetTemplateDirectoryPath(t *testing.T) {
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	renderer := NewRenderer(testConfig, NewTemplates())
	if renderer.GetTemplateDirectoryPath() != testConfig.GetRenderTemplateDirectoryPath() {
		t.Errorf("The template directory path is not correct. Expected: %s, got: %s", testConfig.GetRenderTemplateDirectoryPath(), renderer.GetTemplateDirectoryPath())
	}
}

// TestGetStaticDirectoryPath is a test function for the GetStaticDirectoryPath function.
func TestGetStaticDirectoryPath(t *testing.T) {
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	renderer := NewRenderer(testConfig, NewTemplates())
	if renderer.GetStaticDirectoryPath() != testConfig.GetStaticDirectoryPath() {
		t.Errorf("The static directory path is not correct. Expected: %s, got: %s", testConfig.GetStaticDirectoryPath(), renderer.GetStaticDirectoryPath())
	}
}

// TestJSON is a test function for the JSON function.
func TestJSON(t *testing.T) {
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	renderer := NewRenderer(testConfig, NewTemplates())
	// Test the JSON function with a nil value.
	// The function should return an error.
	for _, code := range httpStatusCodes {
		w := httptest.NewRecorder()
		testValue := make(map[string]string)
		renderer.JSON(w, code, testValue)
		if w.Code != code {
			t.Errorf("The status is not correct. Expected: %d, got: %d", code, w.Code)
		}
		if w.Body.String() != "{}\n" {
			t.Errorf("The body is not correct. Expected: '{}\n', got: '%s'", w.Body.String())
		}
	}

	// unsupported type
	w := httptest.NewRecorder()
	testValue := make(chan int)
	renderer.JSON(w, 200, testValue)
	if w.Code != 500 {
		t.Errorf("The status is not correct. Expected: %d, got: %d", 500, w.Code)
	}
	body := w.Body.String()
	expected := "Internal server error"
	if !strings.Contains(body, expected) {
		t.Errorf("The body is not correct. Expected: '%s', got: '%s'", expected, body)
	}
}

// TestStatus is a test function for the Status function.
func TestStatus(t *testing.T) {
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	renderer := NewRenderer(testConfig, NewTemplates())
	// Test the Status function with different status codes.
	for _, code := range httpStatusCodes {
		w := httptest.NewRecorder()
		renderer.Status(w, code)
		if w.Code != code {
			t.Errorf("The status is not correct. Expected: %d, got: %d", code, w.Code)
		}
		// body is empty
		if w.Body.String() != "" {
			t.Errorf("The body is not correct. Expected: '', got: '%s'", w.Body.String())
		}
	}
}

// TestErrorWithoutDetails is a test function for the Error function.
func TestErrorWithoutDetails(t *testing.T) {
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	renderer := NewRenderer(testConfig, NewTemplates())
	// Test the Error function with different status codes and messages.
	for _, code := range httpStatusCodes {
		w := httptest.NewRecorder()
		testMessage := "test message"
		renderer.Error(w, code, testMessage, nil)
		if w.Code != code {
			t.Errorf("The status is not correct. Expected: %d, got: %d", code, w.Code)
		}
		if w.Body.String() != testMessage+"\n" {
			t.Errorf("The body is not correct. Expected: '%s', got: '%s'", testMessage, w.Body.String())
		}
	}
}

// TestErrorWithDetails is a test function for the Error function.
func TestErrorWithDetails(t *testing.T) {
	testConfig := config.NewEnvironment(testhelper.TestConfigData)
	renderer := NewRenderer(testConfig, NewTemplates())
	// Test the Error function with different status codes and messages.
	for _, code := range httpStatusCodes {
		w := httptest.NewRecorder()
		testMessage := "test message"
		details := errors.New("test error")
		renderer.Error(w, code, testMessage, details)
		if w.Code != code {
			t.Errorf("The status is not correct. Expected: %d, got: %d", code, w.Code)
		}
		expected := testMessage + " " + details.Error()
		if w.Body.String() != expected+"\n" {
			t.Errorf("The body is not correct. Expected: '%s', got: '%s'", expected, w.Body.String())
		}
	}
}
