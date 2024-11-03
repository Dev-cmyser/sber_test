package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/Dev-cmyser/calc_ipoteka/internal/entity"
	"github.com/Dev-cmyser/calc_ipoteka/internal/entity/mortgage"
	"github.com/Dev-cmyser/calc_ipoteka/pkg/logger"
	"github.com/gin-gonic/gin"
)

type UseCase interface {
	Execute(context.Context, mortgage.Request) (entity.Mortgage, error)
	Cache(context.Context) ([]entity.CachedMortgage, error)
}

type ipotecaRoutes struct {
	uc UseCase
	l  logger.Interface
}

func newMortgageRoutes(handler *gin.RouterGroup, uc UseCase, l logger.Interface) {
	routers := &ipotecaRoutes{uc, l}

	h := handler.Group("/mortgage")
	{
		h.GET("/cache", routers.cache)
		h.POST("/execute", routers.execute)
	}
}

// @Summary     Get all cache
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /mogrtgage/cache [get]
func (r *ipotecaRoutes) cache(c *gin.Context) {

	res := []entity.CachedMortgage{
		{
			ID: 0,
			Params: mortgage.Params{
				ObjectCost:     5000000,
				InitialPayment: 1000000,
				Months:         240,
			},
			Program: mortgage.Program{
				Salary: true,
			},
			Aggregates: mortgage.Aggregates{
				Rate:            8,
				LoanSum:         4000000,
				MonthlyPayment:  33458,
				Overpayment:     4029920,
				LastPaymentDate: time.Date(2044, time.February, 18, 0, 0, 0, 0, time.UTC),
			},
		},
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
// @Failure     500 {object} response
// @Router      /mogrtgage/execute [get]
func (r *ipotecaRoutes) execute(c *gin.Context) {
	c.JSON(http.StatusOK, executeResponse{})
	return
}
