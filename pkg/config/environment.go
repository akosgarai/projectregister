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
}

// DefaultEnvironment creates a new instance of the environment with default values.
func DefaultEnvironment() *Environment {
	return &Environment{
		serverWriteTimeout: DefaultServerWriteTimeout,
		serverReadTimeout:  DefaultServerReadTimeout,
		serverIdleTimeout:  DefaultServerIdleTimeout,
		serverAddr:         DefaultServerAddr,
		serverPort:         DefaultServerPort,
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
