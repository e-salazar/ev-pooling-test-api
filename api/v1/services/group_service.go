package services

import (
	"sync"

	"ev-pooling-test-api/api/v1/models/entities"
	"ev-pooling-test-api/api/v1/repositories"
)

type GroupService interface {
	GetAllGroups() []*entities.Group
	AddGroup(group *entities.Group) bool
	RemoveGroupByID(id int) bool
	RemoveAll() bool
}

type GroupServiceImpl struct {
	repository repositories.GroupRepository
	mutex      *sync.Mutex
}

func NewGroupService(repository repositories.GroupRepository) *GroupServiceImpl {
	return &GroupServiceImpl{
		repository: repository,
		mutex:      &sync.Mutex{},
	}
}

func (service *GroupServiceImpl) GetAllGroups() []*entities.Group {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	groups, err := service.repository.FindAll()
	if err != nil {
		return nil
	}

	return groups
}

func (service *GroupServiceImpl) AddGroup(group *entities.Group) bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.Add(group)
	return err == nil
}

func (service *GroupServiceImpl) RemoveGroupByID(id int) bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.RemoveByID(id)
	return err == nil
}

func (service *GroupServiceImpl) RemoveAll() bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.RemoveAll()
	return err == nil
}
