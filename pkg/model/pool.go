package model

// Pool type
type Pool struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string
}

// Pools type is a slice of Pool
type Pools []*Pool

// ToMap converts the Pools to a map
// The key is the pool ID
// The value is the pool Name
func (p Pools) ToMap() map[int64]string {
	result := make(map[int64]string)
	for _, pool := range p {
		result[pool.ID] = pool.Name
	}
	return result
}

// PoolFilter type is the filter for the pools
// It contains the name filter
type PoolFilter struct {
	Name string
}

// NewPoolFilter creates a new pool filter
func NewPoolFilter() *PoolFilter {
	return &PoolFilter{
		Name: "",
	}
}

// PoolRepository interface
type PoolRepository interface {
	CreatePool(name string) (*Pool, error)
	GetPoolByName(name string) (*Pool, error)
	GetPoolByID(id int64) (*Pool, error)
	UpdatePool(client *Pool) error
	DeletePool(id int64) error
	GetPools(filters *PoolFilter) (*Pools, error)
}
