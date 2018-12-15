package main

import (
	"sandbox-api/controller"

	"github.com/gin-gonic/gin"
)

func Routing(v1 *gin.RouterGroup) {
	// Users
	users := v1.Group("/users")
	users.POST("/signup", controller.SignUp)
	users.POST("/signin", controller.SignIn)
}
