package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/render"
	"github.com/akosgarai/projectregister/pkg/session"
)

// I want the router creation to be here.

// New creates a new instance of the router gorilla/mux router.
func New(
	userRepository model.UserRepository,
	roleRepository model.RoleRepository,
	resourceRepository model.ResourceRepository,
	clientRepository model.ClientRepository,
	projectRepository model.ProjectRepository,
	domainRepository model.DomainRepository,
	environmentRepository model.EnvironmentRepository,
	sessionStore *session.Store,
	renderer *render.Renderer,
) *mux.Router {
	r := mux.NewRouter()
	// handle the static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(renderer.GetStaticDirectoryPath()))))
	routerController := controller.New(
		userRepository,
		roleRepository,
		resourceRepository,
		clientRepository,
		projectRepository,
		domainRepository,
		environmentRepository,
		sessionStore,
		renderer,
	)
	r.HandleFunc("/health", routerController.HealthController)
	r.HandleFunc("/login", routerController.LoginPageController)
	r.HandleFunc("/auth/login", routerController.LoginActionController).Methods("POST")
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(routerController.AuthMiddleware)
	adminRouter.HandleFunc("/dashboard", routerController.DashboardController)
	adminRouter.HandleFunc("/user/create", routerController.UserCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/user/view/{userId}", routerController.UserViewController)
	adminRouter.HandleFunc("/user/update/{userId}", routerController.UserUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/user/delete/{userId}", routerController.UserDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/user/list", routerController.UserListViewController)
	adminRouter.HandleFunc("/role/create", routerController.RoleCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/role/view/{roleId}", routerController.RoleViewController)
	adminRouter.HandleFunc("/role/update/{roleId}", routerController.RoleUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/role/delete/{roleId}", routerController.RoleDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/role/list", routerController.RoleListViewController)

	adminRouter.HandleFunc("/client/create", routerController.ClientCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/client/view/{clientId}", routerController.ClientViewController)
	adminRouter.HandleFunc("/client/update/{clientId}", routerController.ClientUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/client/delete/{clientId}", routerController.ClientDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/client/list", routerController.ClientListViewController)

	adminRouter.HandleFunc("/project/create", routerController.ProjectCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/project/view/{projectId}", routerController.ProjectViewController)
	adminRouter.HandleFunc("/project/update/{projectId}", routerController.ProjectUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/project/delete/{projectId}", routerController.ProjectDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/project/list", routerController.ProjectListViewController)

	adminRouter.HandleFunc("/domain/create", routerController.DomainCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/domain/view/{domainId}", routerController.DomainViewController)
	adminRouter.HandleFunc("/domain/update/{domainId}", routerController.DomainUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/domain/delete/{domainId}", routerController.DomainDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/domain/list", routerController.DomainListViewController)

	adminRouter.HandleFunc("/environment/create", routerController.EnvironmentCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/environment/view/{environmentId}", routerController.EnvironmentViewController)
	adminRouter.HandleFunc("/environment/update/{environmentId}", routerController.EnvironmentUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/environment/delete/{environmentId}", routerController.EnvironmentDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/environment/list", routerController.EnvironmentListViewController)

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(routerController.AuthMiddleware)
	apiRouter.HandleFunc("/user/create", routerController.UserCreateAPIController).Methods("POST")
	apiRouter.HandleFunc("/user/view/{userId}", routerController.UserViewAPIController)
	apiRouter.HandleFunc("/user/update/{userId}", routerController.UserUpdateAPIController).Methods("POST")
	apiRouter.HandleFunc("/user/delete/{userId}", routerController.UserDeleteAPIController).Methods("DELETE")
	apiRouter.HandleFunc("/user/list", routerController.UserListAPIController)
	return r
}
