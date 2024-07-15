package model

// Runtime type
type Runtime struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string
}

// RuntimeRepository interface
type RuntimeRepository interface {
	CreateRuntime(name string) (*Runtime, error)
	GetRuntimeByName(name string) (*Runtime, error)
	GetRuntimeByID(id int64) (*Runtime, error)
	UpdateRuntime(client *Runtime) error
	DeleteRuntime(id int64) error
	GetRuntimes() ([]*Runtime, error)
}
