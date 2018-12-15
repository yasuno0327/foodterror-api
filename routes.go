package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Routing(v1 *gin.RouterGroup) {
	// Users
	users := v1.Group("/users")
	fmt.Println(users)
	//users.POST("/signup", controller.)
}
