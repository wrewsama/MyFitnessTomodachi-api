package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wrewsama/MyFitnessTomodachi-api/initialisers"
	"github.com/wrewsama/MyFitnessTomodachi-api/models"
)

func GetAllFoods(c *gin.Context) {
	var foods []models.Food

	initialisers.DB.Find(&foods)

	c.JSON(200, gin.H{
		"foods": foods,
	})
}