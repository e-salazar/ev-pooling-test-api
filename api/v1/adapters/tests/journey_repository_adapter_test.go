package adapters_tests

import (
	"ev-pooling-test-api/api/v1/adapters"
	"ev-pooling-test-api/api/v1/models/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJourneyRepositoryAdapter(t *testing.T) {
	repository := adapters.NewJourneyRepositoryAdapter()
	assert.NotNil(t, repository)
}

func TestJourneyAdd(t *testing.T) {
	repository := adapters.NewJourneyRepositoryAdapter()
	created := &entities.Journey{Vehicle: &entities.Vehicle{ID: 1}, Group: &entities.Group{ID: 1}}

	err := repository.Add(created)
	assert.Nil(t, err)

	found, err := repository.FindByGroupId(1)
	assert.Nil(t, err)
	assert.Equal(t, created, found)
}

func TestJourneyFindById(t *testing.T) {
	repository := adapters.NewJourneyRepositoryAdapter()
	created := &entities.Journey{Vehicle: &entities.Vehicle{ID: 1}, Group: &entities.Group{ID: 1}}

	repository.Add(created)

	found, err := repository.FindByGroupId(1)
	assert.Nil(t, err)
	assert.Equal(t, created, found)

	notFound, err := repository.FindByGroupId(2)
	assert.Nil(t, err)
	assert.Nil(t, notFound)
}

func TestJourneyFindAll(t *testing.T) {
	repository := adapters.NewJourneyRepositoryAdapter()
	created1 := &entities.Journey{Vehicle: &entities.Vehicle{ID: 1}, Group: &entities.Group{ID: 1}}
	created2 := &entities.Journey{Vehicle: &entities.Vehicle{ID: 1}, Group: &entities.Group{ID: 1}}
	repository.Add(created1)
	repository.Add(created2)

	found, err := repository.FindAll()
	assert.Nil(t, err)
	assert.Len(t, found, 2)
	assert.Contains(t, found, created1)
	assert.Contains(t, found, created2)
}

func TestJourneyRemoveByID(t *testing.T) {
	repository := adapters.NewJourneyRepositoryAdapter()
	created := &entities.Journey{Vehicle: &entities.Vehicle{ID: 1}, Group: &entities.Group{ID: 1}}
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
	repository := adapters.NewJourneyRepositoryAdapter()
	created1 := &entities.Journey{Vehicle: &entities.Vehicle{ID: 1}, Group: &entities.Group{ID: 1}}
	created2 := &entities.Journey{Vehicle: &entities.Vehicle{ID: 1}, Group: &entities.Group{ID: 1}}
	repository.Add(created1)
	repository.Add(created2)

	err := repository.RemoveAll()
	assert.Nil(t, err)

	found, err := repository.FindAll()
	assert.Nil(t, err)
	assert.Len(t, found, 0)
}
