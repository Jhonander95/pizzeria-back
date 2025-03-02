package main

import (
	"pizzeria-api/config"
	"pizzeria-api/routes"
)

func main() {
	config.ConnectDatabase() // Conectar a la base de datos
	r := routes.SetupRouter()
	r.Run(":8080") // Iniciar el servidor en el puerto 8080
}

/*package main

 import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Products []Product `json:"products" gorm:"many2many:pedido_productos;"`
	Total    float64   `json:"total"`
}

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=ADMIN dbname=pizzeria port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar con la base de datos")
	}
	db = database
	db.AutoMigrate(&Product{}, &Order{})

	r := gin.Default()

	// Aplicar middleware CORS a todas las rutas
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Manejo de OPTIONS (preflight)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	r.POST("/order", createOrder)
	r.GET("/products", listProducts)
	r.POST("/products", addProduct)

	r.Run(":8080")
}

func createOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&order)
	c.JSON(http.StatusOK, order)
}

func listProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func addProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&product)
	c.JSON(http.StatusOK, product)
}
*/
