package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type integrationServiceMock struct {
	mock.Mock
}

func (s *integrationServiceMock) GetIrrigation(ID uint) (map[string]any, error) {
	args := s.Called(ID)

	return args.Get(0).(map[string]any), args.Error(1)
}

func TestIrrigationController_GetIrrigation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	serviceMock := new(integrationServiceMock)

	serviceResult := map[string]any{"id": 1}
	serviceMock.On("GetIrrigation", uint(1)).Return(serviceResult, nil)

	ctr := NewIrrigationController(serviceMock)

	r.GET("/irrigations/:id", ctr.GetIrrigation)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/irrigations/1", nil)

	r.ServeHTTP(w, req)

	serviceMock.AssertCalled(t, "GetIrrigation", uint(1))

	assert.Equal(t, `{"id":1}`, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}
