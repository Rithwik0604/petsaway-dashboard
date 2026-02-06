package main

import (
	"log"
	"petsaway/internal/auth"
	"petsaway/internal/database"
	"petsaway/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	// setup database
	database.SetupDatabase()
	log.Println("Setup Database")

	// setup auth
	auth.SetupAuth()
	log.Println("Setup Auth")

	// setup server
	server.SetupServer()
	log.Println("Setup Server")
}
