package controller

import (
	"sandbox-api/model"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) model.User {
	val, _ := c.Get("current_user")
	return val.(model.User)
}
