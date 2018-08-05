package configs

import (
	"os"
	"strings"
	"upay/envs"
)

type Database struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

var db *Database

func init() {
	mode := os.Getenv("DISPATCH_ENV")
	var config map[string]string
	if strings.ToUpper(mode) == "PRODUCTION" {
		config = envs.Production
	} else {
		config = envs.Development
	}
	db = LoadDatabaseConfiguration(config)
}

func LoadDatabaseConfiguration(config map[string]string) *Database {
	database := Database{
		Host:     config["DB_HOST"],
		Port:     config["DB_PORT"],
		Database: config["DB_NAME"],
		Username: config["DB_USERNAME"],
		Password: config["DB_PASSWORD"]}
	return &database
}

func DbConfig() *Database {
	return db
}
