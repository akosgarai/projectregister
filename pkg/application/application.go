package application

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// App is a struct that holds the application configuration.
type App struct {
	Server *http.Server
	Router *mux.Router
}

// New creates a new instance of the application.
func New() *App {
	return &App{}
}

// Initialize initializes the application, runs the database migrations, and sets up the routes.
func (a *App) Initialize() {
	// TODO: Implement it.
	// create a new router
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	// create a new server
	a.Server = &http.Server{
		Addr: "0.0.0.0:8090",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      a.Router, // Pass our instance of gorilla/mux in.
	}
}

// Run starts the application. Returns an error if the server fails to start.
func (a *App) Run() error {
	if err := a.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
