package main

import (
	"log"

	"upay/models"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

// db setting
const (
	DBHOST = "tcp(127.0.0.1:3306)"
	DBNAME = "dev"
	DBUSER = "root"
	DBPASS = "123456"
)

func initdb() (db *gorm.DB, err error) {
	dsn := DBUSER + ":" + DBPASS + "@" + DBHOST + "/" + DBNAME + "?charset=utf8"
	return gorm.Open("mysql", dsn)
}

func main() {
	db, err := initdb()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	var deposit models.Deposit
	db.First(&deposit)
	log.Println("Hello world")
}
