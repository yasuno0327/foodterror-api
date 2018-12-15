package main

import (
	"log"
	"net/http"
	"os"
	"sandbox-api/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	db, err := gorm.Open("mysql", dbUser+":"+dbPass+"@tcp(db)/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
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
