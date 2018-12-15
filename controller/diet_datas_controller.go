package controller

import (
	"net/http"
	"sandbox-api/config"
	"sandbox-api/model"

	"github.com/gin-gonic/gin"
)

func CreateData(c *gin.Context) {
	db := config.GetDB()
	user := CurrentUser(c)
	dietData := &model.DietData{}
	if err := BindRequest(c, dietData); err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	if err := db.Model(user).Association("DietDatas").Append(dietData).Error; err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	c.JSON(http.StatusOK, dietData)
}

func CreateFoodData(c *gin.Context) {
}

func CreateMotionData(c *gin.Context) {
}
