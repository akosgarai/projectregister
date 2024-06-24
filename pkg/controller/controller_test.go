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
	roleRepository := &testhelper.RoleRepositoryMock{}
	resourceRepository := &testhelper.ResourceRepositoryMock{}
	clientRepository := &testhelper.ClientRepositoryMock{}
	projectRepository := &testhelper.ProjectRepositoryMock{}
	sessionStore := session.NewStore(config.DefaultEnvironment())
	renderer := render.NewRenderer(config.DefaultEnvironment())
	c := New(
		userRepository,
		roleRepository,
		resourceRepository,
		clientRepository,
		projectRepository,
		sessionStore,
		renderer,
	)
	if c.userRepository != userRepository {
		t.Errorf("UserRepository field is not the same as the input.")
	}
	if c.roleRepository != roleRepository {
		t.Errorf("RoleRepository field is not the same as the input.")
	}
	if c.resourceRepository != resourceRepository {
		t.Errorf("ResourceRepository field is not the same as the input.")
	}
	if c.clientRepository != clientRepository {
		t.Errorf("ClientRepository field is not the same as the input.")
	}
	if c.projectRepository != projectRepository {
		t.Errorf("ProjectRepository field is not the same as the input.")
	}
	if c.sessionStore != sessionStore {
		t.Errorf("SessionStore field is not the same as the input.")
	}
	if c.renderer != renderer {
		t.Errorf("Renderer field is not the same as the input.")
	}
}
