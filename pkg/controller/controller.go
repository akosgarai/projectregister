package controller

import (
	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/database/repository"
	"github.com/akosgarai/projectregister/pkg/session"
)

// Controller type for controller
// it holds the dependencies for the controller
type Controller struct {
	db *database.DB

	userRepository *repository.UserRepository

	sessionStore *session.Store
}

// New creates a new controller
func New(db *database.DB, sessionStore *session.Store) *Controller {
	return &Controller{
		db: db,

		userRepository: repository.NewUserRepository(db),

		sessionStore: sessionStore,
	}
}
