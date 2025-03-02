package controllers

import (
	"net/http"
	"pizzeria-api/config"
	"pizzeria-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

// Definimos la estructura exacta que esperamos recibir
type OrderInput struct {
	Products []struct {
		Product  models.Product `json:"product"`
		Quantity int            `json:"quantity"`
	} `json:"products"`
	Total float64 `json:"total"`
}

func CreateOrder(c *gin.Context) {
	var input OrderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear la orden
	order := models.Order{
		Total:     input.Total,
		CreatedAt: time.Now(),
	}

	// Primero crear la orden
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Luego crear los productos de la orden
	for _, item := range input.Products {
		orderProduct := models.OrderProduct{
			OrderID:   order.ID,
			ProductID: item.Product.ID,
			Quantity:  item.Quantity,
		}

		if err := config.DB.Create(&orderProduct).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Cargar la orden completa con sus productos
	if err := config.DB.Preload("Products.Product").First(&order, order.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func GetOrders(c *gin.Context) {
	var orders []models.Order
	if err := config.DB.Preload("Products.Product").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
