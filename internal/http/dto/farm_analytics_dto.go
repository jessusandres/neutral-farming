package dto

import "time"

type FarmAnalyticsDto struct {
	FarmID          uint              `json:"farm_id"`
	Period          Period            `json:"period"`
	Aggregation     string            `json:"aggregation"`
	Metrics         Metrics           `json:"metrics"`
	TimeSeries      []TimeSeries      `json:"time_series"`
	SectorBreakdown []SectorBreakdown `json:"sector_breakdown"`
}

type Metrics struct {
	TotalIrrigationVolumeMm float64          `json:"total_irrigation_volume_mm"`
	TotalIrrigationEvents   int64            `json:"total_irrigation_events"`
	AverageEfficiency       float64          `json:"average_efficiency"`
	EfficiencyRange         EfficiencyRange  `json:"efficiency_range"`
	SamePeriod1             SamePeriod       `json:"same_period_-1"`
	SamePeriod2             SamePeriod       `json:"same_period_-2"`
	PeriodComparison        PeriodComparison `json:"period_comparison"`
}

type EfficiencyRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type PeriodComparison struct {
	VsSamePeriod1 VsSamePeriod `json:"vs_same_period_-1"`
	VsSamePeriod2 VsSamePeriod `json:"vs_same_period_-2"`
}

type VsSamePeriod struct {
	VolumeChangePercent     float64 `json:"volume_change_percent"`
	EventsChangePercent     float64 `json:"events_change_percent"`
	EfficiencyChangePercent float64 `json:"efficiency_change_percent"`
}

type SamePeriod struct {
	TotalIrrigationVolumeMm float64         `json:"total_irrigation_volume_mm"`
	TotalIrrigationEvents   int64           `json:"total_irrigation_events"`
	AverageEfficiency       float64         `json:"average_efficiency"`
	EfficiencyRange         EfficiencyRange `json:"efficiency_range"`
}

type Period struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type SectorBreakdown struct {
	SectorID          uint    `json:"sector_id"`
	SectorName        string  `json:"sector_name"`
	TotalVolumeMm     float64 `json:"total_volume_mm"`
	AverageEfficiency float64 `json:"average_efficiency"`
}

type TimeSeries struct {
	Date            string  `json:"date"`
	NominalAmountMm float64 `json:"nominal_amount_mm"`
	RealAmountMm    float64 `json:"real_amount_mm"`
	Efficiency      float64 `json:"efficiency"`
	EventCount      int64   `json:"event_count"`
}
