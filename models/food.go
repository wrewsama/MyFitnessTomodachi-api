package models

import(
	"gorm.io/gorm"
)

var db *gorm.DB

type Food struct {
	ID uint `gorm:"primaryKey"` 
	Name string  `gorm:"not null"`
	Unit string `gorm:"not null"`
	Calories uint 
	Protein uint 
	Carbohydrates uint 
	Fat uint 
}



