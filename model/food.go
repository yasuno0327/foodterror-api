package model

import (
	"github.com/jinzhu/gorm"
)

type Food struct {
	gorm.Model
	Name      string     `json"name,omitempty"`
	Intake    float32    `json:"intake,omitempty"` // Intake calories
	ImageURL  string     `json:"image_url,omitempty"`
	DietDatas []DietData `gorm:"polymophic:Data;" json:"diet_datas,omitempty"`
}
