package services_test

import (
	"ev-pooling-test-api/internal/application/services"
	"ev-pooling-test-api/internal/domain/group"
	group_mocks "ev-pooling-test-api/internal/domain/group/mocks"
	"ev-pooling-test-api/internal/domain/journey"
	journey_mocks "ev-pooling-test-api/internal/domain/journey/mocks"
	"ev-pooling-test-api/internal/domain/vehicle"
	vehicle_mocks "ev-pooling-test-api/internal/domain/vehicle/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJourneyServiceGetAllJourneys(t *testing.T) {
	mockRepository := new(journey_mocks.MockJourneyRepository)
	mockVehicleService := new(vehicle_mocks.MockVehicleService)
	mockGroupService := new(group_mocks.MockGroupService)
	service := services.NewJourneyService(mockRepository, mockVehicleService, mockGroupService)

	expected := []*journey.Journey{
		{Vehicle: &vehicle.Vehicle{ID: 1}, Group: &group.Group{ID: 1}},
		{Vehicle: &vehicle.Vehicle{ID: 2}, Group: &group.Group{ID: 2}},
	}

	mockRepository.On("FindAll").Return(expected, nil)

	result := service.GetAllJourneys()
	assert.Equal(t, expected, result)

	mockRepository.AssertExpectations(t)
}

func TestJourneyServiceGetJourneyByGroupID(t *testing.T) {
	mockRepository := new(journey_mocks.MockJourneyRepository)
	mockVehicleService := new(vehicle_mocks.MockVehicleService)
	mockGroupService := new(group_mocks.MockGroupService)
	service := services.NewJourneyService(mockRepository, mockVehicleService, mockGroupService)

	expected := &journey.Journey{Vehicle: &vehicle.Vehicle{ID: 1}, Group: &group.Group{ID: 1}}

	mockRepository.On("FindByGroupId", 1).Return(expected, nil)

	result := service.GetJourneyByGroupID(1)
	assert.Equal(t, expected, result)

	mockRepository.AssertExpectations(t)
}

func TestJourneyServiceAddJourney(t *testing.T) {
	mockRepository := new(journey_mocks.MockJourneyRepository)
	mockVehicleService := new(vehicle_mocks.MockVehicleService)
	mockGroupService := new(group_mocks.MockGroupService)
	service := services.NewJourneyService(mockRepository, mockVehicleService, mockGroupService)

	created := &journey.Journey{Vehicle: &vehicle.Vehicle{ID: 1}, Group: &group.Group{ID: 1}}

	mockRepository.On("Add", created).Return(nil)

	result := service.AddJourney(created)
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}

func TestJourneyServiceRemoveJourneyByGroupID(t *testing.T) {
	mockRepository := new(journey_mocks.MockJourneyRepository)
	mockVehicleService := new(vehicle_mocks.MockVehicleService)
	mockGroupService := new(group_mocks.MockGroupService)
	service := services.NewJourneyService(mockRepository, mockVehicleService, mockGroupService)

	mockRepository.On("RemoveByGroupId", 1).Return(nil)

	result := service.RemoveJourneyByGroupID(1)
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}

func TestJourneyServiceRemoveAllJourneys(t *testing.T) {
	mockRepository := new(journey_mocks.MockJourneyRepository)
	mockVehicleService := new(vehicle_mocks.MockVehicleService)
	mockGroupService := new(group_mocks.MockGroupService)
	service := services.NewJourneyService(mockRepository, mockVehicleService, mockGroupService)

	mockRepository.On("RemoveAll").Return(nil)

	result := service.RemoveAll()
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}

func TestJourneyServiceGetInfoByGroupID(t *testing.T) {
	mockRepository := new(journey_mocks.MockJourneyRepository)
	mockVehicleService := new(vehicle_mocks.MockVehicleService)
	mockGroupService := new(group_mocks.MockGroupService)
	service := services.NewJourneyService(mockRepository, mockVehicleService, mockGroupService)

	created := &journey.Journey{Vehicle: &vehicle.Vehicle{ID: 1}, Group: &group.Group{ID: 1}}
	groups := []*group.Group{
		{ID: 1, People: 3},
	}

	mockRepository.On("FindAll").Return([]*journey.Journey{created}, nil)
	mockGroupService.On("GetAllGroups").Return(groups)

	found, result := service.GetInfoByGroupID(1)
	assert.True(t, found)
	assert.Equal(t, created, result)

	mockRepository.AssertExpectations(t)
	mockGroupService.AssertExpectations(t)
}
