package models

import(
	"gorm.io/gorm"
)

var db *gorm.DB

type Food struct {
	gorm.Model
	Name string `json:"name"`
	Calories int `json:"calories"`
	Protein int `json:"protein"`
	Carbohydrates int `json:"carbohydrates"`
	Fat int `json:"fat"`
}



