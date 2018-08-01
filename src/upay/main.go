package main

import (
	"fmt"
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
	d := NewDispatcher(5)
	d.Run()
	var deposit models.Deposit
	db.First(&deposit)
	job := models.Job{Payload: &deposit}
	for i := 0; i < 100000; i++ {
		d.PushJob(job)
	}

	var input string
	fmt.Scanln(&input)
}
