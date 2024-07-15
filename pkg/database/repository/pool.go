package repository

import (
	"time"

	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/model"
)

// PoolRepository type
type PoolRepository struct {
	db *database.DB
}

// NewPoolRepository creates a new pool repository
func NewPoolRepository(db *database.DB) *PoolRepository {
	return &PoolRepository{
		db: db,
	}
}

// CreatePool creates a new pool
// the input parameter is the name
// it returns the created pool and an error
func (r *PoolRepository) CreatePool(name string) (*model.Pool, error) {
	var pool model.Pool
	query := "INSERT INTO pools (name) VALUES ($1) RETURNING *"
	err := r.db.QueryRow(query, name).Scan(&pool.ID, &pool.Name, &pool.CreatedAt, &pool.UpdatedAt)

	return &pool, err
}

// GetPoolByName gets a pool by name
// the input parameter is the pool name
// it returns the pool and an error
func (r *PoolRepository) GetPoolByName(name string) (*model.Pool, error) {
	var pool model.Pool
	query := "SELECT * FROM pools WHERE name = $1"
	err := r.db.QueryRow(query, name).Scan(&pool.ID, &pool.Name, &pool.CreatedAt, &pool.UpdatedAt)

	return &pool, err
}

// GetPoolByID gets a pool by id
// the input parameter is the pool id
// it returns the pool and an error
func (r *PoolRepository) GetPoolByID(id int64) (*model.Pool, error) {
	var pool model.Pool
	query := "SELECT * FROM pools WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&pool.ID, &pool.Name, &pool.CreatedAt, &pool.UpdatedAt)

	return &pool, err
}

// UpdatePool updates a pool
// the input parameter is the pool
// it returns an error
func (r *PoolRepository) UpdatePool(pool *model.Pool) error {
	query := "UPDATE pools SET name = $1, updated_at = $2 WHERE id = $3"
	now := time.Now().Format("2006-01-02 15:04:05.999999-07:00")
	_, err := r.db.Exec(query, pool.Name, now, pool.ID)

	return err
}

// DeletePool deletes a pool
// the input parameter is the pool id
// it returns an error
func (r *PoolRepository) DeletePool(id int64) error {
	query := "DELETE FROM pools WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

// GetPools gets all pools
// it returns the pools and an error
func (r *PoolRepository) GetPools() ([]*model.Pool, error) {
	// get all pools
	var pools []*model.Pool
	query := "SELECT * FROM pools"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var pool model.Pool
		err = rows.Scan(&pool.ID, &pool.Name, &pool.CreatedAt, &pool.UpdatedAt)
		if err != nil {
			return nil, err
		}
		pools = append(pools, &pool)
	}
	return pools, nil
}
