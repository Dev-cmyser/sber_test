package v1

import (
	"context"
	"net/http"

	"github.com/Dev-cmyser/calc_ipoteka/internal/entity"
	"github.com/Dev-cmyser/calc_ipoteka/internal/entity/mortgage"
	"github.com/Dev-cmyser/calc_ipoteka/pkg/logger"
	"github.com/gin-gonic/gin"
)

type UseCase interface {
	Execute(context.Context, mortgage.Request) (entity.Mortgage, error)
	Cache(context.Context) ([]entity.CachedMortgage, error)
}

type mortgageRoutes struct {
	uc UseCase
	l  logger.Interface
}

func newMortgageRoutes(handler *gin.RouterGroup, uc UseCase, l logger.Interface) {
	routers := &mortgageRoutes{uc, l}

	h := handler.Group("/mortgage")
	{
		h.GET("/cache", routers.cache)
		h.POST("/execute", routers.execute)
	}
}

// @Summary     Get all cache
// @Accept      json
// @Produce     json
// @Success     200 {object} []entity.CachedMortgage
// @Failure     404 {object} error
// @Router      /mortgage/cache [get].
func (r *mortgageRoutes) cache(c *gin.Context) {
	res, err := r.uc.Cache(c)
	if err != nil {
		checkHttpErr(c, err, []HttpSignalError{ErrEmpty})
		return

	}
	c.JSON(http.StatusOK, res)
	return
}

type executeResponse struct {
	Result entity.Mortgage `json:"result"`
}

// @Summary     Execute credit
// @Accept      json
// @Produce     json
// @Success     200 {object} executeResponse
// @Failure     500 {object} error
// @Router      /mortgage/execute [post].
func (r *mortgageRoutes) execute(c *gin.Context) {
	res, err := r.uc.Execute(c, mortgage.Request{})
	if err != nil {
		checkHttpErr(c, err, []HttpSignalError{ErrChoosing, ErrLowInitPay, ErrOnlyOneProgram})
		return

	}
	c.JSON(http.StatusOK, executeResponse{res})
	return
}
