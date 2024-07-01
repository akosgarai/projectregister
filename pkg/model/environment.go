package model

// Environment type
type Environment struct {
	ID          int64
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
}

// EnvironmentRepository interface
type EnvironmentRepository interface {
	CreateEnvironment(name, description string) (*Environment, error)
	GetEnvironmentByName(name string) (*Environment, error)
	GetEnvironmentByID(id int64) (*Environment, error)
	UpdateEnvironment(client *Environment) error
	DeleteEnvironment(id int64) error
	GetEnvironments() ([]*Environment, error)
}
