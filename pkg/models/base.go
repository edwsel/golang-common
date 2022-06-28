package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	Id uuid.UUID `gorm:"column:id;type:UUID;default:UUID();primaryKey"`
}

type Time struct {
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(6);not null;autoCreateTime:nano"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(6);not null;autoUpdateTime:nano"`
}

type SoftDelete struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(6)"`
}
