package services_test

import (
	"ev-pooling-test-api/api/v1/models/entities"
	"ev-pooling-test-api/api/v1/repositories/tests/mocks"
	"ev-pooling-test-api/api/v1/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupServiceGetAllGroups(t *testing.T) {
	mockRepository := new(mocks.MockGroupRepository)
	service := services.NewGroupService(mockRepository)

	expected := []*entities.Group{
		{ID: 1, People: 5},
		{ID: 2, People: 3},
	}

	mockRepository.On("FindAll").Return(expected, nil)

	result := service.GetAllGroups()
	assert.Equal(t, expected, result)

	mockRepository.AssertExpectations(t)
}

func TestGroupServiceAddGroup(t *testing.T) {
	mockRepository := new(mocks.MockGroupRepository)
	service := services.NewGroupService(mockRepository)

	created := &entities.Group{ID: 1, People: 5}

	mockRepository.On("Add", created).Return(nil)

	result := service.AddGroup(created)
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}

func TestGroupServiceRemoveGroupByID(t *testing.T) {
	mockRepository := new(mocks.MockGroupRepository)
	service := services.NewGroupService(mockRepository)

	mockRepository.On("RemoveByID", 1).Return(nil)

	result := service.RemoveGroupByID(1)
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}

func TestGroupServiceRemoveAllGroups(t *testing.T) {
	mockRepository := new(mocks.MockGroupRepository)
	service := services.NewGroupService(mockRepository)

	mockRepository.On("RemoveAll").Return(nil)

	result := service.RemoveAll()
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}
