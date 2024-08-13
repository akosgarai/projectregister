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
	repositoryContainer := testhelper.NewRepositoryContainerMock()
	sessionStore := session.NewStore(config.DefaultEnvironment())
	csvStorage := testhelper.CSVStorageMock{}
	renderer := render.NewRenderer(config.DefaultEnvironment(), render.NewTemplates())
	c := New(
		repositoryContainer,
		sessionStore,
		csvStorage,
		renderer,
	)
	if c.repositoryContainer.GetUserRepository() != repositoryContainer.Users {
		t.Errorf("UserRepository field is not the same as the input.")
	}
	if c.repositoryContainer.GetRoleRepository() != repositoryContainer.Roles {
		t.Errorf("RoleRepository field is not the same as the input.")
	}
	if c.repositoryContainer.GetResourceRepository() != repositoryContainer.Resources {
		t.Errorf("ResourceRepository field is not the same as the input.")
	}
	if c.repositoryContainer.GetClientRepository() != repositoryContainer.Clients {
		t.Errorf("ClientRepository field is not the same as the input.")
	}
	if c.repositoryContainer.GetProjectRepository() != repositoryContainer.Projects {
		t.Errorf("ProjectRepository field is not the same as the input.")
	}
	if c.repositoryContainer.GetDomainRepository() != repositoryContainer.Domains {
		t.Errorf("DomainRepository field is not the same as the input.")
	}
	if c.repositoryContainer.GetEnvironmentRepository() != repositoryContainer.Environments {
		t.Errorf("EnvironmentRepository field is not the same as the input.")
	}
	if c.repositoryContainer.GetRuntimeRepository() != repositoryContainer.Runtimes {
		t.Errorf("RunTimeRepository field is not the same as the input.")
	}
	if c.repositoryContainer.GetPoolRepository() != repositoryContainer.Pools {
		t.Errorf("PoolRepository field is not the same as the input.")
	}
	if c.repositoryContainer.GetDatabaseRepository() != repositoryContainer.Databases {
		t.Errorf("DatabaseRepository field is not the same as the input.")
	}
	if c.repositoryContainer.GetServerRepository() != repositoryContainer.Servers {
		t.Errorf("ServerRepository field is not the same as the input.")
	}
	if c.repositoryContainer.GetApplicationRepository() != repositoryContainer.Applications {
		t.Errorf("ApplicationRepository field is not the same as the input.")
	}
	if c.sessionStore != sessionStore {
		t.Errorf("SessionStore field is not the same as the input.")
	}
	if c.renderer != renderer {
		t.Errorf("Renderer field is not the same as the input.")
	}
}
