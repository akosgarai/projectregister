package controller

var (
	// ApplicationApplicationIDInvalidErrorMessage is the error message prefix for the invalid application id.
	ApplicationApplicationIDInvalidErrorMessage = "Invalid application id"
	// ApplicationCreateClientIDInvalidErrorMessage is the error message for the invalid client id in the application form.
	ApplicationCreateClientIDInvalidErrorMessage = "Invalid client id"
	// ApplicationCreateCreateApplicationErrorMessage is the error message for the failed application creation.
	ApplicationCreateCreateApplicationErrorMessage = "Failed to create the application"
	// ApplicationCreateDatabaseIDInvalidErrorMessage is the error message for the invalid database id in the application form.
	ApplicationCreateDatabaseIDInvalidErrorMessage = "Invalid database id"
	// ApplicationCreateDomainIDInvalidErrorMessage is the error message for the invalid domain id in the application form.
	ApplicationCreateDomainIDInvalidErrorMessage = "Invalid domain id"
	// ApplicationCreateEnvironmentIDInvalidErrorMessage is the error message for the invalid environment id in the application form.
	ApplicationCreateEnvironmentIDInvalidErrorMessage = "Invalid environment id"
	// ApplicationCreateFailedToGetClientsErrorMessage is the error message for the failed clients get.
	ApplicationCreateFailedToGetClientsErrorMessage = "Failed to get clients"
	// ApplicationCreateFailedToGetDatabasesErrorMessage is the error message for the failed databases get.
	ApplicationCreateFailedToGetDatabasesErrorMessage = "Failed to get databases"
	// ApplicationCreateFailedToGetDomainsErrorMessage is the error message for the failed domains get.
	ApplicationCreateFailedToGetDomainsErrorMessage = "Failed to get domains"
	// ApplicationCreateFailedToGetEnvironmentsErrorMessage is the error message for the failed environments get.
	ApplicationCreateFailedToGetEnvironmentsErrorMessage = "Failed to get environments"
	// ApplicationCreateFailedToGetFrameworksErrorMessage is the error message for the failed frameworks get.
	ApplicationCreateFailedToGetFrameworksErrorMessage = "Failed to get frameworks"
	// ApplicationCreateFailedToGetPoolsErrorMessage is the error message for the failed pools get.
	ApplicationCreateFailedToGetPoolsErrorMessage = "Failed to get pools"
	// ApplicationCreateFailedToGetProjectsErrorMessage is the error message for the failed projects get.
	ApplicationCreateFailedToGetProjectsErrorMessage = "Failed to get projects"
	// ApplicationCreateFailedToGetRuntimesErrorMessage is the error message for the failed runtimes get.
	ApplicationCreateFailedToGetRuntimesErrorMessage = "Failed to get runtimes"
	// ApplicationCreateFrameworkIDInvalidErrorMessage is the error message for the invalid framework id in the application form.
	ApplicationCreateFrameworkIDInvalidErrorMessage = "Invalid framework id"
	// ApplicationCreatePoolIDInvalidErrorMessage is the error message for the invalid pool id in the application form.
	ApplicationCreatePoolIDInvalidErrorMessage = "Invalid pool id"
	// ApplicationCreateProjectIDInvalidErrorMessage is the error message for the invalid project id in the application form.
	ApplicationCreateProjectIDInvalidErrorMessage = "Invalid project id"
	// ApplicationCreateRequiredFieldMissing is the error message for the required fields in the application create.
	ApplicationCreateRequiredFieldMissing = "Client, project, environment are required"
	// ApplicationCreateRuntimeIDInvalidErrorMessage is the error message for the invalid runtime id in the application form.
	ApplicationCreateRuntimeIDInvalidErrorMessage = "Invalid runtime id"
	// ApplicationDeleteFailedToDeleteErrorMessage is the error message for the failed application deletion.
	ApplicationDeleteFailedToDeleteErrorMessage = "Failed to delete the application"
	// ApplicationFailedToGetApplicationErrorMessage is the error message for the failed application get.
	ApplicationFailedToGetApplicationErrorMessage = "Failed to get application data"
	// ApplicationImportFailedToGetEnvironmentErrorMessage is the error message for the failed environment get.
	ApplicationImportFailedToGetEnvironmentErrorMessage = "Failed to get environment"
	// ApplicationImportFailedToSaveFileErrorMessage is the error message for the failed file save.
	ApplicationImportFailedToSaveFileErrorMessage = "Failed to save the file"
	// ApplicationImportInvalidEnvironmentIDErrorMessage is the error message for the invalid environment id in the application import form.
	ApplicationImportInvalidEnvironmentIDErrorMessage = "Invalid environment id"
	// ApplicationListFailedToGetApplicationsErrorMessage is the error message for the failed applications get.
	ApplicationListFailedToGetApplicationsErrorMessage = "Failed to get applications"
	// ApplicationUpdateUpdateApplicationErrorMessage is the error message for the failed application update.
	ApplicationUpdateUpdateApplicationErrorMessage = "Failed to update the application"
	// AuthFailedToGenerateSessionKeyErrorMessage is the error message for the failed session key generation.
	AuthFailedToGenerateSessionKeyErrorMessage = "Failed to generate session key"
	// ClientClientIDInvalidErrorMessage is the error message prefix for the invalid client id.
	ClientClientIDInvalidErrorMessage = "Invalid client id"
	// ClientCreateCreateClientErrorMessage is the error message for the failed client creation.
	ClientCreateCreateClientErrorMessage = "Failed to create the client"
	// ClientCreateRequiredFieldMissing is the error message for the required fields in the client create.
	ClientCreateRequiredFieldMissing = "Name is required"
	// ClientDeleteFailedToDeleteErrorMessage is the error message for the failed client deletion.
	ClientDeleteFailedToDeleteErrorMessage = "Failed to delete the client"
	// ClientFailedToGetClientErrorMessage is the error message for the failed client get.
	ClientFailedToGetClientErrorMessage = "Failed to get client data"
	// ClientListFailedToGetClientsErrorMessage is the error message for the failed clients get.
	ClientListFailedToGetClientsErrorMessage = "Failed to get clients"
	// ClientUpdateRequiredFieldMissing is the error message for the required fields in the client update.
	ClientUpdateRequiredFieldMissing = "Name is required"
	// ClientUpdateUpdateClientErrorMessage is the error message for the failed client update.
	ClientUpdateUpdateClientErrorMessage = "Failed to update the client"
	// DatabaseCreateCreateDatabaseErrorMessage is the error message for the failed database creation.
	DatabaseCreateCreateDatabaseErrorMessage = "Failed to create the database"
	// DatabaseCreateRequiredFieldMissing is the error message for the required fields in the database create.
	DatabaseCreateRequiredFieldMissing = "Name is required"
	// DatabaseDeleteFailedToDeleteErrorMessage is the error message for the failed database deletion.
	DatabaseDeleteFailedToDeleteErrorMessage = "Failed to delete the database"
	// DatabaseDatabaseIDInvalidErrorMessage is the error message prefix for the invalid database id.
	DatabaseDatabaseIDInvalidErrorMessage = "Invalid database id"
	// DatabaseFailedToGetDatabaseErrorMessage is the error message for the failed database get.
	DatabaseFailedToGetDatabaseErrorMessage = "Failed to get database data"
	// DatabaseListFailedToGetDatabasesErrorMessage is the error message for the failed databases get.
	DatabaseListFailedToGetDatabasesErrorMessage = "Failed to get databases"
	// DatabaseUpdateRequiredFieldMissing is the error message for the required fields in the database update.
	DatabaseUpdateRequiredFieldMissing = "Name is required"
	// DatabaseUpdateUpdateDatabaseErrorMessage is the error message for the failed database update.
	DatabaseUpdateUpdateDatabaseErrorMessage = "Failed to update the database"
	// DomainCheckSSLFailedToUpdateDomainErrorMessage is the error message for the failed domain update.
	DomainCheckSSLFailedToUpdateDomainErrorMessage = "Failed to update the domain"
	// DomainCreateCreateDomainErrorMessage is the error message for the failed domain creation.
	DomainCreateCreateDomainErrorMessage = "Failed to create the domain"
	// DomainCreateRequiredFieldMissing is the error message for the required fields in the domain create.
	DomainCreateRequiredFieldMissing = "Name is required"
	// DomainDeleteFailedToDeleteErrorMessage is the error message for the failed domain deletion.
	DomainDeleteFailedToDeleteErrorMessage = "Failed to delete the domain"
	// DomainDomainIDInvalidErrorMessage is the error message prefix for the invalid domain id.
	DomainDomainIDInvalidErrorMessage = "Invalid domain id"
	// DomainFailedToGetDomainErrorMessage is the error message for the failed domain get.
	DomainFailedToGetDomainErrorMessage = "Failed to get domain data"
	// DomainListFailedToGetDomainsErrorMessage is the error message for the failed domains get.
	DomainListFailedToGetDomainsErrorMessage = "Failed to get domains"
	// DomainUpdateRequiredFieldMissing is the error message for the required fields in the domain update.
	DomainUpdateRequiredFieldMissing = "Name is required"
	// DomainUpdateUpdateDomainErrorMessage is the error message for the failed domain update.
	DomainUpdateUpdateDomainErrorMessage = "Failed to update the domain"
	// EnvironmentCreateCreateEnvironmentErrorMessage is the error message for the failed environment creation.
	EnvironmentCreateCreateEnvironmentErrorMessage = "Failed to create the environment"
	// EnvironmentCreateDatabaseIDInvalidErrorMessage is the error message for the invalid database id in the environment form.
	EnvironmentCreateDatabaseIDInvalidErrorMessage = "Invalid database id"
	// EnvironmentCreateFailedToGetDatabasesErrorMessage is the error message for the failed databases get.
	EnvironmentCreateFailedToGetDatabasesErrorMessage = "Failed to get databases"
	// EnvironmentCreateFailedToGetServersErrorMessage is the error message for the failed servers get
	EnvironmentCreateFailedToGetServersErrorMessage = "Failed to get projects"
	// EnvironmentCreateServerIDInvalidErrorMessage is the error message for the invalid server id in the environment form.
	EnvironmentCreateServerIDInvalidErrorMessage = "Invalid server id"
	// EnvironmentCreateRequiredFieldMissing is the error message for the required fields in the environment create.
	EnvironmentCreateRequiredFieldMissing = "Name is required"
	// EnvironmentDeleteFailedToDeleteErrorMessage is the error message for the failed environment deletion.
	EnvironmentDeleteFailedToDeleteErrorMessage = "Failed to delete the environment"
	// EnvironmentEnvironmentIDInvalidErrorMessage is the error message prefix for the invalid environment id.
	EnvironmentEnvironmentIDInvalidErrorMessage = "Invalid environment id"
	// EnvironmentFailedToGetEnvironmentErrorMessage is the error message for the failed environment get.
	EnvironmentFailedToGetEnvironmentErrorMessage = "Failed to get environment data"
	// EnvironmentListFailedToGetServersErrorMessage is the error message for the failed servers get.
	EnvironmentListFailedToGetServersErrorMessage = "Failed to get servers"
	// EnvironmentListFailedToGetDatabasesErrorMessage is the error message for the failed databases get.
	EnvironmentListFailedToGetDatabasesErrorMessage = "Failed to get databases"
	// EnvironmentListFailedToGetEnvironmentsErrorMessage is the error message for the failed environments get.
	EnvironmentListFailedToGetEnvironmentsErrorMessage = "Failed to get environments"
	// EnvironmentListDatabaseIDInvalidErrorMessage is the error message for the invalid database id in the environment list form.
	EnvironmentListDatabaseIDInvalidErrorMessage = "Invalid database id"
	// EnvironmentListServerIDInvalidErrorMessage is the error message for the invalid server id in the environment list form.
	EnvironmentListServerIDInvalidErrorMessage = "Invalid server id"
	// EnvironmentUpdateDatabaseIDInvalidErrorMessage is the error message for the invalid database id in the environment form.
	EnvironmentUpdateDatabaseIDInvalidErrorMessage = "Invalid database id"
	// EnvironmentUpdateFailedToGetDatabasesErrorMessage is the error message for the failed databases get.
	EnvironmentUpdateFailedToGetDatabasesErrorMessage = "Failed to get databases"
	// EnvironmentUpdateFailedToGetServersErrorMessage is the error message for the failed servers get.
	EnvironmentUpdateFailedToGetServersErrorMessage = "Failed to get servers"
	// EnvironmentUpdateRequiredFieldMissing is the error message for the required fields in the environment update.
	EnvironmentUpdateRequiredFieldMissing = "Name is required"
	// EnvironmentUpdateServerIDInvalidErrorMessage is the error message for the invalid server id in the environment form.
	EnvironmentUpdateServerIDInvalidErrorMessage = "Invalid server id"
	// EnvironmentUpdateUpdateEnvironmentErrorMessage is the error message for the failed environment update.
	EnvironmentUpdateUpdateEnvironmentErrorMessage = "Failed to update the environment"

	// FrameworkCreateCreateFrameworkErrorMessage is the error message for the failed framework creation.
	FrameworkCreateCreateFrameworkErrorMessage = "Failed to create the framework"
	// FrameworkCreateRequiredFieldMissing is the error message for the required fields in the framework create.
	FrameworkCreateRequiredFieldMissing = "Name is required"
	// FrameworkDeleteFailedToDeleteErrorMessage is the error message for the failed framework deletion.
	FrameworkDeleteFailedToDeleteErrorMessage = "Failed to delete the framework"
	// FrameworkFailedToGetFrameworkErrorMessage is the error message for the failed framework get.
	FrameworkFailedToGetFrameworkErrorMessage = "Failed to get framework data"
	// FrameworkFrameworkIDInvalidErrorMessage is the error message prefix for the invalid framework id.
	FrameworkFrameworkIDInvalidErrorMessage = "Invalid framework id"
	// FrameworkListFailedToGetFrameworksErrorMessage is the error message for the failed frameworks get.
	FrameworkListFailedToGetFrameworksErrorMessage = "Failed to get frameworks"
	// FrameworkUpdateRequiredFieldMissing is the error message for the required fields in the framework update.
	FrameworkUpdateRequiredFieldMissing = "Name is required"
	// FrameworkUpdateUpdateFrameworkErrorMessage is the error message for the failed framework update.
	FrameworkUpdateUpdateFrameworkErrorMessage = "Failed to update the framework"

	// PoolCreateCreatePoolErrorMessage is the error message for the failed pool creation.
	PoolCreateCreatePoolErrorMessage = "Failed to create the pool"
	// PoolCreateRequiredFieldMissing is the error message for the required fields in the pool create.
	PoolCreateRequiredFieldMissing = "Name is required"
	// PoolDeleteFailedToDeleteErrorMessage is the error message for the failed pool deletion.
	PoolDeleteFailedToDeleteErrorMessage = "Failed to delete the pool"
	// PoolPoolIDInvalidErrorMessage is the error message prefix for the invalid pool id.
	PoolPoolIDInvalidErrorMessage = "Invalid pool id"
	// PoolFailedToGetPoolErrorMessage is the error message for the failed pool get.
	PoolFailedToGetPoolErrorMessage = "Failed to get pool data"
	// PoolListFailedToGetPoolsErrorMessage is the error message for the failed pools get.
	PoolListFailedToGetPoolsErrorMessage = "Failed to get pools"
	// PoolUpdateRequiredFieldMissing is the error message for the required fields in the pool update.
	PoolUpdateRequiredFieldMissing = "Name is required"
	// PoolUpdateUpdatePoolErrorMessage is the error message for the failed pool update.
	PoolUpdateUpdatePoolErrorMessage = "Failed to update the pool"
	// ProjectCreateCreateProjectErrorMessage is the error message for the failed project creation.
	ProjectCreateCreateProjectErrorMessage = "Failed to create the project"
	// ProjectCreateRequiredFieldMissing is the error message for the required fields in the project create.
	ProjectCreateRequiredFieldMissing = "Name is required"
	// ProjectDeleteFailedToDeleteErrorMessage is the error message for the failed project deletion.
	ProjectDeleteFailedToDeleteErrorMessage = "Failed to delete the project"
	// ProjectFailedToGetProjectErrorMessage is the error message for the failed project get.
	ProjectFailedToGetProjectErrorMessage = "Failed to get project data"
	// ProjectListFailedToGetProjectsErrorMessage is the error message for the failed projects get.
	ProjectListFailedToGetProjectsErrorMessage = "Failed to get projects"
	// ProjectProjectIDInvalidErrorMessage is the error message prefix for the invalid project id.
	ProjectProjectIDInvalidErrorMessage = "Invalid project id"
	// ProjectUpdateRequiredFieldMissing is the error message for the required fields in the project update.
	ProjectUpdateRequiredFieldMissing = "Name is required"
	// ProjectUpdateUpdateProjectErrorMessage is the error message for the failed project update.
	ProjectUpdateUpdateProjectErrorMessage = "Failed to update the project"
	// RoleFailedToGetRoleErrorMessage is the error message for the failed role get.
	RoleFailedToGetRoleErrorMessage = "Failed to get role data"
	// RoleFailedToGetResourcesErrorMessage is the error message for the failed resources get.
	RoleFailedToGetResourcesErrorMessage = "Failed to get resources"
	// RoleRoleIDInvalidErrorMessage is the error message prefix for the invalid role id.
	RoleRoleIDInvalidErrorMessage = "Invalid role id"
	// RoleResourceIDInvalidErrorMessage is the error message for the invalid resource id in the role form.
	RoleResourceIDInvalidErrorMessage = "Invalid resource id"
	// RoleCreateRequiredFieldMissing is the error message for the required fields in the role create.
	RoleCreateRequiredFieldMissing = "Name is required"
	// RoleCreateCreateRoleErrorMessage is the error message for the failed role creation.
	RoleCreateCreateRoleErrorMessage = "Failed to create the role"
	// RoleDeleteFailedToDeleteErrorMessage is the error message for the failed role deletion.
	RoleDeleteFailedToDeleteErrorMessage = "Failed to delete the role"
	// RoleUpdateRequiredFieldMissing is the error message for the required fields in the role update.
	RoleUpdateRequiredFieldMissing = "Name is required"
	// RoleUpdateUpdateRoleErrorMessage is the error message for the failed role update.
	RoleUpdateUpdateRoleErrorMessage = "Failed to update the role"
	// RoleListFailedToGetRolesErrorMessage is the error message for the failed roles get.
	RoleListFailedToGetRolesErrorMessage = "Failed to get roles"
	// RuntimeCreateCreateRuntimeErrorMessage is the error message for the failed runtime creation.
	RuntimeCreateCreateRuntimeErrorMessage = "Failed to create the runtime"
	// RuntimeCreateRequiredFieldMissing is the error message for the required fields in the runtime create.
	RuntimeCreateRequiredFieldMissing = "Name is required"
	// RuntimeDeleteFailedToDeleteErrorMessage is the error message for the failed runtime deletion.
	RuntimeDeleteFailedToDeleteErrorMessage = "Failed to delete the runtime"
	// RuntimeRuntimeIDInvalidErrorMessage is the error message prefix for the invalid runtime id.
	RuntimeRuntimeIDInvalidErrorMessage = "Invalid runtime id"
	// RuntimeFailedToGetRuntimeErrorMessage is the error message for the failed runtime get.
	RuntimeFailedToGetRuntimeErrorMessage = "Failed to get runtime data"
	// RuntimeListFailedToGetRuntimesErrorMessage is the error message for the failed runtimes get.
	RuntimeListFailedToGetRuntimesErrorMessage = "Failed to get runtimes"
	// RuntimeUpdateRequiredFieldMissing is the error message for the required fields in the runtime update.
	RuntimeUpdateRequiredFieldMissing = "Name is required"
	// RuntimeUpdateUpdateRuntimeErrorMessage is the error message for the failed runtime update.
	RuntimeUpdateUpdateRuntimeErrorMessage = "Failed to update the runtime"
	// ServerCreateCreateServerErrorMessage is the error message for the failed server creation.
	ServerCreateCreateServerErrorMessage = "Failed to create the server"
	// ServerCreateRequiredFieldMissing is the error message for the required fields in the server create.
	ServerCreateRequiredFieldMissing = "Name and remote address are required"
	// ServerCreatePoolIDInvalidErrorMessage is the error message for the invalid pool id in the server form.
	ServerCreatePoolIDInvalidErrorMessage = "Invalid pool id"
	// ServerCreateFailedToGetRuntimesErrorMessage is the error message for the failed runtimes get.
	ServerCreateFailedToGetRuntimesErrorMessage = "Failed to get runtimes"
	// ServerCreateFailedToGetPoolsErrorMessage is the error message for the failed pools get.
	ServerCreateFailedToGetPoolsErrorMessage = "Failed to get pools"
	// ServerCreateRuntimeIDInvalidErrorMessage is the error message for the invalid runtime id in the server form.
	ServerCreateRuntimeIDInvalidErrorMessage = "Invalid runtime id"
	// ServerDeleteFailedToDeleteErrorMessage is the error message for the failed server deletion.
	ServerDeleteFailedToDeleteErrorMessage = "Failed to delete the server"
	// ServerServerIDInvalidErrorMessage is the error message prefix for the invalid server id.
	ServerServerIDInvalidErrorMessage = "Invalid server id"
	// ServerFailedToGetServerErrorMessage is the error message for the failed server get.
	ServerFailedToGetServerErrorMessage = "Failed to get server data"
	// ServerListFailedToGetPoolsErrorMessage is the error message for the failed pools get.
	ServerListFailedToGetPoolsErrorMessage = "Failed to get pools"
	// ServerListFailedToGetRuntimesErrorMessage is the error message for the failed runtimes get.
	ServerListFailedToGetRuntimesErrorMessage = "Failed to get runtimes"
	// ServerListFailedToGetServersErrorMessage is the error message for the failed servers get.
	ServerListFailedToGetServersErrorMessage = "Failed to get servers"
	// ServerListPoolIDInvalidErrorMessage is the error message for the invalid pool id in the server list form.
	ServerListPoolIDInvalidErrorMessage = "Invalid pool id"
	// ServerListRuntimeIDInvalidErrorMessage is the error message for the invalid runtime id in the server list form.
	ServerListRuntimeIDInvalidErrorMessage = "Invalid runtime id"
	// ServerUpdateRequiredFieldMissing is the error message for the required fields in the server update.
	ServerUpdateRequiredFieldMissing = "Name and remote address are required"
	// ServerUpdateUpdateServerErrorMessage is the error message for the failed server update.
	ServerUpdateUpdateServerErrorMessage = "Failed to update the server"
	// ServerUpdatePoolIDInvalidErrorMessage is the error message for the invalid pool id in the server form.
	ServerUpdatePoolIDInvalidErrorMessage = "Invalid pool id"
	// ServerUpdateRuntimeIDInvalidErrorMessage is the error message for the invalid runtime id in the server form.
	ServerUpdateRuntimeIDInvalidErrorMessage = "Invalid runtime id"
	// UserCreateRequiredFieldMissing is the error message for the required fields in the user create.
	UserCreateRequiredFieldMissing = "Name, email, password and role are required"
	// UserCreateCreateUserErrorMessagePrefix is the error message prefix for the failed user creation.
	UserCreateCreateUserErrorMessagePrefix = "Internal server error - failed to create the user"
	// UserUpdateFailedToGetUserErrorMessage is the error message for the failed user get.
	UserUpdateFailedToGetUserErrorMessage = "Internal server error - failed to get the user"
	// UserUpdateRequiredFieldMissing is the error message for the required fields in the user update.
	UserUpdateRequiredFieldMissing = "Name, email and role are required"
	// UserUpdateFailedToUpdateUserErrorMessage is the error message for the failed user update.
	UserUpdateFailedToUpdateUserErrorMessage = "Internal server error - failed to update the user"
	// UserDeleteFailedToDeleteErrorMessage is the error message for the failed user deletion.
	UserDeleteFailedToDeleteErrorMessage = "Internal server error - failed to delete the user"
	// UserListFailedToGetUsersErrorMessage is the error message for the failed users get.
	UserListFailedToGetUsersErrorMessage = "Internal server error - failed to get the users"
	// UserPasswordEncriptionFailedErrorMessage is the error message for the failed password encryption.
	UserPasswordEncriptionFailedErrorMessage = "Internal server error - failed to encrypt the password"
	// UserRoleIDInvalidErrorMessagePrefix is the error message prefix for the invalid role id.
	UserRoleIDInvalidErrorMessagePrefix = "Invalid role id"
	// UserUserIDInvalidErrorMessagePrefix is the error message prefix for the invalid user id.
	UserUserIDInvalidErrorMessagePrefix = "Invalid user id"
	// UserFailedToGetUserErrorMessage is the error message for the failed user get.
	UserFailedToGetUserErrorMessage = "Failed to get user data"
	// UserFailedToGetRolesErrorMessage is the error message for the failed roles get.
	UserFailedToGetRolesErrorMessage = "Failed to get roles"
)
