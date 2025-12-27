package service

import (
	"fmt"
	"log/slog"

	"looker.com/neutral-farming/internal/repository"
)

type IrrigationService struct {
	irrigationRepository repository.IrrigationDataRepository
	logger               *slog.Logger
}

func NewIrrigationService(dataRepository repository.IrrigationDataRepository) *IrrigationService {
	return &IrrigationService{
		irrigationRepository: dataRepository,
		logger:               slog.With("component", "IrrigationService"),
	}
}

func (service *IrrigationService) GetIrrigation(id uint) (map[string]any, error) {
	service.logger.Info(fmt.Sprintf("IrrigationService::GetIrrigation: %d", id))

	instance, err := service.irrigationRepository.FindByID(id)

	if err != nil {
		return nil, err
	}

	return map[string]any{
		"id":         instance.ID,
		"farm_id":    instance.FarmID,
		"sector_id":  instance.IrrigationSectorID,
		"start_time": instance.StartTime,
		"end_time":   instance.EndTime,
	}, err
}
