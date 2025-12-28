package repository

import (
	"time"

	"looker.com/neutral-farming/internal/domain"
	"looker.com/neutral-farming/internal/domain/read_models"
)

type FarmRepository interface {
	FindByID(ID uint) (*domain.Farm, error)
	YearOverYearAnalytics(farmID uint, sectorID uint, startDate, endDate time.Time) ([]*read_models.IrrigationAnalytics, error)
	TimeSeriesByAggregation(farmID uint, sectorID uint, startDate, endDate time.Time, aggregation string) ([]*read_models.TimeSeriesAnalytics, error)
	SectorBreakdownAnalytics(farmID uint, sectorID uint, startDate, endDate time.Time) ([]*read_models.BreakdownAnalytics, error)
}
