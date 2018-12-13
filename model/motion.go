package model

import "github.com/jinzhu/gorm"

type Motion struct {
	gorm.Model
	Name      string     // Motion name ex) スクワット
	Burned    float32    // Burned calories
	DietDatas []DietData `gorm:"polymophic:Data;"`
}
