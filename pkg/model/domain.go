package model

// Domain type
type Domain struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string
}

// Domains type is a slice of Domain
type Domains []*Domain

// ToMap converts the Domains to a map
// The key is the domain ID
// The value is the domain Name
func (d Domains) ToMap() map[int64]string {
	result := make(map[int64]string)
	for _, domain := range d {
		result[domain.ID] = domain.Name
	}
	return result
}

// DomainRepository interface
type DomainRepository interface {
	CreateDomain(name string) (*Domain, error)
	GetDomainByName(name string) (*Domain, error)
	GetDomainByID(id int64) (*Domain, error)
	UpdateDomain(client *Domain) error
	DeleteDomain(id int64) error
	GetDomains() (*Domains, error)
	GetFreeDomains() (*Domains, error)
}
