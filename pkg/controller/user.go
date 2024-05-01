package controller

import (
	"golang.org/x/crypto/bcrypt"
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
		http.Error(w, "Invalid form "+err.Error(), http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal server error - failed to encrypt the password "+err.Error(), http.StatusInternalServerError)
		return
	}
	// create the user
	u := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		ID:       0,
	}
	// insert the user into the database
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err = c.db.QueryRow(query, u.Name, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		http.Error(w, "Internal server error - failed to insert data to the database "+err.Error(), http.StatusInternalServerError)
		return
	}
	// return the user as JSON
	render.JSON(w, http.StatusOK, u)
}
