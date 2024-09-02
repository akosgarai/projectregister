package repository

import (
	"time"

	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/model"
)

// FrameworkRepository type
type FrameworkRepository struct {
	db *database.DB
}

// NewFrameworkRepository creates a new framework repository
func NewFrameworkRepository(db *database.DB) *FrameworkRepository {
	return &FrameworkRepository{
		db: db,
	}
}

// CreateFramework creates a new framework
// the input parameter is the name
// it returns the created framework and an error
func (r *FrameworkRepository) CreateFramework(name string, score int) (*model.Framework, error) {
	var framework model.Framework
	query := "INSERT INTO frameworks (name, score) VALUES ($1, $2) RETURNING *"
	err := r.db.QueryRow(query, name, score).Scan(&framework.ID, &framework.Name, &framework.Score, &framework.CreatedAt, &framework.UpdatedAt)

	return &framework, err
}

// GetFrameworkByName gets a framework by name
// the input parameter is the framework name
// it returns the framework and an error
func (r *FrameworkRepository) GetFrameworkByName(name string) (*model.Framework, error) {
	var framework model.Framework
	query := "SELECT * FROM frameworks WHERE name = $1"
	err := r.db.QueryRow(query, name).Scan(&framework.ID, &framework.Name, &framework.Score, &framework.CreatedAt, &framework.UpdatedAt)

	return &framework, err
}

// GetFrameworkByID gets a framework by id
// the input parameter is the framework id
// it returns the framework and an error
func (r *FrameworkRepository) GetFrameworkByID(id int64) (*model.Framework, error) {
	var framework model.Framework
	query := "SELECT * FROM frameworks WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&framework.ID, &framework.Name, &framework.Score, &framework.CreatedAt, &framework.UpdatedAt)

	return &framework, err
}

// UpdateFramework updates a framework
// the input parameter is the framework
// it returns an error
func (r *FrameworkRepository) UpdateFramework(framework *model.Framework) error {
	query := "UPDATE frameworks SET name = $1, score = $2, updated_at = $3 WHERE id = $4"
	now := time.Now().Format("2006-01-02 15:04:05.999999-07:00")
	_, err := r.db.Exec(query, framework.Name, framework.Score, now, framework.ID)

	return err
}

// DeleteFramework deletes a framework
// the input parameter is the framework id
// it returns an error
func (r *FrameworkRepository) DeleteFramework(id int64) error {
	query := "DELETE FROM frameworks WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

// GetFrameworks gets all frameworks
// it returns the frameworks and an error
func (r *FrameworkRepository) GetFrameworks(filters *model.FrameworkFilter) (*model.Frameworks, error) {
	// get all frameworks
	var frameworks model.Frameworks
	query := "SELECT * FROM frameworks"
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
		var framework model.Framework
		err = rows.Scan(&framework.ID, &framework.Name, &framework.Score, &framework.CreatedAt, &framework.UpdatedAt)
		if err != nil {
			return nil, err
		}
		frameworks = append(frameworks, &framework)
	}
	return &frameworks, nil
}
