package gorm

import (
	"gorm.io/gorm"

	"looker.com/neutral-farming/internal/domain"
	"looker.com/neutral-farming/internal/model"
	"looker.com/neutral-farming/internal/repository"
)

type FarmRepo struct {
	db *gorm.DB
}

func NewFarmRepo(db *gorm.DB) repository.FarmRepository {
	return &FarmRepo{db}
}

func (repo *FarmRepo) FindByID(ID uint) (*domain.Farm, error) {

	domainInstance := &model.Farm{}
	result := repo.db.First(domainInstance, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &domain.Farm{
		ID:   domainInstance.ID,
		Name: domainInstance.Name,
	}, nil
}
