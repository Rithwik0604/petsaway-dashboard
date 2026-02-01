package main

import (
	"petsaway/internal/database"
	"petsaway/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	// setup database
	database.SetupDatabase()

	// setup server
	server.SetupServer()
}
