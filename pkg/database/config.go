package database

import (
	"github.com/akosgarai/projectregister/pkg/config"
)

// get the database url
func getDatabaseURL(envConfig *config.Environment) string {
	return "postgres://" + envConfig.GetDatabaseUser() + ":" + envConfig.GetDatabasePassword() + "@" + envConfig.GetDatabaseHost() + ":" + envConfig.GetDatabasePort() + "/" + envConfig.GetDatabaseName() + "?sslmode=disable"
}
