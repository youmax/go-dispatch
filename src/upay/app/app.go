package app

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {

}

func CreateApplication() *gin.Engine {
	router := gin.Default()
	return router
}
