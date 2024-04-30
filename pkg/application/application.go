package application

import (
	"net/http"
	"time"

	"github.com/golang-migrate/migrate/v4"
	// import the postgres driver
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/config"
	"github.com/akosgarai/projectregister/pkg/router"
)

// App is a struct that holds the application configuration.
type App struct {
	envConfig *config.Environment
	Server    *http.Server
	Router    *mux.Router
}

// New creates a new instance of the application.
func New(envConfig map[string]string) *App {
	return &App{
		envConfig: config.NewEnvironment(envConfig),
	}
}

// Initialize initializes the application, runs the database migrations, and sets up the routes.
func (a *App) Initialize() {
	// execute the migrations
	a.executeMigrations()
	// create a new router
	a.Router = router.New()
	// create a new server
	a.Server = &http.Server{
		Addr: a.envConfig.GetServerAddr() + ":" + a.envConfig.GetServerPort(),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * time.Duration(a.envConfig.GetServerWriteTimeout()),
		ReadTimeout:  time.Second * time.Duration(a.envConfig.GetServerReadTimeout()),
		IdleTimeout:  time.Second * time.Duration(a.envConfig.GetServerIdleTimeout()),
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

// execute the migrations
func (a *App) executeMigrations() {
	m, err := migrate.New(
		"file://db/migrations",
		"postgres://projectregister:password@db:5432/projectregister_development?sslmode=disable")
	if err != nil {
		panic(err)
	}
	m.Up()
}
