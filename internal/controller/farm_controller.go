package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	
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

	id := context.Param("id")

	intID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		pkg.AbortWithError(context, types.NewBadRequestError("Wrong farm ID format."))
		return
	}

	data, err := controller.FarmService.GetFarm(uint(intID))

	if abort := pkg.AbortWithError(context, err); abort {
		return
	}

	context.JSON(200, data)
}
