package app

import (
	"upay/configs"
	database "upay/databases"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func CreateApplication() *gin.Engine {

	configs.Init()
	database.Init(configs.DB())
	router := gin.Default()
	return router
}
