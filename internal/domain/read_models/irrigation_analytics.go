package read_models

type IrrigationAnalytics struct {
	PeriodTag     string
	TotalVolumeMM float64
	TotalEvents   int64
	AvgEfficiency float64
	MinEfficiency float64
	MaxEfficiency float64
}

type TimeSeriesAnalytics struct {
	DatePeriod      string
	NominalAmountMM float64
	RealAmountMM    float64
	Efficiency      float64
	EventCount      int64
}

type BreakdownAnalytics struct {
	SectorID          uint
	SectorName        string
	TotalVolumeMM     float64
	AverageEfficiency float64
}
