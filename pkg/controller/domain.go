package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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
	content := struct {
		Title       string
		Domain      *model.Domain
		CurrentUser *model.User
	}{
		Title:       "Domain View",
		Domain:      domain,
		CurrentUser: currentUser,
	}
	err = c.renderer.Template.RenderTemplate(w, "domain-view.html", content)
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
	domain, err := c.domainRepository.GetDomainByID(domainID)
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
		content := struct {
			Title       string
			CurrentUser *model.User
		}{
			Title:       "Domain Create",
			CurrentUser: currentUser,
		}
		err := c.renderer.Template.RenderTemplate(w, "domain-create.html", content)
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

		_, err := c.domainRepository.CreateDomain(name)
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
	domain, err := c.domainRepository.GetDomainByID(domainID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, DomainFailedToGetDomainErrorMessage, err)
		return
	}

	if r.Method == http.MethodGet {
		content := struct {
			Title       string
			Domain      *model.Domain
			CurrentUser *model.User
		}{
			Title:       "Domain Update",
			Domain:      domain,
			CurrentUser: currentUser,
		}
		err = c.renderer.Template.RenderTemplate(w, "domain-update.html", content)
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
		err = c.domainRepository.UpdateDomain(domain)
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
	err = c.domainRepository.DeleteDomain(domainID)
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
	// get all domains
	domains, err := c.domainRepository.GetDomains()
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, DomainListFailedToGetDomainsErrorMessage, err)
		return
	}
	content := struct {
		Title       string
		Domains     []*model.Domain
		CurrentUser *model.User
	}{
		Title:       "Domain List",
		Domains:     domains,
		CurrentUser: currentUser,
	}
	err = c.renderer.Template.RenderTemplate(w, "domain-list.html", content)
	if err != nil {
		panic(err)
	}
}
