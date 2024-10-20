package manager

import "apifoodweb/api/service"

type ServiceManager interface {
	UserService() service.UserService
}

type serviceManager struct {
	repoManager RepoManager
}

func (u *serviceManager) UserService() service.UserService {
	return service.NewUserService(u.repoManager.UserRepo())
}

func NewServiceManager(rm RepoManager) ServiceManager {
	return &serviceManager{
		repoManager: rm,
	}
}
