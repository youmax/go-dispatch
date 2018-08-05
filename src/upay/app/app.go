package app

import (
	"log"
	"upay/databases"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func CreateApplication() *gin.Engine {
	router := gin.Default()
	return router
}

func Exception() {
	if err := recover(); err != nil {
		log.Fatal(err)
	}
}

func Close() {
	if database.Mysql() != nil {
		defer database.Mysql().Close()
	}
}
