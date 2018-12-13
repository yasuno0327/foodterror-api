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
	Gender    int
	DietDatas []DietData
	Coupons   []Coupon `gorm:"many2many:user_coupons;"`
}
