package manager

import "apifoodweb/api/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infraManager.DbConn())
}
