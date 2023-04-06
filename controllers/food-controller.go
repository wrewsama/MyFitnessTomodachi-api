package controllers

import "github.com/gin-gonic/gin"

func GetFood(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "hi",
	})
}