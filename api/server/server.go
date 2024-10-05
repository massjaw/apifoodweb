package server

import (
	"apifoodweb/src/dto/resp"
	"apifoodweb/src/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type SetupServer struct {
	Port   string
	Engine *gin.Engine
	Path   string
	Env    string
}

func InitApplicationServer() *SetupServer {

	route := gin.New()
	env := viper.GetString("Environment")

	if env == "Production" {
		gin.SetMode(gin.ReleaseMode)
	}

	route.HandleMethodNotAllowed = true
	route.NoMethod(func(c *gin.Context) {

		resp.MethodNotAllowed().HandleResponse(c)
	})
	route.NoRoute(func(c *gin.Context) {

		resp.NoRoute(c.Request.RequestURI).HandleResponse(c)
	})

	return &SetupServer{
		Port:   viper.GetString("Port"),
		Engine: route,
		Path:   viper.GetString("Release_Host"),
		Env:    env,
	}
}

func (s *SetupServer) Run() {
	// userRouter := s.Engine.Group("/user")
	// routes.SetupRouteUser(userRouter)

	if err := s.Engine.Run(s.Port); err != nil {
		util.SystemLog("Server-Setup", "failed to start server on port: "+s.Port, err).Fatal()
	}
}
