package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"upay/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var mutex = &sync.Mutex{}
var mysql *gorm.DB = nil

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return
	}
	file := dir + "/db.json"
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var conf models.DbConf
	json.Unmarshal(byteValue, &conf)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	// log.Printf("connect to %s", dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return
	}
	mysql = db
}

func Mysql() *gorm.DB {
	mutex.Lock()
	defer mutex.Unlock()
	return mysql
}
