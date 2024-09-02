package repository

import (
	"time"

	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/model"
)

// RuntimeRepository type
type RuntimeRepository struct {
	db *database.DB
}

// NewRuntimeRepository creates a new runtime repository
func NewRuntimeRepository(db *database.DB) *RuntimeRepository {
	return &RuntimeRepository{
		db: db,
	}
}

// CreateRuntime creates a new runtime
// the input parameter is the name
// it returns the created runtime and an error
func (r *RuntimeRepository) CreateRuntime(name string, score int) (*model.Runtime, error) {
	var runtime model.Runtime
	query := "INSERT INTO runtimes (name, score) VALUES ($1, $2) RETURNING *"
	err := r.db.QueryRow(query, name, score).Scan(&runtime.ID, &runtime.Name, &runtime.CreatedAt, &runtime.UpdatedAt, &runtime.Score)

	return &runtime, err
}

// GetRuntimeByName gets a runtime by name
// the input parameter is the runtime name
// it returns the runtime and an error
func (r *RuntimeRepository) GetRuntimeByName(name string) (*model.Runtime, error) {
	var runtime model.Runtime
	query := "SELECT * FROM runtimes WHERE name = $1"
	err := r.db.QueryRow(query, name).Scan(&runtime.ID, &runtime.Name, &runtime.CreatedAt, &runtime.UpdatedAt, &runtime.Score)

	return &runtime, err
}

// GetRuntimeByID gets a runtime by id
// the input parameter is the runtime id
// it returns the runtime and an error
func (r *RuntimeRepository) GetRuntimeByID(id int64) (*model.Runtime, error) {
	var runtime model.Runtime
	query := "SELECT * FROM runtimes WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&runtime.ID, &runtime.Name, &runtime.CreatedAt, &runtime.UpdatedAt, &runtime.Score)

	return &runtime, err
}

// UpdateRuntime updates a runtime
// the input parameter is the runtime
// it returns an error
func (r *RuntimeRepository) UpdateRuntime(runtime *model.Runtime) error {
	query := "UPDATE runtimes SET name = $1, score = $2, updated_at = $3 WHERE id = $4"
	now := time.Now().Format("2006-01-02 15:04:05.999999-07:00")
	_, err := r.db.Exec(query, runtime.Name, runtime.Score, now, runtime.ID)

	return err
}

// DeleteRuntime deletes a runtime
// the input parameter is the runtime id
// it returns an error
func (r *RuntimeRepository) DeleteRuntime(id int64) error {
	query := "DELETE FROM runtimes WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

// GetRuntimes gets all runtimes
// it returns the runtimes and an error
func (r *RuntimeRepository) GetRuntimes(filters *model.RuntimeFilter) (*model.Runtimes, error) {
	// get all runtimes
	var runtimes model.Runtimes
	query := "SELECT * FROM runtimes"
	params := []interface{}{}
	if filters.Name != "" {
		query += " WHERE name LIKE '%' || $1 || '%'"
		params = append(params, filters.Name)
	}
	rows, err := r.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var runtime model.Runtime
		err = rows.Scan(&runtime.ID, &runtime.Name, &runtime.CreatedAt, &runtime.UpdatedAt, &runtime.Score)
		if err != nil {
			return nil, err
		}
		runtimes = append(runtimes, &runtime)
	}
	return &runtimes, nil
}
