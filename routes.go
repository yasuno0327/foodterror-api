package main

import (
	"sandbox-api/controller"

	"github.com/gin-gonic/gin"
)

func Routing(v1 *gin.RouterGroup) {
	// Users
	users := v1.Group("/users")
	userRouting(users)
	// DietDatas
	dietDatas := v1.Group("/diet_datas")
	dataRouting(dietDatas)
	// Foods
	foods := v1.Group("/foods")
	foodRouting(foods)
}

// Users
func userRouting(users *gin.RouterGroup) {
	users.POST("/signup", controller.SignUp)
	users.POST("/signin", controller.SignIn)
	//Auth sample
	users.Use(AuthMiddleware())
	users.GET("/auth/sample", controller.AuthSample)
}

// DietDatas
func dataRouting(dietDatas *gin.RouterGroup) {
	dietDatas.Use(AuthMiddleware())
	dietDatas.POST("/foods", controller.CreateFoodData)
	dietDatas.POST("/motions", controller.CreateMotionData)
}

// Foods
func foodRouting(foods *gin.RouterGroup) {
	foods.GET("", controller.AllFoods)
}
