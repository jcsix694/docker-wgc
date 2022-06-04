package structures

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wrestler struct {
	ID        uint           `gorm:"primaryKey;" json:"-"`
	UUID      uuid.UUID      `gorm:"not null;unique;size:36;" json:"uuid"`
	Name      string         `gorm:"not null;unique;" json:"name"`
	CreatedAt time.Time      `gorm:"autoCreateTime:true;" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoCreateTime:true;" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index;" json:"-"`
}

func (w *Wrestler) TableName() string {
	return "wrestlers"
}
