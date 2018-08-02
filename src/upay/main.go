package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"upay/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

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

func main() {
	db, err := initdb()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	d := NewDispatcher(5)
	d.Run()
	var deposit models.Deposit
	db.First(&deposit)
	job := models.Job{Payload: &deposit}
	for i := 0; i < 10; i++ {
		go d.PushJob(job)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
