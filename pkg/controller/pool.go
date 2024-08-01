package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller/response"
	"github.com/akosgarai/projectregister/pkg/model"
)

// PoolViewController is the controller for the pool view page.
// GET /admin/pool/view/{poolId}
// It renders the pool view page.
func (c *Controller) PoolViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("pools.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	pool, statusCode, err := c.poolViewData(r)
	if err != nil {
		c.renderer.Error(w, statusCode, PoolFailedToGetPoolErrorMessage, err)
		return
	}
	content := response.NewPoolDetailResponse(currentUser, pool)
	err = c.renderer.Template.RenderTemplate(w, "pool-view.html", content)
	if err != nil {
		panic(err)
	}
}

// poolViewData gets the request as input, and returns the pool data, status code and error.
func (c *Controller) poolViewData(r *http.Request) (*model.Pool, int, error) {
	vars := mux.Vars(r)
	poolIDVariable := vars["poolId"]
	// it has to be converted to int64
	poolID, err := strconv.ParseInt(poolIDVariable, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	pool, err := c.poolRepository.GetPoolByID(poolID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return pool, http.StatusOK, nil
}

// PoolCreateViewController is the controller for the pool create view.
// On case of get request, it returns the pool create page.
// On case of post request, it creates the pool and redirects to the list page.
func (c *Controller) PoolCreateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("pools.create") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	if r.Method == http.MethodGet {
		content := response.NewPoolFormResponse("Pool Create", currentUser, &model.Pool{})
		err := c.renderer.Template.RenderTemplate(w, "pool-create.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the pool
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, PoolCreateRequiredFieldMissing, nil)
			return
		}

		_, err := c.poolRepository.CreatePool(name)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, PoolCreateCreatePoolErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/pool/list", http.StatusSeeOther)
		return
	}
}

// PoolUpdateViewController is the controller for the pool update view.
// On case of get request, it returns the pool update page.
// On case of post request, it updates the pool and redirects to the list page.
func (c *Controller) PoolUpdateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("pools.update") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	poolIDVariable := vars["poolId"]
	// it has to be converted to int64
	poolID, err := strconv.ParseInt(poolIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, PoolPoolIDInvalidErrorMessage, err)
		return
	}

	// get the pool
	pool, err := c.poolRepository.GetPoolByID(poolID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, PoolFailedToGetPoolErrorMessage, err)
		return
	}

	if r.Method == http.MethodGet {
		content := response.NewPoolFormResponse("Pool Update", currentUser, pool)
		err = c.renderer.Template.RenderTemplate(w, "pool-update.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the pool
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, PoolUpdateRequiredFieldMissing, nil)
			return
		}

		// update the pool
		pool.Name = name
		err = c.poolRepository.UpdatePool(pool)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, PoolUpdateUpdatePoolErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/pool/list", http.StatusSeeOther)
		return
	}
}

// PoolDeleteViewController is the controller for the pool delete form.
// It is responsible for deleting a pool.
// It redirects to the pool list page.
func (c *Controller) PoolDeleteViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("pools.delete") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	poolIDVariable := vars["poolId"]
	// it has to be converted to int64
	poolID, err := strconv.ParseInt(poolIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, PoolPoolIDInvalidErrorMessage, err)
		return
	}
	// delete the pool
	err = c.poolRepository.DeletePool(poolID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, PoolDeleteFailedToDeleteErrorMessage, err)
		return
	}
	// redirect to the pool list
	http.Redirect(w, r, "/admin/pool/list", http.StatusSeeOther)
}

// PoolListViewController is the controller for the pool list view.
func (c *Controller) PoolListViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("pools.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	// get all pools
	pools, err := c.poolRepository.GetPools()
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, PoolListFailedToGetPoolsErrorMessage, err)
		return
	}
	content := response.NewPoolListResponse(currentUser, pools)
	err = c.renderer.Template.RenderTemplate(w, "listing-page.html", content)
	if err != nil {
		panic(err)
	}
}
