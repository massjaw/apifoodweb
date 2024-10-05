package routes

import "github.com/gin-gonic/gin"

func SetupRouteUser(g *gin.RouterGroup) {
	g.Use()             // add middleware
	g.POST("/register") //add handler
	g.POST("/login")    //add handler
	g.POST("/update")   //add handler
}
