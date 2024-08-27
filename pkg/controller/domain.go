package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller/response"
	"github.com/akosgarai/projectregister/pkg/model"
)

// DomainViewController is the controller for the domain view page.
// GET /admin/domain/view/{domainId}
// It renders the domain view page.
func (c *Controller) DomainViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("domains.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	domain, statusCode, err := c.domainViewData(r)
	if err != nil {
		c.renderer.Error(w, statusCode, DomainFailedToGetDomainErrorMessage, err)
		return
	}
	content := response.NewDomainDetailResponse(currentUser, domain)
	err = c.renderer.Template.RenderTemplate(w, "detail-page.html", content)
	if err != nil {
		panic(err)
	}
}

// domainViewData gets the request as input, and returns the domain data, status code and error.
func (c *Controller) domainViewData(r *http.Request) (*model.Domain, int, error) {
	vars := mux.Vars(r)
	domainIDVariable := vars["domainId"]
	// it has to be converted to int64
	domainID, err := strconv.ParseInt(domainIDVariable, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	domain, err := c.repositoryContainer.GetDomainRepository().GetDomainByID(domainID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return domain, http.StatusOK, nil
}

// DomainCreateViewController is the controller for the domain create view.
// On case of get request, it returns the domain create page.
// On case of post request, it creates the domain and redirects to the list page.
func (c *Controller) DomainCreateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("domains.create") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	if r.Method == http.MethodGet {
		content := response.NewCreateDomainResponse(currentUser)
		err := c.renderer.Template.RenderTemplate(w, "form-page.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the domain
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, DomainCreateRequiredFieldMissing, nil)
			return
		}

		_, err := c.repositoryContainer.GetDomainRepository().CreateDomain(name)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, DomainCreateCreateDomainErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/domain/list", http.StatusSeeOther)
		return
	}
}

// DomainUpdateViewController is the controller for the domain update view.
// On case of get request, it returns the domain update page.
// On case of post request, it updates the domain and redirects to the list page.
func (c *Controller) DomainUpdateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("domains.update") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	domainIDVariable := vars["domainId"]
	// it has to be converted to int64
	domainID, err := strconv.ParseInt(domainIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, DomainDomainIDInvalidErrorMessage, err)
		return
	}

	// get the domain
	domain, err := c.repositoryContainer.GetDomainRepository().GetDomainByID(domainID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, DomainFailedToGetDomainErrorMessage, err)
		return
	}

	if r.Method == http.MethodGet {
		content := response.NewUpdateDomainResponse(currentUser, domain)
		err = c.renderer.Template.RenderTemplate(w, "form-page.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the domain
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, DomainUpdateRequiredFieldMissing, nil)
			return
		}

		// update the domain
		domain.Name = name
		err = c.repositoryContainer.GetDomainRepository().UpdateDomain(domain)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, DomainUpdateUpdateDomainErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/domain/list", http.StatusSeeOther)
		return
	}
}

// DomainDeleteViewController is the controller for the domain delete form.
// It is responsible for deleting a domain.
// It redirects to the domain list page.
func (c *Controller) DomainDeleteViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("domains.delete") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	domainIDVariable := vars["domainId"]
	// it has to be converted to int64
	domainID, err := strconv.ParseInt(domainIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, DomainDomainIDInvalidErrorMessage, err)
		return
	}
	// delete the domain
	err = c.repositoryContainer.GetDomainRepository().DeleteDomain(domainID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, DomainDeleteFailedToDeleteErrorMessage, err)
		return
	}
	// redirect to the domain list
	http.Redirect(w, r, "/admin/domain/list", http.StatusSeeOther)
}

// DomainListViewController is the controller for the domain list view.
func (c *Controller) DomainListViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("domains.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	filter := model.NewDomainFilter()
	if r.Method == http.MethodPost {
		filter.Name = r.FormValue("name")
	}
	// get all domains
	domains, err := c.repositoryContainer.GetDomainRepository().GetDomains(filter)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, DomainListFailedToGetDomainsErrorMessage, err)
		return
	}
	content := response.NewDomainListResponse(currentUser, domains, filter)
	err = c.renderer.Template.RenderTemplate(w, "listing-page.html", content)
	if err != nil {
		panic(err)
	}
}
