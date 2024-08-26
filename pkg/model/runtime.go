package model

// Runtime type
type Runtime struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string
}

// Runtimes type is a slice of Runtime
type Runtimes []*Runtime

// ToMap converts the Runtimes to a map
// The key is the runtime ID
// The value is the runtime Name
func (r Runtimes) ToMap() map[int64]string {
	result := make(map[int64]string)
	for _, runtime := range r {
		result[runtime.ID] = runtime.Name
	}
	return result
}

// RuntimeFilter type is the filter for the runtimes
// It contains the name filter
type RuntimeFilter struct {
	Name string
}

// NewRuntimeFilter creates a new runtime filter
func NewRuntimeFilter() *RuntimeFilter {
	return &RuntimeFilter{
		Name: "",
	}
}

// RuntimeRepository interface
type RuntimeRepository interface {
	CreateRuntime(name string) (*Runtime, error)
	GetRuntimeByName(name string) (*Runtime, error)
	GetRuntimeByID(id int64) (*Runtime, error)
	UpdateRuntime(client *Runtime) error
	DeleteRuntime(id int64) error
	GetRuntimes(filter *RuntimeFilter) (*Runtimes, error)
}
