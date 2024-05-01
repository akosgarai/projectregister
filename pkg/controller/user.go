package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/render"
)

// UserViewController is the controller for the user view page.
func (c *Controller) UserViewController(w http.ResponseWriter, r *http.Request) {
	template := render.BuildTemplate("login", []string{"web/template/user/view.html.tmpl"})
	vars := mux.Vars(r)
	userIDVariable := vars["userId"]
	// it has to be converted to int
	userID, err := strconv.ParseInt(userIDVariable, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	u := &model.User{
		Name:  "admin",
		Email: "admin@admin.com",
		ID:    userID,
	}
	content := struct {
		Title string
		User  *model.User
	}{
		Title: "User View",
		User:  u,
	}
	err = template.ExecuteTemplate(w, "base.html", content)
	if err != nil {
		panic(err)
	}
}

// UserCreateAPIController is the controller for the user create API.
// It is responsible for creating a new user.
// It returns the created user as JSON.
// Example request:
// curl -X POST http://localhost:8090/api/user/create -d "name=Bob&email=bob@bob"
func (c *Controller) UserCreateAPIController(w http.ResponseWriter, r *http.Request) {
	// it is a POST request, so we can parse the form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	// create the user
	u := &model.User{
		Name:  name,
		Email: email,
		ID:    0,
	}
	// return the user as JSON
	render.JSON(w, http.StatusOK, u)
}
