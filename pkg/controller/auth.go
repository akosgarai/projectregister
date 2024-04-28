package controller

import (
	"net/http"

	"github.com/akosgarai/projectregister/pkg/render"
)

// LoginPageController is the login controller.
// It returns the login page.
func LoginPageController(w http.ResponseWriter, r *http.Request) {
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
