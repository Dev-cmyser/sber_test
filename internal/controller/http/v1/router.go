// Package v1 contains HTTP handlers and routes for version 1 of the API.
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
// @BasePath    /.

// NewRouter s.
func NewRouter(handler *gin.Engine, l logger.Interface, uc useCase) {
	handler.Use(LoggerMiddleware())
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	handler.GET("/docs/*any", swaggerHandler)

	docs.SwaggerInfo.BasePath = "/v1"
	h := handler.Group("/v1")

	newMortgageRoutes(h, uc, l)
}
