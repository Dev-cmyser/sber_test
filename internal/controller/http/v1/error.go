package v1

import (
	"errors"
	"net/http"

	"github.com/Dev-cmyser/calc_ipoteka/internal/usecase"
	"github.com/gin-gonic/gin"
)

// HTTPSignalError s.
type HTTPSignalError interface {
	error
	Status() int
	unWrap() error
}

type errorResponse struct {
	Message string `json:"message,omitempty"`
}

// HTTPError s.
type HTTPError struct {
	error
	status int
}

func newHTTPError(err error, status int) *HTTPError {
	return &HTTPError{
		error:  err,
		status: status,
	}
}

// Error s.
func (e *HTTPError) Error() string {
	return e.error.Error()
}

// Status s.
func (e *HTTPError) Status() int {
	return e.status
}

func (e *HTTPError) unWrap() error {
	return e.error
}

func checkHTTPErr(c *gin.Context, err error, signalErrors []HTTPSignalError) {
	for _, sigerr := range signalErrors {
		if errors.Is(err, sigerr.unWrap()) {
			c.AbortWithStatusJSON(sigerr.Status(), errorResponse{sigerr.Error()})
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
}

// usecase.
var (
	ErrEmpty          = newHTTPError(usecase.ErrEmpty, http.StatusNotFound)
	ErrChoosing       = newHTTPError(usecase.ErrChoosing, http.StatusBadRequest)
	ErrOnlyOneProgram = newHTTPError(usecase.ErrOnlyOneProgram, http.StatusBadRequest)
	ErrLowInitPay     = newHTTPError(usecase.ErrLowInitPay, http.StatusBadRequest)
)
