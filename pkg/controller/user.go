package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/passwd"
)

var (
	// UserCreateRequiredFieldMissing is the error message for the required fields in the user create.
	UserCreateRequiredFieldMissing = "Name, email, password and role are required"
	// UserCreateCreateUserErrorMessagePrefix is the error message prefix for the failed user creation.
	UserCreateCreateUserErrorMessagePrefix = "Internal server error - failed to create the user "
	// UserUpdateFailedToGetUserErrorMessage is the error message for the failed user get.
	UserUpdateFailedToGetUserErrorMessage = "Internal server error - failed to get the user "
	// UserUpdateRequiredFieldMissing is the error message for the required fields in the user update.
	UserUpdateRequiredFieldMissing = "Name, email and role are required"
	// UserUpdateFailedToUpdateUserErrorMessage is the error message for the failed user update.
	UserUpdateFailedToUpdateUserErrorMessage = "Internal server error - failed to update the user "
	// UserDeleteFailedToDeleteErrorMessage is the error message for the failed user deletion.
	UserDeleteFailedToDeleteErrorMessage = "Internal server error - failed to delete the user "
	// UserListFailedToGetUsersErrorMessage is the error message for the failed users get.
	UserListFailedToGetUsersErrorMessage = "Internal server error - failed to get the users "
	// UserPasswordEncriptionFailedErrorMessage is the error message for the failed password encryption.
	UserPasswordEncriptionFailedErrorMessage = "Internal server error - failed to encrypt the password "
	// UserRoleIDInvalidErrorMessagePrefix is the error message prefix for the invalid role id.
	UserRoleIDInvalidErrorMessagePrefix = "Invalid role id "
	// UserUserIDInvalidErrorMessagePrefix is the error message prefix for the invalid user id.
	UserUserIDInvalidErrorMessagePrefix = "Invalid user id "
	// UserFailedToGetUserErrorMessage is the error message for the failed user get.
	UserFailedToGetUserErrorMessage = "Failed to get user data "
)

// UserViewController is the controller for the user view page.
func (c *Controller) UserViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("users.view") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	template := c.renderer.BuildTemplate("login", []string{c.renderer.GetTemplateDirectoryPath() + "/user/view.html.tmpl"})
	u, statusCode, err := c.userViewData(r)
	if err != nil {
		http.Error(w, UserFailedToGetUserErrorMessage+err.Error(), statusCode)
		return
	}
	content := struct {
		Title       string
		User        *model.User
		CurrentUser *model.User
	}{
		Title:       "User View",
		User:        u,
		CurrentUser: currentUser,
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
		http.Error(w, UserFailedToGetUserErrorMessage+err.Error(), statusCode)
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
// On case of post request, it creates the user and redirects to the list page.
func (c *Controller) UserCreateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("users.create") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	if r.Method == http.MethodGet {
		template := c.renderer.BuildTemplate("user-create", []string{c.renderer.GetTemplateDirectoryPath() + "/user/create.html.tmpl"})
		// get all roles
		roles, err := c.roleRepository.GetRoles()
		if err != nil {
			http.Error(w, "Failed to get role data "+err.Error(), http.StatusInternalServerError)
			return
		}
		content := struct {
			Title       string
			Roles       []*model.Role
			CurrentUser *model.User
		}{
			Title:       "User Create",
			Roles:       roles,
			CurrentUser: currentUser,
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
		roleIDRaw := r.FormValue("role")

		// if the name or email is empty, return an error
		if name == "" || email == "" || password == "" || roleIDRaw == "" {
			http.Error(w, UserCreateRequiredFieldMissing, http.StatusBadRequest)
			return
		}

		password, err := passwd.HashPassword(password)
		if err != nil {
			http.Error(w, UserPasswordEncriptionFailedErrorMessage+err.Error(), http.StatusInternalServerError)
			return
		}
		// it has to be converted to int64
		roleID, err := strconv.ParseInt(roleIDRaw, 10, 64)
		if err != nil {
			http.Error(w, UserRoleIDInvalidErrorMessagePrefix+err.Error(), http.StatusBadRequest)
			return
		}
		_, err = c.userRepository.CreateUser(name, email, string(password), roleID)
		if err != nil {
			http.Error(w, UserCreateCreateUserErrorMessagePrefix+err.Error(), http.StatusInternalServerError)
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
// curl -X POST http://localhost:8090/api/user/create -d "name=Admin&email=system@admin&password=admin"
// Example response:
// {"ID":1,"Name":"Admin","Email":"system@admin","CreatedAt":"2024-05-27T15:12:24.894037Z","UpdatedAt":"2024-05-27T15:12:24.894037Z","Password":"$2a$10$8QIzpaZqZEEI3RVKKjGnh.GJ3DqLEIewuuRMGGCnRD3VW3v7ZodUW"}
func (c *Controller) UserCreateAPIController(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	roleIDRaw := r.FormValue("role")
	hashedPassword, err := passwd.HashPassword(password)
	if err != nil {
		http.Error(w, UserPasswordEncriptionFailedErrorMessage+err.Error(), http.StatusInternalServerError)
		return
	}
	// it has to be converted to int64
	roleID, err := strconv.ParseInt(roleIDRaw, 10, 64)
	if err != nil {
		http.Error(w, UserRoleIDInvalidErrorMessagePrefix+err.Error(), http.StatusBadRequest)
		return
	}
	// create the user
	user, err := c.userRepository.CreateUser(name, email, string(hashedPassword), roleID)
	if err != nil {
		http.Error(w, UserCreateCreateUserErrorMessagePrefix+err.Error(), http.StatusInternalServerError)
		return
	}
	// return the user as JSON
	c.renderer.JSON(w, http.StatusOK, user)
}

// UserUpdateViewController is the controller for the user update view.
// On case of get request, it returns the user update page.
// On case of post request, it updates the user and redirects to the list page.
func (c *Controller) UserUpdateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("users.update") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	vars := mux.Vars(r)
	userIDVariable := vars["userId"]
	// it has to be converted to int64
	userID, err := strconv.ParseInt(userIDVariable, 10, 64)
	if err != nil {
		http.Error(w, UserUserIDInvalidErrorMessagePrefix+err.Error(), http.StatusBadRequest)
		return
	}

	// get the user
	user, err := c.userRepository.GetUserByID(userID)
	if err != nil {
		http.Error(w, UserUpdateFailedToGetUserErrorMessage+err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		template := c.renderer.BuildTemplate("user-update", []string{c.renderer.GetTemplateDirectoryPath() + "/user/update.html.tmpl"})
		// get all roles
		roles, err := c.roleRepository.GetRoles()
		if err != nil {
			http.Error(w, "Failed to get role data "+err.Error(), http.StatusInternalServerError)
			return
		}
		content := struct {
			Title       string
			User        *model.User
			Roles       []*model.Role
			CurrentUser *model.User
		}{
			Title:       "User Update",
			User:        user,
			Roles:       roles,
			CurrentUser: currentUser,
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
		roleIDRaw := r.FormValue("role")

		// if the name or email is empty, return an error
		if name == "" || email == "" || roleIDRaw == "" {
			http.Error(w, UserUpdateRequiredFieldMissing, http.StatusBadRequest)
			return
		}

		// if the password is not empty, encrypt it
		if password != "" {
			password, err := passwd.HashPassword(password)
			if err != nil {
				http.Error(w, UserPasswordEncriptionFailedErrorMessage+err.Error(), http.StatusInternalServerError)
				return
			}
			user.Password = password
		}
		// update the user
		user.Name = name
		user.Email = email
		// roleID has to be converted to int64
		roleID, err := strconv.ParseInt(roleIDRaw, 10, 64)
		if err != nil {
			http.Error(w, UserRoleIDInvalidErrorMessagePrefix+err.Error(), http.StatusBadRequest)
			return
		}
		user.Role.ID = roleID
		err = c.userRepository.UpdateUser(user)
		if err != nil {
			http.Error(w, UserUpdateFailedToUpdateUserErrorMessage+err.Error(), http.StatusInternalServerError)
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
		http.Error(w, UserUserIDInvalidErrorMessagePrefix+err.Error(), http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	roleIDRaw := r.FormValue("role")
	// get the user
	user, err := c.userRepository.GetUserByID(userID)
	if err != nil {
		http.Error(w, UserUpdateFailedToGetUserErrorMessage+err.Error(), http.StatusInternalServerError)
		return
	}
	hashedPassword, err := passwd.HashPassword(password)
	if err != nil {
		http.Error(w, UserPasswordEncriptionFailedErrorMessage+err.Error(), http.StatusInternalServerError)
		return
	}
	// update the user
	user.Name = name
	user.Email = email
	user.Password = hashedPassword
	// roleID has to be converted to int64
	roleID, err := strconv.ParseInt(roleIDRaw, 10, 64)
	if err != nil {
		http.Error(w, UserRoleIDInvalidErrorMessagePrefix+err.Error(), http.StatusBadRequest)
		return
	}
	user.Role.ID = roleID
	err = c.userRepository.UpdateUser(user)
	if err != nil {
		http.Error(w, UserUpdateFailedToUpdateUserErrorMessage+err.Error(), http.StatusInternalServerError)
		return
	}
	// return the updated user as JSON
	c.renderer.JSON(w, http.StatusOK, user)
}

// UserDeleteViewController is the controller for the user delete view.
// It is responsible for deleting a user.
// It redirects to the user list page.
func (c *Controller) UserDeleteViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("users.delete") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	vars := mux.Vars(r)
	userIDVariable := vars["userId"]
	// it has to be converted to int64
	userID, err := strconv.ParseInt(userIDVariable, 10, 64)
	if err != nil {
		http.Error(w, UserUserIDInvalidErrorMessagePrefix+err.Error(), http.StatusBadRequest)
		return
	}
	// delete the user
	err = c.userRepository.DeleteUser(userID)
	if err != nil {
		http.Error(w, UserDeleteFailedToDeleteErrorMessage+err.Error(), http.StatusInternalServerError)
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
		http.Error(w, UserUserIDInvalidErrorMessagePrefix+err.Error(), http.StatusBadRequest)
		return
	}
	// delete the user
	err = c.userRepository.DeleteUser(userID)
	if err != nil {
		http.Error(w, UserDeleteFailedToDeleteErrorMessage+err.Error(), http.StatusInternalServerError)
		return
	}
	// return success
	c.renderer.JSON(w, http.StatusOK, "Success")
}

// UserListViewController is the controller for the user list view.
func (c *Controller) UserListViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("users.view") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	// get all users
	users, err := c.userRepository.GetUsers()
	if err != nil {
		http.Error(w, UserFailedToGetUserErrorMessage+err.Error(), http.StatusInternalServerError)
		return
	}
	template := c.renderer.BuildTemplate("user-list", []string{c.renderer.GetTemplateDirectoryPath() + "/user/list.html.tmpl"})
	content := struct {
		Title       string
		Users       []*model.User
		CurrentUser *model.User
	}{
		Title:       "User List",
		Users:       users,
		CurrentUser: currentUser,
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
		http.Error(w, UserListFailedToGetUsersErrorMessage+err.Error(), http.StatusInternalServerError)
		return
	}
	// return the users as JSON
	c.renderer.JSON(w, http.StatusOK, users)
}
