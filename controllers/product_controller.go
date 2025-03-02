package controllers

import (
	"net/http"
	"pizzeria-api/config"
	"pizzeria-api/models"

	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context) {
	var productos []models.Product
	config.DB.Find(&productos)
	c.JSON(http.StatusOK, productos)
}

func AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&product)
	c.JSON(http.StatusOK, product)
}

func EditProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	config.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado correctamente"})
}
