package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller/response"
	"github.com/akosgarai/projectregister/pkg/model"
)

// RuntimeViewController is the controller for the runtime view page.
// GET /admin/runtime/view/{runtimeId}
// It renders the runtime view page.
func (c *Controller) RuntimeViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("runtimes.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	runtime, statusCode, err := c.runtimeViewData(r)
	if err != nil {
		c.renderer.Error(w, statusCode, RuntimeFailedToGetRuntimeErrorMessage, err)
		return
	}
	content := response.NewRuntimeDetailResponse(currentUser, runtime)
	err = c.renderer.Template.RenderTemplate(w, "detail-page.html", content)
	if err != nil {
		panic(err)
	}
}

// runtimeViewData gets the request as input, and returns the runtime data, status code and error.
func (c *Controller) runtimeViewData(r *http.Request) (*model.Runtime, int, error) {
	vars := mux.Vars(r)
	runtimeIDVariable := vars["runtimeId"]
	// it has to be converted to int64
	runtimeID, err := strconv.ParseInt(runtimeIDVariable, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	runtime, err := c.runtimeRepository.GetRuntimeByID(runtimeID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return runtime, http.StatusOK, nil
}

// RuntimeCreateViewController is the controller for the runtime create view.
// On case of get request, it returns the runtime create page.
// On case of post request, it creates the runtime and redirects to the list page.
func (c *Controller) RuntimeCreateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("runtimes.create") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	if r.Method == http.MethodGet {
		content := response.NewCreateRuntimeResponse(currentUser)
		err := c.renderer.Template.RenderTemplate(w, "form-page.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the runtime
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, RuntimeCreateRequiredFieldMissing, nil)
			return
		}

		_, err := c.runtimeRepository.CreateRuntime(name)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, RuntimeCreateCreateRuntimeErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/runtime/list", http.StatusSeeOther)
		return
	}
}

// RuntimeUpdateViewController is the controller for the runtime update view.
// On case of get request, it returns the runtime update page.
// On case of post request, it updates the runtime and redirects to the list page.
func (c *Controller) RuntimeUpdateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("runtimes.update") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	runtimeIDVariable := vars["runtimeId"]
	// it has to be converted to int64
	runtimeID, err := strconv.ParseInt(runtimeIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, RuntimeRuntimeIDInvalidErrorMessage, err)
		return
	}

	// get the runtime
	runtime, err := c.runtimeRepository.GetRuntimeByID(runtimeID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, RuntimeFailedToGetRuntimeErrorMessage, err)
		return
	}

	if r.Method == http.MethodGet {
		content := response.NewUpdateRuntimeResponse(currentUser, runtime)
		err = c.renderer.Template.RenderTemplate(w, "form-page.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the runtime
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, RuntimeUpdateRequiredFieldMissing, nil)
			return
		}

		// update the runtime
		runtime.Name = name
		err = c.runtimeRepository.UpdateRuntime(runtime)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, RuntimeUpdateUpdateRuntimeErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/runtime/list", http.StatusSeeOther)
		return
	}
}

// RuntimeDeleteViewController is the controller for the runtime delete form.
// It is responsible for deleting a runtime.
// It redirects to the runtime list page.
func (c *Controller) RuntimeDeleteViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("runtimes.delete") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	runtimeIDVariable := vars["runtimeId"]
	// it has to be converted to int64
	runtimeID, err := strconv.ParseInt(runtimeIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, RuntimeRuntimeIDInvalidErrorMessage, err)
		return
	}
	// delete the runtime
	err = c.runtimeRepository.DeleteRuntime(runtimeID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, RuntimeDeleteFailedToDeleteErrorMessage, err)
		return
	}
	// redirect to the runtime list
	http.Redirect(w, r, "/admin/runtime/list", http.StatusSeeOther)
}

// RuntimeListViewController is the controller for the runtime list view.
func (c *Controller) RuntimeListViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("runtimes.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	// get all runtimes
	runtimes, err := c.runtimeRepository.GetRuntimes()
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, RuntimeListFailedToGetRuntimesErrorMessage, err)
		return
	}
	content := response.NewRuntimeListResponse(currentUser, runtimes)
	err = c.renderer.Template.RenderTemplate(w, "listing-page.html", content)
	if err != nil {
		panic(err)
	}
}
