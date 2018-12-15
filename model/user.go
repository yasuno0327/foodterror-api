package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name              string `json:"name,omitempty" gorm:"not null" binding:"exists"`
	Email             string `json:"email,omitempty" gorm:"not null" binding:"exists"`
	EncryptedPassword string `json:"password,omitempty" gorm:"not null"`
	Age               int    `json:"age,omitempty" gorm:"not null"`
	Sex               int    `json:"sex,omitempty" gorm:"not null"`
	DietDatas         []DietData
	Coupons           []Coupon `gorm:"many2many:user_coupons;"`
}
