package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller/response"
	"github.com/akosgarai/projectregister/pkg/model"
)

// FrameworkViewController is the controller for the framework view page.
// GET /admin/framework/view/{frameworkId}
// It renders the framework view page.
func (c *Controller) FrameworkViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("frameworks.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	framework, statusCode, err := c.frameworkViewData(r)
	if err != nil {
		c.renderer.Error(w, statusCode, FrameworkFailedToGetFrameworkErrorMessage, err)
		return
	}
	content := response.NewFrameworkDetailResponse(currentUser, framework)
	err = c.renderer.Template.RenderTemplate(w, "detail-page.html", content)
	if err != nil {
		panic(err)
	}
}

// frameworkViewData gets the request as input, and returns the framework data, status code and error.
func (c *Controller) frameworkViewData(r *http.Request) (*model.Framework, int, error) {
	vars := mux.Vars(r)
	frameworkIDVariable := vars["frameworkId"]
	// it has to be converted to int64
	frameworkID, err := strconv.ParseInt(frameworkIDVariable, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	framework, err := c.repositoryContainer.GetFrameworkRepository().GetFrameworkByID(frameworkID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return framework, http.StatusOK, nil
}

// FrameworkCreateViewController is the controller for the framework create view.
// On case of get request, it returns the framework create page.
// On case of post request, it creates the framework and redirects to the list page.
func (c *Controller) FrameworkCreateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("frameworks.create") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	if r.Method == http.MethodGet {
		content := response.NewCreateFrameworkResponse(currentUser)
		err := c.renderer.Template.RenderTemplate(w, "form-page.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the framework
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, FrameworkCreateRequiredFieldMissing, nil)
			return
		}
		scoreRaw := r.FormValue("score")
		score, err := strconv.Atoi(scoreRaw)
		if err != nil {
			score = 0
		}

		_, err = c.repositoryContainer.GetFrameworkRepository().CreateFramework(name, score)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, FrameworkCreateCreateFrameworkErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/framework/list", http.StatusSeeOther)
		return
	}
}

// FrameworkUpdateViewController is the controller for the framework update view.
// On case of get request, it returns the framework update page.
// On case of post request, it updates the framework and redirects to the list page.
func (c *Controller) FrameworkUpdateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("frameworks.update") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	frameworkIDVariable := vars["frameworkId"]
	// it has to be converted to int64
	frameworkID, err := strconv.ParseInt(frameworkIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, FrameworkFrameworkIDInvalidErrorMessage, err)
		return
	}

	// get the framework
	framework, err := c.repositoryContainer.GetFrameworkRepository().GetFrameworkByID(frameworkID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, FrameworkFailedToGetFrameworkErrorMessage, err)
		return
	}

	if r.Method == http.MethodGet {
		content := response.NewUpdateFrameworkResponse(currentUser, framework)
		err = c.renderer.Template.RenderTemplate(w, "form-page.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the framework
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, FrameworkUpdateRequiredFieldMissing, nil)
			return
		}
		scoreRaw := r.FormValue("score")
		score, err := strconv.Atoi(scoreRaw)
		if err != nil {
			score = 0
		}

		// update the framework
		framework.Name = name
		framework.Score = score
		err = c.repositoryContainer.GetFrameworkRepository().UpdateFramework(framework)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, FrameworkUpdateUpdateFrameworkErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/framework/list", http.StatusSeeOther)
		return
	}
}

// FrameworkDeleteViewController is the controller for the framework delete form.
// It is responsible for deleting a framework.
// It redirects to the framework list page.
func (c *Controller) FrameworkDeleteViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("frameworks.delete") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	frameworkIDVariable := vars["frameworkId"]
	// it has to be converted to int64
	frameworkID, err := strconv.ParseInt(frameworkIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, FrameworkFrameworkIDInvalidErrorMessage, err)
		return
	}
	// delete the framework
	err = c.repositoryContainer.GetFrameworkRepository().DeleteFramework(frameworkID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, FrameworkDeleteFailedToDeleteErrorMessage, err)
		return
	}
	// redirect to the framework list
	http.Redirect(w, r, "/admin/framework/list", http.StatusSeeOther)
}

// FrameworkListViewController is the controller for the framework list view.
func (c *Controller) FrameworkListViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("frameworks.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	filter := model.NewFrameworkFilter()
	if r.Method == http.MethodPost {
		filter.Name = r.FormValue("name")
	}
	// get all frameworks
	frameworks, err := c.repositoryContainer.GetFrameworkRepository().GetFrameworks(filter)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, FrameworkListFailedToGetFrameworksErrorMessage, err)
		return
	}
	content := response.NewFrameworkListResponse(currentUser, frameworks, filter)
	err = c.renderer.Template.RenderTemplate(w, "listing-page.html", content)
	if err != nil {
		panic(err)
	}
}
