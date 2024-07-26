package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/model"
)

// ServerViewController is the controller for the server view page.
// GET /admin/server/view/{serverId}
// It renders the server view page.
func (c *Controller) ServerViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("servers.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	server, statusCode, err := c.serverViewData(r)
	if err != nil {
		c.renderer.Error(w, statusCode, ServerFailedToGetServerErrorMessage, err)
		return
	}
	content := struct {
		Title       string
		Server      *model.Server
		CurrentUser *model.User
	}{
		Title:       "Server View",
		Server:      server,
		CurrentUser: currentUser,
	}
	err = c.renderer.Template.RenderTemplate(w, "server-view.html", content)
	if err != nil {
		panic(err)
	}
}

// serverViewData gets the request as input, and returns the server data, status code and error.
func (c *Controller) serverViewData(r *http.Request) (*model.Server, int, error) {
	vars := mux.Vars(r)
	serverIDVariable := vars["serverId"]
	// it has to be converted to int64
	serverID, err := strconv.ParseInt(serverIDVariable, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	server, err := c.serverRepository.GetServerByID(serverID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return server, http.StatusOK, nil
}

// ServerCreateViewController is the controller for the server create view.
// On case of get request, it returns the server create page.
// On case of post request, it creates the server and redirects to the list page.
func (c *Controller) ServerCreateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("servers.create") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	if r.Method == http.MethodGet {
		runtimes, err := c.runtimeRepository.GetRuntimes()
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, ServerCreateFailedToGetRuntimesErrorMessage, err)
			return
		}
		pools, err := c.poolRepository.GetPools()
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, ServerCreateFailedToGetPoolsErrorMessage, err)
			return
		}
		content := struct {
			Title       string
			Runtimes    []*model.Runtime
			Pools       []*model.Pool
			CurrentUser *model.User
		}{
			Title:       "Server Create",
			Runtimes:    runtimes,
			Pools:       pools,
			CurrentUser: currentUser,
		}
		err = c.renderer.Template.RenderTemplate(w, "server-create.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		description := r.FormValue("description")
		remoteAddress := r.FormValue("remote_address")

		// if the name or remote address is empty, return an error
		if name == "" || remoteAddress == "" {
			c.renderer.Error(w, http.StatusBadRequest, ServerCreateRequiredFieldMissing, nil)
			return
		}
		runtimeIDsRaw := r.Form["runtimes"]
		var runtimeIDs []int64
		for _, v := range runtimeIDsRaw {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				c.renderer.Error(w, http.StatusBadRequest, ServerCreateRuntimeIDInvalidErrorMessage, err)
				return
			}
			runtimeIDs = append(runtimeIDs, id)
		}
		poolIDsRaw := r.Form["pools"]
		var poolIDs []int64
		for _, v := range poolIDsRaw {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				c.renderer.Error(w, http.StatusBadRequest, ServerCreatePoolIDInvalidErrorMessage, err)
				return
			}
			poolIDs = append(poolIDs, id)
		}

		_, err := c.serverRepository.CreateServer(name, description, remoteAddress, runtimeIDs, poolIDs)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, ServerCreateCreateServerErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/server/list", http.StatusSeeOther)
		return
	}
}

// ServerUpdateViewController is the controller for the server update view.
// On case of get request, it returns the server update page.
// On case of post request, it updates the server and redirects to the list page.
func (c *Controller) ServerUpdateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("servers.update") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	serverIDVariable := vars["serverId"]
	// it has to be converted to int64
	serverID, err := strconv.ParseInt(serverIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, ServerServerIDInvalidErrorMessage, err)
		return
	}

	// get the server
	server, err := c.serverRepository.GetServerByID(serverID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ServerFailedToGetServerErrorMessage, err)
		return
	}

	if r.Method == http.MethodGet {
		runtimes, err := c.runtimeRepository.GetRuntimes()
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, ServerCreateFailedToGetRuntimesErrorMessage, err)
			return
		}
		pools, err := c.poolRepository.GetPools()
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, ServerCreateFailedToGetPoolsErrorMessage, err)
			return
		}
		content := struct {
			Title       string
			Server      *model.Server
			Runtimes    []*model.Runtime
			Pools       []*model.Pool
			CurrentUser *model.User
		}{
			Title:       "Server Update",
			Server:      server,
			Runtimes:    runtimes,
			Pools:       pools,
			CurrentUser: currentUser,
		}
		err = c.renderer.Template.RenderTemplate(w, "server-update.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the server
		name := r.FormValue("name")
		description := r.FormValue("description")
		remoteAddress := r.FormValue("remote_address")
		runtimeIDsRaw := r.Form["runtimes"]
		poolIDsRaw := r.Form["pools"]

		// if the name or remote address is empty, return an error
		if name == "" || remoteAddress == "" {
			c.renderer.Error(w, http.StatusBadRequest, ServerUpdateRequiredFieldMissing, nil)
			return
		}

		// update the server
		server.Name = name
		server.Description = description
		server.RemoteAddr = remoteAddress
		// update the server runtime relations
		server.Runtimes = make([]*model.Runtime, len(runtimeIDsRaw))
		for i, v := range runtimeIDsRaw {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				c.renderer.Error(w, http.StatusBadRequest, ServerUpdateRuntimeIDInvalidErrorMessage, err)
				return
			}
			server.Runtimes[i] = &model.Runtime{ID: id}
		}
		// update the server pool relations
		server.Pools = make([]*model.Pool, len(poolIDsRaw))
		for i, v := range poolIDsRaw {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				c.renderer.Error(w, http.StatusBadRequest, ServerUpdatePoolIDInvalidErrorMessage, err)
				return
			}
			server.Pools[i] = &model.Pool{ID: id}
		}
		err = c.serverRepository.UpdateServer(server)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, ServerUpdateUpdateServerErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/server/list", http.StatusSeeOther)
		return
	}
}

// ServerDeleteViewController is the controller for the server delete form.
// It is responsible for deleting a server.
// It redirects to the server list page.
func (c *Controller) ServerDeleteViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("servers.delete") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	serverIDVariable := vars["serverId"]
	// it has to be converted to int64
	serverID, err := strconv.ParseInt(serverIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, ServerServerIDInvalidErrorMessage, err)
		return
	}
	// delete the server
	err = c.serverRepository.DeleteServer(serverID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ServerDeleteFailedToDeleteErrorMessage, err)
		return
	}
	// redirect to the server list
	http.Redirect(w, r, "/admin/server/list", http.StatusSeeOther)
}

// ServerListViewController is the controller for the server list view.
func (c *Controller) ServerListViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("servers.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	// get all servers
	servers, err := c.serverRepository.GetServers()
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ServerListFailedToGetServersErrorMessage, err)
		return
	}
	content := struct {
		Title       string
		Servers     []*model.Server
		CurrentUser *model.User
	}{
		Title:       "Server List",
		Servers:     servers,
		CurrentUser: currentUser,
	}
	err = c.renderer.Template.RenderTemplate(w, "server-list.html", content)
	if err != nil {
		panic(err)
	}
}
