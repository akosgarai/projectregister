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
func (r *EnvironmentRepository) CreateEnvironment(name, description string, serverIDs, databaseIDs []int64) (*model.Environment, error) {
	var environment model.Environment
	query := "INSERT INTO environments (name, description) VALUES ($1, $2) RETURNING *"
	err := r.db.QueryRow(query, name, description).Scan(&environment.ID, &environment.Name, &environment.Description, &environment.CreatedAt, &environment.UpdatedAt)
	if err != nil {
		return nil, err
	}
	// insert the server ids to the environment_to_servers table
	for _, id := range serverIDs {
		query = "INSERT INTO environment_to_servers (environment_id, server_id) VALUES ($1, $2)"
		_, err = r.db.Exec(query, environment.ID, id)
		if err != nil {
			return nil, err
		}
	}
	// insert the database ids to the environment_to_databases table
	for _, id := range databaseIDs {
		query = "INSERT INTO environment_to_databases (environment_id, database_id) VALUES ($1, $2)"
		_, err = r.db.Exec(query, environment.ID, id)
		if err != nil {
			return nil, err
		}
	}

	return r.withRelations(&environment)
}

// GetEnvironmentByName gets a environment by name
// the input parameter is the environment name
// it returns the environment and an error
func (r *EnvironmentRepository) GetEnvironmentByName(name string) (*model.Environment, error) {
	var environment model.Environment
	query := "SELECT * FROM environments WHERE name = $1"
	err := r.db.QueryRow(query, name).Scan(&environment.ID, &environment.Name, &environment.Description, &environment.CreatedAt, &environment.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return r.withRelations(&environment)
}

// GetEnvironmentByID gets a environment by id
// the input parameter is the environment id
// it returns the environment and an error
func (r *EnvironmentRepository) GetEnvironmentByID(id int64) (*model.Environment, error) {
	var environment model.Environment
	query := "SELECT * FROM environments WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&environment.ID, &environment.Name, &environment.Description, &environment.CreatedAt, &environment.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return r.withRelations(&environment)
}

// UpdateEnvironment updates a environment
// the input parameter is the environment
// it returns an error
func (r *EnvironmentRepository) UpdateEnvironment(environment *model.Environment) error {
	query := "UPDATE environments SET name = $1,  description = $2, updated_at = $3 WHERE id = $4"
	now := time.Now().Format("2006-01-02 15:04:05.999999-07:00")
	_, err := r.db.Exec(query, environment.Name, environment.Description, now, environment.ID)

	// delete the environment ids from the environment_to_servers table
	query = "DELETE FROM environment_to_servers WHERE environment_id = $1"
	_, err = r.db.Exec(query, environment.ID)
	if err != nil {
		return err
	}
	// insert the server ids to the environment_to_servers table
	for _, server := range environment.Servers {
		query = "INSERT INTO environment_to_servers (environment_id, server_id) VALUES ($1, $2)"
		_, err = r.db.Exec(query, environment.ID, server.ID)
		if err != nil {
			return err
		}
	}
	// delete the environment ids from the environment_to_databases table
	query = "DELETE FROM environment_to_databases WHERE environment_id = $1"
	_, err = r.db.Exec(query, environment.ID)
	if err != nil {
		return err
	}
	// insert the database ids to the environment_to_databases table
	for _, database := range environment.Databases {
		query = "INSERT INTO environment_to_databases (environment_id, database_id) VALUES ($1, $2)"
		_, err = r.db.Exec(query, environment.ID, database.ID)
		if err != nil {
			return err
		}
	}

	return err
}

// DeleteEnvironment deletes a environment
// the input parameter is the environment id
// it returns an error
func (r *EnvironmentRepository) DeleteEnvironment(id int64) error {
	// delete the environment ids from the environment_to_servers table
	query := "DELETE FROM environment_to_servers WHERE environment_id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	// delete the environment ids from the environment_to_databases table
	query = "DELETE FROM environment_to_databases WHERE environment_id = $1"
	_, err = r.db.Exec(query, id)
	if err != nil {
		return err
	}

	query = "DELETE FROM environments WHERE id = $1"
	_, err = r.db.Exec(query, id)
	return err
}

// GetEnvironments gets all environments
// it returns the environments and an error
func (r *EnvironmentRepository) GetEnvironments() (*model.Environments, error) {
	// get all environments
	var environments model.Environments
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
		environmentWithRelations, err := r.withRelations(&environment)
		if err != nil {
			return nil, err
		}
		environments = append(environments, environmentWithRelations)
	}
	return &environments, nil
}

// withRelations function gets a environment as input and returns a environment with the relations
func (r *EnvironmentRepository) withRelations(environment *model.Environment) (*model.Environment, error) {
	query := "SELECT server_id FROM environment_to_servers WHERE environment_id = $1"
	rows, err := r.db.Query(query, environment.ID)
	if err != nil {
		return nil, err
	}
	serverRepository := NewServerRepository(r.db)
	defer rows.Close()
	for rows.Next() {
		var serverID int64
		err = rows.Scan(&serverID)
		if err != nil {
			return nil, err
		}
		server, err := serverRepository.GetServerByID(serverID)
		if err != nil {
			return nil, err
		}
		environment.Servers = append(environment.Servers, server)
	}

	query = "SELECT database_id FROM environment_to_databases WHERE environment_id = $1"
	rows, err = r.db.Query(query, environment.ID)
	if err != nil {
		return nil, err
	}
	databaseRepository := NewDatabaseRepository(r.db)
	defer rows.Close()
	for rows.Next() {
		var databaseID int64
		err = rows.Scan(&databaseID)
		if err != nil {
			return nil, err
		}
		database, err := databaseRepository.GetDatabaseByID(databaseID)
		if err != nil {
			return nil, err
		}
		environment.Databases = append(environment.Databases, database)
	}
	return environment, nil
}
