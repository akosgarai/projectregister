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
)
