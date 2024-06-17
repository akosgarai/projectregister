package controller

var (
	// AuthFailedToGenerateSessionKeyErrorMessage is the error message for the failed session key generation.
	AuthFailedToGenerateSessionKeyErrorMessage = "Failed to generate session key"
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
)
