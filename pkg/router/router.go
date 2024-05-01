package router

import (
	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller"
)

// I want the router creation to be here.

// New creates a new instance of the router gorilla/mux router.
func New() *mux.Router {
	r := mux.NewRouter()
	routerController := controller.New()
	r.HandleFunc("/health", routerController.HealthController)
	r.HandleFunc("/login", routerController.LoginPageController)
	r.HandleFunc("/auth/login", routerController.LoginActionController).Methods("POST")
	r.HandleFunc("/dashboard", routerController.DashboardController)
	r.HandleFunc("/user/view/{userId}", routerController.UserViewController)

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/user/create", routerController.UserCreateAPIController).Methods("POST")
	return r
}
