package mocks

import (
	"ev-pooling-test-api/api/v1/models/entities"

	"github.com/stretchr/testify/mock"
)

// MockJourneyService es un mock del servicio de journeys
type MockJourneyService struct {
	mock.Mock
}

func (m *MockJourneyService) Status() {
	m.Called()
}

func (m *MockJourneyService) GetAllJourneys() []*entities.Journey {
	args := m.Called()
	return args.Get(0).([]*entities.Journey)
}

func (m *MockJourneyService) AddJourney(journey *entities.Journey) bool {
	args := m.Called(journey)
	return args.Bool(0)
}

func (m *MockJourneyService) GetJourneyByGroupID(groupID int) *entities.Journey {
	args := m.Called(groupID)
	return args.Get(0).(*entities.Journey)
}

func (m *MockJourneyService) RemoveJourneyByGroupID(groupID int) bool {
	args := m.Called(groupID)
	return args.Bool(0)
}

func (m *MockJourneyService) RemoveAll() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockJourneyService) CreateJourneys() {
	m.Called()
}

func (m *MockJourneyService) GetInfoByGroupID(groupID int) (bool, *entities.Journey) {
	args := m.Called(groupID)
	return args.Bool(0), args.Get(1).(*entities.Journey)
}
