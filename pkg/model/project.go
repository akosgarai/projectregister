package model

// Project type
type Project struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string
}

// Projects type is a slice of Project
type Projects []*Project

// ToMap converts the Projects to a map
// The key is the string project ID
// The value is the project Name
func (p Projects) ToMap() map[int64]string {
	result := make(map[int64]string)
	for _, project := range p {
		result[project.ID] = project.Name
	}
	return result
}

// ProjectFilter type is the filter for the projects
// It contains the name filter
type ProjectFilter struct {
	Name string
}

// NewProjectFilter creates a new project filter
func NewProjectFilter() *ProjectFilter {
	return &ProjectFilter{
		Name: "",
	}
}

// ProjectRepository interface
type ProjectRepository interface {
	CreateProject(name string) (*Project, error)
	GetProjectByName(name string) (*Project, error)
	GetProjectByID(id int64) (*Project, error)
	UpdateProject(client *Project) error
	DeleteProject(id int64) error
	GetProjects(filters *ProjectFilter) (*Projects, error)
}
