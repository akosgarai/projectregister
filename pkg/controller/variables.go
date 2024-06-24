package controller

var (
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
