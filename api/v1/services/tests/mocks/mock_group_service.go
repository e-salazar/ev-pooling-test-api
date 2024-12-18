package mocks

import (
	"ev-pooling-test-api/api/v1/models/entities"

	"github.com/stretchr/testify/mock"
)

type MockGroupService struct {
	mock.Mock
}

func (mock *MockGroupService) GetAllGroups() []*entities.Group {
	args := mock.Called()
	return args.Get(0).([]*entities.Group)
}

func (mock *MockGroupService) AddGroup(vehicle *entities.Group) bool {
	args := mock.Called(vehicle)
	return args.Bool(0)
}

func (mock *MockGroupService) RemoveGroupByID(id int) bool {
	args := mock.Called(id)
	return args.Bool(0)
}

func (mock *MockGroupService) RemoveAll() bool {
	args := mock.Called()
	return args.Bool(0)
}
