package model

import (
	"github.com/jinzhu/gorm"
)

type Food struct {
	gorm.Model
	Name      string
	Intake    float32    // Intake calories
	ImageURL  string     `json:"image_url"`
	DietDatas []DietData `gorm:"polymophic:Data;"`
}
