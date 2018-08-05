package controllers

import (
	db "upay/databases"
	m "upay/models"
	r "upay/responses"

	"github.com/gin-gonic/gin"
)

type JobController struct{}

var JobCtrl JobController = JobController{}

func (ctrl *JobController) Pong(c *gin.Context) {
	var d m.Deposit
	db.Mysql().First(&d)
	c.JSON(r.NewResponse(200, map[string]interface{}{
		"message": d.Custid,
	}))
}
