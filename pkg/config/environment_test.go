package config

import (
	"testing"
)

// TestDefaultEnvironment tests the default environment values.
func TestDefaultEnvironment(t *testing.T) {
	env := DefaultEnvironment()
	if env.GetServerWriteTimeout() != DefaultServerWriteTimeout {
		t.Errorf("Expected %d, got %d", DefaultServerWriteTimeout, env.GetServerWriteTimeout())
	}
	if env.GetServerReadTimeout() != DefaultServerReadTimeout {
		t.Errorf("Expected %d, got %d", DefaultServerReadTimeout, env.GetServerReadTimeout())
	}
	if env.GetServerIdleTimeout() != DefaultServerIdleTimeout {
		t.Errorf("Expected %d, got %d", DefaultServerIdleTimeout, env.GetServerIdleTimeout())
	}
	if env.GetServerAddr() != DefaultServerAddr {
		t.Errorf("Expected %s, got %s", DefaultServerAddr, env.GetServerAddr())
	}
	if env.GetServerPort() != DefaultServerPort {
		t.Errorf("Expected %s, got %s", DefaultServerPort, env.GetServerPort())
	}
	if env.GetMigrationDirectoryPath() != DefaultMigrationDirectoryPath {
		t.Errorf("Expected %s, got %s", DefaultMigrationDirectoryPath, env.GetMigrationDirectoryPath())
	}
	if env.GetDatabaseHost() != DefaultDatabaseHost {
		t.Errorf("Expected %s, got %s", DefaultDatabaseHost, env.GetDatabaseHost())
	}
	if env.GetDatabasePort() != DefaultDatabasePort {
		t.Errorf("Expected %s, got %s", DefaultDatabasePort, env.GetDatabasePort())
	}
	if env.GetDatabaseUser() != DefaultDatabaseUser {
		t.Errorf("Expected %s, got %s", DefaultDatabaseUser, env.GetDatabaseUser())
	}
	if env.GetDatabasePassword() != DefaultDatabasePassword {
		t.Errorf("Expected %s, got %s", DefaultDatabasePassword, env.GetDatabasePassword())
	}
	if env.GetDatabaseName() != DefaultDatabaseName {
		t.Errorf("Expected %s, got %s", DefaultDatabaseName, env.GetDatabaseName())
	}
	if env.GetSessionNameLength() != DefaultSessionNameLength {
		t.Errorf("Expected %d, got %d", DefaultSessionNameLength, env.GetSessionNameLength())
	}
	if env.GetSessionLength() != DefaultSessionLength {
		t.Errorf("Expected %d, got %d", DefaultSessionLength, env.GetSessionLength())
	}
	if env.GetSessionNameAlphabet() != DefaultSessionNameAlphabet {
		t.Errorf("Expected %s, got %s", DefaultSessionNameAlphabet, env.GetSessionNameAlphabet())
	}
	if env.GetRenderTemplateDirectoryPath() != DefaultRenderTemplateDirectoryPath {
		t.Errorf("Expected %s, got %s", DefaultRenderTemplateDirectoryPath, env.GetRenderTemplateDirectoryPath())
	}
	if env.GetRenderBaseTemplate() != DefaultRenderBaseTemplate {
		t.Errorf("Expected %s, got %s", DefaultRenderBaseTemplate, env.GetRenderBaseTemplate())
	}
	if env.GetStaticDirectoryPath() != DefaultStaticDirectoryPath {
		t.Errorf("Expected %s, got %s", DefaultStaticDirectoryPath, env.GetStaticDirectoryPath())
	}
	if env.GetUploadDirectoryPath() != DefaultUploadDirectoryPath {
		t.Errorf("Expected %s, got %s", DefaultUploadDirectoryPath, env.GetUploadDirectoryPath())
	}
}

// TestNewEnvironmentWithEmptyValues tests the NewEnvironment function with empty values.
func TestNewEnvironmentWithEmptyValues(t *testing.T) {
	emptyEnvList := make(map[string]string)
	env := NewEnvironment(emptyEnvList)
	if env.GetServerWriteTimeout() != DefaultServerWriteTimeout {
		t.Errorf("Expected empty string, got %s", env.GetDatabaseUser())
	}
	if env.GetServerWriteTimeout() != DefaultServerWriteTimeout {
		t.Errorf("Expected %d, got %d", DefaultServerWriteTimeout, env.GetServerWriteTimeout())
	}
	if env.GetServerReadTimeout() != DefaultServerReadTimeout {
		t.Errorf("Expected %d, got %d", DefaultServerReadTimeout, env.GetServerReadTimeout())
	}
	if env.GetServerIdleTimeout() != DefaultServerIdleTimeout {
		t.Errorf("Expected %d, got %d", DefaultServerIdleTimeout, env.GetServerIdleTimeout())
	}
	if env.GetServerAddr() != DefaultServerAddr {
		t.Errorf("Expected %s, got %s", DefaultServerAddr, env.GetServerAddr())
	}
	if env.GetServerPort() != DefaultServerPort {
		t.Errorf("Expected %s, got %s", DefaultServerPort, env.GetServerPort())
	}
	if env.GetMigrationDirectoryPath() != DefaultMigrationDirectoryPath {
		t.Errorf("Expected %s, got %s", DefaultMigrationDirectoryPath, env.GetMigrationDirectoryPath())
	}
	if env.GetDatabaseHost() != DefaultDatabaseHost {
		t.Errorf("Expected %s, got %s", DefaultDatabaseHost, env.GetDatabaseHost())
	}
	if env.GetDatabasePort() != DefaultDatabasePort {
		t.Errorf("Expected %s, got %s", DefaultDatabasePort, env.GetDatabasePort())
	}
	if env.GetDatabaseUser() != DefaultDatabaseUser {
		t.Errorf("Expected %s, got %s", DefaultDatabaseUser, env.GetDatabaseUser())
	}
	if env.GetDatabasePassword() != DefaultDatabasePassword {
		t.Errorf("Expected %s, got %s", DefaultDatabasePassword, env.GetDatabasePassword())
	}
	if env.GetDatabaseName() != DefaultDatabaseName {
		t.Errorf("Expected %s, got %s", DefaultDatabaseName, env.GetDatabaseName())
	}
	if env.GetSessionNameLength() != DefaultSessionNameLength {
		t.Errorf("Expected %d, got %d", DefaultSessionNameLength, env.GetSessionNameLength())
	}
	if env.GetSessionLength() != DefaultSessionLength {
		t.Errorf("Expected %d, got %d", DefaultSessionLength, env.GetSessionLength())
	}
	if env.GetSessionNameAlphabet() != DefaultSessionNameAlphabet {
		t.Errorf("Expected %s, got %s", DefaultSessionNameAlphabet, env.GetSessionNameAlphabet())
	}
	if env.GetRenderTemplateDirectoryPath() != DefaultRenderTemplateDirectoryPath {
		t.Errorf("Expected %s, got %s", DefaultRenderTemplateDirectoryPath, env.GetRenderTemplateDirectoryPath())
	}
	if env.GetRenderBaseTemplate() != DefaultRenderBaseTemplate {
		t.Errorf("Expected %s, got %s", DefaultRenderBaseTemplate, env.GetRenderBaseTemplate())
	}
	if env.GetStaticDirectoryPath() != DefaultStaticDirectoryPath {
		t.Errorf("Expected %s, got %s", DefaultStaticDirectoryPath, env.GetStaticDirectoryPath())
	}
	if env.GetUploadDirectoryPath() != DefaultUploadDirectoryPath {
		t.Errorf("Expected %s, got %s", DefaultUploadDirectoryPath, env.GetUploadDirectoryPath())
	}
}

// TestNewEnvironmentServerWriteTimeout tests the NewEnvironment function with a server write timeout value.
func TestNewEnvironmentServerWriteTimeout(t *testing.T) {
	envList := make(map[string]string)
	envList[ServerWriteTimeoutEnvName] = "10"
	env := NewEnvironment(envList)
	if env.GetServerWriteTimeout() != 10 {
		t.Errorf("Expected 10, got %d", env.GetServerWriteTimeout())
	}
}

// TestNewEnvironmentServerReadTimeout tests the NewEnvironment function with a server read timeout value.
func TestNewEnvironmentServerReadTimeout(t *testing.T) {
	envList := make(map[string]string)
	envList[ServerReadTimeoutEnvName] = "10"
	env := NewEnvironment(envList)
	if env.GetServerReadTimeout() != 10 {
		t.Errorf("Expected 10, got %d", env.GetServerReadTimeout())
	}
}

// TestNewEnvironmentServerIdleTimeout tests the NewEnvironment function with a server idle timeout value.
func TestNewEnvironmentServerIdleTimeout(t *testing.T) {
	envList := make(map[string]string)
	envList[ServerIdleTimeoutEnvName] = "10"
	env := NewEnvironment(envList)
	if env.GetServerIdleTimeout() != 10 {
		t.Errorf("Expected 10, got %d", env.GetServerIdleTimeout())
	}
}

// TestNewEnvironmentServerAddr tests the NewEnvironment function with a server address value.
func TestNewEnvironmentServerAddr(t *testing.T) {
	envList := make(map[string]string)
	serverAddr := "10.20.30.40"
	envList[ServerAddrEnvName] = serverAddr
	env := NewEnvironment(envList)
	if env.GetServerAddr() != serverAddr {
		t.Errorf("Expected %s, got %s", serverAddr, env.GetServerAddr())
	}
}

// TestNewEnvironmentServerPort tests the NewEnvironment function with a server port value.
func TestNewEnvironmentServerPort(t *testing.T) {
	envList := make(map[string]string)
	serverPort := "8080"
	envList[ServerPortEnvName] = serverPort
	env := NewEnvironment(envList)
	if env.GetServerPort() != serverPort {
		t.Errorf("Expected %s, got %s", serverPort, env.GetServerPort())
	}
}

// TestNewEnvironmentMigrationDirectoryPath tests the NewEnvironment function with a migration directory path value.
func TestNewEnvironmentMigrationDirectoryPath(t *testing.T) {
	envList := make(map[string]string)
	migrationDirectoryPath := "db/migrations/updated/path"
	envList[MigrationDirectoryPathEnvName] = migrationDirectoryPath
	env := NewEnvironment(envList)
	if env.GetMigrationDirectoryPath() != migrationDirectoryPath {
		t.Errorf("Expected %s, got %s", migrationDirectoryPath, env.GetMigrationDirectoryPath())
	}
}

// TestNewEnvironmentDatabaseHost tests the NewEnvironment function with a database host value.
func TestNewEnvironmentDatabaseHost(t *testing.T) {
	envList := make(map[string]string)
	databaseHost := "dbhost"
	envList[DatabaseHostEnvName] = databaseHost
	env := NewEnvironment(envList)
	if env.GetDatabaseHost() != databaseHost {
		t.Errorf("Expected %s, got %s", databaseHost, env.GetDatabaseHost())
	}
}

// TestNewEnvironmentDatabasePort tests the NewEnvironment function with a database port value.
func TestNewEnvironmentDatabasePort(t *testing.T) {
	envList := make(map[string]string)
	databasePort := "5433"
	envList[DatabasePortEnvName] = databasePort
	env := NewEnvironment(envList)
	if env.GetDatabasePort() != databasePort {
		t.Errorf("Expected %s, got %s", databasePort, env.GetDatabasePort())
	}
}

// TestNewEnvironmentDatabaseUser tests the NewEnvironment function with a database user value.
func TestNewEnvironmentDatabaseUser(t *testing.T) {
	envList := make(map[string]string)
	databaseUser := "dbuser-updated"
	envList[DatabaseUserEnvName] = databaseUser
	env := NewEnvironment(envList)
	if env.GetDatabaseUser() != databaseUser {
		t.Errorf("Expected %s, got %s", databaseUser, env.GetDatabaseUser())
	}
}

// TestNewEnvironmentDatabasePassword tests the NewEnvironment function with a database password value.
func TestNewEnvironmentDatabasePassword(t *testing.T) {
	envList := make(map[string]string)
	databasePassword := "password-updated"
	envList[DatabasePasswordEnvName] = databasePassword
	env := NewEnvironment(envList)
	if env.GetDatabasePassword() != databasePassword {
		t.Errorf("Expected %s, got %s", databasePassword, env.GetDatabasePassword())
	}
}

// TestNewEnvironmentDatabaseName tests the NewEnvironment function with a database name value.
func TestNewEnvironmentDatabaseName(t *testing.T) {
	envList := make(map[string]string)
	databaseName := "dbname-updated"
	envList[DatabaseNameEnvName] = databaseName
	env := NewEnvironment(envList)
	if env.GetDatabaseName() != databaseName {
		t.Errorf("Expected %s, got %s", databaseName, env.GetDatabaseName())
	}
}

// TestNewEnvironmentSessionNameLength tests the NewEnvironment function with a session name length value.
func TestNewEnvironmentSessionNameLength(t *testing.T) {
	envList := make(map[string]string)
	sessionNameLength := "10"
	envList[SessionNameLengthEnvName] = sessionNameLength
	env := NewEnvironment(envList)
	if env.GetSessionNameLength() != 10 {
		t.Errorf("Expected 10, got %d", env.GetSessionNameLength())
	}
}

// TestNewEnvironmentSessionLength tests the NewEnvironment function with a session length value.
func TestNewEnvironmentSessionLength(t *testing.T) {
	envList := make(map[string]string)
	sessionLength := "10"
	envList[SessionLengthEnvName] = sessionLength
	env := NewEnvironment(envList)
	if env.GetSessionLength() != 10 {
		t.Errorf("Expected 10, got %d", env.GetSessionLength())
	}
}

// TestNewEnvironmentSessionNameAlphabet tests the NewEnvironment function with a session name alphabet value.
func TestNewEnvironmentSessionNameAlphabet(t *testing.T) {
	envList := make(map[string]string)
	sessionNameAlphabet := "abc"
	envList[SessionNameAlphabetEnvName] = sessionNameAlphabet
	env := NewEnvironment(envList)
	if env.GetSessionNameAlphabet() != sessionNameAlphabet {
		t.Errorf("Expected %s, got %s", sessionNameAlphabet, env.GetSessionNameAlphabet())
	}
}

// TestNewEnvironmentRenderTemplateDirectoryPath tests the NewEnvironment function with a render template directory path value.
func TestNewEnvironmentRenderTemplateDirectoryPath(t *testing.T) {
	envList := make(map[string]string)
	renderTemplateDirectoryPath := "web/template/updated/path"
	envList[RenderTemplateDirectoryPathEnvName] = renderTemplateDirectoryPath
	env := NewEnvironment(envList)
	if env.GetRenderTemplateDirectoryPath() != renderTemplateDirectoryPath {
		t.Errorf("Expected %s, got %s", renderTemplateDirectoryPath, env.GetRenderTemplateDirectoryPath())
	}
}

// TestNewEnvironmentRenderBaseTemplate tests the NewEnvironment function with a render base template value.
func TestNewEnvironmentRenderBaseTemplate(t *testing.T) {
	envList := make(map[string]string)
	renderBaseTemplate := "base.html"
	envList[RenderBaseTemplateEnvName] = renderBaseTemplate
	env := NewEnvironment(envList)
	if env.GetRenderBaseTemplate() != renderBaseTemplate {
		t.Errorf("Expected %s, got %s", renderBaseTemplate, env.GetRenderBaseTemplate())
	}
}

// TestNewEnvironmentStaticDirectoryPath tests the NewEnvironment function with a static directory path value.
func TestNewEnvironmentStaticDirectoryPath(t *testing.T) {
	envList := make(map[string]string)
	staticDirectoryPath := "web/static/updated/path"
	envList[StaticDirectoryPathEnvName] = staticDirectoryPath
	env := NewEnvironment(envList)
	if env.GetStaticDirectoryPath() != staticDirectoryPath {
		t.Errorf("Expected %s, got %s", staticDirectoryPath, env.GetStaticDirectoryPath())
	}
}

// TestNewEnvironmentUploadDirectoryPath tests the NewEnvironment function with an upload directory path value.
func TestNewEnvironmentUploadDirectoryPath(t *testing.T) {
	envList := make(map[string]string)
	uploadDirectoryPath := "uploads/updated/path"
	envList[UploadDirectoryPathEnvName] = uploadDirectoryPath
	env := NewEnvironment(envList)
	if env.GetUploadDirectoryPath() != uploadDirectoryPath {
		t.Errorf("Expected %s, got %s", uploadDirectoryPath, env.GetUploadDirectoryPath())
	}
}

// TestWrongServerWriteTimeout tests the NewEnvironment function with a wrong server write timeout value.
func TestWrongServerWriteTimeout(t *testing.T) {
	envList := make(map[string]string)
	envList[ServerWriteTimeoutEnvName] = "wrong"
	var fallbackTimeout int64
	env := NewEnvironment(envList)
	if env.GetServerWriteTimeout() != fallbackTimeout {
		t.Errorf("Expected %d, got %d", fallbackTimeout, env.GetServerWriteTimeout())
	}
}
