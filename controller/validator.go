package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func BindRequest(c *gin.Context, model interface{}) (err error) {
	b := binding.Default(c.Request.Method, c.ContentType())
	err = c.ShouldBindWith(model, b)
	return
}
