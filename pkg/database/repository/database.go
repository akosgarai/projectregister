package repository

import (
	"time"

	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/model"
)

// DatabaseRepository type
type DatabaseRepository struct {
	db *database.DB
}

// NewDatabaseRepository creates a new database repository
func NewDatabaseRepository(db *database.DB) *DatabaseRepository {
	return &DatabaseRepository{
		db: db,
	}
}

// CreateDatabase creates a new database
// the input parameter is the name
// it returns the created database and an error
func (r *DatabaseRepository) CreateDatabase(name string) (*model.Database, error) {
	var database model.Database
	query := "INSERT INTO databases (name) VALUES ($1) RETURNING *"
	err := r.db.QueryRow(query, name).Scan(&database.ID, &database.Name, &database.CreatedAt, &database.UpdatedAt)

	return &database, err
}

// GetDatabaseByName gets a database by name
// the input parameter is the database name
// it returns the database and an error
func (r *DatabaseRepository) GetDatabaseByName(name string) (*model.Database, error) {
	var database model.Database
	query := "SELECT * FROM databases WHERE name = $1"
	err := r.db.QueryRow(query, name).Scan(&database.ID, &database.Name, &database.CreatedAt, &database.UpdatedAt)

	return &database, err
}

// GetDatabaseByID gets a database by id
// the input parameter is the database id
// it returns the database and an error
func (r *DatabaseRepository) GetDatabaseByID(id int64) (*model.Database, error) {
	var database model.Database
	query := "SELECT * FROM databases WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&database.ID, &database.Name, &database.CreatedAt, &database.UpdatedAt)

	return &database, err
}

// UpdateDatabase updates a database
// the input parameter is the database
// it returns an error
func (r *DatabaseRepository) UpdateDatabase(database *model.Database) error {
	query := "UPDATE databases SET name = $1, updated_at = $2 WHERE id = $3"
	now := time.Now().Format("2006-01-02 15:04:05.999999-07:00")
	_, err := r.db.Exec(query, database.Name, now, database.ID)

	return err
}

// DeleteDatabase deletes a database
// the input parameter is the database id
// it returns an error
func (r *DatabaseRepository) DeleteDatabase(id int64) error {
	query := "DELETE FROM databases WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

// GetDatabases gets all databases
// it returns the databases and an error
func (r *DatabaseRepository) GetDatabases() (*model.Databases, error) {
	// get all databases
	var databases model.Databases
	query := "SELECT * FROM databases"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var database model.Database
		err = rows.Scan(&database.ID, &database.Name, &database.CreatedAt, &database.UpdatedAt)
		if err != nil {
			return nil, err
		}
		databases = append(databases, &database)
	}
	return &databases, nil
}
