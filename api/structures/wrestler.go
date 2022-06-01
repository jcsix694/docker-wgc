package structures

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Wrestler struct {
	ID        uint           `gorm:"primaryKey;unique;type:uuid;default:uuid_generate_v4()" json:"id"`
	UUID      uuid.UUID      `gorm:"not null;unique;" json:"uuid"`
	Name      string         `gorm:"not null;" json:"name"`
	CreatedAt time.Time      `gorm:"autoCreateTime:true;" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoCreateTime:true;" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index;" json:"deletedAt"`
}

func (w *Wrestler) TableName() string {
	return "wrestlers"
}
