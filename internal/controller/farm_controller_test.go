package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"looker.com/neutral-farming/internal/http/dto"
)

type farmService struct {
	mock.Mock
}

func (ctr *farmService) GetFarm(id uint) (map[string]any, error) {
	args := ctr.Called(id)
	return args.Get(0).(map[string]any), args.Error(1)
}

func (ctr *farmService) RetrieveAnalytics(farmID uint, _ uint, _ string, _ string, _ string) (*dto.FarmAnalyticsDto, error) {
	args := ctr.Called(farmID)
	return args.Get(0).(*dto.FarmAnalyticsDto), args.Error(1)
}

func TestFarmController_GetFarm(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	mockService := new(farmService)
	mockService.On("GetFarm", uint(1)).Return(map[string]any{"id": 1}, nil)

	ctr := NewFarmController(mockService)

	r.GET("/farms/:farm_id", ctr.GetFarm)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/farms/1", nil)

	r.ServeHTTP(w, req)

	mockService.AssertCalled(t, "GetFarm", uint(1))

	assert.Equal(t, `{"id":1}`, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFarmController_AnalyticsByFarmAndSectorAndDates(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	svc := new(farmService)
	svc.
		On("RetrieveAnalytics", uint(1)).
		Return(
			&dto.FarmAnalyticsDto{
				FarmID: 1,
			},
			nil,
		)

	ctr := NewFarmController(svc)
	r.GET("/farms/:farm_id/analytics", ctr.AnalyticsByFarmAndSectorAndDates)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/farms/1/analytics", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	svc.AssertCalled(t, "RetrieveAnalytics", uint(1))
}
