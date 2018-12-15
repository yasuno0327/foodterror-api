package controller

import (
	"sandbox-api/model"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	u := &model.User{}
	if err := BindRequest(c, u); err != nil {
		c.JSON(400, NewError(err))
		return
	}

}
