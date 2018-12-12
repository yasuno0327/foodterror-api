package model

import "github.com/jinzhu/gorm"

type Coupon struct {
	gorm.Model
	Name   string
	Intake float32 // Intake calories
	Users  []User  `gorm:"many2many:user_coupons;"`
}
