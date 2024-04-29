package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/render"
)

// UserViewController is the controller for the user view page.
func UserViewController(w http.ResponseWriter, r *http.Request) {
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
