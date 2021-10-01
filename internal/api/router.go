package api

import (
	"github.com/akhmettolegen/onex/internal/api/handlers"
	"github.com/akhmettolegen/onex/pkg/application"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New - creates new instance of gin.Engine
func New(app application.Application) (*gin.Engine, error) {
	router := gin.Default()
	handler := handlers.Get(app)

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, "pong")
		})

		v1.GET("/welcome", handler.Welcome)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router, nil
}
