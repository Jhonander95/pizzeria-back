package routes

import (
	"pizzeria-api/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	products := router.Group("/products")
	{
		products.GET("", controllers.ListProducts)
		products.POST("", controllers.AddProduct)
		products.PUT(":id", controllers.EditProduct)
		products.DELETE(":id", controllers.DeleteProduct)
	}
}
