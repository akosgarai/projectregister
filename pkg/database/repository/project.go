package repository

import (
	"time"

	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/model"
)

// ProjectRepository type
type ProjectRepository struct {
	db *database.DB
}

// NewProjectRepository creates a new project repository
func NewProjectRepository(db *database.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

// CreateProject creates a new project
// the input parameter is the name
// it returns the created project and an error
func (r *ProjectRepository) CreateProject(name string) (*model.Project, error) {
	var project model.Project
	query := "INSERT INTO projects (name) VALUES ($1) RETURNING *"
	err := r.db.QueryRow(query, name).Scan(&project.ID, &project.Name, &project.CreatedAt, &project.UpdatedAt)

	return &project, err
}

// GetProjectByName gets a project by name
// the input parameter is the project name
// it returns the project and an error
func (r *ProjectRepository) GetProjectByName(name string) (*model.Project, error) {
	var project model.Project
	query := "SELECT * FROM projects WHERE name = $1"
	err := r.db.QueryRow(query, name).Scan(&project.ID, &project.Name, &project.CreatedAt, &project.UpdatedAt)

	return &project, err
}

// GetProjectByID gets a project by id
// the input parameter is the project id
// it returns the project and an error
func (r *ProjectRepository) GetProjectByID(id int64) (*model.Project, error) {
	var project model.Project
	query := "SELECT * FROM projects WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&project.ID, &project.Name, &project.CreatedAt, &project.UpdatedAt)

	return &project, err
}

// UpdateProject updates a project
// the input parameter is the project
// it returns an error
func (r *ProjectRepository) UpdateProject(project *model.Project) error {
	query := "UPDATE projects SET name = $1, updated_at = $2 WHERE id = $3"
	now := time.Now().Format("2006-01-02 15:04:05.999999-07:00")
	_, err := r.db.Exec(query, project.Name, now, project.ID)

	return err
}

// DeleteProject deletes a project
// the input parameter is the project id
// it returns an error
func (r *ProjectRepository) DeleteProject(id int64) error {
	query := "DELETE FROM projects WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

// GetProjects gets all projects
// it returns the projects and an error
func (r *ProjectRepository) GetProjects() (*model.Projects, error) {
	// get all projects
	var projects model.Projects
	query := "SELECT * FROM projects"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var project model.Project
		err = rows.Scan(&project.ID, &project.Name, &project.CreatedAt, &project.UpdatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}
	return &projects, nil
}
