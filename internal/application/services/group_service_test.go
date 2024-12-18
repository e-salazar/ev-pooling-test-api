package services

import (
	"ev-pooling-test-api/internal/domain/group"
	"ev-pooling-test-api/internal/domain/group/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupServiceGetAllGroups(t *testing.T) {
	mockRepository := new(mocks.MockGroupRepository)
	service := NewGroupService(mockRepository)

	expected := []*group.Group{
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
	service := NewGroupService(mockRepository)

	created := &group.Group{ID: 1, People: 5}

	mockRepository.On("Add", created).Return(nil)

	result := service.AddGroup(created)
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}

func TestGroupServiceRemoveGroupByID(t *testing.T) {
	mockRepository := new(mocks.MockGroupRepository)
	service := NewGroupService(mockRepository)

	mockRepository.On("RemoveByID", 1).Return(nil)

	result := service.RemoveGroupByID(1)
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}

func TestGroupServiceRemoveAllGroups(t *testing.T) {
	mockRepository := new(mocks.MockGroupRepository)
	service := NewGroupService(mockRepository)

	mockRepository.On("RemoveAll").Return(nil)

	result := service.RemoveAll()
	assert.True(t, result)

	mockRepository.AssertExpectations(t)
}
