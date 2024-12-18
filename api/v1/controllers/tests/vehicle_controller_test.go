package controllers_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ev-pooling-test-api/api/v1/controllers"
	"ev-pooling-test-api/api/v1/models/converters"
	"ev-pooling-test-api/api/v1/models/dtos"
	"ev-pooling-test-api/api/v1/models/entities"
	"ev-pooling-test-api/api/v1/services/tests/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	mockVehicleService := new(mocks.MockVehicleService)
	mockGroupService := new(mocks.MockGroupService)
	mockJourneyService := new(mocks.MockJourneyService)
	controller := controllers.NewVehicleController(mockVehicleService, mockGroupService, mockJourneyService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/status", controller.Status)

	mockJourneyService.On("Status").Return()

	request, _ := http.NewRequest(http.MethodGet, "/status", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	mockJourneyService.AssertExpectations(t)
}

func TestUpdateVehicles(t *testing.T) {
	mockVehicleService := new(mocks.MockVehicleService)
	mockGroupService := new(mocks.MockGroupService)
	mockJourneyService := new(mocks.MockJourneyService)
	controller := controllers.NewVehicleController(mockVehicleService, mockGroupService, mockJourneyService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/vehicles", controller.UpdateVehicles)

	vehiclesDTO := []*dtos.VehicleDTO{
		{ID: 1, Seats: 5},
		{ID: 2, Seats: 3},
	}
	vehicles := converters.ToVehicleEntities(vehiclesDTO)

	mockVehicleService.On("RemoveAll").Return(true)
	mockGroupService.On("RemoveAll").Return(true)
	mockJourneyService.On("RemoveAll").Return(true)
	mockVehicleService.On("AddVehicles", vehicles).Return(true)
	mockJourneyService.On("Status").Return()

	body, _ := json.Marshal(vehiclesDTO)
	request, _ := http.NewRequest(http.MethodPost, "/vehicles", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	mockVehicleService.AssertExpectations(t)
	mockGroupService.AssertExpectations(t)
	mockJourneyService.AssertExpectations(t)
}

func TestAddGroup(t *testing.T) {
	mockVehicleService := new(mocks.MockVehicleService)
	mockGroupService := new(mocks.MockGroupService)
	mockJourneyService := new(mocks.MockJourneyService)
	controller := controllers.NewVehicleController(mockVehicleService, mockGroupService, mockJourneyService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/groups", controller.AddGroup)

	groupDTO := &dtos.GroupDTO{ID: 1, People: 3}
	group := converters.ToGroupEntity(groupDTO)

	mockGroupService.On("AddGroup", group).Return(true)
	mockJourneyService.On("Status").Return()
	mockJourneyService.On("CreateJourneys").Return()

	body, _ := json.Marshal(groupDTO)
	request, _ := http.NewRequest(http.MethodPost, "/groups", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	mockGroupService.AssertExpectations(t)
	mockJourneyService.AssertExpectations(t)
}

func TestDropOffGroup(t *testing.T) {
	mockVehicleService := new(mocks.MockVehicleService)
	mockGroupService := new(mocks.MockGroupService)
	mockJourneyService := new(mocks.MockJourneyService)
	controller := controllers.NewVehicleController(mockVehicleService, mockGroupService, mockJourneyService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/groups/dropoff", controller.DropOffGroup)

	groupDTO := &dtos.GroupDTO{ID: 1, People: 3}
	group := converters.ToGroupEntity(groupDTO)
	journey := &entities.Journey{Vehicle: &entities.Vehicle{ID: 1, Seats: 5}, Group: group}

	mockGroupService.On("RemoveGroupByID", group.ID).Return(false)
	mockJourneyService.On("GetJourneyByGroupID", group.ID).Return(journey)
	mockVehicleService.On("AddVehicle", journey.Vehicle).Return(true)
	mockJourneyService.On("RemoveJourneyByGroupID", group.ID).Return(true)
	mockJourneyService.On("Status").Return()
	mockJourneyService.On("CreateJourneys").Return()

	body, _ := json.Marshal(groupDTO)
	request, _ := http.NewRequest(http.MethodPost, "/groups/dropoff", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	mockGroupService.AssertExpectations(t)
	mockJourneyService.AssertExpectations(t)
	mockVehicleService.AssertExpectations(t)
}

func TestLocateGroup(t *testing.T) {
	mockVehicleService := new(mocks.MockVehicleService)
	mockGroupService := new(mocks.MockGroupService)
	mockJourneyService := new(mocks.MockJourneyService)
	controller := controllers.NewVehicleController(mockVehicleService, mockGroupService, mockJourneyService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/groups/locate", controller.LocateGroup)

	groupDTO := &dtos.GroupDTO{ID: 1, People: 3}
	group := converters.ToGroupEntity(groupDTO)
	journey := &entities.Journey{Vehicle: &entities.Vehicle{ID: 1, Seats: 5}, Group: group}

	mockJourneyService.On("GetInfoByGroupID", group.ID).Return(true, journey)
	mockJourneyService.On("Status").Return()

	body, _ := json.Marshal(groupDTO)
	request, _ := http.NewRequest(http.MethodPost, "/groups/locate", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	mockJourneyService.AssertExpectations(t)
}
