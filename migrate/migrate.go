package main

import (
	"github.com/wrewsama/MyFitnessTomodachi-api/initialisers"
	"github.com/wrewsama/MyFitnessTomodachi-api/models"
)

func init() {
	initialisers.LoadEnvVars()
	initialisers.Connect()
}

func main() {
	initialisers.DB.AutoMigrate(&models.Food{})
}