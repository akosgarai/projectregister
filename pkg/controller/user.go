package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/passwd"
)

// UserViewController is the controller for the user view page.
func (c *Controller) UserViewController(w http.ResponseWriter, r *http.Request) {
	template := c.renderer.BuildTemplate("login", []string{c.renderer.GetTemplateDirectoryPath() + "/user/view.html.tmpl"})
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
	c.renderer.JSON(w, statusCode, u)
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

// UserCreateViewController is the controller for the user create view.
// On case of get request, it returns the user create page.
// On case of put request, it creates the user and redirects to the list page.
func (c *Controller) UserCreateViewController(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		template := c.renderer.BuildTemplate("user-create", []string{c.renderer.GetTemplateDirectoryPath() + "/user/create.html.tmpl"})
		content := struct {
			Title string
		}{
			Title: "User Create",
		}
		err := template.ExecuteTemplate(w, "base.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the user
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		// if the name or email is empty, return an error
		if name == "" || email == "" || password == "" {
			http.Error(w, "Name and email and password are required", http.StatusBadRequest)
			return
		}

		password, err := passwd.HashPassword(password)
		if err != nil {
			http.Error(w, "Internal server error - failed to encrypt the password "+err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = c.userRepository.CreateUser(name, email, string(password))
		if err != nil {
			http.Error(w, "Internal server error - failed to create the user "+err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/user/list", http.StatusSeeOther)
		return
	}
}

// UserCreateAPIController is the controller for the user create API.
// It is responsible for creating a new user.
// It returns the created user as JSON.
// Example request:
// curl -X POST http://localhost:8090/api/user/create -d "name=Bob&email=bob@bob"
func (c *Controller) UserCreateAPIController(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	hashedPassword, err := passwd.HashPassword(password)
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
	c.renderer.JSON(w, http.StatusOK, user)
}

// UserUpdateViewController is the controller for the user update view.
// On case of get request, it returns the user update page.
// On case of put request, it updates the user and redirects to the list page.
func (c *Controller) UserUpdateViewController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDVariable := vars["userId"]
	// it has to be converted to int64
	userID, err := strconv.ParseInt(userIDVariable, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id "+err.Error(), http.StatusBadRequest)
		return
	}

	// get the user
	user, err := c.userRepository.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Internal server error - failed to get the user "+err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		template := c.renderer.BuildTemplate("user-update", []string{c.renderer.GetTemplateDirectoryPath() + "/user/update.html.tmpl"})
		content := struct {
			Title string
			User  *model.User
		}{
			Title: "User Update",
			User:  user,
		}
		err = template.ExecuteTemplate(w, "base.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the user
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		// if the name or email is empty, return an error
		if name == "" || email == "" {
			http.Error(w, "Name and email are required", http.StatusBadRequest)
			return
		}

		// if the password is not empty, encrypt it
		if password != "" {
			password, err := passwd.HashPassword(password)
			if err != nil {
				http.Error(w, "Internal server error - failed to encrypt the password "+err.Error(), http.StatusInternalServerError)
				return
			}
			user.Password = password
		}
		// update the user
		user.Name = name
		user.Email = email
		err = c.userRepository.UpdateUser(user)
		if err != nil {
			http.Error(w, "Internal server error - failed to update the user "+err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/user/list", http.StatusSeeOther)
		return
	}
}

// UserUpdateAPIController is the controller for the user update API.
// It is responsible for updating a user.
// It returns the updated user as JSON.
// Example request:
// curl -X POST http://localhost:8090/api/user/update/1 -d "name=Bob&email=bob@bob&password=password"
func (c *Controller) UserUpdateAPIController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDVariable := vars["userId"]
	// it has to be converted to int64
	userID, err := strconv.ParseInt(userIDVariable, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id "+err.Error(), http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	// get the user
	user, err := c.userRepository.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Internal server error - failed to get the user "+err.Error(), http.StatusInternalServerError)
		return
	}
	hashedPassword, err := passwd.HashPassword(password)
	if err != nil {
		http.Error(w, "Internal server error - failed to encrypt the password "+err.Error(), http.StatusInternalServerError)
		return
	}
	// update the user
	user.Name = name
	user.Email = email
	user.Password = hashedPassword
	err = c.userRepository.UpdateUser(user)
	if err != nil {
		http.Error(w, "Internal server error - failed to update the user "+err.Error(), http.StatusInternalServerError)
		return
	}
	// return the updated user as JSON
	c.renderer.JSON(w, http.StatusOK, user)
}

// UserDeleteViewController is the controller for the user delete view.
// It is responsible for deleting a user.
// It redirects to the user list page.
func (c *Controller) UserDeleteViewController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDVariable := vars["userId"]
	// it has to be converted to int64
	userID, err := strconv.ParseInt(userIDVariable, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id "+err.Error(), http.StatusBadRequest)
		return
	}
	// delete the user
	err = c.userRepository.DeleteUser(userID)
	if err != nil {
		http.Error(w, "Internal server error - failed to delete the user "+err.Error(), http.StatusInternalServerError)
		return
	}
	// redirect to the user list
	http.Redirect(w, r, "/admin/user/list", http.StatusSeeOther)
}

// UserDeleteAPIController is the controller for the user delete API.
// It is responsible for deleting a user.
// Example request:
// curl -X DELETE http://localhost:8090/api/user/delete/1
func (c *Controller) UserDeleteAPIController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDVariable := vars["userId"]
	// it has to be converted to int64
	userID, err := strconv.ParseInt(userIDVariable, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id "+err.Error(), http.StatusBadRequest)
		return
	}
	// delete the user
	err = c.userRepository.DeleteUser(userID)
	if err != nil {
		http.Error(w, "Internal server error - failed to delete the user "+err.Error(), http.StatusInternalServerError)
		return
	}
	// return success
	c.renderer.JSON(w, http.StatusOK, "Success")
}

// UserListViewController is the controller for the user list view.
func (c *Controller) UserListViewController(w http.ResponseWriter, r *http.Request) {
	// get all users
	users, err := c.userRepository.GetUsers()
	if err != nil {
		http.Error(w, "Failed to get user data "+err.Error(), http.StatusInternalServerError)
		return
	}
	template := c.renderer.BuildTemplate("user-list", []string{c.renderer.GetTemplateDirectoryPath() + "/user/list.html.tmpl"})
	content := struct {
		Title string
		Users []*model.User
	}{
		Title: "User List",
		Users: users,
	}
	err = template.ExecuteTemplate(w, "base.html", content)
	if err != nil {
		panic(err)
	}
}

// UserListAPIController is the controller for the user list API.
// It is responsible for returning all users.
// Example request:
// curl -X GET http://localhost:8090/api/user/list
func (c *Controller) UserListAPIController(w http.ResponseWriter, r *http.Request) {
	// get all users
	users, err := c.userRepository.GetUsers()
	if err != nil {
		http.Error(w, "Internal server error - failed to get the users "+err.Error(), http.StatusInternalServerError)
		return
	}
	// return the users as JSON
	c.renderer.JSON(w, http.StatusOK, users)
}
