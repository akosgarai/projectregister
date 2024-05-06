package config

import (
	"strconv"
)

// Environment is a struct that holds the environment configuration.
type Environment struct {
	serverWriteTimeout int64
	serverReadTimeout  int64
	serverIdleTimeout  int64
	serverAddr         string
	serverPort         string

	migrationDirectoryPath string

	databaseHost     string
	databasePort     string
	databaseUser     string
	databasePassword string
	databaseName     string

	sessionNameLength   int
	sessionLength       int64
	sessionNameAlphabet string

	renderTemplateDirectoryPath string
	renderBaseTemplate          string
}

// DefaultEnvironment creates a new instance of the environment with default values.
func DefaultEnvironment() *Environment {
	return &Environment{
		serverWriteTimeout: DefaultServerWriteTimeout,
		serverReadTimeout:  DefaultServerReadTimeout,
		serverIdleTimeout:  DefaultServerIdleTimeout,
		serverAddr:         DefaultServerAddr,
		serverPort:         DefaultServerPort,

		migrationDirectoryPath: DefaultMigrationDirectoryPath,

		databaseHost:     DefaultDatabaseHost,
		databasePort:     DefaultDatabasePort,
		databaseUser:     DefaultDatabaseUser,
		databasePassword: DefaultDatabasePassword,
		databaseName:     DefaultDatabaseName,

		sessionNameLength:   DefaultSessionNameLength,
		sessionLength:       DefaultSessionLength,
		sessionNameAlphabet: DefaultSessionNameAlphabet,

		renderTemplateDirectoryPath: DefaultRenderTemplateDirectoryPath,
		renderBaseTemplate:          DefaultRenderBaseTemplate,
	}
}

// GetServerWriteTimeout returns the server write timeout.
func (e *Environment) GetServerWriteTimeout() int64 {
	return e.serverWriteTimeout
}

// GetServerReadTimeout returns the server read timeout.
func (e *Environment) GetServerReadTimeout() int64 {
	return e.serverReadTimeout
}

// GetServerIdleTimeout returns the server idle timeout.
func (e *Environment) GetServerIdleTimeout() int64 {
	return e.serverIdleTimeout
}

// GetServerAddr returns the server address.
func (e *Environment) GetServerAddr() string {
	return e.serverAddr
}

// GetServerPort returns the server port.
func (e *Environment) GetServerPort() string {
	return e.serverPort
}

// GetMigrationDirectoryPath returns the migration directory path.
func (e *Environment) GetMigrationDirectoryPath() string {
	return e.migrationDirectoryPath
}

// GetDatabaseHost returns the database host.
func (e *Environment) GetDatabaseHost() string {
	return e.databaseHost
}

// GetDatabasePort returns the database port.
func (e *Environment) GetDatabasePort() string {
	return e.databasePort
}

// GetDatabaseUser returns the database user.
func (e *Environment) GetDatabaseUser() string {
	return e.databaseUser
}

// GetDatabasePassword returns the database password.
func (e *Environment) GetDatabasePassword() string {
	return e.databasePassword
}

// GetDatabaseName returns the database name.
func (e *Environment) GetDatabaseName() string {
	return e.databaseName
}

// GetSessionNameLength returns the session name length.
func (e *Environment) GetSessionNameLength() int {
	return e.sessionNameLength
}

// GetSessionLength returns the session length.
func (e *Environment) GetSessionLength() int64 {
	return e.sessionLength
}

// GetSessionNameAlphabet returns the session name alphabet.
func (e *Environment) GetSessionNameAlphabet() string {
	return e.sessionNameAlphabet
}

// GetRenderTemplateDirectoryPath returns the render template directory path.
func (e *Environment) GetRenderTemplateDirectoryPath() string {
	return e.renderTemplateDirectoryPath
}

// GetRenderBaseTemplate returns the render base template.
func (e *Environment) GetRenderBaseTemplate() string {
	return e.renderBaseTemplate
}

// NewEnvironment creates a new instance of the environment.
func NewEnvironment(envConfig map[string]string) *Environment {
	env := DefaultEnvironment()
	if val, ok := envConfig[ServerWriteTimeoutEnvName]; ok {
		env.serverWriteTimeout = env.toInt64(val)
	}
	if val, ok := envConfig[ServerReadTimeoutEnvName]; ok {
		env.serverReadTimeout = env.toInt64(val)
	}
	if val, ok := envConfig[ServerIdleTimeoutEnvName]; ok {
		env.serverIdleTimeout = env.toInt64(val)
	}
	if val, ok := envConfig[ServerAddrEnvName]; ok {
		env.serverAddr = val
	}
	if val, ok := envConfig[ServerPortEnvName]; ok {
		env.serverPort = val
	}
	if val, ok := envConfig[MigrationDirectoryPathEnvName]; ok {
		env.migrationDirectoryPath = val
	}
	if val, ok := envConfig[DatabaseHostEnvName]; ok {
		env.databaseHost = val
	}
	if val, ok := envConfig[DatabasePortEnvName]; ok {
		env.databasePort = val
	}
	if val, ok := envConfig[DatabaseUserEnvName]; ok {
		env.databaseUser = val
	}
	if val, ok := envConfig[DatabasePasswordEnvName]; ok {
		env.databasePassword = val
	}
	if val, ok := envConfig[DatabaseNameEnvName]; ok {
		env.databaseName = val
	}
	if val, ok := envConfig[SessionNameLengthEnvName]; ok {
		env.sessionNameLength = int(env.toInt64(val))
	}
	if val, ok := envConfig[SessionLengthEnvName]; ok {
		env.sessionLength = env.toInt64(val)
	}
	if val, ok := envConfig[SessionNameAlphabetEnvName]; ok {
		env.sessionNameAlphabet = val
	}
	if val, ok := envConfig[RenderTemplateDirectoryPathEnvName]; ok {
		env.renderTemplateDirectoryPath = val
	}
	if val, ok := envConfig[RenderBaseTemplateEnvName]; ok {
		env.renderBaseTemplate = val
	}

	return env
}

// toInt converts a string to an integer.
func (e *Environment) toInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
