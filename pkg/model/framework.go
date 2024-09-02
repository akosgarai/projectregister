package model

// Framework type
type Framework struct {
	ID        int64
	Name      string
	Score     int
	CreatedAt string
	UpdatedAt string
}

// Frameworks type is a slice of Framework
type Frameworks []*Framework

// ToMap converts the Frameworks to a map
// The key is the framework ID
// The value is the framework Name
func (r Frameworks) ToMap() map[int64]string {
	result := make(map[int64]string)
	for _, framework := range r {
		result[framework.ID] = framework.Name
	}
	return result
}

// FrameworkFilter type is the filter for the frameworks
// It contains the name filter
type FrameworkFilter struct {
	Name string
}

// NewFrameworkFilter creates a new framework filter
func NewFrameworkFilter() *FrameworkFilter {
	return &FrameworkFilter{
		Name: "",
	}
}

// FrameworkRepository interface
type FrameworkRepository interface {
	CreateFramework(name string, score int) (*Framework, error)
	GetFrameworkByName(name string) (*Framework, error)
	GetFrameworkByID(id int64) (*Framework, error)
	UpdateFramework(client *Framework) error
	DeleteFramework(id int64) error
	GetFrameworks(filter *FrameworkFilter) (*Frameworks, error)
}
