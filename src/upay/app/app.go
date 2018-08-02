package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"upay/dispatch"
	"upay/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var dispatcher *dispatch.Dispatcher
var router *gin.Engine

var dbMux = &sync.Mutex{}
var dispatcherMux = &sync.Mutex{}

func init() {
	err := initdb()
	if err != nil {
		log.Printf("%s", err)
	}
	dispatcher = dispatch.NewDispatcher(5)
	dispatcher.Run()
}

func initdb() (err error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	file := dir + "/db.json"
	jsonFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var conf models.DbConf
	json.Unmarshal(byteValue, &conf)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	// log.Printf("connect to %s", dsn)
	database, err := gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	db = database
	return nil
}

func DB() *gorm.DB {
	dbMux.Lock()
	defer dbMux.Unlock()
	return db
}

func Dispatcher() *dispatch.Dispatcher {
	dispatcherMux.Lock()
	defer dispatcherMux.Unlock()
	return dispatcher
}

func CreateApplication() *gin.Engine {
	router := gin.Default()
	return router
}
