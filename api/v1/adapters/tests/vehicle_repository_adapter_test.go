package adapters_tests

import (
	"ev-pooling-test-api/api/v1/adapters"
	"ev-pooling-test-api/api/v1/models/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVehicleRepositoryAdapter(t *testing.T) {
	repository := adapters.NewVehicleRepositoryAdapter()
	assert.NotNil(t, repository)
}

func TestVehicleAdd(t *testing.T) {
	repository := adapters.NewVehicleRepositoryAdapter()
	vehicle := &entities.Vehicle{ID: 1, Seats: 5}

	err := repository.Add(vehicle)
	assert.Nil(t, err)

	found, err := repository.FindById(1)
	assert.Nil(t, err)
	assert.Equal(t, vehicle, found)
}

func TestVehicleAddAll(t *testing.T) {
	repository := adapters.NewVehicleRepositoryAdapter()
	created := []*entities.Vehicle{
		{ID: 1, Seats: 5},
		{ID: 2, Seats: 5},
	}
	err := repository.AddAll(created)
	assert.Nil(t, err)
	assert.Len(t, created, 2)
}

func TestVehicleFindById(t *testing.T) {
	repository := adapters.NewVehicleRepositoryAdapter()
	created := &entities.Vehicle{ID: 1}

	repository.Add(created)

	found, err := repository.FindById(1)
	assert.Nil(t, err)
	assert.Equal(t, created, found)

	notFound, err := repository.FindById(2)
	assert.Nil(t, err)
	assert.Nil(t, notFound)
}

func TestVehicleFindAll(t *testing.T) {
	repository := adapters.NewVehicleRepositoryAdapter()
	created := []*entities.Vehicle{
		{ID: 1, Seats: 5},
		{ID: 2, Seats: 5},
	}
	repository.AddAll(created)

	found, err := repository.FindAll()
	assert.Nil(t, err)
	assert.Len(t, found, 2)
	assert.Contains(t, found, created[0])
	assert.Contains(t, found, created[1])
}

func TestVehicleRemoveByID(t *testing.T) {
	repository := adapters.NewVehicleRepositoryAdapter()
	created := &entities.Vehicle{ID: 1, Seats: 5}
	repository.Add(created)

	err := repository.RemoveByID(1)
	assert.Nil(t, err)

	found, err := repository.FindById(1)
	assert.Nil(t, err)
	assert.Nil(t, found)

	err = repository.RemoveByID(2)
	assert.NotNil(t, err)
	assert.Equal(t, "vehicle not found", err.Error())
}

func TestVehicleRemoveAll(t *testing.T) {
	repository := adapters.NewVehicleRepositoryAdapter()
	created := []*entities.Vehicle{
		{ID: 1, Seats: 5},
		{ID: 2, Seats: 5},
	}
	repository.AddAll(created)

	err := repository.RemoveAll()
	assert.Nil(t, err)

	found, err := repository.FindAll()
	assert.Nil(t, err)
	assert.Len(t, found, 0)
}
