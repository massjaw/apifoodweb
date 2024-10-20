package resp

import (
	"fmt"
	"net/http"
)

type ErrHandler struct {
	ErrorMessage string
	ErrorType    int
}

func (e *ErrHandler) Error() string {
	return fmt.Sprintf("error type: %d, error message: %s", e.ErrorType, e.ErrorMessage)
}

func InvalidError(msg string) error {
	if msg == "" {
		return &ErrHandler{
			ErrorMessage: "invalid input",
			ErrorType:    http.StatusBadRequest,
		}
	} else {
		return &ErrHandler{
			ErrorMessage: msg,
			ErrorType:    http.StatusBadRequest,
		}
	}
}

func UnauthorizedError(msg string) error {
	if msg == "" {
		return &ErrHandler{
			ErrorMessage: "unauthorized user",
			ErrorType:    http.StatusUnauthorized,
		}
	} else {
		return &ErrHandler{
			ErrorMessage: msg,
			ErrorType:    http.StatusUnauthorized,
		}
	}
}

func DataNotFoundError(msg string) error {
	if msg == "" {
		return &ErrHandler{
			ErrorMessage: "no data found",
		}
	} else {
		return &ErrHandler{
			ErrorMessage: msg,
		}
	}
}

func UnknownError(msg string) error {
	if msg == "" {
		return &ErrHandler{
			ErrorMessage: "something went wrong",
		}
	} else {
		return &ErrHandler{
			ErrorMessage: msg,
		}
	}
}
