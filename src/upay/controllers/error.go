package controllers

import (
	"github.com/gin-gonic/gin"

	r "upay/responses"
)

type ErrorController struct{}

var ErrorCtrl ErrorController = ErrorController{}

// ErrorController handles error messages for wrong routes
func (ctrl *ErrorController) Error(c *gin.Context) {
	c.JSON(r.NewError(r.ErrNotFound))
}
