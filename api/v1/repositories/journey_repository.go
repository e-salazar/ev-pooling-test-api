package repositories

import "ev-pooling-test-api/api/v1/models/entities"

type JourneyRepository interface {
	Add(journey *entities.Journey) error
	FindByGroupId(groupID int) (*entities.Journey, error)
	FindAll() ([]*entities.Journey, error)
	RemoveByGroupId(groupID int) error
	RemoveAll() error
}
