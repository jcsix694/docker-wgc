package database

import (
	"os"
	"wgcapi/logging"
	"wgcapi/structures"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect() *gorm.DB {
	logging.Info("Checking Database Connection")

	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logging.Fatal(err.Error())
	}

	logging.Info("Database Connected!")

	DB = db

	return DB
	// Set connections?
	// Auto Migrate?
}

func Migrate() {
	DB.AutoMigrate(&structures.Book{})
	logging.Info("Database Auto Migration Completed")
}
