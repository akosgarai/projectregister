package controller

import (
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/render"
	"github.com/akosgarai/projectregister/pkg/session"
)

// Controller type for controller
// it holds the dependencies for the controller
type Controller struct {
	userRepository     model.UserRepository
	roleRepository     model.RoleRepository
	resourceRepository model.ResourceRepository
	clientRepository   model.ClientRepository
	projectRepository  model.ProjectRepository
	domainRepository   model.DomainRepository

	environmentRepository model.EnvironmentRepository
	runtimeRepository     model.RuntimeRepository
	poolRepository        model.PoolRepository
	databaseRepository    model.DatabaseRepository
	serverRepository      model.ServerRepository

	sessionStore *session.Store

	renderer *render.Renderer
}

// New creates a new controller
func New(
	userRepository model.UserRepository,
	roleRepository model.RoleRepository,
	resourceRepository model.ResourceRepository,
	clientRepository model.ClientRepository,
	projectRepository model.ProjectRepository,
	domainRepository model.DomainRepository,
	environmentRepository model.EnvironmentRepository,
	runtimeRepository model.RuntimeRepository,
	poolRepository model.PoolRepository,
	databaseRepository model.DatabaseRepository,
	serverRepository model.ServerRepository,
	sessionStore *session.Store,
	renderer *render.Renderer,
) *Controller {
	return &Controller{
		userRepository:        userRepository,
		roleRepository:        roleRepository,
		resourceRepository:    resourceRepository,
		clientRepository:      clientRepository,
		projectRepository:     projectRepository,
		domainRepository:      domainRepository,
		environmentRepository: environmentRepository,
		runtimeRepository:     runtimeRepository,
		databaseRepository:    databaseRepository,
		poolRepository:        poolRepository,
		serverRepository:      serverRepository,

		sessionStore: sessionStore,

		renderer: renderer,
	}
}
