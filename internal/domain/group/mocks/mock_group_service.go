package mocks

import (
	"ev-pooling-test-api/internal/domain/group"

	"github.com/stretchr/testify/mock"
)

type MockGroupService struct {
	mock.Mock
}

func (mock *MockGroupService) GetAllGroups() []*group.Group {
	args := mock.Called()
	return args.Get(0).([]*group.Group)
}

func (mock *MockGroupService) AddGroup(vehicle *group.Group) bool {
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
