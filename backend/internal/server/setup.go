package server

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupServer() {
	port := os.Getenv("PORT")
	debugMode, err := strconv.ParseBool(os.Getenv("DEBUG"))

	if err != nil {
		panic(err)
	}
	if !debugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	Router = gin.Default()
	Router.Use(corsMiddleware())

	setupHandlers()

	Router.Run(port)
}
