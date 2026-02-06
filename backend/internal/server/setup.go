package server

import (
	"log"
	"os"
	"petsaway/internal/clients"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

var Router *gin.Engine

func SetupServer() {
	port := os.Getenv("PORT")
	debugMode, err := strconv.ParseBool(os.Getenv("DEBUG"))

	// set up emailer cron
	setupEmailingCron()

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

func setupEmailingCron() {
	log.Println("Setting up emailing cron job")
	s := gocron.NewScheduler(time.Local)
	// for prod
	// _, err := s.Every(1).Day().At("00:00").Do(clients.RunExpirationCheck)
	// for dev
	_, err := s.Every(1).Day().At("00:00").Do(clients.RunExpirationCheck)
	clients.SetupEnvVars()
	if err != nil {
		log.Printf("Error running email job: %s\n", err.Error())
	}

	log.Println("Email Scheduler Running")
	s.StartAsync()

}
