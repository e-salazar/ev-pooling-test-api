package v1

import (
	"github.com/gin-gonic/gin"

	"ev-pooling-test-api/api/v1/adapters"
	"ev-pooling-test-api/api/v1/controllers"
	"ev-pooling-test-api/api/v1/routes"
	"ev-pooling-test-api/api/v1/services"
)

func SetUp(engine *gin.Engine) {
	vehicleService := services.NewVehicleService(adapters.NewVehicleRepositoryAdapter())
	groupService := services.NewGroupService(adapters.NewGroupRepositoryAdapter())
	journeyService := services.NewJourneyService(adapters.NewJourneyRepositoryAdapter(), vehicleService, groupService)

	vehicleController := controllers.NewVehicleController(vehicleService, groupService, journeyService)

	routes.RegisterVehicleRoutes(engine, vehicleController)
}
