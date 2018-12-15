package main

import (
	"sandbox-api/controller"

	"github.com/gin-gonic/gin"
)

func Routing(v1 *gin.RouterGroup) {
	// Users
	users := v1.Group("/users")
	userRouting(users)
}

// Users
func userRouting(users *gin.RouterGroup) {
	users.POST("/signup", controller.SignUp)
	users.POST("/signin", controller.SignIn)
	//Auth sample
	users.Use(AuthMiddleware())
	users.GET("/auth/sample", controller.AuthSample)
}
