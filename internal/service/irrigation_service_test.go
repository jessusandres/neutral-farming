package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"looker.com/neutral-farming/internal/domain"
)

type irrigationRepo struct {
	mock.Mock
}

func (repo *irrigationRepo) FindByID(irrigationID uint) (*domain.IrrigationData, error) {
	args := repo.Called(irrigationID)
	return args.Get(0).(*domain.IrrigationData), args.Error(1)
}

func TestIrrigationService_GetIrrigation(t *testing.T) {
	now := time.Now()

	mockedRepo := new(irrigationRepo)
	mockedRepo.On("FindByID", uint(1)).Return(&domain.IrrigationData{
		ID:                 1,
		FarmID:             2,
		IrrigationSectorID: 3,
		StartTime:          now,
		EndTime:            now,
	}, nil)

	svc := NewIrrigationService(mockedRepo)
	result, _ := svc.GetIrrigation(uint(1))

	expectedResult := map[string]any{
		"id":         uint(1),
		"farm_id":    uint(2),
		"sector_id":  uint(3),
		"start_time": now,
		"end_time":   now,
	}

	mockedRepo.AssertCalled(t, "FindByID", uint(1))

	assert.Equal(t, expectedResult, result)
}
