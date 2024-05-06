package controller

import (
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/render"
	"github.com/akosgarai/projectregister/pkg/session"
)

// Controller type for controller
// it holds the dependencies for the controller
type Controller struct {
	userRepository model.UserRepository

	sessionStore *session.Store

	renderer *render.Renderer
}

// New creates a new controller
func New(userRepository model.UserRepository, sessionStore *session.Store, renderer *render.Renderer) *Controller {
	return &Controller{
		userRepository: userRepository,

		sessionStore: sessionStore,

		renderer: renderer,
	}
}
