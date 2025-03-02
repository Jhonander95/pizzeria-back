package config

import (
	"log"
	"pizzeria-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=ADMIN dbname=pizzeria port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	DB = database

	// Migrar las tablas
	err = DB.AutoMigrate(&models.Product{}, &models.Order{}, &models.OrderProduct{})
	if err != nil {
		log.Fatal("Error al migrar la base de datos:", err)
	}

	log.Println("Base de datos conectada y migrada exitosamente")
}
