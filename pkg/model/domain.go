package model

// Domain type
type Domain struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string
}

// DomainRepository interface
type DomainRepository interface {
	CreateDomain(name string) (*Domain, error)
	GetDomainByName(name string) (*Domain, error)
	GetDomainByID(id int64) (*Domain, error)
	UpdateDomain(client *Domain) error
	DeleteDomain(id int64) error
	GetDomains() ([]*Domain, error)
	GetFreeDomains() ([]*Domain, error)
}
