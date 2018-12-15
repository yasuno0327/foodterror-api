package controller

import (
	"net/http"
	"sandbox-api/config"
	"sandbox-api/controller/serializer"
	"sandbox-api/model"

	"github.com/gin-gonic/gin"
)

func Matching(c *gin.Context) {
	db := config.GetDB()
	user := CurrentUser(c)
	opponent := &model.User{}
	if err := db.Not(model.User{Email: user.Email}).First(opponent).Error; err != nil {
		c.JSON(http.StatusNotFound, NewError(err))
		return
	}
	c.JSON(http.StatusOK, opponent)
}

func Battle(c *gin.Context) {
	db := config.GetDB()
	user := CurrentUser(c)
	request := serializer.BattleRequest{}
	opponent := &model.User{}
	if err := BindRequest(c, &request); err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	if err := db.Where(&model.User{Email: request.OpponentEmail}).First(opponent).Error; err != nil {
		c.JSON(http.StatusNotFound, NewError(err))
		return
	}
	result := user.Battle(opponent, request.BattleType)
	c.JSON(http.StatusOK, serializer.BattleResult{Result: result})
}
