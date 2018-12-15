package model

import (
	"crypto/sha256"
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name              string  `json:"name,omitempty" gorm:"not null" binding:"exists"`
	Email             string  `json:"email,omitempty" gorm:"not null" binding:"exists"`
	EncryptedPassword string  `json:"password,omitempty" gorm:"not null"`
	Age               int     `json:"age,omitempty" gorm:"not null"`
	Sex               int     `json:"sex,omitempty" gorm:"not null"`
	Weight            float32 `json:"weight,omitempty" gorm:"not null"`
	Height            float32 `json:"height,omitempty" gorm:"not null"`
	VictoryPoint      uint    `json:"victory_point,omitempty" gorm:"not null"`
	DietDatas         []DietData
	FoodTerror        []FoodTerror
}

func (u *User) EncryptPassword() {
	unsafePass := u.EncryptedPassword
	u.EncryptedPassword = ToHash(unsafePass)
}

func ToHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}
