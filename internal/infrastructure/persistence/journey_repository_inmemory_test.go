package persistence

import (
	"ev-pooling-test-api/internal/domain/group"
	"ev-pooling-test-api/internal/domain/journey"
	"ev-pooling-test-api/internal/domain/vehicle"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJourneyRepositoryInMemory(t *testing.T) {
	repository := NewJourneyRepositoryInMemory()
	assert.NotNil(t, repository)
}

func TestJourneyAdd(t *testing.T) {
	repository := NewJourneyRepositoryInMemory()
	created := &journey.Journey{Vehicle: &vehicle.Vehicle{ID: 1}, Group: &group.Group{ID: 1}}

	err := repository.Add(created)
	assert.Nil(t, err)

	found, err := repository.FindByGroupId(1)
	assert.Nil(t, err)
	assert.Equal(t, created, found)
}

func TestJourneyFindById(t *testing.T) {
	repository := NewJourneyRepositoryInMemory()
	created := &journey.Journey{Vehicle: &vehicle.Vehicle{ID: 1}, Group: &group.Group{ID: 1}}

	repository.Add(created)

	found, err := repository.FindByGroupId(1)
	assert.Nil(t, err)
	assert.Equal(t, created, found)

	notFound, err := repository.FindByGroupId(2)
	assert.Nil(t, err)
	assert.Nil(t, notFound)
}

func TestJourneyFindAll(t *testing.T) {
	repository := NewJourneyRepositoryInMemory()
	created1 := &journey.Journey{Vehicle: &vehicle.Vehicle{ID: 1}, Group: &group.Group{ID: 1}}
	created2 := &journey.Journey{Vehicle: &vehicle.Vehicle{ID: 1}, Group: &group.Group{ID: 1}}
	repository.Add(created1)
	repository.Add(created2)

	found, err := repository.FindAll()
	assert.Nil(t, err)
	assert.Len(t, found, 2)
	assert.Contains(t, found, created1)
	assert.Contains(t, found, created2)
}

func TestJourneyRemoveByID(t *testing.T) {
	repository := NewJourneyRepositoryInMemory()
	created := &journey.Journey{Vehicle: &vehicle.Vehicle{ID: 1}, Group: &group.Group{ID: 1}}
	repository.Add(created)

	err := repository.RemoveByGroupId(1)
	assert.Nil(t, err)

	found, err := repository.FindByGroupId(1)
	assert.Nil(t, err)
	assert.Nil(t, found)

	err = repository.RemoveByGroupId(2)
	assert.NotNil(t, err)
	assert.Equal(t, "journey not found", err.Error())
}

func TestJourneyRemoveAll(t *testing.T) {
	repository := NewJourneyRepositoryInMemory()
	created1 := &journey.Journey{Vehicle: &vehicle.Vehicle{ID: 1}, Group: &group.Group{ID: 1}}
	created2 := &journey.Journey{Vehicle: &vehicle.Vehicle{ID: 1}, Group: &group.Group{ID: 1}}
	repository.Add(created1)
	repository.Add(created2)

	err := repository.RemoveAll()
	assert.Nil(t, err)

	found, err := repository.FindAll()
	assert.Nil(t, err)
	assert.Len(t, found, 0)
}
