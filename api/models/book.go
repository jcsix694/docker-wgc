package models

import (
	"fmt"
	"wgcapi/config"
	"wgcapi/logging"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func (b *Book) TableName() string {
	return "book"
}

func GetAllBook(b *[]Book) ([]Book, error) {
	logging.Info("we are here now")
	var book []Book

	if err := config.DB.Find(&book).Error; err != nil {
		fmt.Println(err)
		return book, err
	} else {
		return book, nil
	}
}
