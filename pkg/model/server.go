package model

// Server type
type Server struct {
	ID          int64
	Name        string
	Description string
	RemoteAddr  string
	CreatedAt   string
	UpdatedAt   string

	Runtimes []*Runtime
	Pools    []*Pool
}

// HasRuntime checks if the server has the runtime
func (s *Server) HasRuntime(runtime string) bool {
	for _, rt := range s.Runtimes {
		if rt.Name == runtime {
			return true
		}
	}
	return false
}

// HasPool checks if the server has the pool
func (s *Server) HasPool(pool string) bool {
	for _, p := range s.Pools {
		if p.Name == pool {
			return true
		}
	}
	return false
}

// ServerRepository interface
type ServerRepository interface {
	CreateServer(name, description, remoteAddress string, runtimes, pools []int64) (*Server, error)
	GetServerByName(name string) (*Server, error)
	GetServerByRemoteAddress(remoteAddress string) (*Server, error)
	GetServerByID(id int64) (*Server, error)
	UpdateServer(server *Server) error
	DeleteServer(id int64) error
	GetServers() ([]*Server, error)
}
