package router

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/config"
	"github.com/akosgarai/projectregister/pkg/render"
	"github.com/akosgarai/projectregister/pkg/session"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestNew Tests the New function
// It has to return a new router
// The router has to have the given routes
// The router has to have the given controller
func TestNew(t *testing.T) {
	sessionStore := session.NewStore(config.DefaultEnvironment())
	router := New(
		&testhelper.UserRepositoryMock{},
		&testhelper.RoleRepositoryMock{},
		&testhelper.ResourceRepositoryMock{},
		sessionStore,
		render.NewRenderer(config.DefaultEnvironment()))
	if router == nil {
		t.Error("New router is nil")
	}
	// check the routes
	routesToCheck := []string{
		"/health",
		"/login",
		"/auth/login",

		"/admin/dashboard",
		"/admin/user/create",
		"/admin/user/view/{userId}",
		"/admin/user/update/{userId}",
		"/admin/user/delete/{userId}",
		"/admin/user/list",

		"/admin/role/create",
		"/admin/role/view/{roleId}",
		"/admin/role/update/{roleId}",
		"/admin/role/delete/{roleId}",
		"/admin/role/list",

		"/api/user/create",
		"/api/user/view/{userId}",
		"/api/user/update/{userId}",
		"/api/user/delete/{userId}",
		"/api/user/list",
	}
	for _, route := range routesToCheck {
		if router.Path(route) == nil {
			t.Errorf("Route %s is missing.", route)
		}
	}
}
