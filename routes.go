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
	// Motions
	motions := v1.Group("/motions")
	motionRouting(motions)
	// Battles
	battles := v1.Group("/battles")
	battleRouting(battles)
}

// Users
func userRouting(users *gin.RouterGroup) {
	users.POST("/signup", controller.SignUp)
	users.POST("/signin", controller.SignIn)
	users.GET("", controller.UserList)
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

// Motions
func motionRouting(motions *gin.RouterGroup) {
	motions.GET("", controller.AllMotions)
}

// Battles
func battleRouting(battles *gin.RouterGroup) {
	battles.Use(AuthMiddleware())
	battles.POST("/matching", controller.Matching)
	battles.POST("", controller.Battle)
}
