package main

import (
	"log"
	"upay/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var deposit models.Deposit
	mDb.First(&deposit)
	log.Println(deposit)
}
