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

// ApplicationRepository interface
type ApplicationRepository interface {
	CreateApplication(clientID, projectID, environmentID, databaseID, runtimeID, poolID int64, repository, branch, dbName, dbUser, framework, docRoot string, domains []int64) (*Application, error)
	GetApplicationByID(id int64) (*Application, error)
	UpdateApplication(application *Application) error
	DeleteApplication(id int64) error
	GetApplications() ([]*Application, error)
}
