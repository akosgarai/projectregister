package repository

import (
	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/model"
)

// ContainerRepository type
// It implements the RepositoryContainer interface.
type ContainerRepository struct {
	applications *ApplicationRepository
	clients      *ClientRepository
	databases    *DatabaseRepository
	domains      *DomainRepository
	environments *EnvironmentRepository
	pools        *PoolRepository
	projects     *ProjectRepository
	resources    *ResourceRepository
	roles        *RoleRepository
	runtimes     *RuntimeRepository
	servers      *ServerRepository
	users        *UserRepository
}

// NewContainerRepository creates a new container repository
func NewContainerRepository(db *database.DB) *ContainerRepository {
	return &ContainerRepository{
		applications: NewApplicationRepository(db),
		clients:      NewClientRepository(db),
		databases:    NewDatabaseRepository(db),
		domains:      NewDomainRepository(db),
		environments: NewEnvironmentRepository(db),
		pools:        NewPoolRepository(db),
		projects:     NewProjectRepository(db),
		resources:    NewResourceRepository(db),
		roles:        NewRoleRepository(db),
		runtimes:     NewRuntimeRepository(db),
		servers:      NewServerRepository(db),
		users:        NewUserRepository(db),
	}
}

// GetApplicationRepository returns the application repository
func (r *ContainerRepository) GetApplicationRepository() model.ApplicationRepository {
	return r.applications
}

// GetClientRepository returns the client repository
func (r *ContainerRepository) GetClientRepository() model.ClientRepository {
	return r.clients
}

// GetDatabaseRepository returns the database repository
func (r *ContainerRepository) GetDatabaseRepository() model.DatabaseRepository {
	return r.databases
}

// GetDomainRepository returns the domain repository
func (r *ContainerRepository) GetDomainRepository() model.DomainRepository {
	return r.domains
}

// GetEnvironmentRepository returns the environment repository
func (r *ContainerRepository) GetEnvironmentRepository() model.EnvironmentRepository {
	return r.environments
}

// GetPoolRepository returns the pool repository
func (r *ContainerRepository) GetPoolRepository() model.PoolRepository {
	return r.pools
}

// GetProjectRepository returns the project repository
func (r *ContainerRepository) GetProjectRepository() model.ProjectRepository {
	return r.projects
}

// GetResourceRepository returns the resource repository
func (r *ContainerRepository) GetResourceRepository() model.ResourceRepository {
	return r.resources
}

// GetRoleRepository returns the role repository
func (r *ContainerRepository) GetRoleRepository() model.RoleRepository {
	return r.roles
}

// GetRuntimeRepository returns the runtime repository
func (r *ContainerRepository) GetRuntimeRepository() model.RuntimeRepository {
	return r.runtimes
}

// GetServerRepository returns the server repository
func (r *ContainerRepository) GetServerRepository() model.ServerRepository {
	return r.servers
}

// GetUserRepository returns the user repository
func (r *ContainerRepository) GetUserRepository() model.UserRepository {
	return r.users
}
