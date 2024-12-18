package mocks

import (
	"ev-pooling-test-api/internal/domain/group"

	"github.com/stretchr/testify/mock"
)

type MockGroupRepository struct {
	mock.Mock
}

func (mock *MockGroupRepository) Add(group *group.Group) error {
	args := mock.Called(group)
	return args.Error(0)
}

func (mock *MockGroupRepository) FindById(id int) (*group.Group, error) {
	args := mock.Called(id)
	return args.Get(0).(*group.Group), args.Error(1)
}

func (mock *MockGroupRepository) FindAll() ([]*group.Group, error) {
	args := mock.Called()
	return args.Get(0).([]*group.Group), args.Error(1)
}

func (mock *MockGroupRepository) RemoveByID(id int) error {
	args := mock.Called(id)
	return args.Error(0)
}

func (mock *MockGroupRepository) RemoveAll() error {
	args := mock.Called()
	return args.Error(0)
}
