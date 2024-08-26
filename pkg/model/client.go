package model

// Client type
type Client struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string
}

// Clients type is a slice of Client
type Clients []*Client

// ToMap converts the Clients to a map
// The key is the string client ID
// The value is the client Name
func (c Clients) ToMap() map[int64]string {
	result := make(map[int64]string)
	for _, client := range c {
		result[client.ID] = client.Name
	}
	return result
}

// ClientFilter type is the filter for the clients
// It contains the name filter
type ClientFilter struct {
	Name string
}

// NewClientFilter creates a new client filter
func NewClientFilter() *ClientFilter {
	return &ClientFilter{
		Name: "",
	}
}

// ClientRepository interface
type ClientRepository interface {
	CreateClient(name string) (*Client, error)
	GetClientByName(name string) (*Client, error)
	GetClientByID(id int64) (*Client, error)
	UpdateClient(client *Client) error
	DeleteClient(id int64) error
	GetClients(filters *ClientFilter) (*Clients, error)
}
