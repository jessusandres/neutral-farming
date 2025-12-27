package model

type Farm struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}
