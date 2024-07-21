package repository

import (
	"time"

	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/model"
)

// DomainRepository type
type DomainRepository struct {
	db *database.DB
}

// NewDomainRepository creates a new domain repository
func NewDomainRepository(db *database.DB) *DomainRepository {
	return &DomainRepository{
		db: db,
	}
}

// CreateDomain creates a new domain
// the input parameter is the name
// it returns the created domain and an error
func (r *DomainRepository) CreateDomain(name string) (*model.Domain, error) {
	var domain model.Domain
	query := "INSERT INTO domains (name) VALUES ($1) RETURNING *"
	err := r.db.QueryRow(query, name).Scan(&domain.ID, &domain.Name, &domain.CreatedAt, &domain.UpdatedAt)

	return &domain, err
}

// GetDomainByName gets a domain by name
// the input parameter is the domain name
// it returns the domain and an error
func (r *DomainRepository) GetDomainByName(name string) (*model.Domain, error) {
	var domain model.Domain
	query := "SELECT * FROM domains WHERE name = $1"
	err := r.db.QueryRow(query, name).Scan(&domain.ID, &domain.Name, &domain.CreatedAt, &domain.UpdatedAt)

	return &domain, err
}

// GetDomainByID gets a domain by id
// the input parameter is the domain id
// it returns the domain and an error
func (r *DomainRepository) GetDomainByID(id int64) (*model.Domain, error) {
	var domain model.Domain
	query := "SELECT * FROM domains WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&domain.ID, &domain.Name, &domain.CreatedAt, &domain.UpdatedAt)

	return &domain, err
}

// UpdateDomain updates a domain
// the input parameter is the domain
// it returns an error
func (r *DomainRepository) UpdateDomain(domain *model.Domain) error {
	query := "UPDATE domains SET name = $1, updated_at = $2 WHERE id = $3"
	now := time.Now().Format("2006-01-02 15:04:05.999999-07:00")
	_, err := r.db.Exec(query, domain.Name, now, domain.ID)

	return err
}

// DeleteDomain deletes a domain
// the input parameter is the domain id
// it returns an error
func (r *DomainRepository) DeleteDomain(id int64) error {
	query := "DELETE FROM domains WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

// GetDomains gets all domains
// it returns the domains and an error
func (r *DomainRepository) GetDomains() ([]*model.Domain, error) {
	// get all domains
	var domains []*model.Domain
	query := "SELECT * FROM domains"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var domain model.Domain
		err = rows.Scan(&domain.ID, &domain.Name, &domain.CreatedAt, &domain.UpdatedAt)
		if err != nil {
			return nil, err
		}
		domains = append(domains, &domain)
	}
	return domains, nil
}

// GetFreeDomains gets all domains without any relation
// it returns the domains and an error
func (r *DomainRepository) GetFreeDomains() ([]*model.Domain, error) {
	// get all domains
	var domains []*model.Domain
	query := "SELECT * FROM domains WHERE id NOT IN (SELECT domain_id FROM application_to_domains)"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var domain model.Domain
		err = rows.Scan(&domain.ID, &domain.Name, &domain.CreatedAt, &domain.UpdatedAt)
		if err != nil {
			return nil, err
		}
		domains = append(domains, &domain)
	}
	return domains, nil
}
