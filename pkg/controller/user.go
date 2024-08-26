package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller/response"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/passwd"
)

// UserViewController is the controller for the user view page.
func (c *Controller) UserViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("users.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	u, statusCode, err := c.userViewData(r)
	if err != nil {
		c.renderer.Error(w, statusCode, UserFailedToGetUserErrorMessage, err)
		return
	}
	content := response.NewUserDetailResponse(currentUser, u)
	err = c.renderer.Template.RenderTemplate(w, "detail-page.html", content)
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
		c.renderer.Error(w, statusCode, UserFailedToGetUserErrorMessage, err)
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
	u, err := c.repositoryContainer.GetUserRepository().GetUserByID(userID)
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
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	if r.Method == http.MethodGet {
		// get all roles
		roles, err := c.repositoryContainer.GetRoleRepository().GetRoles(model.NewRoleFilter())
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, UserFailedToGetRolesErrorMessage, err)
			return
		}
		content := response.NewCreateUserResponse(currentUser, roles)
		err = c.renderer.Template.RenderTemplate(w, "form-page.html", content)
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
			c.renderer.Error(w, http.StatusBadRequest, UserCreateRequiredFieldMissing, nil)
			return
		}

		password, err := passwd.HashPassword(password)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, UserPasswordEncriptionFailedErrorMessage, err)
			return
		}
		// it has to be converted to int64
		roleID, err := strconv.ParseInt(roleIDRaw, 10, 64)
		if err != nil {
			c.renderer.Error(w, http.StatusBadRequest, UserRoleIDInvalidErrorMessagePrefix, err)
			return
		}
		_, err = c.repositoryContainer.GetUserRepository().CreateUser(name, email, string(password), roleID)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, UserCreateCreateUserErrorMessagePrefix, err)
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
		c.renderer.Error(w, http.StatusInternalServerError, UserPasswordEncriptionFailedErrorMessage, err)
		return
	}
	// it has to be converted to int64
	roleID, err := strconv.ParseInt(roleIDRaw, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, UserRoleIDInvalidErrorMessagePrefix, err)
		return
	}
	// create the user
	user, err := c.repositoryContainer.GetUserRepository().CreateUser(name, email, string(hashedPassword), roleID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, UserCreateCreateUserErrorMessagePrefix, err)
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
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	userIDVariable := vars["userId"]
	// it has to be converted to int64
	userID, err := strconv.ParseInt(userIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, UserUserIDInvalidErrorMessagePrefix, err)
		return
	}

	// get the user
	user, err := c.repositoryContainer.GetUserRepository().GetUserByID(userID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, UserUpdateFailedToGetUserErrorMessage, err)
		return
	}

	if r.Method == http.MethodGet {
		// get all roles
		roles, err := c.repositoryContainer.GetRoleRepository().GetRoles(model.NewRoleFilter())
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, UserFailedToGetRolesErrorMessage, err)
			return
		}
		content := response.NewUpdateUserResponse(currentUser, user, roles)
		err = c.renderer.Template.RenderTemplate(w, "form-page.html", content)
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
			c.renderer.Error(w, http.StatusBadRequest, UserUpdateRequiredFieldMissing, nil)
			return
		}

		// if the password is not empty, encrypt it
		if password != "" {
			password, err := passwd.HashPassword(password)
			if err != nil {
				c.renderer.Error(w, http.StatusInternalServerError, UserPasswordEncriptionFailedErrorMessage, err)
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
			c.renderer.Error(w, http.StatusBadRequest, UserRoleIDInvalidErrorMessagePrefix, err)
			return
		}
		user.Role.ID = roleID
		err = c.repositoryContainer.GetUserRepository().UpdateUser(user)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, UserUpdateFailedToUpdateUserErrorMessage, err)
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
		c.renderer.Error(w, http.StatusBadRequest, UserUserIDInvalidErrorMessagePrefix, err)
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	roleIDRaw := r.FormValue("role")
	// get the user
	user, err := c.repositoryContainer.GetUserRepository().GetUserByID(userID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, UserUpdateFailedToGetUserErrorMessage, err)
		return
	}
	hashedPassword, err := passwd.HashPassword(password)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, UserPasswordEncriptionFailedErrorMessage, err)
		return
	}
	// update the user
	user.Name = name
	user.Email = email
	user.Password = hashedPassword
	// roleID has to be converted to int64
	roleID, err := strconv.ParseInt(roleIDRaw, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, UserRoleIDInvalidErrorMessagePrefix, err)
		return
	}
	user.Role.ID = roleID
	err = c.repositoryContainer.GetUserRepository().UpdateUser(user)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, UserUpdateFailedToUpdateUserErrorMessage, err)
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
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	userIDVariable := vars["userId"]
	// it has to be converted to int64
	userID, err := strconv.ParseInt(userIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, UserUserIDInvalidErrorMessagePrefix, err)
		return
	}
	// delete the user
	err = c.repositoryContainer.GetUserRepository().DeleteUser(userID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, UserDeleteFailedToDeleteErrorMessage, err)
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
		c.renderer.Error(w, http.StatusBadRequest, UserUserIDInvalidErrorMessagePrefix, err)
		return
	}
	// delete the user
	err = c.repositoryContainer.GetUserRepository().DeleteUser(userID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, UserDeleteFailedToDeleteErrorMessage, err)
		return
	}
	// return success
	c.renderer.JSON(w, http.StatusOK, "Success")
}

// UserListViewController is the controller for the user list view.
func (c *Controller) UserListViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("users.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	// get all users
	users, err := c.repositoryContainer.GetUserRepository().GetUsers()
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, UserFailedToGetUserErrorMessage, err)
		return
	}
	content := response.NewUserListResponse(currentUser, users)
	err = c.renderer.Template.RenderTemplate(w, "listing-page.html", content)
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
	users, err := c.repositoryContainer.GetUserRepository().GetUsers()
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, UserListFailedToGetUsersErrorMessage, err)
		return
	}
	// return the users as JSON
	c.renderer.JSON(w, http.StatusOK, users)
}
