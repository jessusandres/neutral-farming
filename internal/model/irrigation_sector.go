package model

import (
	"time"

	"gorm.io/gorm"
)

type IrrigationSector struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	FarmID    uint   `gorm:"not null;index"`
	Name      string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Farm Farm `gorm:"foreignKey:FarmID;constraint:OnDelete:CASCADE;"`
}
