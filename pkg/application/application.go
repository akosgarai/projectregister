package application

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/config"
	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/database/repository"
	"github.com/akosgarai/projectregister/pkg/render"
	"github.com/akosgarai/projectregister/pkg/router"
	"github.com/akosgarai/projectregister/pkg/session"
)

// App is a struct that holds the application configuration.
type App struct {
	envConfig *config.Environment
	db        *database.DB

	Server *http.Server
	Router *mux.Router
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
	// create a new database
	a.db = database.NewDB(a.envConfig)
	// connect to the database
	err := a.db.Connect()
	if err != nil {
		panic(err)
	}
	// create a new router
	userRepository := repository.NewUserRepository(a.db)
	roleRepository := repository.NewRoleRepository(a.db)
	resourceRepository := repository.NewResourceRepository(a.db)
	clientRepository := repository.NewClientRepository(a.db)
	projectRepository := repository.NewProjectRepository(a.db)
	a.Router = router.New(
		userRepository,
		roleRepository,
		resourceRepository,
		clientRepository,
		projectRepository,
		session.NewStore(a.envConfig),
		render.NewRenderer(a.envConfig),
	)
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
	migration := database.NewMigration(a.envConfig)
	err := migration.Up()
	if err != nil {
		panic(err)
	}
}
