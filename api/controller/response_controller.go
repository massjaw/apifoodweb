package controller

import (
	"apifoodweb/api/dto/resp"

	"github.com/gin-gonic/gin"
)

type HandlerResponse struct{}

func (r *HandlerResponse) Success(c *gin.Context, httpCode int, code string, msg string, data any) {

	resp.NewSuccessJsonResponse(c, httpCode, code, msg, data).Send()
}

func (r *HandlerResponse) Failed(c *gin.Context, httpCode int, code string, msg string, err error) {

	resp.NewErrorJsonResponse(c, httpCode, code, err).Send()
}
