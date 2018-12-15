package main

import (
	"log"
	"sandbox-api/config"
	"sandbox-api/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Setup DB
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Migrate model
	Migrate(db)
	// Setup gin
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	Routing(v1)
	router.Run(":3000")
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Motion{})
	db.AutoMigrate(&model.Coupon{})
	db.AutoMigrate(&model.Food{})
	db.AutoMigrate(&model.DietData{})
}
