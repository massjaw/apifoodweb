package server

import (
	"apifoodweb/api/controller"
	"apifoodweb/api/dto/resp"
	"apifoodweb/api/manager"
	config "apifoodweb/internal"
	"apifoodweb/pkg/util"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type SetupServer struct {
	serviceManager manager.ServiceManager
	Port           string
	Engine         *gin.Engine
	Path           string
	Env            string
}

func (s *SetupServer) route() {
	v1Route := s.Engine.Group("/v1")
	s.userController(v1Route)
}

func (p *SetupServer) userController(rg *gin.RouterGroup) {
	controller.NewUserController(rg, p.serviceManager.UserService())
}

func Server() *SetupServer {

	route := gin.New()
	config := config.NewConfig("DATABASE.APIFOODAPP")

	infraManager := manager.NewInfraManager(config)
	repoManager := manager.NewRepoManager(infraManager)
	serviceManager := manager.NewServiceManager(repoManager)

	env := viper.GetString("ENVIRONMENT")

	if env == "Production" {
		gin.SetMode(gin.ReleaseMode)
	}

	route.HandleMethodNotAllowed = true
	route.NoRoute(func(c *gin.Context) {

		resp.NewErrorJsonResponse(c, http.StatusNotFound, "404", errors.New("invalid endpoint")).Send()
	})
	route.NoMethod(func(c *gin.Context) {

		resp.NewErrorJsonResponse(c, http.StatusMethodNotAllowed, "405", errors.New("method not allowed")).Send()
	})

	return &SetupServer{
		serviceManager: serviceManager,
		Port:           fmt.Sprintf(":%s", config.ApiConfig.ApiPort),
		Engine:         route,
		Path:           viper.GetString("SERVER_HOST"),
		Env:            env,
	}

}

func (s *SetupServer) Run() {

	s.route()
	err := s.Engine.Run(s.Port)

	defer func() {
		if err := recover(); err != nil {
			util.SystemLog("Server-Setup", "Application failed to run", fmt.Errorf("recover: %v", err)).Error()
		}
	}()

	if err != nil {
		util.SystemLog("Server-Setup", "failed to start server on port: "+s.Port, err).Panic()
	}
}
