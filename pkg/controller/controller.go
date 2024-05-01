package controller

import (
	"github.com/akosgarai/projectregister/pkg/database"
)

// Controller type for controller
// it holds the dependencies for the controller
type Controller struct {
	db *database.DB
}

// New creates a new controller
func New(db *database.DB) *Controller {
	return &Controller{
		db: db,
	}
}
