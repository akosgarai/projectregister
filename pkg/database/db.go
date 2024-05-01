package database

import (
	"database/sql"

	// pq is the driver for the postgres database
	_ "github.com/lib/pq"

	"github.com/akosgarai/projectregister/pkg/config"
)

// DB type for database
type DB struct {
	envConfig *config.Environment
	database  *sql.DB
}

// NewDB creates a new database
func NewDB(envConfig *config.Environment) *DB {
	return &DB{
		envConfig: envConfig,
	}
}

// Connect connects to the database
func (d *DB) Connect() error {
	db, err := sql.Open("postgres", getDatabaseURL(d.envConfig))
	if err != nil {
		return err
	}
	d.database = db
	return nil
}

// Close closes the database connection
func (d *DB) Close() error {
	return d.database.Close()
}

// Exec executes a query
func (d *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.database.Exec(query, args...)
}

// QueryRow executes a query that is expected to return at most one row
func (d *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.database.QueryRow(query, args...)
}
