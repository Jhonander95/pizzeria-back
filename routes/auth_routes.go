package routes

import (
	"pizzeria-api/controllers"
	"pizzeria-api/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	// Rutas públicas de productos
	products := router.Group("/products")
	{
		products.GET("", controllers.ListProducts) // Ruta pública
	}

	// Rutas de autenticación
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// Rutas protegidas de administración
	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	{
		admin.POST("/products", controllers.AddProduct)
		admin.PUT("/products/:id", controllers.EditProduct)
		admin.DELETE("/products/:id", controllers.DeleteProduct)
	}
}
