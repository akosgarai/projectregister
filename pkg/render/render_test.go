package render

import (
	"net/http/httptest"
	"testing"

	"github.com/akosgarai/projectregister/pkg/config"
)

var (
	testConfigData = map[string]string{
		"RENDER_TEMPLATE_DIRECTORY_PATH": "../../web/template",
	}
)

// TestNewRenderer is a test function for the NewRenderer function.
func TestNewRenderer(t *testing.T) {
	testConfig := config.NewEnvironment(testConfigData)
	renderer := NewRenderer(testConfig)
	if renderer == nil {
		t.Error("The renderer is nil.")
	}
	if renderer.baseTemplate != testConfig.GetRenderTemplateDirectoryPath()+"/"+testConfig.GetRenderBaseTemplate() {
		t.Errorf("The base template is not correct. Expected: %s, got: %s", testConfig.GetRenderTemplateDirectoryPath()+"/"+testConfig.GetRenderBaseTemplate(), renderer.baseTemplate)
	}
}

// TestBuildTemplate is a test function for the BuildTemplate function.
func TestBuildTemplate(t *testing.T) {
	// Test the template building with empty file list.
	testConfig := config.NewEnvironment(testConfigData)
	renderer := NewRenderer(testConfig)
	emptyList := []string{}
	templateName := "test"
	template := renderer.BuildTemplate(templateName, emptyList)
	if template == nil {
		t.Error("The template is nil.")
	}
	// The template name has to be 'test'
	if template.Name() != templateName {
		t.Errorf("The template name is not correct. Expected: %s, got: %s", templateName, template.Name())
	}
}

// TestJSON is a test function for the JSON function.
func TestJSON(t *testing.T) {
	testConfig := config.NewEnvironment(testConfigData)
	renderer := NewRenderer(testConfig)
	// Test the JSON function with a nil value.
	// The function should return an error.
	httpStatusCodes := []int{200, 201, 202, 203, 204, 205, 206, 207, 208, 226, 300, 301, 302, 303, 304, 305, 306, 307, 308}
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
	if w.Body.String() != "Internal server error\n" {
		t.Errorf("The body is not correct. Expected: 'Internal server error\n', got: '%s'", w.Body.String())
	}
}
