package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller/response"
	"github.com/akosgarai/projectregister/pkg/model"
)

// EnvironmentViewController is the controller for the environment view page.
// GET /admin/environment/view/{environmentId}
// It renders the environment view page.
func (c *Controller) EnvironmentViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("environments.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	environment, statusCode, err := c.environmentViewData(r)
	if err != nil {
		c.renderer.Error(w, statusCode, EnvironmentFailedToGetEnvironmentErrorMessage, err)
		return
	}
	content := response.NewEnvironmentDetailResponse(currentUser, environment)
	err = c.renderer.Template.RenderTemplate(w, "environment-view.html", content)
	if err != nil {
		panic(err)
	}
}

// environmentViewData gets the request as input, and returns the environment data, status code and error.
func (c *Controller) environmentViewData(r *http.Request) (*model.Environment, int, error) {
	vars := mux.Vars(r)
	environmentIDVariable := vars["environmentId"]
	// it has to be converted to int64
	environmentID, err := strconv.ParseInt(environmentIDVariable, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	environment, err := c.environmentRepository.GetEnvironmentByID(environmentID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return environment, http.StatusOK, nil
}

// EnvironmentCreateViewController is the controller for the environment create view.
// On case of get request, it returns the environment create page.
// On case of post request, it creates the environment and redirects to the list page.
func (c *Controller) EnvironmentCreateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("environments.create") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	if r.Method == http.MethodGet {
		servers, err := c.serverRepository.GetServers()
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, EnvironmentCreateFailedToGetServersErrorMessage, err)
			return
		}
		databases, err := c.databaseRepository.GetDatabases()
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, EnvironmentCreateFailedToGetDatabasesErrorMessage, err)
			return
		}
		content := response.NewEnvironmentFormResponse("Environment Create", currentUser, &model.Environment{}, servers, databases)
		err = c.renderer.Template.RenderTemplate(w, "environment-create.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the environment
		name := r.FormValue("name")
		description := r.FormValue("description")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, EnvironmentCreateRequiredFieldMissing, nil)
			return
		}
		serverIDsRaw := r.Form["server_id"]
		var serverIDs []int64
		for _, serverIDRaw := range serverIDsRaw {
			serverID, err := strconv.ParseInt(serverIDRaw, 10, 64)
			if err != nil {
				c.renderer.Error(w, http.StatusBadRequest, EnvironmentCreateServerIDInvalidErrorMessage, err)
				return
			}
			serverIDs = append(serverIDs, serverID)
		}

		databaseIDsRaw := r.Form["database_id"]
		var databaseIDs []int64
		for _, databaseIDRaw := range databaseIDsRaw {
			databaseID, err := strconv.ParseInt(databaseIDRaw, 10, 64)
			if err != nil {
				c.renderer.Error(w, http.StatusBadRequest, EnvironmentCreateDatabaseIDInvalidErrorMessage, err)
				return
			}
			databaseIDs = append(databaseIDs, databaseID)
		}

		_, err := c.environmentRepository.CreateEnvironment(name, description, serverIDs, databaseIDs)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, EnvironmentCreateCreateEnvironmentErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/environment/list", http.StatusSeeOther)
		return
	}
}

// EnvironmentUpdateViewController is the controller for the environment update view.
// On case of get request, it returns the environment update page.
// On case of post request, it updates the environment and redirects to the list page.
func (c *Controller) EnvironmentUpdateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("environments.update") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	environmentIDVariable := vars["environmentId"]
	// it has to be converted to int64
	environmentID, err := strconv.ParseInt(environmentIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, EnvironmentEnvironmentIDInvalidErrorMessage, err)
		return
	}

	// get the environment
	environment, err := c.environmentRepository.GetEnvironmentByID(environmentID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, EnvironmentFailedToGetEnvironmentErrorMessage, err)
		return
	}

	if r.Method == http.MethodGet {
		servers, err := c.serverRepository.GetServers()
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, EnvironmentUpdateFailedToGetServersErrorMessage, err)
			return
		}
		databases, err := c.databaseRepository.GetDatabases()
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, EnvironmentUpdateFailedToGetDatabasesErrorMessage, err)
			return
		}
		content := response.NewEnvironmentFormResponse("Environment Update", currentUser, environment, servers, databases)
		err = c.renderer.Template.RenderTemplate(w, "environment-update.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the environment
		name := r.FormValue("name")
		description := r.FormValue("description")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, EnvironmentUpdateRequiredFieldMissing, nil)
			return
		}

		serverIDsRaw := r.Form["server_id"]
		environment.Servers = make([]*model.Server, len(serverIDsRaw))
		for i, serverIDRaw := range serverIDsRaw {
			serverID, err := strconv.ParseInt(serverIDRaw, 10, 64)
			if err != nil {
				c.renderer.Error(w, http.StatusBadRequest, EnvironmentUpdateServerIDInvalidErrorMessage, err)
				return
			}
			environment.Servers[i] = &model.Server{ID: serverID}
		}

		databaseIDsRaw := r.Form["database_id"]
		environment.Databases = make([]*model.Database, len(databaseIDsRaw))
		for i, databaseIDRaw := range databaseIDsRaw {
			databaseID, err := strconv.ParseInt(databaseIDRaw, 10, 64)
			if err != nil {
				c.renderer.Error(w, http.StatusBadRequest, EnvironmentUpdateDatabaseIDInvalidErrorMessage, err)
				return
			}
			environment.Databases[i] = &model.Database{ID: databaseID}
		}

		// update the environment
		environment.Name = name
		environment.Description = description
		err = c.environmentRepository.UpdateEnvironment(environment)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, EnvironmentUpdateUpdateEnvironmentErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/environment/list", http.StatusSeeOther)
		return
	}
}

// EnvironmentDeleteViewController is the controller for the environment delete form.
// It is responsible for deleting a environment.
// It redirects to the environment list page.
func (c *Controller) EnvironmentDeleteViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("environments.delete") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	environmentIDVariable := vars["environmentId"]
	// it has to be converted to int64
	environmentID, err := strconv.ParseInt(environmentIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, EnvironmentEnvironmentIDInvalidErrorMessage, err)
		return
	}
	// delete the environment
	err = c.environmentRepository.DeleteEnvironment(environmentID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, EnvironmentDeleteFailedToDeleteErrorMessage, err)
		return
	}
	// redirect to the environment list
	http.Redirect(w, r, "/admin/environment/list", http.StatusSeeOther)
}

// EnvironmentListViewController is the controller for the environment list view.
func (c *Controller) EnvironmentListViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("environments.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	// get all environments
	environments, err := c.environmentRepository.GetEnvironments()
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, EnvironmentListFailedToGetEnvironmentsErrorMessage, err)
		return
	}
	content := response.NewEnvironmentListResponse(currentUser, environments)
	err = c.renderer.Template.RenderTemplate(w, "listing-page.html", content)
	if err != nil {
		panic(err)
	}
}
