package repository

import (
	"time"

	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/model"
)

// ApplicationRepository type
type ApplicationRepository struct {
	db *database.DB
}

// NewApplicationRepository creates a new application repository
func NewApplicationRepository(db *database.DB) *ApplicationRepository {
	return &ApplicationRepository{
		db: db,
	}
}

// CreateApplication creates a new application
// the input parameter is the name
// it returns the created application and an error
func (a *ApplicationRepository) CreateApplication(
	clientID, projectID, environmentID, databaseID, runtimeID, poolID int64,
	repository, branch, dbName, dbUser, framework, docRoot string,
	domains []int64) (*model.Application, error) {
	var appID int64
	query := "INSERT INTO applications (client_id, project_id, env_id, database_id, runtime_id, pool_id, repository, branch, db_name, db_user, framework, document_root) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id"
	err := a.db.QueryRow(query, clientID, projectID, environmentID, databaseID, runtimeID, poolID, repository, branch, dbName, dbUser, framework, docRoot).Scan(&appID)
	if err != nil {
		return nil, err
	}
	// create the application domain relations
	for _, domainID := range domains {
		query = "INSERT INTO application_to_domains (application_id, domain_id) VALUES ($1, $2)"
		_, err = a.db.Exec(query, appID, domainID)
		if err != nil {
			return nil, err
		}
	}

	return a.GetApplicationByID(appID)
}

// GetApplicationByID gets a application by id
// the input parameter is the application id
// it returns the application and an error
func (a *ApplicationRepository) GetApplicationByID(id int64) (*model.Application, error) {
	var application model.Application
	application.Client = &model.Client{}
	application.Project = &model.Project{}
	application.Environment = &model.Environment{}
	application.Database = &model.Database{}
	application.Runtime = &model.Runtime{}
	application.Pool = &model.Pool{}
	application.ID = id
	query := "SELECT clients.*, projects.*, environments.*, databases.*, runtimes.*, pools.*, a.repository, a.branch, a.db_name, a.db_user, a.framework, a.document_root, a.created_at, a.updated_at FROM applications a JOIN clients ON a.client_id = clients.id JOIN projects ON a.project_id = projects.id JOIN environments ON a.env_id = environments.id JOIN databases ON a.database_id = databases.id JOIN runtimes ON a.runtime_id = runtimes.id JOIN pools ON a.pool_id = pools.id WHERE a.id = $1"
	err := a.db.QueryRow(query, id).Scan(
		&application.Client.ID, &application.Client.Name, &application.Client.CreatedAt, &application.Client.UpdatedAt,
		&application.Project.ID, &application.Project.Name, &application.Project.CreatedAt, &application.Project.UpdatedAt,
		&application.Environment.ID, &application.Environment.Name, &application.Environment.Description, &application.Environment.CreatedAt, &application.Environment.UpdatedAt,
		&application.Database.ID, &application.Database.Name, &application.Database.CreatedAt, &application.Database.UpdatedAt,
		&application.Runtime.ID, &application.Runtime.Name, &application.Runtime.CreatedAt, &application.Runtime.UpdatedAt,
		&application.Pool.ID, &application.Pool.Name, &application.Pool.CreatedAt, &application.Pool.UpdatedAt,
		&application.Repository, &application.Branch, &application.DBName, &application.DBUser, &application.Framework, &application.DocumentRoot, &application.CreatedAt, &application.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return a.withRelations(&application)
}

// UpdateApplication updates a application
// the input parameter is the application
// it returns an error
func (a *ApplicationRepository) UpdateApplication(application *model.Application) error {
	query := "UPDATE applications SET client_id = $1, project_id = $2, env_id = $3, database_id = $4, runtime_id = $5, pool_id = $6, repository = $7, branch = $8, db_name = $9, db_user = $10, framework = $11, document_root = $12, updated_at = $13 WHERE id = $14"
	now := time.Now().Format("2006-01-02 15:04:05.999999-07:00")
	_, err := a.db.Exec(query, application.Client.ID, application.Project.ID, application.Environment.ID, application.Database.ID, application.Runtime.ID, application.Pool.ID, application.Repository, application.Branch, application.DBName, application.DBUser, application.Framework, application.DocumentRoot, now, application.ID)

	if err != nil {
		return err
	}

	// update the application domain relations
	query = "DELETE FROM application_to_domains WHERE application_id = $1"
	_, err = a.db.Exec(query, application.ID)
	if err != nil {
		return err
	}
	for _, domain := range application.Domains {
		query = "INSERT INTO application_to_domains (application_id, domain_id) VALUES ($1, $2)"
		_, err = a.db.Exec(query, application.ID, domain.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteApplication deletes a application
// the input parameter is the application id
// it returns an error
func (a *ApplicationRepository) DeleteApplication(id int64) error {
	query := "DELETE FROM application_to_domains WHERE application_id = $1"
	_, err := a.db.Exec(query, id)
	if err != nil {
		return err
	}
	query = "DELETE FROM applications WHERE id = $1"
	_, err = a.db.Exec(query, id)
	return err
}

// GetApplications gets all applications
// it returns the applications and an error
func (a *ApplicationRepository) GetApplications() ([]*model.Application, error) {
	// get all applications
	var applications []*model.Application
	query := "SELECT id FROM applications"
	rows, err := a.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		applicationWithRelations, err := a.GetApplicationByID(id)
		if err != nil {
			return nil, err
		}
		applications = append(applications, applicationWithRelations)
	}
	return applications, nil
}

// withRelations function gets a application as input and returns a application with the relations
func (a *ApplicationRepository) withRelations(application *model.Application) (*model.Application, error) {
	domainRepository := NewDomainRepository(a.db)
	// get the application domains
	query := "SELECT domain_id FROM application_to_domains WHERE application_id = $1"
	rows, err := a.db.Query(query, application.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var domainID int64
		err = rows.Scan(&domainID)
		if err != nil {
			return nil, err
		}
		domain, err := domainRepository.GetDomainByID(domainID)
		if err != nil {
			return nil, err
		}
		application.Domains = append(application.Domains, domain)
	}

	return application, nil
}
