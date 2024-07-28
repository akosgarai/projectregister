package resources

const (
	// UserResource is the resource name for the user.
	UserResource = "user"
	// RoleResource is the resource name for the role.
	RoleResource = "role"
	// ClientResource is the resource name for the client.
	ClientResource = "client"
	// ProjectResource is the resource name for the project.
	ProjectResource = "project"
	// DomainResource is the resource name for the domain.
	DomainResource = "domain"
	// EnvironmentResource is the resource name for the environment.
	EnvironmentResource = "environment"
	// RuntimeResource is the resource name for the runtime.
	RuntimeResource = "runtime"
	// PoolResource is the resource name for the pool.
	PoolResource = "pool"
	// DatabaseResource is the resource name for the database.
	DatabaseResource = "database"
	// ServerResource is the resource name for the server.
	ServerResource = "server"
	// ApplicationResource is the resource name for the application.
	ApplicationResource = "application"

	// UsersPrivilege is the privilege name for the users.
	UsersPrivilege = "users"
	// RolesPrivilege is the privilege name for the roles.
	RolesPrivilege = "roles"
	// ClientsPrivilege is the privilege name for the clients.
	ClientsPrivilege = "clients"
	// ProjectsPrivilege is the privilege name for the projects.
	ProjectsPrivilege = "projects"
	// DomainsPrivilege is the privilege name for the domains.
	DomainsPrivilege = "domains"
	// EnvironmentsPrivilege is the privilege name for the environments.
	EnvironmentsPrivilege = "environments"
	// RuntimesPrivilege is the privilege name for the runtimes.
	RuntimesPrivilege = "runtimes"
	// PoolsPrivilege is the privilege name for the pools.
	PoolsPrivilege = "pools"
	// DatabasesPrivilege is the privilege name for the databases.
	DatabasesPrivilege = "databases"
	// ServersPrivilege is the privilege name for the servers.
	ServersPrivilege = "servers"
	// ApplicationsPrivilege is the privilege name for the applications.
	ApplicationsPrivilege = "applications"
)

var (
	// ResourcePrivileges is a map for the resources and the necessary privileges.
	ResourcePrivileges = map[string]string{
		UserResource:        UsersPrivilege,
		RoleResource:        RolesPrivilege,
		ClientResource:      ClientsPrivilege,
		ProjectResource:     ProjectsPrivilege,
		DomainResource:      DomainsPrivilege,
		EnvironmentResource: EnvironmentsPrivilege,
		RuntimeResource:     RuntimesPrivilege,
		PoolResource:        PoolsPrivilege,
		DatabaseResource:    DatabasesPrivilege,
		ServerResource:      ServersPrivilege,
		ApplicationResource: ApplicationsPrivilege,
	}

	// Resources is a slice of the resource names.
	Resources = []string{
		UserResource, RoleResource, ClientResource, ProjectResource, DomainResource,
		EnvironmentResource, RuntimeResource, PoolResource, DatabaseResource,
		ServerResource, ApplicationResource,
	}
)
