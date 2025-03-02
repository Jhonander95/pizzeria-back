package routes

import (
	"pizzeria-api/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	orders := router.Group("/orders")
	{
		orders.POST("", controllers.CreateOrder)
	}
}
