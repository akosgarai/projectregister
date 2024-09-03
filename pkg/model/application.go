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
	Framework   *Framework

	Branch       string
	DBName       string
	DBUser       string
	DocumentRoot string
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
	FrameworkIDs   []string

	Domain     string
	Branch     string
	DBName     string
	DBUser     string
	DocRoot    string
	Repository string
	Score      int

	VisibleColumns []int64
	allColumns     map[int64]string
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
		FrameworkIDs:   []string{},
		Domain:         "",
		Branch:         "",
		DBName:         "",
		DBUser:         "",
		DocRoot:        "",
		Repository:     "",
		Score:          0,

		VisibleColumns: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
		allColumns: map[int64]string{
			0: "ID", 1: "Client", 2: "Project",
			3: "Environment", 4: "Database", 5: "Runtime",
			6: "Pool", 7: "Codebase", 8: "Framework",
			9: "Document Root", 10: "Score", 11: "Domains",
			12: "Created At", 13: "Updated At", 14: "Actions"},
	}
}

// GetAllColumns returns all the columns
func (f *ApplicationFilter) GetAllColumns() map[int64]string {
	return f.allColumns
}

// IsVisibleColumn checks if the column is visible
func (f *ApplicationFilter) IsVisibleColumn(column string) bool {
	for _, v := range f.VisibleColumns {
		if f.allColumns[v] == column {
			return true
		}
	}
	return false
}

// ApplicationRepository interface
type ApplicationRepository interface {
	CreateApplication(clientID, projectID, environmentID, databaseID, runtimeID, poolID, frameworkID int64, repository, branch, dbName, dbUser, docRoot string, domains []int64) (*Application, error)
	GetApplicationByID(id int64) (*Application, error)
	UpdateApplication(application *Application) error
	DeleteApplication(id int64) error
	GetApplications(filter *ApplicationFilter) (*Applications, error)
}
