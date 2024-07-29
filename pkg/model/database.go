package model

// Database type
type Database struct {
	ID        int64
	Name      string
	CreatedAt string
	UpdatedAt string
}

// Databases type is a slice of Database
type Databases []*Database

// ToMap converts the Databases to a map
// The key is the database ID
// The value is the database Name
func (d Databases) ToMap() map[int64]string {
	result := make(map[int64]string)
	for _, database := range d {
		result[database.ID] = database.Name
	}
	return result
}

// DatabaseRepository interface
type DatabaseRepository interface {
	CreateDatabase(name string) (*Database, error)
	GetDatabaseByName(name string) (*Database, error)
	GetDatabaseByID(id int64) (*Database, error)
	UpdateDatabase(client *Database) error
	DeleteDatabase(id int64) error
	GetDatabases() (*Databases, error)
}
