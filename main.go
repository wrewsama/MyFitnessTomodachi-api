package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wrewsama/MyFitnessTomodachi-api/initialisers"
	"github.com/wrewsama/MyFitnessTomodachi-api/routes"
)

func init() {
	initialisers.LoadEnvVars()
	initialisers.Connect()
}

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run()
}