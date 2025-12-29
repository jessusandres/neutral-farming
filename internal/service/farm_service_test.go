package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"looker.com/neutral-farming/internal/domain"
	"looker.com/neutral-farming/internal/domain/read_models"
)

type farmRepository struct {
	mock.Mock
}

func (repo *farmRepository) FindByID(ID uint) (*domain.Farm, error) {
	args := repo.Called(ID)
	return args.Get(0).(*domain.Farm), args.Error(1)
}

func (repo *farmRepository) YearOverYearAnalytics(farmID uint, _ uint, _, _ time.Time) ([]*read_models.IrrigationAnalytics, error) {
	args := repo.Called(farmID)
	return args.Get(0).([]*read_models.IrrigationAnalytics), args.Error(1)
}

func (repo *farmRepository) TimeSeriesByAggregation(farmID uint, _ uint, _, _ time.Time, _ string) ([]*read_models.TimeSeriesAnalytics, error) {
	args := repo.Called(farmID)
	return args.Get(0).([]*read_models.TimeSeriesAnalytics), args.Error(1)
}

func (repo *farmRepository) SectorBreakdownAnalytics(farmID uint, _ uint, _, _ time.Time) ([]*read_models.BreakdownAnalytics, error) {
	args := repo.Called(farmID)
	return args.Get(0).([]*read_models.BreakdownAnalytics), args.Error(1)
}

func TestFarmService_GetFarm(t *testing.T) {
	mockedRepo := new(farmRepository)
	mockedRepo.On("FindByID", uint(1)).Return(&domain.Farm{ID: 1, Name: "Test"}, nil)

	svc := NewFarmService(mockedRepo)

	result, _ := svc.GetFarm(uint(1))

	expectedResult := map[string]any{"id": uint(1), "name": "Test"}

	mockedRepo.AssertCalled(t, "FindByID", uint(1))
	assert.Equal(t, expectedResult, result)
}

func TestFarmService_RetrieveAnalytics(t *testing.T) {
	mockedRepo := new(farmRepository)
	mockedRepo.On("YearOverYearAnalytics", uint(1)).Return([]*read_models.IrrigationAnalytics{}, nil)
	mockedRepo.On("TimeSeriesByAggregation", uint(1)).Return([]*read_models.TimeSeriesAnalytics{}, nil)
	mockedRepo.On("SectorBreakdownAnalytics", uint(1)).Return([]*read_models.BreakdownAnalytics{}, nil)

	svc := NewFarmService(mockedRepo)

	result, _ := svc.RetrieveAnalytics(uint(1), uint(1), "", "", "")
	fmt.Printf("Result: %v\n", result)

	mockedRepo.AssertNumberOfCalls(t, "YearOverYearAnalytics", 1)
	mockedRepo.AssertNumberOfCalls(t, "TimeSeriesByAggregation", 1)
	mockedRepo.AssertNumberOfCalls(t, "SectorBreakdownAnalytics", 1)

	assert.Equal(t, result.FarmID, uint(1))
}
