package controller

import (
	"net/http"

	"github.com/akosgarai/projectregister/pkg/passwd"
	"github.com/akosgarai/projectregister/pkg/render"
)

// LoginPageController is the login controller.
// It returns the login page.
func (c *Controller) LoginPageController(w http.ResponseWriter, r *http.Request) {
	template := render.BuildTemplate("login", []string{"web/template/auth/login.html.tmpl"})
	content := struct {
		Title string
	}{
		Title: "Login",
	}
	err := template.ExecuteTemplate(w, "base.html", content)
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
	user, err := c.userRepository.GetUserByEmail(username)
	if err != nil {
		http.Error(w, "Failed to get user data "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !passwd.ComparePassword(password, user.Password) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
