package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	c "upay/controllers"
	"upay/models"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

var mDb *gorm.DB
var mDispatcher *Dispatcher
var mRouter *gin.Engine

func init() {
	mDb, err := initdb()
	if err != nil {
		log.Printf("%s", err)
		return
	}
	c.Db = mDb
	defer mDb.Close()
	mDispatcher = NewDispatcher(5)
	mDispatcher.Run()
	mRouter = initRouter()
	mRouter.Run()
}

func initdb() (db *gorm.DB, err error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}
	file := dir + "/db.json"
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var conf models.DbConf
	json.Unmarshal(byteValue, &conf)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	// log.Printf("connect to %s", dsn)
	return gorm.Open("mysql", dsn)
}

func initRouter() (r *gin.Engine) {
	r = gin.Default()
	r.NoRoute(c.ErrorController)
	r.GET("/job/ping", c.Pong)
	return
}
