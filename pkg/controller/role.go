package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/model"
)

// RoleViewController is the controller for the role view page.
// GET /admin/role/view/{roleId}
// It renders the role view page.
func (c *Controller) RoleViewController(w http.ResponseWriter, r *http.Request) {
	template := c.renderer.BuildTemplate("role-view", []string{c.renderer.GetTemplateDirectoryPath() + "/role/view.html.tmpl"})
	role, statusCode, err := c.roleViewData(r)
	if err != nil {
		http.Error(w, "Failed to get role data "+err.Error(), statusCode)
		return
	}
	content := struct {
		Title string
		Role  *model.Role
	}{
		Title: "Role View",
		Role:  role,
	}
	err = template.ExecuteTemplate(w, "base.html", content)
	if err != nil {
		panic(err)
	}
}

// roleViewData gets the request as input, and returns the role data, status code and error.
func (c *Controller) roleViewData(r *http.Request) (*model.Role, int, error) {
	vars := mux.Vars(r)
	roleIDVariable := vars["roleId"]
	// it has to be converted to int64
	roleID, err := strconv.ParseInt(roleIDVariable, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	role, err := c.roleRepository.GetRoleByID(roleID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return role, http.StatusOK, nil
}

// RoleCreateViewController is the controller for the role create view.
// On case of get request, it returns the role create page.
// On case of post request, it creates the role and redirects to the list page.
func (c *Controller) RoleCreateViewController(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		template := c.renderer.BuildTemplate("role-create", []string{c.renderer.GetTemplateDirectoryPath() + "/role/create.html.tmpl"})
		content := struct {
			Title string
		}{
			Title: "Role Create",
		}
		err := template.ExecuteTemplate(w, "base.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the role
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		_, err := c.roleRepository.CreateRole(name, []int64{})
		if err != nil {
			http.Error(w, "Internal server error - failed to create the role "+err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/role/list", http.StatusSeeOther)
		return
	}
}

// RoleUpdateViewController is the controller for the role update view.
// On case of get request, it returns the role update page.
// On case of post request, it updates the role and redirects to the list page.
func (c *Controller) RoleUpdateViewController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleIDVariable := vars["roleId"]
	// it has to be converted to int64
	roleID, err := strconv.ParseInt(roleIDVariable, 10, 64)
	if err != nil {
		http.Error(w, "Invalid role id "+err.Error(), http.StatusBadRequest)
		return
	}

	// get the role
	role, err := c.roleRepository.GetRoleByID(roleID)
	if err != nil {
		http.Error(w, "Internal server error - failed to get the role "+err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		template := c.renderer.BuildTemplate("user-role", []string{c.renderer.GetTemplateDirectoryPath() + "/role/update.html.tmpl"})
		content := struct {
			Title string
			Role  *model.Role
		}{
			Title: "Role Update",
			Role:  role,
		}
		err = template.ExecuteTemplate(w, "base.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the role
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		// update the us
		role.Name = name
		err = c.roleRepository.UpdateRole(role)
		if err != nil {
			http.Error(w, "Internal server error - failed to update the role "+err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/role/list", http.StatusSeeOther)
		return
	}
}

// RoleDeleteViewController is the controller for the role delete form.
// It is responsible for deleting a role.
// It redirects to the role list page.
func (c *Controller) RoleDeleteViewController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleIDVariable := vars["roleId"]
	// it has to be converted to int64
	roleID, err := strconv.ParseInt(roleIDVariable, 10, 64)
	if err != nil {
		http.Error(w, "Invalid role id "+err.Error(), http.StatusBadRequest)
		return
	}
	// delete the role
	err = c.roleRepository.DeleteRole(roleID)
	if err != nil {
		http.Error(w, "Internal server error - failed to delete the role "+err.Error(), http.StatusInternalServerError)
		return
	}
	// redirect to the role list
	http.Redirect(w, r, "/admin/role/list", http.StatusSeeOther)
}

// RoleListViewController is the controller for the role list view.
func (c *Controller) RoleListViewController(w http.ResponseWriter, r *http.Request) {
	// get all roles
	roles, err := c.roleRepository.GetRoles()
	if err != nil {
		http.Error(w, "Failed to get role data "+err.Error(), http.StatusInternalServerError)
		return
	}
	template := c.renderer.BuildTemplate("role-list", []string{c.renderer.GetTemplateDirectoryPath() + "/role/list.html.tmpl"})
	content := struct {
		Title string
		Roles []*model.Role
	}{
		Title: "User List",
		Roles: roles,
	}
	err = template.ExecuteTemplate(w, "base.html", content)
	if err != nil {
		panic(err)
	}
}
