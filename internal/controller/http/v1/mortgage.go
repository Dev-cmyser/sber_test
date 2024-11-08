package v1

import (
	"context"
	"net/http"

	"github.com/Dev-cmyser/calc_ipoteka/internal/entity"
	"github.com/Dev-cmyser/calc_ipoteka/internal/entity/mortgage"
	"github.com/Dev-cmyser/calc_ipoteka/pkg/logger"
	"github.com/gin-gonic/gin"
)

//go:generate mockery --dir=../../../../pkg/logger --all --output=./mocks --case=underscore

//go:generate mockery --dir=. --all --output=./mocks --case=underscore

type useCase interface {
	Execute(context.Context, mortgage.Request) (entity.Mortgage, error)
	Cache(context.Context) ([]entity.CachedMortgage, error)
}

type mortgageRoutes struct {
	uc useCase
	l  logger.Interface
}

func newMortgageRoutes(handler *gin.RouterGroup, uc useCase, l logger.Interface) {
	routers := &mortgageRoutes{uc, l}

	h := handler.Group("/mortgage")

	h.GET("/cache", routers.cache)
	h.POST("/execute", routers.execute)
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
		checkHTTPErr(c, err, []HTTPSignalError{ErrEmpty})
		return
	}
	c.JSON(http.StatusOK, res)
}

type executeResponse struct {
	Result entity.Mortgage `json:"result"`
}

// Execute performs a mortgage calculation based on input details and selected credit program.
// @Summary     Mortgage Calculation
// @Description Calculates mortgage payments and provides a summary of the payment plan based on the input details and selected credit program.
// @Tags        Mortgage
// @Accept      json
// @Produce     json
// @Param       request body mortgage.Request true "Mortgage calculation request payload"
// @Success     200 {object} executeResponse "Successful mortgage calculation with loan details"
// @Failure     500 {object} error "Internal server error"
// @Router      /mortgage/execute [post].
func (r *mortgageRoutes) execute(c *gin.Context) {
	var req mortgage.Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := r.uc.Execute(c, req)
	if err != nil {
		checkHTTPErr(c, err, []HTTPSignalError{ErrChoosing, ErrLowInitPay, ErrOnlyOneProgram})
		return
	}

	c.JSON(http.StatusOK, executeResponse{res})
}
