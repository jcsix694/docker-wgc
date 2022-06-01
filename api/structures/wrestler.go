package structures

import (
	"gorm.io/gorm"
	"time"
)

type Wrestler struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `gorm:"autoCreateTime:true" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoCreateTime:true" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (w *Wrestler) TableName() string {
	return "wrestlers"
}
