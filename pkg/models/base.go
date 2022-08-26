package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseIdByUUID struct {
	Id uuid.UUID `gorm:"column:id;type:UUID;primaryKey"`
}

type BaseIdByInt64 struct {
	Id *uint64 `gorm:"column:id;primaryKey;autoIncrement"`
}

type Time struct {
	CreatedAt time.Time `gorm:"column:created_at;not null;autoCreateTime:nano"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;autoUpdateTime:nano"`
}

type SoftDelete struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;"`
}
