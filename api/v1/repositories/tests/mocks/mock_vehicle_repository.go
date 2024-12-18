package mocks

import (
	"ev-pooling-test-api/api/v1/models/entities"

	"github.com/stretchr/testify/mock"
)

type MockVehicleRepository struct {
	mock.Mock
}

func (mock *MockVehicleRepository) Add(vehicle *entities.Vehicle) error {
	args := mock.Called(vehicle)
	return args.Error(0)
}

func (mock *MockVehicleRepository) AddAll(vehicles []*entities.Vehicle) error {
	args := mock.Called(vehicles)
	return args.Error(0)
}

func (mock *MockVehicleRepository) FindById(id int) (*entities.Vehicle, error) {
	args := mock.Called(id)
	return args.Get(0).(*entities.Vehicle), args.Error(1)
}

func (mock *MockVehicleRepository) FindAll() ([]*entities.Vehicle, error) {
	args := mock.Called()
	return args.Get(0).([]*entities.Vehicle), args.Error(1)
}

func (mock *MockVehicleRepository) RemoveByID(id int) error {
	args := mock.Called(id)
	return args.Error(0)
}

func (mock *MockVehicleRepository) RemoveAll() error {
	args := mock.Called()
	return args.Error(0)
}
