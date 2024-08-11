package model

// RepositoryContainer interface
type RepositoryContainer interface {
	GetApplicationRepository() ApplicationRepository
	GetClientRepository() ClientRepository
	GetDatabaseRepository() DatabaseRepository
	GetDomainRepository() DomainRepository
	GetEnvironmentRepository() EnvironmentRepository
	GetPoolRepository() PoolRepository
	GetProjectRepository() ProjectRepository
	GetResourceRepository() ResourceRepository
	GetRoleRepository() RoleRepository
	GetRuntimeRepository() RuntimeRepository
	GetServerRepository() ServerRepository
	GetUserRepository() UserRepository
}
