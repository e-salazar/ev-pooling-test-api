package mocks

import (
	"ev-pooling-test-api/internal/domain/vehicle"

	"github.com/stretchr/testify/mock"
)

type MockVehicleRepository struct {
	mock.Mock
}

func (mock *MockVehicleRepository) Add(vehicle *vehicle.Vehicle) error {
	args := mock.Called(vehicle)
	return args.Error(0)
}

func (mock *MockVehicleRepository) AddAll(vehicles []*vehicle.Vehicle) error {
	args := mock.Called(vehicles)
	return args.Error(0)
}

func (mock *MockVehicleRepository) FindById(id int) (*vehicle.Vehicle, error) {
	args := mock.Called(id)
	return args.Get(0).(*vehicle.Vehicle), args.Error(1)
}

func (mock *MockVehicleRepository) FindAll() ([]*vehicle.Vehicle, error) {
	args := mock.Called()
	return args.Get(0).([]*vehicle.Vehicle), args.Error(1)
}

func (mock *MockVehicleRepository) RemoveByID(id int) error {
	args := mock.Called(id)
	return args.Error(0)
}

func (mock *MockVehicleRepository) RemoveAll() error {
	args := mock.Called()
	return args.Error(0)
}
