package controllers

import (
	"upay/models"
	r "upay/responses"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	var deposit models.Deposit
	Db.First(&deposit)
	c.JSON(r.NewResponse(200, map[string]interface{}{
		"message": "hello world",
	}))
	return
}
