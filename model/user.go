package model

import (
	"crypto/sha256"
	"fmt"

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
	FoodTerror        []FoodTerror `gorm:"many2many:user_coupons;"`
}

func (u *User) EncryptPassword() {
	unsafePass := u.EncryptedPassword
	u.EncryptedPassword = ToHash(unsafePass)
}

func ToHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}
