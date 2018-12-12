package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	Age       int
	DietDatas []DietData
	Motions   []Motion `gorm:"many2many:user_motions;"`
	Coupons   []Coupon `gorm:"many2many:user_coupons;"`
}
