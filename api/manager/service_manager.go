package manager

import "apifoodweb/api/service"

type ServiceManager interface {
	UserUsecase() service.UserService
}

type serviceManager struct {
	repoManager RepoManager
}

func (u *serviceManager) UserUsecase() service.UserService {
	return service.NewUserService(u.repoManager.UserRepo())
}

func NewServiceManager(rm RepoManager) ServiceManager {
	return &serviceManager{
		repoManager: rm,
	}
}
