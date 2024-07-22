package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/model"
)

// ApplicationFormResponse is the struct for the application form responses.
type ApplicationFormResponse struct {
	Title        string
	Application  *model.Application
	Clients      []*model.Client
	Projects     []*model.Project
	Environments []*model.Environment
	Databases    []*model.Database
	Runtimes     []*model.Runtime
	Pools        []*model.Pool
	Domains      []*model.Domain
	CurrentUser  *model.User
}

// ApplicationViewController is the controller for the application view page.
// GET /admin/application/view/{applicationId}
// It renders the application view page.
func (c *Controller) ApplicationViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("applications.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	template := c.renderer.BuildTemplate("application-view", []string{c.renderer.GetTemplateDirectoryPath() + "/application/view.html.tmpl"})
	application, statusCode, err := c.applicationViewData(r)
	if err != nil {
		c.renderer.Error(w, statusCode, ApplicationFailedToGetApplicationErrorMessage, err)
		return
	}
	content := struct {
		Title       string
		Application *model.Application
		CurrentUser *model.User
	}{
		Title:       "Application View",
		Application: application,
		CurrentUser: currentUser,
	}
	err = template.ExecuteTemplate(w, "base.html", content)
	if err != nil {
		panic(err)
	}
}

// applicationViewData gets the request as input, and returns the application data, status code and error.
func (c *Controller) applicationViewData(r *http.Request) (*model.Application, int, error) {
	vars := mux.Vars(r)
	applicationIDVariable := vars["applicationId"]
	// it has to be converted to int64
	applicationID, err := strconv.ParseInt(applicationIDVariable, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	application, err := c.applicationRepository.GetApplicationByID(applicationID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return application, http.StatusOK, nil
}

// ApplicationCreateViewController is the controller for the application create view.
// On case of get request, it returns the application create page.
// On case of post request, it creates the application and redirects to the list page.
func (c *Controller) ApplicationCreateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("applications.create") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	if r.Method == http.MethodGet {
		template := c.renderer.BuildTemplate("application-create", []string{c.renderer.GetTemplateDirectoryPath() + "/application/create.html.tmpl"})
		// TODO: add the free domains also.
		content, errorMessage, err := c.createApplicationFormResponse("Application Create", currentUser, &model.Application{})
		if errorMessage != "" {
			c.renderer.Error(w, http.StatusInternalServerError, errorMessage, err)
			return
		}
		err = template.ExecuteTemplate(w, "base.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		app, domainIDS, errorMessage, err := c.validateApplicationForm(r)
		if errorMessage != "" {
			c.renderer.Error(w, http.StatusBadRequest, errorMessage, err)
			return
		}
		_, err = c.applicationRepository.CreateApplication(app.Client.ID, app.Project.ID, app.Environment.ID, app.Database.ID, app.Runtime.ID, app.Pool.ID, app.Repository, app.Branch, app.DBName, app.DBUser, app.Framework, app.DocumentRoot, domainIDS)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, ApplicationCreateCreateApplicationErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/application/list", http.StatusSeeOther)
		return
	}
}

// ApplicationUpdateViewController is the controller for the application update view.
// On case of get request, it returns the application update page.
// On case of post request, it updates the application and redirects to the list page.
func (c *Controller) ApplicationUpdateViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("applications.update") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	applicationIDVariable := vars["applicationId"]
	// it has to be converted to int64
	applicationID, err := strconv.ParseInt(applicationIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, ApplicationApplicationIDInvalidErrorMessage, err)
		return
	}

	// get the application
	application, err := c.applicationRepository.GetApplicationByID(applicationID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ApplicationFailedToGetApplicationErrorMessage, err)
		return
	}

	if r.Method == http.MethodGet {
		template := c.renderer.BuildTemplate("user-application", []string{c.renderer.GetTemplateDirectoryPath() + "/application/update.html.tmpl"})
		content, errorMessage, err := c.createApplicationFormResponse("Application Update", currentUser, application)
		if errorMessage != "" {
			c.renderer.Error(w, http.StatusInternalServerError, errorMessage, err)
			return
		}
		err = template.ExecuteTemplate(w, "base.html", content)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == http.MethodPost {
		app, _, errorMessage, err := c.validateApplicationForm(r)
		if errorMessage != "" {
			c.renderer.Error(w, http.StatusBadRequest, errorMessage, err)
			return
		}
		app.ID = applicationID

		err = c.applicationRepository.UpdateApplication(app)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, ApplicationUpdateUpdateApplicationErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/application/list", http.StatusSeeOther)
		return
	}
}

// ApplicationDeleteViewController is the controller for the application delete form.
// It is responsible for deleting a application.
// It redirects to the application list page.
func (c *Controller) ApplicationDeleteViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("applications.delete") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	vars := mux.Vars(r)
	applicationIDVariable := vars["applicationId"]
	// it has to be converted to int64
	applicationID, err := strconv.ParseInt(applicationIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, ApplicationApplicationIDInvalidErrorMessage, err)
		return
	}
	// delete the application
	err = c.applicationRepository.DeleteApplication(applicationID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ApplicationDeleteFailedToDeleteErrorMessage, err)
		return
	}
	// redirect to the application list
	http.Redirect(w, r, "/admin/application/list", http.StatusSeeOther)
}

// ApplicationListViewController is the controller for the application list view.
func (c *Controller) ApplicationListViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("applications.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	// get all applications
	applications, err := c.applicationRepository.GetApplications()
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ApplicationListFailedToGetApplicationsErrorMessage, err)
		return
	}
	template := c.renderer.BuildTemplate("application-list", []string{c.renderer.GetTemplateDirectoryPath() + "/application/list.html.tmpl"})
	content := struct {
		Title        string
		Applications []*model.Application
		CurrentUser  *model.User
	}{
		Title:        "Application List",
		Applications: applications,
		CurrentUser:  currentUser,
	}
	err = template.ExecuteTemplate(w, "base.html", content)
	if err != nil {
		panic(err)
	}
}

// ApplicationImportToEnvironmentFormController is the controller for the application import to environment form.
// It is responsible for handling the forms that guides you throught the import process.
func (c *Controller) ApplicationImportToEnvironmentFormController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("applications.create") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	// get the environment id from the url
	vars := mux.Vars(r)
	environmentIDVariable := vars["environmentId"]
	// it has to be converted to int64
	environmentID, err := strconv.ParseInt(environmentIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, ApplicationImportInvalidEnvironmentIDErrorMessage, err)
		return
	}
	// load the environment
	environment, err := c.environmentRepository.GetEnvironmentByID(environmentID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ApplicationImportFailedToGetEnvironmentErrorMessage, err)
		return
	}
	// On case of get method load the form template
	if r.Method == http.MethodGet {
		content := struct {
			Title       string
			Environment *model.Environment
			CurrentUser *model.User
		}{
			Title:       "Import Application to Environment",
			Environment: environment,
			CurrentUser: currentUser,
		}
		template := c.renderer.BuildTemplate("application-import", []string{c.renderer.GetTemplateDirectoryPath() + "/application/import.html.tmpl"})
		err = template.ExecuteTemplate(w, "base.html", content)
		if err != nil {
			panic(err)
		}
	}
	// On case of post method, store the csv file content in file and redirect to the mapping page.
	if r.Method == http.MethodPost {
		filename, err := c.storeImportCSVFile(r)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, ApplicationImportFailedToSaveFileErrorMessage, err)
			return
		}
		http.Redirect(w, r, "/admin/application/mapping-to-environment/"+environmentIDVariable+"/"+filename, http.StatusSeeOther)
	}
}

// ApplicationMappingToEnvironmentFormController is the controller for the application mapping to environment form.
// It is responsible for handling the forms that guides you throught the mapping process.
func (c *Controller) ApplicationMappingToEnvironmentFormController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("applications.create") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	// get the environment id from the url
	vars := mux.Vars(r)
	environmentIDVariable := vars["environmentId"]
	// it has to be converted to int64
	environmentID, err := strconv.ParseInt(environmentIDVariable, 10, 64)
	if err != nil {
		c.renderer.Error(w, http.StatusBadRequest, ApplicationImportInvalidEnvironmentIDErrorMessage, err)
		return
	}
	fileID := vars["fileId"]
	// load the environment
	environment, err := c.environmentRepository.GetEnvironmentByID(environmentID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ApplicationImportFailedToGetEnvironmentErrorMessage, err)
		return
	}
	// On case of get method load the form template
	if r.Method == http.MethodGet {
		content := struct {
			Title       string
			Environment *model.Environment
			FileID      string
			CurrentUser *model.User
		}{
			Title:       "Import Mapping to Environment",
			Environment: environment,
			CurrentUser: currentUser,
			FileID:      fileID,
		}
		template := c.renderer.BuildTemplate("application-import", []string{c.renderer.GetTemplateDirectoryPath() + "/application/mapping.html.tmpl"})
		err = template.ExecuteTemplate(w, "base.html", content)
		if err != nil {
			panic(err)
		}
	}
}

// It creates the content for the application forms.
func (c *Controller) createApplicationFormResponse(title string, currentUser *model.User, application *model.Application) (ApplicationFormResponse, string, error) {
	runtimes, err := c.runtimeRepository.GetRuntimes()
	if err != nil {
		return ApplicationFormResponse{}, ApplicationCreateFailedToGetRuntimesErrorMessage, err
	}
	pools, err := c.poolRepository.GetPools()
	if err != nil {
		return ApplicationFormResponse{}, ApplicationCreateFailedToGetPoolsErrorMessage, err
	}
	clients, err := c.clientRepository.GetClients()
	if err != nil {
		return ApplicationFormResponse{}, ApplicationCreateFailedToGetClientsErrorMessage, err
	}
	projects, err := c.projectRepository.GetProjects()
	if err != nil {
		return ApplicationFormResponse{}, ApplicationCreateFailedToGetProjectsErrorMessage, err
	}
	environments, err := c.environmentRepository.GetEnvironments()
	if err != nil {
		return ApplicationFormResponse{}, ApplicationCreateFailedToGetEnvironmentsErrorMessage, err
	}
	databases, err := c.databaseRepository.GetDatabases()
	if err != nil {
		return ApplicationFormResponse{}, ApplicationCreateFailedToGetDatabasesErrorMessage, err
	}
	domains, err := c.domainRepository.GetFreeDomains()
	if err != nil {
		return ApplicationFormResponse{}, ApplicationCreateFailedToGetDomainsErrorMessage, err
	}

	return ApplicationFormResponse{
		Title:        title,
		Clients:      clients,
		Projects:     projects,
		Environments: environments,
		Databases:    databases,
		Runtimes:     runtimes,
		Pools:        pools,
		Domains:      domains,
		CurrentUser:  currentUser,
		Application:  application,
	}, "", nil
}

// It validates the application form data. Returns the Application with the validated data, the domain ids, and an error message and error.
func (c *Controller) validateApplicationForm(r *http.Request) (*model.Application, []int64, string, error) {
	clientIDRaw := r.FormValue("client")
	projectIDRaw := r.FormValue("project")
	envIDRaw := r.FormValue("environment")
	dbIDRaw := r.FormValue("database")
	runtimeIDRaw := r.FormValue("runtime")
	poolIDRaw := r.FormValue("pool")
	dbName := r.FormValue("db_name")
	dbUser := r.FormValue("db_user")
	repository := r.FormValue("repository")
	branch := r.FormValue("branch")
	framework := r.FormValue("framework")
	documentRoot := r.FormValue("document_root")
	domainIDsRaw := r.Form["domain"]
	var domainIDS []int64

	// if the clientID, projectID, envID, dbID, runtimeID, poolID is empty, return an error
	if clientIDRaw == "" || projectIDRaw == "" || envIDRaw == "" {
		return nil, domainIDS, ApplicationCreateRequiredFieldMissing, nil
	}
	// convert the clientID, projectID, envID, dbID, runtimeID, poolID to int64
	clientID, err := strconv.ParseInt(clientIDRaw, 10, 64)
	if err != nil {
		return nil, domainIDS, ApplicationCreateClientIDInvalidErrorMessage, err
	}
	projectID, err := strconv.ParseInt(projectIDRaw, 10, 64)
	if err != nil {
		return nil, domainIDS, ApplicationCreateProjectIDInvalidErrorMessage, err
	}
	envID, err := strconv.ParseInt(envIDRaw, 10, 64)
	if err != nil {
		return nil, domainIDS, ApplicationCreateEnvironmentIDInvalidErrorMessage, err
	}
	dbID, err := strconv.ParseInt(dbIDRaw, 10, 64)
	if err != nil {
		return nil, domainIDS, ApplicationCreateDatabaseIDInvalidErrorMessage, err
	}
	runtimeID, err := strconv.ParseInt(runtimeIDRaw, 10, 64)
	if err != nil {
		return nil, domainIDS, ApplicationCreateRuntimeIDInvalidErrorMessage, err
	}
	poolID, err := strconv.ParseInt(poolIDRaw, 10, 64)
	if err != nil {
		return nil, domainIDS, ApplicationCreatePoolIDInvalidErrorMessage, err
	}
	app := &model.Application{
		Client:       &model.Client{ID: clientID},
		Project:      &model.Project{ID: projectID},
		Environment:  &model.Environment{ID: envID},
		Database:     &model.Database{ID: dbID},
		Runtime:      &model.Runtime{ID: runtimeID},
		Pool:         &model.Pool{ID: poolID},
		Repository:   repository,
		Branch:       branch,
		DBName:       dbName,
		DBUser:       dbUser,
		Framework:    framework,
		DocumentRoot: documentRoot,
	}
	for _, domainIDRaw := range domainIDsRaw {
		domainID, err := strconv.ParseInt(domainIDRaw, 10, 64)
		if err != nil {
			return nil, domainIDS, ApplicationCreateDomainIDInvalidErrorMessage, err
		}
		domainIDS = append(domainIDS, domainID)
		app.Domains = append(app.Domains, &model.Domain{ID: domainID})
	}
	// add the domain ids
	return app, domainIDS, "", nil

}

// storeImportCSVFile stores the csv file in the file system.
// It returns the filename and an error.
func (c *Controller) storeImportCSVFile(r *http.Request) (string, error) {
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("csvfile")
	if err != nil {
		return "", err
	}
	defer file.Close()
	// store the file
	return c.csvStorage.Save(file)
}
