package mocks

import (
	"ev-pooling-test-api/internal/domain/journey"

	"github.com/stretchr/testify/mock"
)

type MockJourneyRepository struct {
	mock.Mock
}

func (mock *MockJourneyRepository) Add(journey *journey.Journey) error {
	args := mock.Called(journey)
	return args.Error(0)
}

func (mock *MockJourneyRepository) FindById(id int) (*journey.Journey, error) {
	args := mock.Called(id)
	return args.Get(0).(*journey.Journey), args.Error(1)
}

func (mock *MockJourneyRepository) FindAll() ([]*journey.Journey, error) {
	args := mock.Called()
	return args.Get(0).([]*journey.Journey), args.Error(1)
}

func (mock *MockJourneyRepository) FindByGroupId(groupID int) (*journey.Journey, error) {
	args := mock.Called(groupID)
	return args.Get(0).(*journey.Journey), args.Error(1)
}

func (mock *MockJourneyRepository) RemoveByGroupId(groupID int) error {
	args := mock.Called(groupID)
	return args.Error(0)
}

func (mock *MockJourneyRepository) RemoveAll() error {
	args := mock.Called()
	return args.Error(0)
}
