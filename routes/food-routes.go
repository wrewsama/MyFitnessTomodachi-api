package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wrewsama/MyFitnessTomodachi-api/controllers"
)

var RegisterRoutes = func(router *gin.Engine) {
	router.GET("/food", controllers.GetAllFoods)
	router.GET("/food/:id", controllers.GetFoodById)
	router.POST("/food", controllers.CreateFood)
	router.DELETE("/food/:id", controllers.DeleteFood)
	router.PUT("/food/:id", controllers.UpdateFood)
}