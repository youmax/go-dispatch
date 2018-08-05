package database

import (
	"fmt"
	"sync"
	"upay/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var mutex = &sync.Mutex{}
var mysql *gorm.DB = nil

func init() {
	mysql = SetupDBConnection(configs.DbConfig())
}

func SetupDBConnection(config *configs.Database) *gorm.DB {
	if config == nil {
		panic("database config is nil")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", config.Username,
		config.Password, config.Host, config.Port, config.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func Mysql() *gorm.DB {
	mutex.Lock()
	defer mutex.Unlock()
	return mysql
}
