package config

const (
	// DefaultServerWriteTimeout is the default server write timeout.
	DefaultServerWriteTimeout = 15
	// DefaultServerReadTimeout is the default server read timeout.
	DefaultServerReadTimeout = 15
	// DefaultServerIdleTimeout is the default server idle timeout.
	DefaultServerIdleTimeout = 60
	// DefaultServerAddr is the default server address.
	DefaultServerAddr = "0.0.0.0"
	// DefaultServerPort is the default server port.
	DefaultServerPort = "8090"
	// DefaultMigrationDirectoryPath is the default migration directory path.
	DefaultMigrationDirectoryPath = "db/migrations"
	// DefaultDatabaseUser is the default database user.
	DefaultDatabaseUser = "projectregister"
	// DefaultDatabasePassword is the default database password.
	DefaultDatabasePassword = "password"
	// DefaultDatabaseName is the default database name.
	DefaultDatabaseName = "projectregister_development"
	// DefaultDatabaseHost is the default database host.
	DefaultDatabaseHost = "db"
	// DefaultDatabasePort is the default database port.
	DefaultDatabasePort = "5432"
	// DefaultSessionNameLength is the default session name length.
	DefaultSessionNameLength = 32
	// DefaultSessionLength is the default session length in minutes.
	DefaultSessionLength = 30
	// DefaultSessionNameAlphabet is the default session name alphabet.
	DefaultSessionNameAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	// DefaultRenderTemplateDirectoryPath is the default render template directory path.
	DefaultRenderTemplateDirectoryPath = "./web/template"
	// DefaultRenderBaseTemplate is the default render base template.
	DefaultRenderBaseTemplate = "base.html.tmpl"
	// DefaultStaticDirectoryPath is the default render static directory path.
	DefaultStaticDirectoryPath = "./web/public"
	// DefaultUploadDirectoryPath is the default upload directory path.
	DefaultUploadDirectoryPath = "./uploads"

	// environment variables

	// ServerWriteTimeoutEnvName is the server write timeout environment variable name.
	ServerWriteTimeoutEnvName = "SERVER_WRITE_TIMEOUT"
	// ServerReadTimeoutEnvName is the server read timeout environment variable name.
	ServerReadTimeoutEnvName = "SERVER_READ_TIMEOUT"
	// ServerIdleTimeoutEnvName is the server idle timeout environment variable name.
	ServerIdleTimeoutEnvName = "SERVER_IDLE_TIMEOUT"
	// ServerAddrEnvName is the server address environment variable name.
	ServerAddrEnvName = "SERVER_ADDR"
	// ServerPortEnvName is the server port environment variable name.
	ServerPortEnvName = "SERVER_PORT"
	// MigrationDirectoryPathEnvName is the migration directory path environment variable name.
	MigrationDirectoryPathEnvName = "MIGRATION_DIRECTORY_PATH"
	// DatabaseUserEnvName is the database user environment variable name.
	DatabaseUserEnvName = "DATABASE_USER"
	// DatabasePasswordEnvName is the database password environment variable name.
	DatabasePasswordEnvName = "DATABASE_PASSWORD"
	// DatabaseNameEnvName is the database name environment variable name.
	DatabaseNameEnvName = "DATABASE_NAME"
	// DatabaseHostEnvName is the database host environment variable name.
	DatabaseHostEnvName = "DATABASE_HOST"
	// DatabasePortEnvName is the database port environment variable name.
	DatabasePortEnvName = "DATABASE_PORT"
	// SessionNameLengthEnvName is the session name length environment variable name.
	SessionNameLengthEnvName = "SESSION_NAME_LENGTH"
	// SessionLengthEnvName is the session length environment variable name.
	SessionLengthEnvName = "SESSION_LENGTH"
	// SessionNameAlphabetEnvName is the session name alphabet environment variable name.
	SessionNameAlphabetEnvName = "SESSION_NAME_ALPHABET"
	// RenderTemplateDirectoryPathEnvName is the render template directory path environment variable name.
	RenderTemplateDirectoryPathEnvName = "RENDER_TEMPLATE_DIRECTORY_PATH"
	// RenderBaseTemplateEnvName is the render base template environment variable name.
	RenderBaseTemplateEnvName = "RENDER_BASE_TEMPLATE"
	// StaticDirectoryPathEnvName is the static directory path environment variable name.
	StaticDirectoryPathEnvName = "STATIC_DIRECTORY_PATH"
	// UploadDirectoryPathEnvName is the upload directory path environment variable name.
	UploadDirectoryPathEnvName = "UPLOAD_DIRECTORY_PATH"
)
