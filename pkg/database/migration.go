package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	// import the postgres driver
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/akosgarai/projectregister/pkg/config"
)

// Migration type
type Migration struct {
	envConfig *config.Environment
}

// NewMigration creates a new instance of the migration.
func NewMigration(envConfig *config.Environment) *Migration {
	return &Migration{
		envConfig: envConfig,
	}
}

// Up executes the migrations
func (m *Migration) Up() error {
	migration, err := migrate.New(
		"file://"+m.envConfig.GetMigrationDirectoryPath(),
		m.getDatabaseURL())
	if err != nil {
		return err
	}
	migration.Up()
	return nil
}

// get the database url
func (m *Migration) getDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		m.envConfig.GetDatabaseUser(),
		m.envConfig.GetDatabasePassword(),
		m.envConfig.GetDatabaseHost(),
		m.envConfig.GetDatabasePort(),
		m.envConfig.GetDatabaseName())
}
