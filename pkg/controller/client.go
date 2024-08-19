package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller/response"
	"github.com/akosgarai/projectregister/pkg/model"
)

// ClientViewController is the controller for the client view page.
// GET /admin/client/view/{clientId}
// It renders the client view page.
func (c *Controller) ClientViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("clients.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	client, statusCode, err := c.clientViewData(r)
	if err != nil {
		c.renderer.Error(w, statusCode, ClientFailedToGetClientErrorMessage, err)
		return
	}
	content := response.NewClientDetailResponse(currentUser, client)
	err = c.renderer.Template.RenderTemplate(w, "detail-page.html", content)
	if err != nil {
		panic(err)
	}
}

// clientViewData gets the request as input, and returns the client data, status code and error.
func (c *Controller) clientViewData(r *http.Request) (*model.Client, int, error) {
	vars := mux.Vars(r)
	clientIDVariable := vars["clientId"]
	// it has to be converted to int64
	clientID, err := strconv.ParseInt(clientIDVariable, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	client, err := c.repositoryContainer.GetClientRepository().GetClientByID(clientID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return client, http.StatusOK, nil
}

// ClientCreateViewController is the controller for the client create view.
// On case of get request, it returns the client create page.
// On case of post request, it creates the client and redirects to the list page.
func (c *Controller) ClientCreateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("clients.create") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	if r.Method == http.MethodGet {
		content := response.NewCreateClientResponse(currentUser)
		err := c.renderer.Template.RenderTemplate(w, "form-page.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the client
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, ClientCreateRequiredFieldMissing, nil)
			return
		}

		_, err := c.repositoryContainer.GetClientRepository().CreateClient(name)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, ClientCreateCreateClientErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/client/list", http.StatusSeeOther)
		return
	}
}

// ClientUpdateViewController is the controller for the client update view.
// On case of get request, it returns the client update page.
// On case of post request, it updates the client and redirects to the list page.
func (c *Controller) ClientUpdateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("clients.update") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	clientIDVariable := vars["clientId"]
	// it has to be converted to int64
	clientID, err := strconv.ParseInt(clientIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, ClientClientIDInvalidErrorMessage, err)
		return
	}

	// get the client
	client, err := c.repositoryContainer.GetClientRepository().GetClientByID(clientID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ClientFailedToGetClientErrorMessage, err)
		return
	}

	if r.Method == http.MethodGet {
		content := response.NewUpdateClientResponse(currentUser, client)
		err = c.renderer.Template.RenderTemplate(w, "form-page.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the client
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, ClientUpdateRequiredFieldMissing, nil)
			return
		}

		// update the client
		client.Name = name
		err = c.repositoryContainer.GetClientRepository().UpdateClient(client)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, ClientUpdateUpdateClientErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/client/list", http.StatusSeeOther)
		return
	}
}

// ClientDeleteViewController is the controller for the client delete form.
// It is responsible for deleting a client.
// It redirects to the client list page.
func (c *Controller) ClientDeleteViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("clients.delete") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	clientIDVariable := vars["clientId"]
	// it has to be converted to int64
	clientID, err := strconv.ParseInt(clientIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, ClientClientIDInvalidErrorMessage, err)
		return
	}
	// delete the client
	err = c.repositoryContainer.GetClientRepository().DeleteClient(clientID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ClientDeleteFailedToDeleteErrorMessage, err)
		return
	}
	// redirect to the client list
	http.Redirect(w, r, "/admin/client/list", http.StatusSeeOther)
}

// ClientListViewController is the controller for the client list view.
func (c *Controller) ClientListViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("clients.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	// get all clients
	clients, err := c.repositoryContainer.GetClientRepository().GetClients()
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ClientListFailedToGetClientsErrorMessage, err)
		return
	}
	content := response.NewClientListResponse(currentUser, clients)
	err = c.renderer.Template.RenderTemplate(w, "listing-page.html", content)
	if err != nil {
		panic(err)
	}
}
