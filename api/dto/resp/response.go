package resp

import (
	"errors"
	"net/http"
)

const (
	DefaultSuccessCode    = "00"
	DefaultSuccessStatus  = "Success"
	DefaultSuccessMessage = "Success"

	DefaultErrorCode    = "X00"
	DefaultErrorStatus  = "Failed"
	DefaultErrorMessage = "something went wrong"
)

type ApiResponse struct {
	StatusCode    string      `json:"statusCode,omitempty"`
	Status        string      `json:"status,omitempty"`
	StatusMessage string      `json:"statusMessage,omitempty"`
	Result        interface{} `json:"result,omitempty"`
}

func NewSuccessMessage(httpCode int, code string, msg string, data interface{}) (httpStatusCode int, apiResponse ApiResponse) {

	if httpCode == 0 {
		httpStatusCode = http.StatusOK
	} else {
		httpStatusCode = httpCode
	}

	if code == "" {
		code = DefaultSuccessCode
	}

	if msg == "" {
		msg = DefaultSuccessMessage
	}

	apiResponse = ApiResponse{
		code,
		DefaultSuccessStatus,
		msg,
		data,
	}

	return
}

func NewFailedMessage(httpCode int, code string, err error) (httpStatusCode int, apiResponse ApiResponse) {

	if httpCode == 0 {
		httpStatusCode = http.StatusInternalServerError
	} else {
		httpStatusCode = httpCode
	}

	if code == "" {
		code = DefaultSuccessCode
	}

	var errHandler *ErrHandler

	if errors.As(err, &errHandler) {
		apiResponse = ApiResponse{
			code,
			DefaultErrorStatus,
			errHandler.ErrorMessage,
			nil,
		}
	} else {
		apiResponse = ApiResponse{
			code,
			DefaultErrorStatus,
			DefaultErrorMessage,
			nil,
		}
	}

	return
}
