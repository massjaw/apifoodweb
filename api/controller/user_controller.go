package controller

import (
	"apifoodweb/api/dto/req"
	"apifoodweb/api/service"
	"apifoodweb/pkg/database/model"
	"apifoodweb/pkg/util"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	HandlerResponse
	router  *gin.RouterGroup
	service service.UserService
}

func NewUserController(route *gin.RouterGroup, service service.UserService) *UserController {

	controller := UserController{
		router:  route,
		service: service,
	}

	UserGroup := route.Group("/user")
	UserGroup.Use() //add middleware
	UserGroup.POST("/register", controller.RegisterUser)
	UserGroup.POST("/login", controller.LoginUser)

	return &controller
}

func (c *UserController) RegisterUser(ctx *gin.Context) {

	var request req.UsersRequestAPI

	if err := ctx.BindJSON(&request); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "400", "failed to bind request to json", err)
		return
	}

	if request.Password != request.RePassword {
		c.Failed(ctx, http.StatusBadRequest, "400", "password not match", errors.New("first input password and secon input password not match"))
		return
	}

	user := &model.Users{
		Username:       request.Username,
		Email:          request.Email,
		HashedPassword: util.HashPassword(request.Password),
	}

	res, err := c.service.RegisterUser(user)
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "500", "failed to create new user", err)
		return
	}

	c.Success(ctx, http.StatusOK, "200", "succesfully registered user", res)
}

func (c *UserController) LoginUser(ctx *gin.Context) {

	var request req.UsersRequestAPI

	if err := ctx.BindJSON(&request); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "400", "failed to bind request to json", err)
		return
	}

	user, err := c.service.LoginUser(request.Username)
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "500", "failed to get user info", err)
		return
	}

	if user.HashedPassword != util.HashPassword(request.Password) {
		c.Failed(ctx, http.StatusUnauthorized, "401", "password not match", errors.New("password didn't match with existing password"))
		return
	}

	c.Success(ctx, http.StatusOK, "200", "succesfully login", nil)
}
