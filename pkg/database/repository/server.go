package repository

import (
	"time"

	"github.com/akosgarai/projectregister/pkg/database"
	"github.com/akosgarai/projectregister/pkg/model"
)

// ServerRepository type
type ServerRepository struct {
	db *database.DB
}

// NewServerRepository creates a new server repository
func NewServerRepository(db *database.DB) *ServerRepository {
	return &ServerRepository{
		db: db,
	}
}

// CreateServer creates a new server
// the input parameter is the name
// it returns the created server and an error
func (r *ServerRepository) CreateServer(name, description, remoteAddress string, runtimes, pools []int64) (*model.Server, error) {
	var server model.Server
	query := "INSERT INTO servers (name, description, remote_address) VALUES ($1, $2, $3) RETURNING *"
	err := r.db.QueryRow(query, name, description, remoteAddress).Scan(&server.ID, &server.Name, &server.Description, &server.RemoteAddr, &server.CreatedAt, &server.UpdatedAt)
	if err != nil {
		return nil, err
	}
	// create the server runtime relations
	for _, runtimeID := range runtimes {
		query = "INSERT INTO server_to_runtime (server_id, runtime_id) VALUES ($1, $2)"
		_, err = r.db.Exec(query, server.ID, runtimeID)
		if err != nil {
			return nil, err
		}
	}
	// create the server pool relations
	for _, poolID := range pools {
		query = "INSERT INTO server_to_pool (server_id, pool_id) VALUES ($1, $2)"
		_, err = r.db.Exec(query, server.ID, poolID)
		if err != nil {
			return nil, err
		}
	}

	return r.withRelations(&server)
}

// GetServerByName gets a server by name
// the input parameter is the server name
// it returns the server and an error
func (r *ServerRepository) GetServerByName(name string) (*model.Server, error) {
	var server model.Server
	query := "SELECT * FROM servers WHERE name = $1"
	err := r.db.QueryRow(query, name).Scan(&server.ID, &server.Name, &server.Description, &server.RemoteAddr, &server.CreatedAt, &server.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return r.withRelations(&server)
}

// GetServerByRemoteAddress gets a server by remote address
// the input parameter is the server remote address
// it returns the server and an error
func (r *ServerRepository) GetServerByRemoteAddress(remoteAddress string) (*model.Server, error) {
	var server model.Server
	query := "SELECT * FROM servers WHERE remote_address = $1"
	err := r.db.QueryRow(query, remoteAddress).Scan(&server.ID, &server.Name, &server.Description, &server.RemoteAddr, &server.CreatedAt, &server.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return r.withRelations(&server)
}

// GetServerByID gets a server by id
// the input parameter is the server id
// it returns the server and an error
func (r *ServerRepository) GetServerByID(id int64) (*model.Server, error) {
	var server model.Server
	query := "SELECT * FROM servers WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&server.ID, &server.Name, &server.Description, &server.RemoteAddr, &server.CreatedAt, &server.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return r.withRelations(&server)
}

// UpdateServer updates a server
// the input parameter is the server
// it returns an error
func (r *ServerRepository) UpdateServer(server *model.Server) error {
	query := "UPDATE servers SET name = $1, updated_at = $2 WHERE id = $3"
	now := time.Now().Format("2006-01-02 15:04:05.999999-07:00")
	_, err := r.db.Exec(query, server.Name, now, server.ID)

	if err != nil {
		return err
	}

	// update the server runtime relations
	query = "DELETE FROM server_to_runtime WHERE server_id = $1"
	_, err = r.db.Exec(query, server.ID)
	if err != nil {
		return err
	}
	for _, runtime := range server.Runtimes {
		query = "INSERT INTO server_to_runtime (server_id, runtime_id) VALUES ($1, $2)"
		_, err = r.db.Exec(query, server.ID, runtime.ID)
		if err != nil {
			return err
		}
	}

	// update the server pool relations
	query = "DELETE FROM server_to_pool WHERE server_id = $1"
	_, err = r.db.Exec(query, server.ID)
	if err != nil {
		return err
	}
	for _, pool := range server.Pools {
		query = "INSERT INTO server_to_pool (server_id, pool_id) VALUES ($1, $2)"
		_, err = r.db.Exec(query, server.ID, pool.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteServer deletes a server
// the input parameter is the server id
// it returns an error
func (r *ServerRepository) DeleteServer(id int64) error {
	query := "DELETE FROM server_to_runtime WHERE server_id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	query = "DELETE FROM server_to_pool WHERE server_id = $1"
	_, err = r.db.Exec(query, id)
	if err != nil {
		return err
	}
	query = "DELETE FROM servers WHERE id = $1"
	_, err = r.db.Exec(query, id)
	return err
}

// GetServers gets all servers
// it returns the servers and an error
func (r *ServerRepository) GetServers() (*model.Servers, error) {
	// get all servers
	var servers model.Servers
	query := "SELECT * FROM servers"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var server model.Server
		err = rows.Scan(&server.ID, &server.Name, &server.Description, &server.RemoteAddr, &server.CreatedAt, &server.UpdatedAt)
		if err != nil {
			return nil, err
		}
		serverWithRelations, err := r.withRelations(&server)
		if err != nil {
			return nil, err
		}
		servers = append(servers, serverWithRelations)
	}
	return &servers, nil
}

// withRelations function gets a server as input and returns a server with the relations
func (r *ServerRepository) withRelations(server *model.Server) (*model.Server, error) {
	runtimeRepository := NewRuntimeRepository(r.db)
	poolRepository := NewPoolRepository(r.db)
	// get the server runtimes
	query := "SELECT runtime_id FROM server_to_runtime WHERE server_id = $1"
	rows, err := r.db.Query(query, server.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var runtimeID int64
		err = rows.Scan(&runtimeID)
		if err != nil {
			return nil, err
		}
		runtime, err := runtimeRepository.GetRuntimeByID(runtimeID)
		if err != nil {
			return nil, err
		}
		server.Runtimes = append(server.Runtimes, runtime)
	}

	// get the server pools
	query = "SELECT pool_id FROM server_to_pool WHERE server_id = $1"
	rows, err = r.db.Query(query, server.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var poolID int64
		err = rows.Scan(&poolID)
		if err != nil {
			return nil, err
		}
		pool, err := poolRepository.GetPoolByID(poolID)
		if err != nil {
			return nil, err
		}
		server.Pools = append(server.Pools, pool)
	}

	return server, nil
}
