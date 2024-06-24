package model

// Project type
type Project struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string
}

// ProjectRepository interface
type ProjectRepository interface {
	CreateProject(name string) (*Project, error)
	GetProjectByName(name string) (*Project, error)
	GetProjectByID(id int64) (*Project, error)
	UpdateProject(client *Project) error
	DeleteProject(id int64) error
	GetProjects() ([]*Project, error)
}
