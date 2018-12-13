package main

import (
	"log"
	"net/http"
	"sandbox-api/model"

	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "sandbox:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	Migrate(db)
	http.ListenAndServe(":3000", nil)
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Motion{})
	db.AutoMigrate(&model.Coupon{})
	db.AutoMigrate(&model.Food{})
	db.AutoMigrate(&model.DietData{})
}
