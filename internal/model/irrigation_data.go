package model

import "time"

type IrrigationData struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement"`
	FarmID             uint      `gorm:"not null;index:idx_irrigation_farm_time,priority:1"`
	IrrigationSectorID uint      `gorm:"not null;index:idx_irrigation_sector_time,priority:1"`
	StartTime          time.Time `gorm:"not null;index:idx_irrigation_farm_time,priority:2;index:idx_irrigation_sector_time,priority:2"`
	EndTime            time.Time `gorm:"not null"`
	NominalAmount      float64   `gorm:"type:numeric(10,2)"`
	RealAmount         float64   `gorm:"type:numeric(10,2)"`
	CreatedAt          time.Time
	UpdatedAt          time.Time

	Farm             Farm             `gorm:"foreignKey:FarmID"`
	IrrigationSector IrrigationSector `gorm:"foreignKey:IrrigationSectorID"`
}
