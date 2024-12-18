package persistence

import (
	"errors"
	"ev-pooling-test-api/internal/domain/journey"
)

type JourneyRepositoryInMemory struct {
	journeys []*journey.Journey
}

func NewJourneyRepositoryInMemory() journey.JourneyRepository {
	return &JourneyRepositoryInMemory{
		journeys: make([]*journey.Journey, 0),
	}
}

func (repository *JourneyRepositoryInMemory) Add(journey *journey.Journey) error {
	repository.journeys = append(repository.journeys, journey)
	return nil
}

func (repository *JourneyRepositoryInMemory) FindByGroupId(groupID int) (*journey.Journey, error) {
	for _, journey := range repository.journeys {
		if journey.Group.ID == groupID {
			return journey, nil
		}
	}

	return nil, nil
}

func (repository *JourneyRepositoryInMemory) FindAll() ([]*journey.Journey, error) {
	return repository.journeys, nil
}

func (repository *JourneyRepositoryInMemory) RemoveByGroupId(groupID int) error {
	for i, journey := range repository.journeys {
		if journey.Group.ID == groupID {
			repository.journeys = append(repository.journeys[:i], repository.journeys[i+1:]...)
			return nil
		}
	}

	return errors.New("journey not found")
}

func (repository *JourneyRepositoryInMemory) RemoveAll() error {
	repository.journeys = make([]*journey.Journey, 0)
	return nil
}
