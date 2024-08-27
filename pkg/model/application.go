package model

// Application type
type Application struct {
	ID          int64
	Client      *Client
	Project     *Project
	Environment *Environment
	Database    *Database
	Runtime     *Runtime
	Pool        *Pool

	Branch       string
	DBName       string
	DBUser       string
	DocumentRoot string
	Framework    string
	Repository   string

	CreatedAt string
	UpdatedAt string

	Domains []*Domain
}

// HasDomain checks if the application has the domain
func (a *Application) HasDomain(domain string) bool {
	for _, d := range a.Domains {
		if d.Name == domain {
			return true
		}
	}
	return false
}

// Applications type is a slice of Application
type Applications []*Application

// ApplicationFilter type is the filter for the applications
// It contains the name filter
type ApplicationFilter struct {
	ClientIDs      []string
	ProjectIDs     []string
	EnvironmentIDs []string
	DatabaseIDs    []string
	RuntimeIDs     []string
	PoolIDs        []string

	Domain     string
	Branch     string
	DBName     string
	DBUser     string
	DocRoot    string
	Framework  string
	Repository string
}

// NewApplicationFilter creates a new application filter
func NewApplicationFilter() *ApplicationFilter {
	return &ApplicationFilter{
		ClientIDs:      []string{},
		ProjectIDs:     []string{},
		EnvironmentIDs: []string{},
		DatabaseIDs:    []string{},
		RuntimeIDs:     []string{},
		PoolIDs:        []string{},
		Domain:         "",
		Branch:         "",
		DBName:         "",
		DBUser:         "",
		DocRoot:        "",
		Framework:      "",
		Repository:     "",
	}
}

// ApplicationRepository interface
type ApplicationRepository interface {
	CreateApplication(clientID, projectID, environmentID, databaseID, runtimeID, poolID int64, repository, branch, dbName, dbUser, framework, docRoot string, domains []int64) (*Application, error)
	GetApplicationByID(id int64) (*Application, error)
	UpdateApplication(application *Application) error
	DeleteApplication(id int64) error
	GetApplications(filter *ApplicationFilter) (*Applications, error)
}
