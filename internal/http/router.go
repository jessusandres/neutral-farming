package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"looker.com/neutral-farming/internal/config"
	"looker.com/neutral-farming/internal/controller"
)

func SetupRouter(router *gin.Engine, controllers *controller.Controllers) {

	api := router.Group(config.EnvConfig.ApiPrefix)
	{
		api.GET("", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Hello from Gin API!",
			})
		})

		v1 := api.Group("v1")
		{
			// farms
			farmApi := v1.Group("farms")
			{
				farmApi.GET("/:farm_id/irrigation/analytics", controllers.FarmController.AnalyticsByFarmAndSectorAndDates)
				farmApi.GET("/:farm_id", controllers.FarmController.GetFarm)
			}

			// sectors
			sectorApi := v1.Group("sectors")
			{
				sectorApi.GET("/:id", controllers.SectorController.GetSector)
			}

			// irrigation data
			irrigationApi := v1.Group("irrigations")
			{
				irrigationApi.GET("/:id", controllers.IrrigationController.GetIrrigation)
			}
		}

	}
}
