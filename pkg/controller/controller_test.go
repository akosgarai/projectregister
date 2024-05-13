package controller

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/config"
	"github.com/akosgarai/projectregister/pkg/render"
	"github.com/akosgarai/projectregister/pkg/session"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestNew tests the New function.
// It creates a new controller and checks the fields.
func TestNew(t *testing.T) {
	userRepository := &testhelper.UserRepositoryMock{}
	sessionStore := session.NewStore(config.DefaultEnvironment())
	renderer := render.NewRenderer(config.DefaultEnvironment())
	c := New(userRepository, sessionStore, renderer)
	if c.userRepository != userRepository {
		t.Errorf("UserRepository field is not the same as the input.")
	}
	if c.sessionStore != sessionStore {
		t.Errorf("SessionStore field is not the same as the input.")
	}
	if c.renderer != renderer {
		t.Errorf("Renderer field is not the same as the input.")
	}
}
