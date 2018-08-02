package controllers

import (
	"github.com/gin-gonic/gin"

	r "upay/responses"
)

// ErrorController handles error messages for wrong routes
func ErrorController(c *gin.Context) {

	c.JSON(r.ErrorMessage(r.ErrNotFound))

	return

}
