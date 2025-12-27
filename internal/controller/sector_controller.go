package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"looker.com/neutral-farming/internal/service"
	"looker.com/neutral-farming/internal/types"
	"looker.com/neutral-farming/pkg"
)

type SectorController struct {
	SectorService *service.SectorService
}

func NewSectorController(sectorService *service.SectorService) *SectorController {
	return &SectorController{sectorService}
}

func (controller *SectorController) GetSector(context *gin.Context) {

	id := context.Param("id")

	idInt, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		pkg.AbortWithError(context, types.NewBadRequestError("Wrong sector ID format."))
		return
	}

	data, err := controller.SectorService.GetSector(uint(idInt))

	if abort := pkg.AbortWithError(context, err); abort {
		return
	}

	context.JSON(200, data)
}
