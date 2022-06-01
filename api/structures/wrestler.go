package structures

import (
	"gorm.io/gorm"
)

type Wrestler struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	gorm.Model
}

func (w *Wrestler) TableName() string {
	return "wrestlers"
}
