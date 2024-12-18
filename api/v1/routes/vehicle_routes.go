package routes

import (
	"github.com/gin-gonic/gin"

	"ev-pooling-test-api/api/v1/controllers"
)

func RegisterVehicleRoutes(engine *gin.Engine, controller *controllers.VehicleController) {
	router := engine.Group("/api/v1")
	{
		router.GET("/status", controller.Status)
		router.PUT("/evs", controller.UpdateVehicles)
		router.POST("/journey", controller.AddGroup)
		router.POST("/dropoff", controller.DropOffGroup)
		router.POST("/locate", controller.LocateGroup)
	}
}
