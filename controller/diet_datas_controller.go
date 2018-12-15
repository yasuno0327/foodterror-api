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
	db := config.GetDB()
	user := CurrentUser(c)
	food := &model.Food{}
	if err := BindRequest(c, food); err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	if err := db.Where(food).First(food).Error; err != nil {
		c.JSON(http.StatusNotFound, NewError(err))
		return
	}
	if err := db.Model(&user).Association("DietDatas").Append(model.DietData{Intake: food.Intake, DataID: food.ID, DataType: "food"}).Error; err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	c.JSON(http.StatusOK, model.DietData{Intake: food.Intake, DataID: food.ID, DataType: "food"})
}

func CreateMotionData(c *gin.Context) {
	db := config.GetDB()
	user := CurrentUser(c)
	motion := &model.Motion{}
	if err := BindRequest(c, motion); err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	if err := db.Where(motion).First(motion).Error; err != nil {
		c.JSON(http.StatusNotFound, NewError(err))
		return
	}
	if err := db.Model(&user).Association("DietDatas").Append(model.DietData{Burned: motion.Burned, DataID: motion.ID, DataType: "motion"}).Error; err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	c.JSON(http.StatusOK, model.DietData{Burned: motion.Burned, DataID: motion.ID, DataType: "motion"})
}
