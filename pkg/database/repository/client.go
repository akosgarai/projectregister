package repository

import (
	"time"

	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/model"
)

// ClientRepository type
type ClientRepository struct {
	db *database.DB
}

// NewClientRepository creates a new client repository
func NewClientRepository(db *database.DB) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

// CreateClient creates a new client
// the input parameter is the name
// it returns the created client and an error
func (r *ClientRepository) CreateClient(name string) (*model.Client, error) {
	var client model.Client
	query := "INSERT INTO clients (name) VALUES ($1) RETURNING *"
	err := r.db.QueryRow(query, name).Scan(&client.ID, &client.Name, &client.CreatedAt, &client.UpdatedAt)

	return &client, err
}

// GetClientByName gets a client by name
// the input parameter is the client name
// it returns the client and an error
func (r *ClientRepository) GetClientByName(name string) (*model.Client, error) {
	var client model.Client
	query := "SELECT * FROM clients WHERE name = $1"
	err := r.db.QueryRow(query, name).Scan(&client.ID, &client.Name, &client.CreatedAt, &client.UpdatedAt)

	return &client, err
}

// GetClientByID gets a client by id
// the input parameter is the client id
// it returns the client and an error
func (r *ClientRepository) GetClientByID(id int64) (*model.Client, error) {
	var client model.Client
	query := "SELECT * FROM clients WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&client.ID, &client.Name, &client.CreatedAt, &client.UpdatedAt)

	return &client, err
}

// UpdateClient updates a client
// the input parameter is the client
// it returns an error
func (r *ClientRepository) UpdateClient(client *model.Client) error {
	query := "UPDATE clients SET name = $1, updated_at = $2 WHERE id = $3"
	now := time.Now().Format("2006-01-02 15:04:05.999999-07:00")
	_, err := r.db.Exec(query, client.Name, now, client.ID)

	return err
}

// DeleteClient deletes a client
// the input parameter is the client id
// it returns an error
func (r *ClientRepository) DeleteClient(id int64) error {
	query := "DELETE FROM clients WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

// GetClients gets all clients
// it returns the clients and an error
func (r *ClientRepository) GetClients() ([]*model.Client, error) {
	// get all clients
	var clients []*model.Client
	query := "SELECT * FROM clients"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var client model.Client
		err = rows.Scan(&client.ID, &client.Name, &client.CreatedAt, &client.UpdatedAt)
		if err != nil {
			return nil, err
		}
		clients = append(clients, &client)
	}
	return clients, nil
}
