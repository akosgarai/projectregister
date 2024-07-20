package model

// Database type
type Database struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string
}

// DatabaseRepository interface
type DatabaseRepository interface {
	CreateDatabase(name string) (*Database, error)
	GetDatabaseByName(name string) (*Database, error)
	GetDatabaseByID(id int64) (*Database, error)
	UpdateDatabase(client *Database) error
	DeleteDatabase(id int64) error
	GetDatabases() ([]*Database, error)
}
