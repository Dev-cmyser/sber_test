package v1

import (
	"github.com/Dev-cmyser/calc_ipoteka/docs"
	"github.com/Dev-cmyser/calc_ipoteka/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Swagger spec:
// @title       Calculate Ipoteca API
// @version     1.0
// @host        localhost:8080
// @BasePath    /
func NewRouter(handler *gin.Engine, l logger.Interface, uc UseCase) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	handler.GET("/docs/*any", swaggerHandler)

	docs.SwaggerInfo.BasePath = "/v1"
	h := handler.Group("/v1")
	{
		newMortgageRoutes(h, uc, l)
	}
}
