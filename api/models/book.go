package models

import (
	"fmt"
	"wgcapi/database"
	"wgcapi/logging"
	"wgcapi/structures"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllBook(b *[]structures.Book) ([]structures.Book, error) {
	logging.Info("we are here now")
	var book []structures.Book

	if err := database.DB.Find(&book).Error; err != nil {
		fmt.Println(book)
		fmt.Println(err)
		return book, err
	} else {
		fmt.Println(book)
		fmt.Println(book)
		fmt.Println(book)
		fmt.Println(book)
		fmt.Println(book)
		fmt.Println(book)
		return book, nil
	}
}
