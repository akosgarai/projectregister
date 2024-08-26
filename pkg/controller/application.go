package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller/response"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/parser"
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
	application, err := c.repositoryContainer.GetApplicationRepository().GetApplicationByID(applicationID)
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
		_, err = c.repositoryContainer.GetApplicationRepository().CreateApplication(app.Client.ID, app.Project.ID, app.Environment.ID, app.Database.ID, app.Runtime.ID, app.Pool.ID, app.Repository, app.Branch, app.DBName, app.DBUser, app.Framework, app.DocumentRoot, domainIDS)
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
	application, err := c.repositoryContainer.GetApplicationRepository().GetApplicationByID(applicationID)
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

		err = c.repositoryContainer.GetApplicationRepository().UpdateApplication(app)
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
	err = c.repositoryContainer.GetApplicationRepository().DeleteApplication(applicationID)
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
	applications, err := c.repositoryContainer.GetApplicationRepository().GetApplications()
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
	environment, err := c.repositoryContainer.GetEnvironmentRepository().GetEnvironmentByID(environmentID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ApplicationImportFailedToGetEnvironmentErrorMessage, err)
		return
	}
	// On case of get method load the form template
	if r.Method == http.MethodGet {
		content := response.NewApplicationImportToEnvironmentFormResponse(currentUser, environment)
		err = c.renderer.Template.RenderTemplate(w, "form-page.html", content)
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
		hasHeader := r.FormValue("has_header")
		redirectURL := "/admin/application/mapping-to-environment/" + environmentIDVariable + "/" + filename
		if hasHeader == "1" {
			redirectURL += "?has_header=true"
		}
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
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
	environment, err := c.repositoryContainer.GetEnvironmentRepository().GetEnvironmentByID(environmentID)
	if err != nil {
		c.renderer.Error(w, http.StatusInternalServerError, ApplicationImportFailedToGetEnvironmentErrorMessage, err)
		return
	}
	// On case of get method load the form template
	if r.Method == http.MethodGet {
		// get the file content
		csvData, err := c.csvStorage.Read(fileID + ".csv")
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, "Failed to read the CSV.", err)
			return
		}
		hasHeader := r.URL.Query().Get("has_header")
		header := []string{}
		if hasHeader == "true" {
			// if the first line is the header, remove it
			header = csvData[0]
			csvData = csvData[1:]
		}
		content := response.NewApplicationMappingToEnvironmentFormResponse(currentUser, environment, fileID, csvData, header)
		err = c.renderer.Template.RenderTemplate(w, "application-import-mapping.html", content)
		if err != nil {
			panic(err)
		}
	}
	// On case of post process the mapping form and execute the import process.
	if r.Method == http.MethodPost {
		fileName := r.FormValue("file_id")
		environmentIDRaw := r.FormValue("environment_id")
		// it has to be converted to int64
		environmentID, err := strconv.ParseInt(environmentIDRaw, 10, 64)
		if err != nil {
			c.renderer.Error(w, http.StatusBadRequest, ApplicationImportInvalidEnvironmentIDErrorMessage, err)
			return
		}
		mappingRules, err := c.getImportApplicationToEnvironmentMapping(r)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, "Failed to get the mapping rules", err)
			return
		}
		results, err := c.importApplicationToEnvironment(environmentID, fileName, mappingRules)
		if err != nil {
			c.renderer.Error(w, http.StatusInternalServerError, "Failed to import", err)
			return
		}
		content := response.NewApplicationImportToEnvironmentListResponse(currentUser, environment, fileName, results)
		err = c.renderer.Template.RenderTemplate(w, "listing-page.html", content)
		if err != nil {
			panic(err)
		}
	}
}

// getImportApplicationToEnvironmentMapping returns the mapping rules for the import process.
func (c *Controller) getImportApplicationToEnvironmentMapping(r *http.Request) (parser.ApplicationImportMapping, error) {
	mappingRules := parser.NewApplicationImportMapping()
	parameterNames := []string{"client", "project", "domains", "runtime", "pool", "framework", "database", "database_name", "database_user", "doc_root", "repository", "branch"}
	for _, parameterName := range parameterNames {
		parameterIndexRaw := r.FormValue(parameterName)
		parameterCustomValue := r.FormValue(parameterName + "_custom")
		if parameterIndexRaw != "" {
			parameterInt, err := strconv.Atoi(parameterIndexRaw)
			if err != nil {
				return parser.ApplicationImportMapping{}, err
			}
			mappingRules[parameterName].ColumnIndex = parameterInt
		} else {
			mappingRules[parameterName].CustomValue = parameterCustomValue
		}
	}
	return mappingRules, nil
}

// It creates the content for the application forms.
func (c *Controller) createApplicationFormResponse(currentUser *model.User, application *model.Application) (*response.FormResponse, string, error) {
	runtimes, err := c.repositoryContainer.GetRuntimeRepository().GetRuntimes(model.NewRuntimeFilter())
	if err != nil {
		return &response.FormResponse{}, ApplicationCreateFailedToGetRuntimesErrorMessage, err
	}
	pools, err := c.repositoryContainer.GetPoolRepository().GetPools()
	if err != nil {
		return &response.FormResponse{}, ApplicationCreateFailedToGetPoolsErrorMessage, err
	}
	clients, err := c.repositoryContainer.GetClientRepository().GetClients(model.NewClientFilter())
	if err != nil {
		return &response.FormResponse{}, ApplicationCreateFailedToGetClientsErrorMessage, err
	}
	projects, err := c.repositoryContainer.GetProjectRepository().GetProjects(model.NewProjectFilter())
	if err != nil {
		return &response.FormResponse{}, ApplicationCreateFailedToGetProjectsErrorMessage, err
	}
	environments, err := c.repositoryContainer.GetEnvironmentRepository().GetEnvironments()
	if err != nil {
		return &response.FormResponse{}, ApplicationCreateFailedToGetEnvironmentsErrorMessage, err
	}
	databases, err := c.repositoryContainer.GetDatabaseRepository().GetDatabases()
	if err != nil {
		return &response.FormResponse{}, ApplicationCreateFailedToGetDatabasesErrorMessage, err
	}
	domains, err := c.repositoryContainer.GetDomainRepository().GetFreeDomains()
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
func (c *Controller) importApplicationToEnvironment(environmentID int64, fileName string, mappingRules parser.ApplicationImportMapping) (*parser.ApplicationImportResult, error) {
	results := parser.NewApplicationImportResult()
	// get the file content
	csvData, err := c.csvStorage.Read(fileName + ".csv")
	if err != nil {
		return &results, err
	}

	// parse the file content every line represents an application
	for rowIndex, line := range csvData {
		importRow := mappingRules.MapRow(line)
		// set the response to empty string. it means process went well
		results[rowIndex] = importRow
		clientName := importRow.RowData["client"]
		projectName := importRow.RowData["project"]
		runtimeName := importRow.RowData["runtime"]
		poolName := importRow.RowData["pool"]
		databaseTypeName := importRow.RowData["database"]

		domainsRaw := importRow.RowData["domains"]
		framework := importRow.RowData["framework"]
		databaseName := importRow.RowData["database_name"]
		if databaseName == "-" {
			databaseName = ""
		}
		databaseUser := importRow.RowData["database_user"]
		if databaseUser == "-" {
			databaseUser = ""
		}
		docRoot := importRow.RowData["doc_root"]
		repository := importRow.RowData["repository"]
		branch := importRow.RowData["branch"]

		// if the client name does not exist, create it
		client, err := c.repositoryContainer.GetClientRepository().GetClientByName(clientName)
		if err != nil {
			client, err = c.repositoryContainer.GetClientRepository().CreateClient(clientName)
			if err != nil {
				currentData := results[rowIndex]
				currentData.ErrorMessage = err.Error()
				results[rowIndex] = currentData
				continue
			}
		}

		// if the project name does not exist, create it
		project, err := c.repositoryContainer.GetProjectRepository().GetProjectByName(projectName)
		if err != nil {
			project, err = c.repositoryContainer.GetProjectRepository().CreateProject(projectName)
			if err != nil {
				currentData := results[rowIndex]
				currentData.ErrorMessage = err.Error()
				results[rowIndex] = currentData
				continue
			}
		}
		// if the runtime name does not exist, create it
		runtime, err := c.repositoryContainer.GetRuntimeRepository().GetRuntimeByName(runtimeName)
		if err != nil {
			runtime, err = c.repositoryContainer.GetRuntimeRepository().CreateRuntime(runtimeName)
			if err != nil {
				currentData := results[rowIndex]
				currentData.ErrorMessage = err.Error()
				results[rowIndex] = currentData
				continue
			}
		}
		// if the pool name does not exist, create it
		pool, err := c.repositoryContainer.GetPoolRepository().GetPoolByName(poolName)
		if err != nil {
			pool, err = c.repositoryContainer.GetPoolRepository().CreatePool(poolName)
			if err != nil {
				currentData := results[rowIndex]
				currentData.ErrorMessage = err.Error()
				results[rowIndex] = currentData
				continue
			}
		}
		// if the database name does not exist, create it
		database, err := c.repositoryContainer.GetDatabaseRepository().GetDatabaseByName(databaseTypeName)
		if err != nil {
			database, err = c.repositoryContainer.GetDatabaseRepository().CreateDatabase(databaseTypeName)
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
			domain, err := c.repositoryContainer.GetDomainRepository().GetDomainByName(domainName)
			if err != nil {
				domain, err = c.repositoryContainer.GetDomainRepository().CreateDomain(domainName)
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
		app, err := c.repositoryContainer.GetApplicationRepository().CreateApplication(client.ID, project.ID, environmentID, database.ID, runtime.ID, pool.ID, repository, branch, databaseName, databaseUser, framework, docRoot, domainIDs)
		currentData := results[rowIndex]
		if err != nil {
			currentData.ErrorMessage = err.Error()
			results[rowIndex] = currentData
			continue
		}
		currentData.Application = app
		results[rowIndex] = currentData
	}
	return &results, nil
}
