package model

import "github.com/jinzhu/gorm"

type Motion struct {
	gorm.Model
	Name      string     `json:"name"`
	Burned    float32    `json:"burned"`
	DietDatas []DietData `json:"diet_datas,omitempty" gorm:"polymophic:Data;"`
}
