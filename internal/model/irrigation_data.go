package model

import "time"

type IrrigationData struct {
	ID                 uint      `gorm:"primaryKey"`
	FarmID             uint      `gorm:"not null;index"`
	IrrigationSectorID uint      `gorm:"not null;index"`
	StartTime          time.Time `gorm:"not null;index"`
	EndTime            time.Time `gorm:"not null"`
	NominalAmount      float32   `gorm:"type:numeric(10,2)"` // in mm
	RealAmount         float32   `gorm:"type:numeric(10,2)"` // in mm
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Farm               Farm             `gorm:"foreignKey:FarmID"`
	IrrigationSector   IrrigationSector `gorm:"foreignKey:IrrigationSectorID"`
}
