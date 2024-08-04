package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller/response"
	"github.com/akosgarai/projectregister/pkg/model"
)

// ApplicationViewController is the controller for the application view page.
// GET /admin/application/view/{applicationId}
// It renders the application view page.
func (c *Controller) ApplicationViewController(w http.ResponseWriter, r *http.Request) {
	currentUser := c.CurrentUser(r)
	if !currentUser.HasPrivilege("applications.view") {
		c.renderer.Error(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	application, statusCode, err := c.applicationViewData(r)
	if err != nil {
		c.renderer.Error(w, statusCode, ApplicationFailedToGetApplicationErrorMessage, err)
		return
	}
	content := response.NewApplicationDetailResponse(currentUser, application)
	err = c.renderer.Template.RenderTemplate(w, "detail-page.html", content)
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
		content, errorMessage, err := c.createApplicationFormResponse(currentUser, nil)
		if errorMessage != "" {
			c.renderer.Error(w, http.StatusInternalServerError, errorMessage, err)
			return
		}
		err = c.renderer.Template.RenderTemplate(w, "form-page.html", content)
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
		content, errorMessage, err := c.createApplicationFormResponse(currentUser, application)
		if errorMessage != "" {
			c.renderer.Error(w, http.StatusInternalServerError, errorMessage, err)
			return
		}
		err = c.renderer.Template.RenderTemplate(w, "form-page.html", content)
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
	content := response.NewApplicationListResponse(currentUser, applications)
	err = c.renderer.Template.RenderTemplate(w, "listing-page.html", content)
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
		content := response.NewApplicationImportToEnvironmentFormResponse(currentUser, environment)
		err = c.renderer.Template.RenderTemplate(w, "application-import.html", content)
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
		content := response.NewApplicationMappingToEnvironmentFormResponse(currentUser, environment, fileID)
		err = c.renderer.Template.RenderTemplate(w, "application-mapping.html", content)
		if err != nil {
			panic(err)
		}
	}
	// On case of post process the mapping form and execute the import process.
	if r.Method == http.MethodPost {
		// TODO: implement the import process
		// Now use hardcoded mapping based on the excel file.
		fileName := r.FormValue("file_id")
		environmentIDRaw := r.FormValue("environment_id")
		// it has to be converted to int64
		environmentID, err := strconv.ParseInt(environmentIDRaw, 10, 64)
		if err != nil {
			c.renderer.Error(w, http.StatusBadRequest, ApplicationImportInvalidEnvironmentIDErrorMessage, err)
			return
		}
		results, err := c.importApplicationToEnvironment(environmentID, fileName)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, "Failed to import", err)
			return
		}
		content := response.NewApplicationImportToEnvironmentListResponse(currentUser, environment, fileName, results)
		err = c.renderer.Template.RenderTemplate(w, "application-import-results.html", content)
		if err != nil {
			panic(err)
		}
	}
}

// It creates the content for the application forms.
func (c *Controller) createApplicationFormResponse(currentUser *model.User, application *model.Application) (*response.FormResponse, string, error) {
	runtimes, err := c.runtimeRepository.GetRuntimes()
	if err != nil {
		return &response.FormResponse{}, ApplicationCreateFailedToGetRuntimesErrorMessage, err
	}
	pools, err := c.poolRepository.GetPools()
	if err != nil {
		return &response.FormResponse{}, ApplicationCreateFailedToGetPoolsErrorMessage, err
	}
	clients, err := c.clientRepository.GetClients()
	if err != nil {
		return &response.FormResponse{}, ApplicationCreateFailedToGetClientsErrorMessage, err
	}
	projects, err := c.projectRepository.GetProjects()
	if err != nil {
		return &response.FormResponse{}, ApplicationCreateFailedToGetProjectsErrorMessage, err
	}
	environments, err := c.environmentRepository.GetEnvironments()
	if err != nil {
		return &response.FormResponse{}, ApplicationCreateFailedToGetEnvironmentsErrorMessage, err
	}
	databases, err := c.databaseRepository.GetDatabases()
	if err != nil {
		return &response.FormResponse{}, ApplicationCreateFailedToGetDatabasesErrorMessage, err
	}
	domains, err := c.domainRepository.GetFreeDomains()
	if err != nil {
		return &response.FormResponse{}, ApplicationCreateFailedToGetDomainsErrorMessage, err
	}

	if application == nil {
		return response.NewCreateApplicationResponse(currentUser, clients, projects, environments, databases, runtimes, pools, domains), "", nil
	}

	return response.NewUpdateApplicationResponse(currentUser, application, clients, projects, environments, databases, runtimes, pools, domains), "", nil
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

// importApplicationToEnvironment imports the applications to the environment.
// It returns the results and an error.
func (c *Controller) importApplicationToEnvironment(environmentID int64, fileName string) (map[int]*response.ApplicationImportRowResult, error) {
	// get the file content
	results := map[int]*response.ApplicationImportRowResult{}
	csvData, err := c.csvStorage.Read(fileName + ".csv")
	if err != nil {
		return results, err
	}
	// parse the file content every line represents an application
	for rowIndex, line := range csvData {
		// set the response to empty string. it means process went well
		results[rowIndex] = &response.ApplicationImportRowResult{ErrorMessage: "", RowData: line, Application: nil}
		clientName := line[0]
		projectName := line[1]
		runtimeName := line[5]
		poolName := line[6]
		databaseTypeName := line[7]

		domainsRaw := line[3]
		framework := line[4]
		databaseName := line[8]
		if databaseName == "-" {
			databaseName = ""
		}
		databaseUser := line[9]
		if databaseUser == "-" {
			databaseUser = ""
		}
		docRoot := line[10]
		repository := line[11]
		branch := line[12]

		// if the client name does not exist, create it
		client, err := c.clientRepository.GetClientByName(clientName)
		if err != nil {
			client, err = c.clientRepository.CreateClient(clientName)
			if err != nil {
				currentData := results[rowIndex]
				currentData.ErrorMessage = err.Error()
				results[rowIndex] = currentData
				continue
			}
		}

		// if the project name does not exist, create it
		project, err := c.projectRepository.GetProjectByName(projectName)
		if err != nil {
			project, err = c.projectRepository.CreateProject(projectName)
			if err != nil {
				currentData := results[rowIndex]
				currentData.ErrorMessage = err.Error()
				results[rowIndex] = currentData
				continue
			}
		}
		// if the runtime name does not exist, create it
		runtime, err := c.runtimeRepository.GetRuntimeByName(runtimeName)
		if err != nil {
			runtime, err = c.runtimeRepository.CreateRuntime(runtimeName)
			if err != nil {
				currentData := results[rowIndex]
				currentData.ErrorMessage = err.Error()
				results[rowIndex] = currentData
				continue
			}
		}
		// if the pool name does not exist, create it
		pool, err := c.poolRepository.GetPoolByName(poolName)
		if err != nil {
			pool, err = c.poolRepository.CreatePool(poolName)
			if err != nil {
				currentData := results[rowIndex]
				currentData.ErrorMessage = err.Error()
				results[rowIndex] = currentData
				continue
			}
		}
		// if the database name does not exist, create it
		database, err := c.databaseRepository.GetDatabaseByName(databaseTypeName)
		if err != nil {
			database, err = c.databaseRepository.CreateDatabase(databaseTypeName)
			if err != nil {
				currentData := results[rowIndex]
				currentData.ErrorMessage = err.Error()
				results[rowIndex] = currentData
				continue
			}
		}

		// handle the domains. The domains are separated by space.
		// If the domain does not exists, create it.
		domainNames := strings.Split(domainsRaw, " ")
		domainIDs := []int64{}
		for _, domainName := range domainNames {
			domain, err := c.domainRepository.GetDomainByName(domainName)
			if err != nil {
				domain, err = c.domainRepository.CreateDomain(domainName)
				if err != nil {
					currentData := results[rowIndex]
					currentData.ErrorMessage = err.Error()
					results[rowIndex] = currentData
					continue
				}
			}
			domainIDs = append(domainIDs, domain.ID)
		}

		// create the application
		app, err := c.applicationRepository.CreateApplication(client.ID, project.ID, environmentID, database.ID, runtime.ID, pool.ID, repository, branch, databaseName, databaseUser, framework, docRoot, domainIDs)
		currentData := results[rowIndex]
		if err != nil {
			currentData.ErrorMessage = err.Error()
			results[rowIndex] = currentData
			continue
		}
		currentData.Application = app
		results[rowIndex] = currentData
	}
	return results, nil
}
