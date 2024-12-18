package services

import (
	"ev-pooling-test-api/internal/domain/group"
	"sync"
)

type GroupService struct {
	repository group.GroupRepository
	mutex      *sync.Mutex
}

func NewGroupService(repository group.GroupRepository) *GroupService {
	return &GroupService{
		repository: repository,
		mutex:      &sync.Mutex{},
	}
}

func (service *GroupService) GetGroupByID(id int) *group.Group {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	group, err := service.repository.FindById(id)
	if err != nil {
		return nil
	}

	return group
}

func (service *GroupService) GetAllGroups() []*group.Group {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	groups, err := service.repository.FindAll()
	if err != nil {
		return nil
	}

	return groups
}

func (service *GroupService) AddGroup(group *group.Group) bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.Add(group)
	return err == nil
}

func (service *GroupService) RemoveGroupByID(id int) bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.RemoveByID(id)
	return err == nil
}

func (service *GroupService) RemoveAll() bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.RemoveAll()
	return err == nil
}
