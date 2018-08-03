package routes

import (
	c "upay/controllers"

	"github.com/gin-gonic/gin"
)

// ApiRoutes
var ApiRoutes *gin.RouterGroup = nil

func SetupApiRoutes(r *gin.Engine) {
	ApiRoutes = r.Group("v1")
	{
		r.GET("/job/ping", c.JobCtrl.Pong)
	}

	r.NoRoute(c.ErrorCtrl.Error)

}
