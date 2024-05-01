package controller

import (
	"net/http"

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
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "admin" && password == "admin" {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
