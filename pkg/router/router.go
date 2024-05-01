package router

import (
	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller"
)

// I want the router creation to be here.

// New creates a new instance of the router gorilla/mux router.
func New() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", controller.HealthController)
	r.HandleFunc("/login", controller.LoginPageController)
	r.HandleFunc("/auth/login", controller.LoginActionController).Methods("POST")
	r.HandleFunc("/dashboard", controller.DashboardController)
	r.HandleFunc("/user/view/{userId}", controller.UserViewController)

	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/user/create", controller.UserCreateAPIController).Methods("POST")
	return r
}
