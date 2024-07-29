package model

// Resource type
type Resource struct {
	ID   int64
	Name string
}

// Resources type is a slice of Resource
type Resources []*Resource

// ToMap converts the Resources to a map
// The key is the string resource ID
// The value is the resource Name
func (r Resources) ToMap() map[int64]string {
	result := make(map[int64]string)
	for _, resource := range r {
		result[resource.ID] = resource.Name
	}
	return result
}

// ResourceRepository interface
type ResourceRepository interface {
	GetResources() (*Resources, error)
}
