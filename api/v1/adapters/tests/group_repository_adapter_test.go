package adapters_tests

import (
	"ev-pooling-test-api/api/v1/adapters"
	"ev-pooling-test-api/api/v1/models/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGroupRepositoryAdapter(t *testing.T) {
	repository := adapters.NewGroupRepositoryAdapter()
	assert.NotNil(t, repository)
}

func TestGroupAdd(t *testing.T) {
	repository := adapters.NewGroupRepositoryAdapter()
	created := &entities.Group{ID: 1, People: 5}

	err := repository.Add(created)
	assert.Nil(t, err)

	found, err := repository.FindById(1)
	assert.Nil(t, err)
	assert.Equal(t, created, found)
}

func TestGroupFindById(t *testing.T) {
	repository := adapters.NewGroupRepositoryAdapter()
	created := &entities.Group{ID: 1, People: 5}

	repository.Add(created)

	found, err := repository.FindById(1)
	assert.Nil(t, err)
	assert.Equal(t, created, found)

	notFound, err := repository.FindById(2)
	assert.Nil(t, err)
	assert.Nil(t, notFound)
}

func TestGroupFindAll(t *testing.T) {
	repository := adapters.NewGroupRepositoryAdapter()
	created1 := &entities.Group{ID: 1, People: 5}
	created2 := &entities.Group{ID: 2, People: 6}
	repository.Add(created1)
	repository.Add(created2)

	found, err := repository.FindAll()
	assert.Nil(t, err)
	assert.Len(t, found, 2)
	assert.Contains(t, found, created1)
	assert.Contains(t, found, created2)
}

func TestGroupRemoveByID(t *testing.T) {
	repository := adapters.NewGroupRepositoryAdapter()
	created := &entities.Group{ID: 1, People: 5}
	repository.Add(created)

	err := repository.RemoveByID(1)
	assert.Nil(t, err)

	found, err := repository.FindById(1)
	assert.Nil(t, err)
	assert.Nil(t, found)

	err = repository.RemoveByID(2)
	assert.NotNil(t, err)
	assert.Equal(t, "group not found", err.Error())
}

func TestGroupRemoveAll(t *testing.T) {
	repository := adapters.NewGroupRepositoryAdapter()
	created1 := &entities.Group{ID: 1, People: 5}
	created2 := &entities.Group{ID: 2, People: 6}
	repository.Add(created1)
	repository.Add(created2)

	err := repository.RemoveAll()
	assert.Nil(t, err)

	found, err := repository.FindAll()
	assert.Nil(t, err)
	assert.Len(t, found, 0)
}
