package controllers

import (
	"net/http"
	"pizzeria-api/config"
	"pizzeria-api/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&order)
	c.JSON(http.StatusOK, order)
}
