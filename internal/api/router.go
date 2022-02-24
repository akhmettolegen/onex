package api

import (
	"github.com/akhmettolegen/texert/internal/api/handlers"
	"github.com/akhmettolegen/texert/pkg/application"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New - creates new instance of gin.Engine
func New(app application.Application) (*gin.Engine, error) {
	router := gin.Default()
	handler := handlers.Get(app)

	router.Use(handler.CORSMiddleware)
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

		base := v1.Group("", handler.CheckChannelToken, handler.FetchMobileUserInfo)

		users := base.Group("/users")
		{
			users.GET("", handler.GetUsers)
			users.GET("/:id", handler.GetUserByID)
			users.POST("", handler.CreateUser)
			users.DELETE("/:id", handler.DeleteUser)
		}

		file := base.Group("/files")
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

		product := base.Group("/products")
		{
			product.GET("", handler.GetProducts)
			product.GET("/:id", handler.GetProductByID)
			product.POST("", handler.CreateProduct)
			product.PUT("/:id", handler.UpdateProductByID)
			product.DELETE("/:id", handler.DeleteProduct)
		}

		category := base.Group("/categories")
		{
			category.GET("", handler.GetCategories)
			category.GET("/:id", handler.GetCategoryByID)
			category.POST("", handler.CreateCategory)
			category.PUT("/:id", handler.UpdateCategoryByID)
			category.DELETE("/:id", handler.DeleteCategory)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router, nil
}
