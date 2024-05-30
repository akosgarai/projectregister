package model

// Resource type
type Resource struct {
	ID   int64
	Name string
}

// ResourceRepository interface
type ResourceRepository interface {
	GetResources() ([]*Resource, error)
}
