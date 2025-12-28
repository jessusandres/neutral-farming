package service

import (
	"fmt"
	"log/slog"
	"time"

	"looker.com/neutral-farming/internal/domain/read_models"
	"looker.com/neutral-farming/internal/http/dto"
	"looker.com/neutral-farming/internal/repository"
	"looker.com/neutral-farming/pkg"
)

type FarmService struct {
	farmRepository repository.FarmRepository
	logger         *slog.Logger
}

func NewFarmService(farmRepo repository.FarmRepository) *FarmService {
	return &FarmService{
		farmRepository: farmRepo,
		logger:         slog.With("component", "FarmService"),
	}
}

func (service *FarmService) GetFarm(id uint) (map[string]any, error) {
	service.logger.Info(fmt.Sprintf("FarmService::GetFarm: %d", id))

	farmInstance, err := service.farmRepository.FindByID(id)

	if err != nil {
		return nil, err
	}

	return map[string]any{
		"id":   farmInstance.ID,
		"name": farmInstance.Name,
	}, nil
}

func (service *FarmService) RetrieveAnalytics(
	farmID uint,
	sectorID uint,
	startDateString string,
	endDateString string,
	aggregation string,
) (*dto.FarmAnalyticsDto, error) {

	service.logger.With("method", "RetrieveAnalytics").Info(
		fmt.Sprintf(
			"Using params farm_id %d - sector_id %d - start_data %s - end_date %s - aggregation %s",
			farmID, sectorID, startDateString, endDateString, aggregation,
		),
	)

	startDate, err := time.Parse("2006-01-02", startDateString)
	endDate, err := time.Parse("2006-01-02", endDateString)

	normalizedStartDate := pkg.StartOfDay(startDate)
	normalizedEndDate := pkg.EndOfDay(endDate)

	yearAnalytics, err := service.farmRepository.YearOverYearAnalytics(farmID, sectorID, normalizedStartDate, normalizedEndDate)

	if err != nil {
		return nil, err
	}

	aggregation = service.mapAggregation(aggregation)

	timeSeriesAnalytics, err := service.farmRepository.TimeSeriesByAggregation(farmID, sectorID, normalizedStartDate, normalizedEndDate, aggregation)

	if err != nil {
		return nil, err
	}

	sectorBreakdownResult, err := service.farmRepository.SectorBreakdownAnalytics(farmID, sectorID, normalizedStartDate, normalizedEndDate)

	if err != nil {
		return nil, err
	}

	periods := make([]*read_models.IrrigationAnalytics, 3)

	for ind, period := range yearAnalytics {
		periods[ind] = period
	}

	periodResume := dto.Period{Start: normalizedStartDate, End: normalizedEndDate}

	initial := periods[0]
	previousPeriod := periods[1]
	minorThreePeriod := periods[2]

	var metrics dto.Metrics
	var previousMetrics dto.SamePeriod
	var minorThreeMetrics dto.SamePeriod

	if initial != nil {
		metrics = dto.Metrics{
			TotalIrrigationVolumeMm: pkg.RoundToDecimals(initial.TotalVolumeMM, 2),
			TotalIrrigationEvents:   initial.TotalEvents,
			AverageEfficiency:       pkg.RoundToDecimals(initial.AvgEfficiency, 2),
			EfficiencyRange: dto.EfficiencyRange{
				Min: pkg.RoundToDecimals(initial.MinEfficiency, 2),
				Max: pkg.RoundToDecimals(initial.MaxEfficiency, 2),
			},
		}
	}

	if previousPeriod != nil {
		previousMetrics = dto.SamePeriod{
			TotalIrrigationVolumeMm: pkg.RoundToDecimals(previousPeriod.TotalVolumeMM, 2),
			TotalIrrigationEvents:   previousPeriod.TotalEvents,
			AverageEfficiency:       pkg.RoundToDecimals(previousPeriod.AvgEfficiency, 2),
			EfficiencyRange: dto.EfficiencyRange{
				Min: pkg.RoundToDecimals(previousPeriod.MinEfficiency, 2),
				Max: pkg.RoundToDecimals(previousPeriod.MaxEfficiency, 2),
			},
		}
	}

	if minorThreePeriod != nil {
		minorThreeMetrics = dto.SamePeriod{
			TotalIrrigationVolumeMm: pkg.RoundToDecimals(minorThreePeriod.TotalVolumeMM, 2),
			TotalIrrigationEvents:   minorThreePeriod.TotalEvents,
			AverageEfficiency:       pkg.RoundToDecimals(minorThreePeriod.AvgEfficiency, 2),
			EfficiencyRange: dto.EfficiencyRange{
				Min: pkg.RoundToDecimals(minorThreePeriod.MinEfficiency, 2),
				Max: pkg.RoundToDecimals(minorThreePeriod.MaxEfficiency, 2),
			},
		}
	}

	var comparativeMinorOne dto.VsSamePeriod
	var comparativeMinorTwo dto.VsSamePeriod

	if initial != nil && previousPeriod != nil {
		comparativeMinorOne = dto.VsSamePeriod{
			VolumeChangePercent:     pkg.AveragePercentForTwoValues(initial.TotalVolumeMM, previousPeriod.TotalVolumeMM),
			EventsChangePercent:     pkg.AveragePercentForTwoValues(initial.TotalEvents, previousPeriod.TotalEvents),
			EfficiencyChangePercent: pkg.AveragePercentForTwoValues(initial.AvgEfficiency, previousPeriod.AvgEfficiency),
		}
	}

	if initial != nil && minorThreePeriod != nil {
		comparativeMinorTwo = dto.VsSamePeriod{
			VolumeChangePercent:     pkg.AveragePercentForTwoValues(initial.TotalVolumeMM, minorThreePeriod.TotalVolumeMM),
			EventsChangePercent:     pkg.AveragePercentForTwoValues(initial.TotalEvents, minorThreePeriod.TotalEvents),
			EfficiencyChangePercent: pkg.AveragePercentForTwoValues(initial.AvgEfficiency, minorThreePeriod.AvgEfficiency),
		}
	}

	periodComparison := dto.PeriodComparison{
		VsSamePeriod1: comparativeMinorOne,
		VsSamePeriod2: comparativeMinorTwo,
	}

	metrics.SamePeriod1 = previousMetrics
	metrics.SamePeriod2 = minorThreeMetrics
	metrics.PeriodComparison = periodComparison

	// Time series analytics

	timeSeries := make([]dto.TimeSeries, len(timeSeriesAnalytics))

	for ind, resume := range timeSeriesAnalytics {
		timeSeries[ind] = dto.TimeSeries{
			Date:            resume.DatePeriod,
			NominalAmountMm: pkg.RoundToDecimals(resume.NominalAmountMM, 2),
			RealAmountMm:    pkg.RoundToDecimals(resume.RealAmountMM, 2),
			Efficiency:      pkg.RoundToDecimals(resume.Efficiency, 2),
			EventCount:      resume.EventCount,
		}
	}

	// Breakdown analytics

	sectorBreakdown := make([]dto.SectorBreakdown, len(sectorBreakdownResult))

	for ind, resume := range sectorBreakdownResult {
		sectorBreakdown[ind] = dto.SectorBreakdown{
			SectorID:          resume.SectorID,
			SectorName:        resume.SectorName,
			TotalVolumeMm:     pkg.RoundToDecimals(resume.TotalVolumeMM, 2),
			AverageEfficiency: pkg.RoundToDecimals(resume.AverageEfficiency, 2),
		}
	}

	return &dto.FarmAnalyticsDto{
		FarmID:          farmID,
		Period:          periodResume,
		Aggregation:     aggregation,
		Metrics:         metrics,
		TimeSeries:      timeSeries,
		SectorBreakdown: sectorBreakdown,
	}, nil
}

func (service *FarmService) mapAggregation(aggregation string) string {
	switch aggregation {
	case "daily":
		return "day"
	case "monthly":
		return "month"
	case "weekly":
		return "week"
	default:
		return "day"
	}
}
