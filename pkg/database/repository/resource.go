package repository

import (
	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/model"
)

// ResourceRepository type
type ResourceRepository struct {
	db *database.DB
}

// NewResourceRepository creates a new resource repository
func NewResourceRepository(db *database.DB) *ResourceRepository {
	return &ResourceRepository{
		db: db,
	}
}

// GetResources gets all resources
// it returns the resources and an error
func (r *ResourceRepository) GetResources() (*model.Resources, error) {
	// get all resources
	var resources model.Resources
	query := "SELECT * FROM resources"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var res model.Resource
		err = rows.Scan(&res.ID, &res.Name)
		if err != nil {
			return nil, err
		}
		resources = append(resources, &res)
	}
	return &resources, nil
}
