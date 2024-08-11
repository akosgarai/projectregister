package controller

import (
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/render"
	"github.com/akosgarai/projectregister/pkg/session"
	"github.com/akosgarai/projectregister/pkg/storage"
)

// Controller type for controller
// it holds the dependencies for the controller
type Controller struct {
	repositoryContainer model.RepositoryContainer

	sessionStore *session.Store
	csvStorage   storage.CSVStorage

	renderer *render.Renderer
}

// New creates a new controller
func New(
	repositoryContainer model.RepositoryContainer,
	sessionStore *session.Store,
	csvStorage storage.CSVStorage,
	renderer *render.Renderer,
) *Controller {
	return &Controller{
		repositoryContainer: repositoryContainer,

		sessionStore: sessionStore,
		csvStorage:   csvStorage,

		renderer: renderer,
	}
}
