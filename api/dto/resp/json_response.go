package resp

import "github.com/gin-gonic/gin"

type HttpResponse interface {
	Send()
}
type JsonResponse struct {
	c              *gin.Context
	httpStatusCode int
	response       ApiResponse
}

func (j *JsonResponse) Send() {

	j.c.JSON(j.httpStatusCode, j.response)
}

func NewSuccessJsonResponse(c *gin.Context, httpCode int, code string, msg string, data interface{}) HttpResponse {

	httpStatusCode, res := NewSuccessMessage(httpCode, code, msg, data)

	return &JsonResponse{
		c,
		httpStatusCode,
		res,
	}
}

func NewErrorJsonResponse(c *gin.Context, httpCode int, code string, err error) HttpResponse {

	httpStatusCode, res := NewFailedMessage(httpCode, code, err)

	return &JsonResponse{
		c,
		httpStatusCode,
		res,
	}
}
