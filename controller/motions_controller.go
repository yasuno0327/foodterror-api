package controller

import (
	"net/http"
	"sandbox-api/config"
	"sandbox-api/controller/serializer"
	"sandbox-api/model"

	"github.com/gin-gonic/gin"
)

func AllMotions(c *gin.Context) {
	db := config.GetDB()
	motions := []model.Motion{}
	if err := db.Find(&motions).Error; err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	c.JSON(http.StatusOK, serializer.Motions{Motions: motions})
}
