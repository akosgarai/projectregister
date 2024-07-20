package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/model"
)

// DatabaseViewController is the controller for the database view page.
// GET /admin/database/view/{databaseId}
// It renders the database view page.
func (c *Controller) DatabaseViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("databases.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	template := c.renderer.BuildTemplate("database-view", []string{c.renderer.GetTemplateDirectoryPath() + "/database/view.html.tmpl"})
	database, statusCode, err := c.databaseViewData(r)
	if err != nil {
		c.renderer.Error(w, statusCode, DatabaseFailedToGetDatabaseErrorMessage, err)
		return
	}
	content := struct {
		Title       string
		Database    *model.Database
		CurrentUser *model.User
	}{
		Title:       "Database View",
		Database:    database,
		CurrentUser: currentUser,
	}
	err = template.ExecuteTemplate(w, "base.html", content)
	if err != nil {
		panic(err)
	}
}

// databaseViewData gets the request as input, and returns the database data, status code and error.
func (c *Controller) databaseViewData(r *http.Request) (*model.Database, int, error) {
	vars := mux.Vars(r)
	databaseIDVariable := vars["databaseId"]
	// it has to be converted to int64
	databaseID, err := strconv.ParseInt(databaseIDVariable, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	database, err := c.databaseRepository.GetDatabaseByID(databaseID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return database, http.StatusOK, nil
}

// DatabaseCreateViewController is the controller for the database create view.
// On case of get request, it returns the database create page.
// On case of post request, it creates the database and redirects to the list page.
func (c *Controller) DatabaseCreateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("databases.create") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	if r.Method == http.MethodGet {
		template := c.renderer.BuildTemplate("database-create", []string{c.renderer.GetTemplateDirectoryPath() + "/database/create.html.tmpl"})
		content := struct {
			Title       string
			CurrentUser *model.User
		}{
			Title:       "Database Create",
			CurrentUser: currentUser,
		}
		err := template.ExecuteTemplate(w, "base.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the database
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, DatabaseCreateRequiredFieldMissing, nil)
			return
		}

		_, err := c.databaseRepository.CreateDatabase(name)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, DatabaseCreateCreateDatabaseErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/database/list", http.StatusSeeOther)
		return
	}
}

// DatabaseUpdateViewController is the controller for the database update view.
// On case of get request, it returns the database update page.
// On case of post request, it updates the database and redirects to the list page.
func (c *Controller) DatabaseUpdateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("databases.update") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	databaseIDVariable := vars["databaseId"]
	// it has to be converted to int64
	databaseID, err := strconv.ParseInt(databaseIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, DatabaseDatabaseIDInvalidErrorMessage, err)
		return
	}

	// get the database
	database, err := c.databaseRepository.GetDatabaseByID(databaseID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, DatabaseFailedToGetDatabaseErrorMessage, err)
		return
	}

	if r.Method == http.MethodGet {
		template := c.renderer.BuildTemplate("user-database", []string{c.renderer.GetTemplateDirectoryPath() + "/database/update.html.tmpl"})
		content := struct {
			Title       string
			Database    *model.Database
			CurrentUser *model.User
		}{
			Title:       "Database Update",
			Database:    database,
			CurrentUser: currentUser,
		}
		err = template.ExecuteTemplate(w, "base.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		// update the database
		name := r.FormValue("name")

		// if the name is empty, return an error
		if name == "" {
			c.renderer.Error(w, http.StatusBadRequest, DatabaseUpdateRequiredFieldMissing, nil)
			return
		}

		// update the database
		database.Name = name
		err = c.databaseRepository.UpdateDatabase(database)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, DatabaseUpdateUpdateDatabaseErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/database/list", http.StatusSeeOther)
		return
	}
}

// DatabaseDeleteViewController is the controller for the database delete form.
// It is responsible for deleting a database.
// It redirects to the database list page.
func (c *Controller) DatabaseDeleteViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("databases.delete") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	databaseIDVariable := vars["databaseId"]
	// it has to be converted to int64
	databaseID, err := strconv.ParseInt(databaseIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, DatabaseDatabaseIDInvalidErrorMessage, err)
		return
	}
	// delete the database
	err = c.databaseRepository.DeleteDatabase(databaseID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, DatabaseDeleteFailedToDeleteErrorMessage, err)
		return
	}
	// redirect to the database list
	http.Redirect(w, r, "/admin/database/list", http.StatusSeeOther)
}

// DatabaseListViewController is the controller for the database list view.
func (c *Controller) DatabaseListViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("databases.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	// get all databases
	databases, err := c.databaseRepository.GetDatabases()
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, DatabaseListFailedToGetDatabasesErrorMessage, err)
		return
	}
	template := c.renderer.BuildTemplate("database-list", []string{c.renderer.GetTemplateDirectoryPath() + "/database/list.html.tmpl"})
	content := struct {
		Title       string
		Databases   []*model.Database
		CurrentUser *model.User
	}{
		Title:       "Database List",
		Databases:   databases,
		CurrentUser: currentUser,
	}
	err = template.ExecuteTemplate(w, "base.html", content)
	if err != nil {
		panic(err)
	}
}
