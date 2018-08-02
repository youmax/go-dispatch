package routes

import (
	c "upay/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	r.NoRoute(c.ErrorController)
	r.GET("/job/ping", c.Pong)
}
