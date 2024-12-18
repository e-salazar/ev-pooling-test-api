package adapters

import (
	"errors"
	"ev-pooling-test-api/api/v1/models/entities"
	"ev-pooling-test-api/api/v1/repositories"
)

type JourneyRepositoryAdapter struct {
	journeys []*entities.Journey
}

func NewJourneyRepositoryAdapter() repositories.JourneyRepository {
	return &JourneyRepositoryAdapter{
		journeys: make([]*entities.Journey, 0),
	}
}

func (adapter *JourneyRepositoryAdapter) Add(journey *entities.Journey) error {
	adapter.journeys = append(adapter.journeys, journey)
	return nil
}

func (adapter *JourneyRepositoryAdapter) FindByGroupId(groupID int) (*entities.Journey, error) {
	for _, journey := range adapter.journeys {
		if journey.Group.ID == groupID {
			return journey, nil
		}
	}

	return nil, nil
}

func (adapter *JourneyRepositoryAdapter) FindAll() ([]*entities.Journey, error) {
	return adapter.journeys, nil
}

func (adapter *JourneyRepositoryAdapter) RemoveByGroupId(groupID int) error {
	for i, journey := range adapter.journeys {
		if journey.Group.ID == groupID {
			adapter.journeys = append(adapter.journeys[:i], adapter.journeys[i+1:]...)
			return nil
		}
	}

	return errors.New("journey not found")
}

func (adapter *JourneyRepositoryAdapter) RemoveAll() error {
	adapter.journeys = make([]*entities.Journey, 0)
	return nil
}
