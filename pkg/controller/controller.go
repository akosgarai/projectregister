package controller

import (
	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/database/repository"
)

// Controller type for controller
// it holds the dependencies for the controller
type Controller struct {
	db *database.DB

	userRepository *repository.UserRepository
}

// New creates a new controller
func New(db *database.DB) *Controller {
	return &Controller{
		db: db,

		userRepository: repository.NewUserRepository(db),
	}
}
