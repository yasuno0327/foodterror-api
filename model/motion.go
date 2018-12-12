package model

import "github.com/jinzhu/gorm"

type Motion struct {
	gorm.Model
	Name   string  // Motion name ex) スクワット
	Burned float32 // Burned calories
	Users  []User  `gorm:"many2many:user_motions;"`
}
