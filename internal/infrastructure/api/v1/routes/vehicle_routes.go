package routes

import (
	"ev-pooling-test-api/internal/infrastructure/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterVehicleRoutes(engine *gin.Engine, handler *handlers.VehicleHandler) {
	router := engine.Group("/api/v1")
	{
		router.GET("/status", handler.Status)
		router.PUT("/evs", handler.UpdateVehicles)
		router.POST("/journey", handler.AddGroup)
		router.POST("/dropoff", handler.DropOffGroup)
		router.POST("/locate", handler.LocateGroup)
	}
}
