package router

import (
	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/render"
	"github.com/akosgarai/projectregister/pkg/session"
)

// I want the router creation to be here.

// New creates a new instance of the router gorilla/mux router.
func New(userRepository model.UserRepository, sessionStore *session.Store, renderer *render.Renderer) *mux.Router {
	r := mux.NewRouter()
	routerController := controller.New(userRepository, sessionStore, renderer)
	r.HandleFunc("/health", routerController.HealthController)
	r.HandleFunc("/login", routerController.LoginPageController)
	r.HandleFunc("/auth/login", routerController.LoginActionController).Methods("POST")
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(routerController.AuthMiddleware)
	adminRouter.HandleFunc("/dashboard", routerController.DashboardController)
	adminRouter.HandleFunc("/user/view/{userId}", routerController.UserViewController)
	adminRouter.HandleFunc("/user/list", routerController.UserListViewController)

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(routerController.AuthMiddleware)
	apiRouter.HandleFunc("/user/create", routerController.UserCreateAPIController).Methods("POST")
	apiRouter.HandleFunc("/user/view/{userId}", routerController.UserViewAPIController)
	apiRouter.HandleFunc("/user/update/{userId}", routerController.UserUpdateAPIController).Methods("POST")
	apiRouter.HandleFunc("/user/delete/{userId}", routerController.UserDeleteAPIController).Methods("DELETE")
	apiRouter.HandleFunc("/user/list", routerController.UserListAPIController)
	return r
}
