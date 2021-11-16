package api

import (
	"fmt"
	"github.com/akhmettolegen/onex/internal/api/handlers"
	"github.com/akhmettolegen/onex/pkg/application"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)


// New - creates new instance of gin.Engine
func New(app application.Application) (*gin.Engine, error) {
	router := gin.Default()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("err", err)
	}

	router.LoadHTMLGlob(fmt.Sprintf("%v/internal/api/templates/*", dir))

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{
			"title": "Main website",
		})
	})

	handler := handlers.Get(app)

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, "pong")
		})

		v1.GET("/welcome", handler.Welcome)

		auth := v1.Group("/auth")
		{
			auth.POST("/sign-up", handler.SignUp)
			auth.POST("/sign-in", handler.SignIn)
		}

		base := v1.Group("")

		users := base.Group("/users")
		{
			users.GET("", handler.GetUsers)
			users.GET("/:id", handler.GetUserByID)
			users.POST("", handler.CreateUser)
			users.DELETE("/:id", handler.DeleteUser)
		}

		file := base.Group("/file")
		{
			file.POST("/upload", handler.Upload)
		}

		order := base.Group("/orders")
		{
			order.GET("", handler.GetOrders)
			order.GET("/:id", handler.GetOrderByID)
			order.POST("", handler.CreateOrder)
			order.PUT("/:id", handler.UpdateOrderByID)
			order.DELETE("/:id", handler.DeleteOrder)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router, nil
}
