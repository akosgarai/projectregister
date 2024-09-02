package model

// Environment type
type Environment struct {
	ID          int64
	Name        string
	Description string
	Score       int
	CreatedAt   string
	UpdatedAt   string
	Servers     []*Server
	Databases   []*Database
}

// HasServer checks if the environment has the server
func (e *Environment) HasServer(server string) bool {
	for _, srv := range e.Servers {
		if srv.Name == server {
			return true
		}
	}
	return false
}

// HasDatabase checks if the environment has the database
func (e *Environment) HasDatabase(database string) bool {
	for _, db := range e.Databases {
		if db.Name == database {
			return true
		}
	}
	return false
}

// Environments type is a slice of Environment
type Environments []*Environment

// ToMap converts the Environments to a map
// The key is the environment ID
// The value is the environment Name
func (e Environments) ToMap() map[int64]string {
	result := make(map[int64]string)
	for _, environment := range e {
		result[environment.ID] = environment.Name
	}
	return result
}

// EnvironmentFilter type is the filter for the environments
// It contains the name filter
type EnvironmentFilter struct {
	Name        string
	Description string
	ServerIDs   []string
	DatabaseIDs []string
}

// NewEnvironmentFilter creates a new environment filter
func NewEnvironmentFilter() *EnvironmentFilter {
	return &EnvironmentFilter{
		Name:        "",
		Description: "",
		ServerIDs:   []string{},
		DatabaseIDs: []string{},
	}
}

// EnvironmentRepository interface
type EnvironmentRepository interface {
	CreateEnvironment(name, description string, serverIDs, databaseIDs []int64, score int) (*Environment, error)
	GetEnvironmentByName(name string) (*Environment, error)
	GetEnvironmentByID(id int64) (*Environment, error)
	UpdateEnvironment(client *Environment) error
	DeleteEnvironment(id int64) error
	GetEnvironments(filter *EnvironmentFilter) (*Environments, error)
}
