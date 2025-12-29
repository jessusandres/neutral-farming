package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"looker.com/neutral-farming/internal/domain"
)

type sectorRepo struct {
	mock.Mock
}

func (repo *sectorRepo) FindByID(ID uint) (*domain.IrrigationSector, error) {
	args := repo.Called(ID)
	return args.Get(0).(*domain.IrrigationSector), args.Error(1)
}

func TestSectorService_GetSector(t *testing.T) {
	mockedRepo := new(sectorRepo)
	mockedRepo.
		On("FindByID", uint(1)).
		Return(
			&domain.IrrigationSector{
				ID: 1, Name: "Test", FarmID: 2,
			},
			nil,
		)

	svc := NewSectorService(mockedRepo)

	result, _ := svc.GetSector(uint(1))

	expectedResult := map[string]any{
		"id":      uint(1),
		"farm_id": uint(2),
		"name":    "Test",
	}

	mockedRepo.AssertCalled(t, "FindByID", uint(1))

	assert.Equal(t, expectedResult, result)
}
