// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"
	"order-app/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/uptrace/bun"

	_ "order-app/docs"
)

// NewRouter -.
// Swagger spec:
// @title       	Order Apps API v1
// @description 	My First API for Order Apps with clean architecture in golang
// @version     	1.0
// @contact.name   	Kevin Christian C.
// @contact.email  	glovenkevincch@gmail.com
// @host        	localhost:8000
// @BasePath    	/api
func NewRouter(handler *gin.Engine, l logger.LoggerInterface, db *bun.DB) {
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/api/v1")
	{
		newAuthRoutes(h, l, db)
		newMenuRoutes(h, l, db)
	}
}
