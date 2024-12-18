package mocks

import (
	"ev-pooling-test-api/internal/domain/vehicle"

	"github.com/stretchr/testify/mock"
)

type MockVehicleService struct {
	mock.Mock
}

func (mock *MockVehicleService) GetAllVehicles() []*vehicle.Vehicle {
	args := mock.Called()
	return args.Get(0).([]*vehicle.Vehicle)
}

func (mock *MockVehicleService) AddVehicle(vehicle *vehicle.Vehicle) bool {
	args := mock.Called(vehicle)
	return args.Bool(0)
}

func (mock *MockVehicleService) AddVehicles(vehicles []*vehicle.Vehicle) bool {
	args := mock.Called(vehicles)
	return args.Bool(0)
}

func (mock *MockVehicleService) RemoveVehicleByID(id int) bool {
	args := mock.Called(id)
	return args.Bool(0)
}

func (mock *MockVehicleService) RemoveAll() bool {
	args := mock.Called()
	return args.Bool(0)
}
