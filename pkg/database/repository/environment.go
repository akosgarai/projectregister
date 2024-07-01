package repository

import (
	"time"

	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/model"
)

// EnvironmentRepository type
type EnvironmentRepository struct {
	db *database.DB
}

// NewEnvironmentRepository creates a new environment repository
func NewEnvironmentRepository(db *database.DB) *EnvironmentRepository {
	return &EnvironmentRepository{
		db: db,
	}
}

// CreateEnvironment creates a new environment
// the input parameter is the name
// it returns the created environment and an error
func (r *EnvironmentRepository) CreateEnvironment(name, description string) (*model.Environment, error) {
	var environment model.Environment
	query := "INSERT INTO environments (name, description) VALUES ($1, $2) RETURNING *"
	err := r.db.QueryRow(query, name, description).Scan(&environment.ID, &environment.Name, &environment.Description, &environment.CreatedAt, &environment.UpdatedAt)

	return &environment, err
}

// GetEnvironmentByName gets a environment by name
// the input parameter is the environment name
// it returns the environment and an error
func (r *EnvironmentRepository) GetEnvironmentByName(name string) (*model.Environment, error) {
	var environment model.Environment
	query := "SELECT * FROM environments WHERE name = $1"
	err := r.db.QueryRow(query, name).Scan(&environment.ID, &environment.Name, &environment.Description, &environment.CreatedAt, &environment.UpdatedAt)

	return &environment, err
}

// GetEnvironmentByID gets a environment by id
// the input parameter is the environment id
// it returns the environment and an error
func (r *EnvironmentRepository) GetEnvironmentByID(id int64) (*model.Environment, error) {
	var environment model.Environment
	query := "SELECT * FROM environments WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&environment.ID, &environment.Name, &environment.Description, &environment.CreatedAt, &environment.UpdatedAt)

	return &environment, err
}

// UpdateEnvironment updates a environment
// the input parameter is the environment
// it returns an error
func (r *EnvironmentRepository) UpdateEnvironment(environment *model.Environment) error {
	query := "UPDATE environments SET name = $1,  description = $2, updated_at = $3 WHERE id = $4"
	now := time.Now().Format("2006-01-02 15:04:05.999999-07:00")
	_, err := r.db.Exec(query, environment.Name, environment.Description, now, environment.ID)

	return err
}

// DeleteEnvironment deletes a environment
// the input parameter is the environment id
// it returns an error
func (r *EnvironmentRepository) DeleteEnvironment(id int64) error {
	query := "DELETE FROM environments WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

// GetEnvironments gets all environments
// it returns the environments and an error
func (r *EnvironmentRepository) GetEnvironments() ([]*model.Environment, error) {
	// get all environments
	var environments []*model.Environment
	query := "SELECT * FROM environments"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var environment model.Environment
		err = rows.Scan(&environment.ID, &environment.Name, &environment.Description, &environment.CreatedAt, &environment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		environments = append(environments, &environment)
	}
	return environments, nil
}
