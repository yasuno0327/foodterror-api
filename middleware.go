package main

import (
	"net/http"
	"sandbox-api/config"
	"sandbox-api/model"
	"sandbox-api/service"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"error"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		userId, err := service.GetUserId(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, &Error{Message: err.Error()})
			c.Abort()
			return
		}
		if err := setCurrentUser(c, userId); err != nil {
			c.JSON(http.StatusUnauthorized, &Error{Message: err.Error()})
			c.Abort()
		}
	}
}

func setCurrentUser(c *gin.Context, userId int) error {
	var currentUser model.User
	db := config.GetDB()
	err := db.First(&currentUser, userId).Error
	if err != nil {
		return err
	}
	c.Set("current_user", currentUser)
	return nil
}
