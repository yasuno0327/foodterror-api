package model

import "github.com/jinzhu/gorm"

type FoodTerror struct {
	gorm.Model
	Name     string
	Intake   float32 // Intake calories
	ImageURL string  `json:"image_url"`
	UserID   uint
}
