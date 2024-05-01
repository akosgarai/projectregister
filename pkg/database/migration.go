package database

import (
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
		getDatabaseURL(m.envConfig))
	if err != nil {
		return err
	}
	migration.Up()
	return nil
}
