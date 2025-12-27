package service

import (
	"fmt"
	"log/slog"

	"looker.com/neutral-farming/internal/repository"
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

	fmt.Printf("Value: %v\n", farmInstance)

	if err != nil {
		return nil, err
	}

	return map[string]any{
		"id":   farmInstance.ID,
		"name": farmInstance.Name,
	}, nil
}
