package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wrewsama/MyFitnessTomodachi-api/initialisers"
	"github.com/wrewsama/MyFitnessTomodachi-api/routes"
)

func init() {
	log.Println("Loading environment variables")
	initialisers.LoadEnvVars()
	log.Println("Connecting to Database")
	initialisers.Connect()
}

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run()
	log.Println("App successfully started!")
}