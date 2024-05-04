package router

import (
	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller"
	"github.com/akosgarai/projectregister/pkg/database"
)

// I want the router creation to be here.

// New creates a new instance of the router gorilla/mux router.
func New(db *database.DB) *mux.Router {
	r := mux.NewRouter()
	routerController := controller.New(db)
	r.HandleFunc("/health", routerController.HealthController)
	r.HandleFunc("/login", routerController.LoginPageController)
	r.HandleFunc("/auth/login", routerController.LoginActionController).Methods("POST")
	r.HandleFunc("/dashboard", routerController.DashboardController)
	r.HandleFunc("/user/view/{userId}", routerController.UserViewController)

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/user/create", routerController.UserCreateAPIController).Methods("POST")
	apiRouter.HandleFunc("/user/view/{userId}", routerController.UserViewAPIController)
	apiRouter.HandleFunc("/user/update/{userId}", routerController.UserUpdateAPIController).Methods("POST")
	apiRouter.HandleFunc("/user/delete/{userId}", routerController.UserDeleteAPIController).Methods("DELETE")
	apiRouter.HandleFunc("/user/list", routerController.UserListAPIController)
	return r
}
