package routers

import (
	"challenge-12/controllers"
	"challenge-12/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.GET("/", controllers.GetAllProduct)
		productRouter.GET("/:productId", controllers.GetProductById)
		productRouter.POST("/", controllers.CreateProduct)

		productRouter.PUT("/:productId", middlewares.CheckUserLevel(), middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE(("/:productId"), middlewares.CheckUserLevel(), middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	return r
}