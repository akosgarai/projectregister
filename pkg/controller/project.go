package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller/response"
	"github.com/akosgarai/projectregister/pkg/model"
)

// ProjectViewController is the controller for the project view page.
// GET /admin/project/view/{projectId}
// It renders the project view page.
func (c *Controller) ProjectViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("projects.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	project, statusCode, err := c.projectViewData(r)
	if err != nil {
		c.renderer.Error(w, statusCode, ProjectFailedToGetProjectErrorMessage, err)
		return
	}
	content := response.NewProjectDetailResponse(currentUser, project)
	err = c.renderer.Template.RenderTemplate(w, "detail-page.html", content)
	if err != nil {
		panic(err)
	}
}

// projectViewData gets the request as input, and returns the project data, status code and error.
func (c *Controller) projectViewData(r *http.Request) (*model.Project, int, error) {
	vars := mux.Vars(r)
	projectIDVariable := vars["projectId"]
	// it has to be converted to int64
	projectID, err := strconv.ParseInt(projectIDVariable, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	project, err := c.projectRepository.GetProjectByID(projectID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return project, http.StatusOK, nil
}

// ProjectCreateViewController is the controller for the project create view.
// On case of get request, it returns the project create page.
// On case of post request, it creates the project and redirects to the list page.
func (c *Controller) ProjectCreateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("projects.create") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	if r.Method == http.MethodGet {
		content := response.NewCreateProjectResponse(currentUser)
		err := c.renderer.Template.RenderTemplate(w, "form-page.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the project
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, ProjectCreateRequiredFieldMissing, nil)
			return
		}

		_, err := c.projectRepository.CreateProject(name)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, ProjectCreateCreateProjectErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/project/list", http.StatusSeeOther)
		return
	}
}

// ProjectUpdateViewController is the controller for the project update view.
// On case of get request, it returns the project update page.
// On case of post request, it updates the project and redirects to the list page.
func (c *Controller) ProjectUpdateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("projects.update") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	projectIDVariable := vars["projectId"]
	// it has to be converted to int64
	projectID, err := strconv.ParseInt(projectIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, ProjectProjectIDInvalidErrorMessage, err)
		return
	}

	// get the project
	project, err := c.projectRepository.GetProjectByID(projectID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ProjectFailedToGetProjectErrorMessage, err)
		return
	}

	if r.Method == http.MethodGet {
		content := response.NewUpdateProjectResponse(currentUser, project)
		err = c.renderer.Template.RenderTemplate(w, "form-page.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the project
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, ProjectUpdateRequiredFieldMissing, nil)
			return
		}

		// update the project
		project.Name = name
		err = c.projectRepository.UpdateProject(project)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, ProjectUpdateUpdateProjectErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/project/list", http.StatusSeeOther)
		return
	}
}

// ProjectDeleteViewController is the controller for the project delete form.
// It is responsible for deleting a project.
// It redirects to the project list page.
func (c *Controller) ProjectDeleteViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("projects.delete") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	projectIDVariable := vars["projectId"]
	// it has to be converted to int64
	projectID, err := strconv.ParseInt(projectIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, ProjectProjectIDInvalidErrorMessage, err)
		return
	}
	// delete the project
	err = c.projectRepository.DeleteProject(projectID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ProjectDeleteFailedToDeleteErrorMessage, err)
		return
	}
	// redirect to the project list
	http.Redirect(w, r, "/admin/project/list", http.StatusSeeOther)
}

// ProjectListViewController is the controller for the project list view.
func (c *Controller) ProjectListViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("projects.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	// get all projects
	projects, err := c.projectRepository.GetProjects()
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ProjectListFailedToGetProjectsErrorMessage, err)
		return
	}
	content := response.NewProjectListResponse(currentUser, projects)
	err = c.renderer.Template.RenderTemplate(w, "listing-page.html", content)
	if err != nil {
		panic(err)
	}
}
