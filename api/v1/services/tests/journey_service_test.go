package services_test

import (
	"ev-pooling-test-api/api/v1/models/entities"
	"ev-pooling-test-api/api/v1/repositories/tests/mocks"
	"ev-pooling-test-api/api/v1/services"
	services_mocks "ev-pooling-test-api/api/v1/services/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJourneyServiceGetAllJourneys(t *testing.T) {
	mockRepository := new(mocks.MockJourneyRepository)
	mockVehicleService := new(services_mocks.MockVehicleService)
	mockGroupService := new(services_mocks.MockGroupService)
	service := services.NewJourneyService(mockRepository, mockVehicleService, mockGroupService)

	expected := []*entities.Journey{
		{Vehicle: &entities.Vehicle{ID: 1}, Group: &entities.Group{ID: 1}},
		{Vehicle: &entities.Vehicle{ID: 2}, Group: &entities.Group{ID: 2}},
	}

	mockRepository.On("FindAll").Return(expected, nil)

	result := service.GetAllJourneys()
	assert.Equal(t, expected, result)

	mockRepository.AssertExpectations(t)
}

func TestJourneyServiceGetJourneyByGroupID(t *testing.T) {
	mockRepository := new(mocks.MockJourneyRepository)
	mockVehicleService := new(services_mocks.MockVehicleService)
	mockGroupService := new(services_mocks.MockGroupService)
	service := services.NewJourneyService(mockRepository, mockVehicleService, mockGroupService)

	expected := &entities.Journey{Vehicle: &entities.Vehicle{ID: 1}, Group: &entities.Group{ID: 1}}

	mockRepository.On("FindByGroupId", 1).Return(expected, nil)

	result := service.GetJourneyByGroupID(1)
	assert.Equal(t, expected, result)

	mockRepository.AssertExpectations(t)
}

func TestJourneyServiceAddJourney(t *testing.T) {
	mockRepository := new(mocks.MockJourneyRepository)
	mockVehicleService := new(services_mocks.MockVehicleService)
	mockGroupService := new(services_mocks.MockGroupService)
	service := services.NewJourneyService(mockRepository, mockVehicleService, mockGroupService)

	created := &entities.Journey{Vehicle: &entities.Vehicle{ID: 1}, Group: &entities.Group{ID: 1}}

	mockRepository.On("Add", created).Return(nil)

	result := service.AddJourney(created)
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}

func TestJourneyServiceRemoveJourneyByGroupID(t *testing.T) {
	mockRepository := new(mocks.MockJourneyRepository)
	mockVehicleService := new(services_mocks.MockVehicleService)
	mockGroupService := new(services_mocks.MockGroupService)
	service := services.NewJourneyService(mockRepository, mockVehicleService, mockGroupService)

	mockRepository.On("RemoveByGroupId", 1).Return(nil)

	result := service.RemoveJourneyByGroupID(1)
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}

func TestJourneyServiceRemoveAllJourneys(t *testing.T) {
	mockRepository := new(mocks.MockJourneyRepository)
	mockVehicleService := new(services_mocks.MockVehicleService)
	mockGroupService := new(services_mocks.MockGroupService)
	service := services.NewJourneyService(mockRepository, mockVehicleService, mockGroupService)

	mockRepository.On("RemoveAll").Return(nil)

	result := service.RemoveAll()
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}

func TestJourneyServiceGetInfoByGroupID(t *testing.T) {
	mockRepository := new(mocks.MockJourneyRepository)
	mockVehicleService := new(services_mocks.MockVehicleService)
	mockGroupService := new(services_mocks.MockGroupService)
	service := services.NewJourneyService(mockRepository, mockVehicleService, mockGroupService)

	created := &entities.Journey{Vehicle: &entities.Vehicle{ID: 1}, Group: &entities.Group{ID: 1}}
	groups := []*entities.Group{
		{ID: 1, People: 3},
	}

	mockRepository.On("FindAll").Return([]*entities.Journey{created}, nil)
	mockGroupService.On("GetAllGroups").Return(groups)

	found, result := service.GetInfoByGroupID(1)
	assert.True(t, found)
	assert.Equal(t, created, result)

	mockRepository.AssertExpectations(t)
	mockGroupService.AssertExpectations(t)
}
