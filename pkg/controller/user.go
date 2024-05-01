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
	u, statusCode, err := c.userViewData(r)
	if err != nil {
		http.Error(w, "Failed to get user data "+err.Error(), statusCode)
		return
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

// UserViewAPIController is the controller for the user view API.
// It is responsible for returning the user data as JSON.
// Example request:
// curl -X GET http://localhost:8090/api/user/view/1
func (c *Controller) UserViewAPIController(w http.ResponseWriter, r *http.Request) {
	u, statusCode, err := c.userViewData(r)
	if err != nil {
		http.Error(w, "Failed to get user data "+err.Error(), statusCode)
		return
	}
	render.JSON(w, statusCode, u)
}

// userViewData gets the request as input, and returns the user data, status code and error.
func (c *Controller) userViewData(r *http.Request) (*model.User, int, error) {
	vars := mux.Vars(r)
	userIDVariable := vars["userId"]
	// it has to be converted to int64
	userID, err := strconv.ParseInt(userIDVariable, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	u, err := c.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return u, http.StatusOK, nil
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
	user, err := c.userRepository.CreateUser(name, email, string(hashedPassword))
	if err != nil {
		http.Error(w, "Internal server error - failed to insert data to the database "+err.Error(), http.StatusInternalServerError)
		return
	}
	// return the user as JSON
	render.JSON(w, http.StatusOK, user)
}
