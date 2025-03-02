package routes

import (
	"pizzeria-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New() // En lugar de gin.Default()
	router.RedirectTrailingSlash = false
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Usa la configuración de desarrollo durante la fase de desarrollo
	router.Use(middleware.CORSMiddleware(middleware.DevelopmentCORSConfig()))

	// Cuando vayas a producción, cambia a:
	// router.Use(middleware.CORSMiddleware(middleware.ProductionCORSConfig([]string{"https://tudominio.com"})))

	// Registrar las rutas de productos
	ProductRoutes(router)

	// Registrar las rutas de pedidos
	OrderRoutes(router)

	return router
}
