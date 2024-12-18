package mocks

import (
	"ev-pooling-test-api/api/v1/models/entities"

	"github.com/stretchr/testify/mock"
)

type MockJourneyRepository struct {
	mock.Mock
}

func (mock *MockJourneyRepository) Add(journey *entities.Journey) error {
	args := mock.Called(journey)
	return args.Error(0)
}

func (mock *MockJourneyRepository) FindById(id int) (*entities.Journey, error) {
	args := mock.Called(id)
	return args.Get(0).(*entities.Journey), args.Error(1)
}

func (mock *MockJourneyRepository) FindAll() ([]*entities.Journey, error) {
	args := mock.Called()
	return args.Get(0).([]*entities.Journey), args.Error(1)
}

func (mock *MockJourneyRepository) FindByGroupId(groupID int) (*entities.Journey, error) {
	args := mock.Called(groupID)
	return args.Get(0).(*entities.Journey), args.Error(1)
}

func (mock *MockJourneyRepository) RemoveByGroupId(groupID int) error {
	args := mock.Called(groupID)
	return args.Error(0)
}

func (mock *MockJourneyRepository) RemoveAll() error {
	args := mock.Called()
	return args.Error(0)
}
