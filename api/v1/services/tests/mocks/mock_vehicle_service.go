package mocks

import (
	"ev-pooling-test-api/api/v1/models/entities"

	"github.com/stretchr/testify/mock"
)

type MockVehicleService struct {
	mock.Mock
}

func (mock *MockVehicleService) GetAllVehicles() []*entities.Vehicle {
	args := mock.Called()
	return args.Get(0).([]*entities.Vehicle)
}

func (mock *MockVehicleService) AddVehicle(vehicle *entities.Vehicle) bool {
	args := mock.Called(vehicle)
	return args.Bool(0)
}

func (mock *MockVehicleService) AddVehicles(vehicles []*entities.Vehicle) bool {
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
