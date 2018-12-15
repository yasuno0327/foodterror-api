package controller

import (
	"net/http"
	"sandbox-api/config"
	"sandbox-api/model"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var db = config.GetDB()
	u := &model.User{}
	if err := BindRequest(c, u); err != nil {
		c.JSON(400, NewError(err))
		return
	}
	u.EncryptPassword() // Encrypt pass
	if err := db.Create(u).Error; err != nil {
		c.JSON(400, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, u)
}
