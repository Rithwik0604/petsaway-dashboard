package main

import (
	"petsaway/internal/auth"
	"petsaway/internal/database"
	"petsaway/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	// setup database
	database.SetupDatabase()

	// setup auth
	auth.SetupAuth()

	// setup server
	server.SetupServer()
}
