package models

import (
	"os"
	"wgcapi/logging"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type databaseInfo struct {
	uname string
	psw   string
	name  string
	host  string
}

var dbConn *gorm.DB

func init() {
	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logging.Fatal(err.Error())
	}

	dbConn = db
}

func DB() *gorm.DB {
	return dbConn
}
