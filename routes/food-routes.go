package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wrewsama/MyFitnessTomodachi-api/controllers"
)

var RegisterRoutes = func(router *gin.Engine) {
	router.GET("/food", controllers.GetFood)
}