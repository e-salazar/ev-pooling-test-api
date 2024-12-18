package v1

import (
	"ev-pooling-test-api/internal/application/services"
	"ev-pooling-test-api/internal/infrastructure/api/v1/handlers"
	"ev-pooling-test-api/internal/infrastructure/api/v1/routes"
	"ev-pooling-test-api/internal/infrastructure/persistence"

	"github.com/gin-gonic/gin"
)

func SetUp(engine *gin.Engine) {
	vehicleService := services.NewVehicleService(persistence.NewVehicleRepositoryInMemory())
	groupService := services.NewGroupService(persistence.NewGroupRepositoryInMemory())
	journeyService := services.NewJourneyService(persistence.NewJourneyRepositoryInMemory(), vehicleService, groupService)

	vehicleHandler := handlers.NewVehicleHandler(vehicleService, groupService, journeyService)

	routes.RegisterVehicleRoutes(engine, vehicleHandler)
}
