package main

import (
	"jwt-auth-go/config"
	"jwt-auth-go/routes"
	"log"
)

func init() {
	config.ConnectDB()
	config.Migrate()
}

func main() {
	router := routes.SetupRoutes()
	log.Fatal(router.Run(":8080"))
}
