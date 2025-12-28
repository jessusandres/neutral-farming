package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"looker.com/neutral-farming/internal/http/dto"
	"looker.com/neutral-farming/internal/service"
	"looker.com/neutral-farming/internal/types"
	"looker.com/neutral-farming/pkg"
)

type FarmController struct {
	FarmService *service.FarmService
}

func NewFarmController(farmService *service.FarmService) *FarmController {
	return &FarmController{farmService}
}

func (controller *FarmController) GetFarm(context *gin.Context) {

	id := context.Param("farm_id")

	intID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		pkg.AbortWithError(context, types.NewBadRequestError("Wrong farm ID format."))
		return
	}

	data, err := controller.FarmService.GetFarm(uint(intID))

	if abort := pkg.AbortWithError(context, err); abort {
		return
	}

	context.JSON(http.StatusOK, data)
}

func (controller *FarmController) AnalyticsByFarmAndSectorAndDates(context *gin.Context) {
	var query dto.FarmAnalyticsQuery

	farmID := context.Param("farm_id")

	intID, err := strconv.ParseUint(farmID, 10, 64)

	if err != nil {
		pkg.AbortWithError(context, types.NewBadRequestError("Wrong farm ID format."))
		return
	}

	if err := context.ShouldBindQuery(&query); err != nil {
		pkg.AbortWithError(context, types.NewBadRequestError(err.Error()))
		return
	}

	query.SetDefaults()

	data, err := controller.FarmService.RetrieveAnalytics(uint(intID), *query.SectorID, *query.StartDate, *query.EndDate, query.Aggregation)

	if abort := pkg.AbortWithError(context, err); abort {
		return
	}

	context.JSON(http.StatusOK, data)
}
