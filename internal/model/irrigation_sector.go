package model

type IrrigationSector struct {
	ID     uint   `gorm:"primaryKey"`
	FarmID uint   `gorm:"not null;index"`
	Name   string `gorm:"not null"`
	Farm   Farm   `gorm:"foreignKey:FarmID"`
}
