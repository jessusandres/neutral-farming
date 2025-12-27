package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"looker.com/neutral-farming/internal/service"
	"looker.com/neutral-farming/internal/types"
	"looker.com/neutral-farming/pkg"
)

type IrrigationController struct {
	IrrigationService *service.IrrigationService
}

func NewIrrigationController(irrigationService *service.IrrigationService) *IrrigationController {
	return &IrrigationController{irrigationService}
}

func (controller *IrrigationController) GetIrrigation(context *gin.Context) {

	id := context.Param("id")

	intID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		pkg.AbortWithError(context, types.NewBadRequestError("Wrong irrigation ID format."))
		return
	}

	data, err := controller.IrrigationService.GetIrrigation(uint(intID))

	if abort := pkg.AbortWithError(context, err); abort {
		return
	}

	context.JSON(200, data)
}
