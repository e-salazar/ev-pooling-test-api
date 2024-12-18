package controllers

import (
	"ev-pooling-test-api/api/v1/models/converters"
	"ev-pooling-test-api/api/v1/models/dtos"
	"ev-pooling-test-api/api/v1/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	vehicleService services.VehicleService
	groupService   services.GroupService
	journeyService services.JourneyService
}

func NewVehicleController(vehicleService services.VehicleService, groupService services.GroupService, journeyService services.JourneyService) *VehicleController {
	return &VehicleController{
		vehicleService: vehicleService,
		groupService:   groupService,
		journeyService: journeyService,
	}
}

func (controller *VehicleController) Status(context *gin.Context) {
	context.String(http.StatusOK, "")

	controller.journeyService.Status()
}

func (controller *VehicleController) UpdateVehicles(context *gin.Context) {
	var vehiclesDTO []*dtos.VehicleDTO
	if err := context.ShouldBindJSON(&vehiclesDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicles := converters.ToVehicleEntities(vehiclesDTO)

	if !controller.vehicleService.RemoveAll() {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "deleting vehicles failed"})
		return
	}

	if !controller.groupService.RemoveAll() {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "deleting groups failed"})
		return
	}

	if !controller.journeyService.RemoveAll() {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "deleting journeys failed"})
		return
	}

	if !controller.vehicleService.AddVehicles(vehicles) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "adding vehicles failed"})
		return
	}

	context.String(http.StatusOK, "")

	controller.journeyService.Status()
}

func (controller *VehicleController) AddGroup(context *gin.Context) {
	var groupDTO *dtos.GroupDTO
	if err := context.ShouldBindJSON(&groupDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group := converters.ToGroupEntity(groupDTO)

	if !controller.groupService.AddGroup(group) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "adding group failed"})
		return
	}

	controller.journeyService.Status()

	controller.journeyService.CreateJourneys()

	context.String(http.StatusOK, "")

	controller.journeyService.Status()
}

func (controller *VehicleController) DropOffGroup(context *gin.Context) {
	var groupDTO *dtos.GroupDTO
	if err := context.ShouldBindJSON(&groupDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group := converters.ToGroupEntity(groupDTO)

	if !controller.groupService.RemoveGroupByID(group.ID) {
		journey := controller.journeyService.GetJourneyByGroupID(group.ID)
		if journey == nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
			return
		}

		if !controller.vehicleService.AddVehicle(journey.Vehicle) {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "adding vehicle failed"})
			return
		}

		if !controller.journeyService.RemoveJourneyByGroupID(group.ID) {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "removing journey failed"})
			return
		}
	}

	controller.journeyService.Status()

	controller.journeyService.CreateJourneys()

	context.String(http.StatusOK, "")

	controller.journeyService.Status()
}

func (controller *VehicleController) LocateGroup(context *gin.Context) {
	var groupDTO *dtos.GroupDTO
	if err := context.ShouldBindJSON(&groupDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group := converters.ToGroupEntity(groupDTO)

	if found, journey := controller.journeyService.GetInfoByGroupID(group.ID); found {
		if journey != nil {
			// Found in a journey
			context.JSON(http.StatusOK, converters.ToVehicleDTO(journey.Vehicle))
		} else {
			// Found in waiting list
			context.String(http.StatusNoContent, "")
		}
	} else {
		// Not found
		context.String(http.StatusNotFound, "")
	}

	controller.journeyService.Status()
}
