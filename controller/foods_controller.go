package controller

import (
	"net/http"
	"sandbox-api/config"
	"sandbox-api/controller/serializer"
	"sandbox-api/model"

	"github.com/gin-gonic/gin"
)

func AllFoods(c *gin.Context) {
	foods := []model.Food{}
	db := config.GetDB()
	if err := db.Find(&foods).Error; err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	c.JSON(http.StatusOK, serializer.Foods{Foods: foods})
}
