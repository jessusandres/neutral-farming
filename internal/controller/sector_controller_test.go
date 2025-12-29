package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"looker.com/neutral-farming/internal/http/middlewares"
	"looker.com/neutral-farming/internal/types"
)

type mockSectorService struct {
	mock.Mock
}

func (ctr *mockSectorService) GetSector(id uint) (map[string]any, error) {
	args := ctr.Called(id)

	return args.Get(0).(map[string]any), args.Error(1)
}

func TestSectorController_GetSector(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	sectorServiceMock := new(mockSectorService)

	serviceResult := map[string]any{
		"id": 1,
	}
	sectorServiceMock.On("GetSector", uint(1)).Return(serviceResult, nil)

	ctrl := NewSectorController(sectorServiceMock)

	r.GET("/sectors/:id", ctrl.GetSector)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sectors/1", nil)

	r.ServeHTTP(w, req)

	stringifyBytes, _ := json.Marshal(serviceResult)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(stringifyBytes), w.Body.String())

	sectorServiceMock.AssertCalled(t, "GetSector", uint(1))
}

func TestSectorController_GetSector_NotID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.Use(middlewares.HandleErr())
	sectorServiceMock := new(mockSectorService)
	serviceResult := map[string]any{"error": map[string]any{"message": "Wrong sector ID format."}}

	ctrl := NewSectorController(sectorServiceMock)

	r.GET("/sectors/:id", ctrl.GetSector)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sectors/not-an-id", nil)

	r.ServeHTTP(w, req)

	serviceResultBytes, _ := json.Marshal(serviceResult)

	assert.Equal(t, string(serviceResultBytes), w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSectorController_GetSector_WithError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.Use(middlewares.HandleErr())
	sectorServiceMock := new(mockSectorService)
	serviceResult := map[string]any{}

	sectorServiceMock.On("GetSector", uint(1)).Return(serviceResult, types.NewBadRequestError("A test error"))

	ctrl := NewSectorController(sectorServiceMock)

	r.GET("/sectors/:id", ctrl.GetSector)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sectors/1", nil)

	r.ServeHTTP(w, req)

	expectedResponse := map[string]any{"error": map[string]any{"message": "A test error"}}
	serviceResultBytes, _ := json.Marshal(expectedResponse)

	assert.Equal(t, string(serviceResultBytes), w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
