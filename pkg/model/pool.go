package model

// Pool type
type Pool struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string
}

// PoolRepository interface
type PoolRepository interface {
	CreatePool(name string) (*Pool, error)
	GetPoolByName(name string) (*Pool, error)
	GetPoolByID(id int64) (*Pool, error)
	UpdatePool(client *Pool) error
	DeletePool(id int64) error
	GetPools() ([]*Pool, error)
}
