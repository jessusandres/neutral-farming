package domain

import "time"

type IrrigationData struct {
	ID                 uint
	FarmID             uint
	IrrigationSectorID uint
	StartTime          time.Time
	EndTime            time.Time
	NominalAmount      float64
	RealAmount         float64
}
