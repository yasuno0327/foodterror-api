package model

import "github.com/jinzhu/gorm"

type DietData struct {
	gorm.Model
	UserID   int     `json:"-" gorm:"not null"`
	Burned   float32 `json:"burned,omitempty"`
	Intake   float32 `json:"intake,omitempty"`
	Weight   float32 `json:"weight,omitempty"`
	DataID   int     `json:"-"`
	DataType string  `json:"-"`
}
