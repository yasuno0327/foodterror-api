package model

import (
	"crypto/sha256"
	"fmt"
	"sandbox-api/config"
	"time"

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

func (u *User) Battle(opponent *User, battle_type string) string {
	db := config.GetDB()
	loc, _ := time.LoadLocation("Asia/Tokyo")
	var start time.Time
	end := time.Now().In(loc)
	switch battle_type {
	case "week":
		start = end.AddDate(0, 0, -7)
	case "month":
		start = end.AddDate(0, -1, 0)
	default:
		start = end.AddDate(0, 0, -1)
	}
	myPower := u.CalculatePower(start, end)
	opponentPower := opponent.CalculatePower(start, end)
	if myPower > opponentPower {
		db.Model(u).Updates(User{VictoryPoint: u.VictoryPoint + 2})
		if opponent.VictoryPoint > 0 {
			db.Model(opponent).Updates(User{VictoryPoint: opponent.VictoryPoint - 1})
		}
		return "Lose"
	} else if myPower < opponentPower {
		db.Model(opponent).Updates(User{VictoryPoint: opponent.VictoryPoint + 2})
		if u.VictoryPoint > 0 {
			db.Model(u).Updates(User{VictoryPoint: u.VictoryPoint - 1})
		}
		return "Win"
	}
	return "Draw"
}

func (u *User) CalculatePower(start, end time.Time) float32 {
	db := config.GetDB()
	diet_datas := []DietData{}
	db.Model(u).Related(diet_datas).Where("created_at BETWEEN ? AND ?", start, end)
	var power float32
	for _, data := range diet_datas {
		power += data.Burned
		power -= data.Intake
	}
	return power
}

func (u *User) EncryptPassword() {
	unsafePass := u.EncryptedPassword
	u.EncryptedPassword = ToHash(unsafePass)
}

func ToHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}
