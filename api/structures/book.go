package structures

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func (b *Book) TableName() string {
	return "book"
}
