package model

// Client type
type Client struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string
}

// ClientRepository interface
type ClientRepository interface {
	CreateClient(name string) (*Client, error)
	GetClientByName(name string) (*Client, error)
	GetClientByID(id int64) (*Client, error)
	UpdateClient(client *Client) error
	DeleteClient(id int64) error
	GetClients() ([]*Client, error)
}
