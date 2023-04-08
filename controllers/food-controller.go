package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wrewsama/MyFitnessTomodachi-api/initialisers"
	"github.com/wrewsama/MyFitnessTomodachi-api/models"
	"github.com/wrewsama/MyFitnessTomodachi-api/utils"
)

func GetFoodById(c *gin.Context) {
	id := c.Param("id")

	var food models.Food

	initialisers.DB.First(&food, id)

	c.JSON(200, gin.H{
		"food": food,
	})
}

func GetFoods(c *gin.Context) {
	var foods []models.Food
	initialisers.DB.Find(&foods)

	query := c.Query("name")

	if query == "" {
		c.JSON(200, gin.H{
			"foods": foods,
			"temp": "temp",
		})
	} else {
		filteredFoods := utils.GetMatches(query, foods) 
	
		c.JSON(200, gin.H{
			"foods": filteredFoods,
		})
	}

}

func CreateFood(c *gin.Context) {
	var body struct {
		Name string 
		Unit string 
		Calories uint 
		Protein uint 
		Carbohydrates uint 
		Fat uint 
	}

	c.Bind(&body)

	food := models.Food{
		Name: body.Name,
		Unit: body.Unit,
		Calories: body.Calories,
		Protein: body.Protein,
		Carbohydrates: body.Carbohydrates,
		Fat: body.Fat,
	}

	result := initialisers.DB.Create(&food)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"msg": "Food Added!",
	})	
}

func DeleteFood(c *gin.Context) {
	id := c.Param("id")

	initialisers.DB.Delete(&models.Food{}, id)

	c.JSON(200, gin.H{
		"msg": "Food Deleted!",
	})	
}

func UpdateFood(c *gin.Context) {
	id := c.Param("id")

	var food models.Food
	initialisers.DB.First(&food, id)

	var body struct {
		Name string 
		Unit string 
		Calories uint 
		Protein uint 
		Carbohydrates uint 
		Fat uint 
	}

	c.Bind(&body)	

	initialisers.DB.Model(&food).Updates(models.Food{
		Name: body.Name,
		Unit: body.Unit,
		Calories: body.Calories,
		Protein: body.Protein,
		Carbohydrates: body.Carbohydrates,
		Fat: body.Fat,
	})

	c.JSON(200, gin.H{
		"msg": "Food Updated!",
	})
}
