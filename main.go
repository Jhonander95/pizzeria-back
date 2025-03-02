package main

import (
	"pizzeria-api/config"
	"pizzeria-api/routes"
)

func main() {
	config.ConnectDatabase()

	r := routes.SetupRouter()
	r.Run(":8080")
}
