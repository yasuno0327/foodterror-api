package config

import (
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() (db *gorm.DB, err error) {
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	db, err = gorm.Open("mysql", dbUser+":"+dbPass+"@tcp(db)/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	DB = db
	return
}
func GetDB() *gorm.DB {
	return DB
}
