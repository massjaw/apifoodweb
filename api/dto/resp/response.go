package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonResponse struct {
	StatusCode    int         `json:"statusCode,omitempty"`
	StatusMessage string      `json:"statusMessage,omitempty"`
	Result        interface{} `json:"result,omitempty"`
}

func (r *CommonResponse) HandleResponse(ctx *gin.Context) {

	ctx.JSON(r.StatusCode, r)
}

// 200
func StatusOK(message string, data interface{}) *CommonResponse {

	return &CommonResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: message,
		Result:        data,
	}
}

func Created(message string, dataResponse interface{}) *CommonResponse {
	return &CommonResponse{
		StatusCode:    http.StatusCreated,
		StatusMessage: message,
		Result:        dataResponse,
	}
}

func Accepted(message string, dataResponse interface{}) *CommonResponse {
	return &CommonResponse{
		StatusCode:    http.StatusAccepted,
		StatusMessage: message,
		Result:        dataResponse,
	}
}

func NoContent(message string, dataResponse interface{}) *CommonResponse {
	return &CommonResponse{
		StatusCode:    http.StatusNoContent,
		StatusMessage: message,
		Result:        dataResponse,
	}
}

// 400

func UserNotFound() *CommonResponse {
	return &CommonResponse{
		StatusCode:    http.StatusNotFound,
		StatusMessage: "user not found",
		Result:        nil,
	}
}

func WrongPassword() *CommonResponse {
	return &CommonResponse{
		StatusCode:    http.StatusUnauthorized,
		StatusMessage: "password is invalid",
		Result:        nil,
	}
}

func PasswordNotMatch() *CommonResponse {
	return &CommonResponse{
		StatusCode:    http.StatusUnauthorized,
		StatusMessage: "password not match",
		Result:        nil,
	}
}

func MethodNotAllowed() *CommonResponse {
	return &CommonResponse{
		StatusCode:    http.StatusMethodNotAllowed,
		StatusMessage: "method not allowed",
		Result:        nil,
	}
}

func NoRoute(url string) *CommonResponse {
	return &CommonResponse{
		StatusCode:    http.StatusNotFound,
		StatusMessage: url + " endpoint not available",
		Result:        nil,
	}
}

// 500
func FailedConnectDatabase(databaseName string) *CommonResponse {
	return &CommonResponse{
		StatusCode:    http.StatusInternalServerError,
		StatusMessage: "failed to connect database " + databaseName,
		Result:        nil,
	}
}

func SomethingWentWrong() *CommonResponse {
	return &CommonResponse{
		StatusCode:    http.StatusBadGateway,
		StatusMessage: "something went wrong",
		Result:        nil,
	}
}
