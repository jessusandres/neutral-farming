package service

import (
	"fmt"
	"log/slog"

	"looker.com/neutral-farming/internal/repository"
)

type ISectorService interface {
	GetSector(id uint) (map[string]any, error)
}

type SectorService struct {
	sectorRepository repository.IrrigationSectorRepository
	logger           *slog.Logger
}

func NewSectorService(repo repository.IrrigationSectorRepository) *SectorService {
	return &SectorService{
		sectorRepository: repo,
		logger:           slog.With("component", "SectorService"),
	}
}

func (service *SectorService) GetSector(id uint) (map[string]any, error) {
	service.logger.Info(fmt.Sprintf("FarmService::GetFarm: %d", id))

	entity, err := service.sectorRepository.FindByID(id)

	if err != nil {
		return nil, err
	}

	return map[string]any{
		"id":      entity.ID,
		"farm_id": entity.FarmID,
		"name":    entity.Name,
	}, nil
}
