package controller

import (
	"net/http"

	"github.com/akosgarai/projectregister/pkg/controller/response"
	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/passwd"
	"github.com/akosgarai/projectregister/pkg/session"
)

// LoginPageController is the login controller.
// It returns the login page.
// in case of the user is already authenticated, it redirects to the dashboard.
func (c *Controller) LoginPageController(w http.ResponseWriter, r *http.Request) {
	// check if the user is authenticated
	// if yes, redirect to the dashboard
	sessionKey, err := r.Cookie("session")
	if err == nil {
		_, err = c.sessionStore.Get(sessionKey.Value)
		if err == nil {
			http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
			return
		}
	}
	headerText := "Login"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	content := response.NewResponse(headerText, &model.User{Role: &model.Role{}}, headerContent)
	err = c.renderer.Template.RenderTemplate(w, "login.html", content)
	if err != nil {
		panic(err)
	}
}

// LoginActionController is the login action controller.
// It is responsible for handling the login action.
func (c *Controller) LoginActionController(w http.ResponseWriter, r *http.Request) {
	// the username is the email as it is unique.
	username := r.FormValue("username")
	password := r.FormValue("password")
	// If the username or password is empty, redirect to the login page.
	if username == "" || password == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	user, err := c.repositoryContainer.GetUserRepository().GetUserByEmail(username)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, UserFailedToGetUserErrorMessage, err)
		return
	}
	if !passwd.ComparePassword(password, user.Password) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	// generate session key
	sessionKey, err := c.sessionStore.GenerateSessionKey()
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, AuthFailedToGenerateSessionKeyErrorMessage, err)
		return
	}
	// set the session
	c.sessionStore.Set(sessionKey, session.New(user))
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: sessionKey,
		Path:  "/",
	})
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

// AuthMiddleware is the authentication middleware.
// It is responsible for checking if the user is authenticated.
func (c *Controller) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if the user is authenticated
		// if not, redirect to the login page
		// if yes, call the next handler
		sessionKey, err := r.Cookie("session")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		_, err = c.sessionStore.Get(sessionKey.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// CurrentUser returns the current user.
func (c *Controller) CurrentUser(r *http.Request) *model.User {
	sessionKey, err := r.Cookie("session")
	if err != nil {
		panic(err)
	}
	session, err := c.sessionStore.Get(sessionKey.Value)
	if err != nil {
		panic(err)
	}
	return session.GetUser()
}
