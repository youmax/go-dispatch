package controllers

import (
	"log"
	app "upay/app"
	"upay/models"
	r "upay/responses"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	var d models.Deposit
	app.DB().First(&d)
	log.Printf("%s", d)
	c.JSON(r.NewResponse(200, map[string]interface{}{
		"message": d.Custid,
	}))
}
