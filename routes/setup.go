package routes

import (
	"pizzeria-api/controllers"
	"pizzeria-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.RedirectTrailingSlash = false
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Configuración CORS
	router.Use(middleware.CORSMiddleware(middleware.DevelopmentCORSConfig()))

	// Configurar todas las rutas aquí
	setupRoutes(router)

	return router
}

func setupRoutes(router *gin.Engine) {
	// Rutas públicas
	publicProducts := router.Group("/products")
	{
		publicProducts.GET("", controllers.ListProducts)
	}

	orders := router.Group("/orders")
	{
		orders.POST("", controllers.CreateOrder)
		orders.GET("", controllers.GetOrders)
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
