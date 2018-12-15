package controller

import (
	"net/http"
	"sandbox-api/config"
	"sandbox-api/controller/serializer"
	"sandbox-api/model"
	"sandbox-api/service"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var db = config.GetDB()
	u := &model.User{}
	if err := BindRequest(c, u); err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	u.EncryptPassword() // Encrypt pass
	if err := db.Create(u).Error; err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, u)
}

func SignIn(c *gin.Context) {
	var db = config.GetDB()
	u := &model.User{}
	if err := BindRequest(c, u); err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
		return
	}
	u.EncryptPassword()
	if err := db.Where(u).First(u).Error; err != nil {
		c.JSON(http.StatusUnauthorized, NewError(err))
		return
	}
	token, err := service.SetSession(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, serializer.SignInResponse{Token: token})
}
