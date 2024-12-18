package handlers

import (
	"ev-pooling-test-api/internal/application/converters"
	"ev-pooling-test-api/internal/application/dtos"
	"ev-pooling-test-api/internal/domain/group"
	"ev-pooling-test-api/internal/domain/journey"
	"ev-pooling-test-api/internal/domain/vehicle"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VehicleHandler struct {
	vehicleService vehicle.VehicleService
	groupService   group.GroupService
	journeyService journey.JourneyService
}

func NewVehicleHandler(vehicleService vehicle.VehicleService, groupService group.GroupService, journeyService journey.JourneyService) *VehicleHandler {
	return &VehicleHandler{
		vehicleService: vehicleService,
		groupService:   groupService,
		journeyService: journeyService,
	}
}

func (handler *VehicleHandler) Status(context *gin.Context) {
	context.String(http.StatusOK, "")
}

func (handler *VehicleHandler) UpdateVehicles(context *gin.Context) {
	var vehiclesDTO []*dtos.VehicleDTO
	if err := context.ShouldBindJSON(&vehiclesDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicles := converters.ToVehicleEntities(vehiclesDTO)

	if !handler.vehicleService.RemoveAll() {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "deleting vehicles failed"})
		return
	}

	if !handler.groupService.RemoveAll() {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "deleting groups failed"})
		return
	}

	if !handler.journeyService.RemoveAll() {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "deleting journeys failed"})
		return
	}

	if !handler.vehicleService.AddVehicles(vehicles) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "adding vehicles failed"})
		return
	}

	context.String(http.StatusOK, "")
}

func (handler *VehicleHandler) AddGroup(context *gin.Context) {
	var groupDTO *dtos.GroupDTO
	if err := context.ShouldBindJSON(&groupDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group := converters.ToGroupEntity(groupDTO)

	if found := handler.groupService.GetGroupByID(group.ID); found != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "group already exists"})
		return
	}

	if found, _ := handler.journeyService.GetInfoByGroupID(group.ID); found {
		context.JSON(http.StatusBadRequest, gin.H{"error": "group already exists"})
		return
	}
	if !handler.groupService.AddGroup(group) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "adding group failed"})
		return
	}

	handler.journeyService.CreateJourneys()

	context.String(http.StatusOK, "")
}

func (handler *VehicleHandler) DropOffGroup(context *gin.Context) {
	var groupDTO *dtos.GroupDTO
	if err := context.ShouldBindJSON(&groupDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group := converters.ToGroupEntity(groupDTO)

	if !handler.groupService.RemoveGroupByID(group.ID) {
		journey := handler.journeyService.GetJourneyByGroupID(group.ID)
		if journey == nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
			return
		}

		if !handler.vehicleService.AddVehicle(journey.Vehicle) {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "adding vehicle failed"})
			return
		}

		if !handler.journeyService.RemoveJourneyByGroupID(group.ID) {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "removing journey failed"})
			return
		}
	}

	handler.journeyService.CreateJourneys()

	context.String(http.StatusOK, "")
}

func (handler *VehicleHandler) LocateGroup(context *gin.Context) {
	var groupDTO *dtos.GroupDTO
	if err := context.ShouldBindJSON(&groupDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group := converters.ToGroupEntity(groupDTO)

	if found, journey := handler.journeyService.GetInfoByGroupID(group.ID); found {
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
}
