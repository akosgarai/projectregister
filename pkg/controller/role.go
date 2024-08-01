package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller/response"
	"github.com/akosgarai/projectregister/pkg/model"
)

// RoleViewController is the controller for the role view page.
// GET /admin/role/view/{roleId}
// It renders the role view page.
func (c *Controller) RoleViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("roles.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	role, statusCode, err := c.roleViewData(r)
	if err != nil {
		c.renderer.Error(w, statusCode, RoleFailedToGetRoleErrorMessage, err)
		return
	}
	content := response.NewRoleDetailResponse(currentUser, role)
	err = c.renderer.Template.RenderTemplate(w, "detail-page.html", content)
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
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("roles.create") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	if r.Method == http.MethodGet {
		resources, err := c.resourceRepository.GetResources()
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, RoleFailedToGetResourcesErrorMessage, err)
			return
		}
		content := response.NewRoleFormResponse("Role Create", currentUser, &model.Role{}, resources)
		err = c.renderer.Template.RenderTemplate(w, "role-create.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the role
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, RoleCreateRequiredFieldMissing, nil)
			return
		}
		resources := r.Form["resources"]
		// transform the resources to int64
		var resourceIDs []int64
		for _, resource := range resources {
			resourceID, err := strconv.ParseInt(resource, 10, 64)
			if err != nil {
				c.renderer.Error(w, http.StatusBadRequest, RoleResourceIDInvalidErrorMessage, err)
				return
			}
			resourceIDs = append(resourceIDs, resourceID)
		}

		_, err := c.roleRepository.CreateRole(name, resourceIDs)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, RoleCreateCreateRoleErrorMessage, err)
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
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("roles.update") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	roleIDVariable := vars["roleId"]
	// it has to be converted to int64
	roleID, err := strconv.ParseInt(roleIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, RoleRoleIDInvalidErrorMessage, err)
		return
	}

	// get the role
	role, err := c.roleRepository.GetRoleByID(roleID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, RoleFailedToGetRoleErrorMessage, err)
		return
	}

	if r.Method == http.MethodGet {
		resources, err := c.resourceRepository.GetResources()
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, RoleFailedToGetResourcesErrorMessage, err)
			return
		}
		content := response.NewRoleFormResponse("Role Update", currentUser, role, resources)
		err = c.renderer.Template.RenderTemplate(w, "role-update.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the role
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, RoleUpdateRequiredFieldMissing, nil)
			return
		}

		// update the role
		role.Name = name
		resources := r.Form["resources"]
		// transform the resources to int64
		var resourceIDs []int64
		for _, resource := range resources {
			resourceID, err := strconv.ParseInt(resource, 10, 64)
			if err != nil {
				c.renderer.Error(w, http.StatusBadRequest, RoleResourceIDInvalidErrorMessage, err)
				return
			}
			resourceIDs = append(resourceIDs, resourceID)
		}
		err = c.roleRepository.UpdateRole(role, resourceIDs)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, RoleUpdateUpdateRoleErrorMessage, err)
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
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("roles.delete") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	roleIDVariable := vars["roleId"]
	// it has to be converted to int64
	roleID, err := strconv.ParseInt(roleIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, RoleRoleIDInvalidErrorMessage, err)
		return
	}
	// delete the role
	err = c.roleRepository.DeleteRole(roleID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, RoleDeleteFailedToDeleteErrorMessage, err)
		return
	}
	// redirect to the role list
	http.Redirect(w, r, "/admin/role/list", http.StatusSeeOther)
}

// RoleListViewController is the controller for the role list view.
func (c *Controller) RoleListViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("roles.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	// get all roles
	roles, err := c.roleRepository.GetRoles()
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, RoleListFailedToGetRolesErrorMessage, err)
		return
	}
	content := response.NewRoleListResponse(currentUser, roles)
	err = c.renderer.Template.RenderTemplate(w, "listing-page.html", content)
	if err != nil {
		panic(err)
	}
}
