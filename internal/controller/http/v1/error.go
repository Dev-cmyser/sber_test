package v1

import (
	"errors"
	"net/http"

	"github.com/Dev-cmyser/calc_ipoteka/internal/usecase"
	"github.com/gin-gonic/gin"
)

type response struct {
	Error string `json:"error" example:"message"`
}

type HttpSignalError interface {
	error
	Status() int
	UnWrap() error
}

type errorResponse struct {
	Message string `json:"message,omitempty"`
}

type HttpError struct {
	error
	status int
}

func NewHttpError(error error, status int) *HttpError {
	return &HttpError{
		error:  error,
		status: status,
	}
}

func (e *HttpError) Error() string {
	return e.error.Error()
}

func (e *HttpError) Status() int {
	return e.status
}

func (e *HttpError) UnWrap() error {
	return e.error
}

func checkHttpErr(c *gin.Context, err error, signalErrors []HttpSignalError) {
	for _, sigerr := range signalErrors {
		if errors.Is(err, sigerr.UnWrap()) {
			c.AbortWithStatusJSON(sigerr.Status(), errorResponse{sigerr.Error()})
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
}

var (
	ErrEmpty          = NewHttpError(usecase.ErrEmpty, http.StatusNotFound)
	ErrChoosing       = NewHttpError(usecase.ErrChoosing, http.StatusBadRequest)
	ErrOnlyOneProgram = NewHttpError(usecase.ErrOnlyOneProgram, http.StatusBadRequest)
	ErrLowInitPay     = NewHttpError(usecase.ErrLowInitPay, http.StatusBadRequest)
)
