package mocks

import (
	"ev-pooling-test-api/api/v1/models/entities"

	"github.com/stretchr/testify/mock"
)

type MockGroupRepository struct {
	mock.Mock
}

func (mock *MockGroupRepository) Add(group *entities.Group) error {
	args := mock.Called(group)
	return args.Error(0)
}

func (mock *MockGroupRepository) FindById(id int) (*entities.Group, error) {
	args := mock.Called(id)
	return args.Get(0).(*entities.Group), args.Error(1)
}

func (mock *MockGroupRepository) FindAll() ([]*entities.Group, error) {
	args := mock.Called()
	return args.Get(0).([]*entities.Group), args.Error(1)
}

func (mock *MockGroupRepository) RemoveByID(id int) error {
	args := mock.Called(id)
	return args.Error(0)
}

func (mock *MockGroupRepository) RemoveAll() error {
	args := mock.Called()
	return args.Error(0)
}
