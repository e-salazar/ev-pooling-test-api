package services_test

import (
	"ev-pooling-test-api/internal/application/services"
	"ev-pooling-test-api/internal/domain/vehicle"
	"ev-pooling-test-api/internal/domain/vehicle/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVehicleServiceGetAllVehicles(t *testing.T) {
	mockRepository := new(mocks.MockVehicleRepository)
	service := services.NewVehicleService(mockRepository)

	expected := []*vehicle.Vehicle{
		{ID: 1, Seats: 5},
		{ID: 2, Seats: 5},
	}

	mockRepository.On("FindAll").Return(expected, nil)

	result := service.GetAllVehicles()
	assert.Equal(t, expected, result)

	mockRepository.AssertExpectations(t)
}

func TestVehicleServiceAddVehicle(t *testing.T) {
	mockRepository := new(mocks.MockVehicleRepository)
	service := services.NewVehicleService(mockRepository)

	created := &vehicle.Vehicle{ID: 1, Seats: 5}

	mockRepository.On("Add", created).Return(nil)

	result := service.AddVehicle(created)
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}

func TestVehicleServiceRemoveVehicleByID(t *testing.T) {
	mockRepository := new(mocks.MockVehicleRepository)
	service := services.NewVehicleService(mockRepository)

	mockRepository.On("RemoveByID", 1).Return(nil)

	result := service.RemoveVehicleByID(1)
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}

func TestVehicleServiceRemoveAllVehicles(t *testing.T) {
	mockRepository := new(mocks.MockVehicleRepository)
	service := services.NewVehicleService(mockRepository)

	mockRepository.On("RemoveAll").Return(nil)

	result := service.RemoveAll()
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}
