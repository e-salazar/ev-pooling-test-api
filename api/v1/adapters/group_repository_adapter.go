package adapters

import (
	"errors"
	"ev-pooling-test-api/api/v1/models/entities"
	"ev-pooling-test-api/api/v1/repositories"
)

type GroupRepositoryAdapter struct {
	groups []*entities.Group
}

func NewGroupRepositoryAdapter() repositories.GroupRepository {
	return &GroupRepositoryAdapter{
		groups: make([]*entities.Group, 0),
	}
}

func (adapter *GroupRepositoryAdapter) Add(group *entities.Group) error {
	adapter.groups = append(adapter.groups, group)
	return nil
}

func (adapter *GroupRepositoryAdapter) FindById(id int) (*entities.Group, error) {
	for _, group := range adapter.groups {
		if group.ID == id {
			return group, nil
		}
	}

	return nil, nil
}

func (adapter *GroupRepositoryAdapter) FindAll() ([]*entities.Group, error) {
	return adapter.groups, nil
}

func (adapter *GroupRepositoryAdapter) RemoveByID(id int) error {
	for i, group := range adapter.groups {
		if group.ID == id {
			adapter.groups = append(adapter.groups[:i], adapter.groups[i+1:]...)
			return nil
		}
	}

	return errors.New("group not found")
}

func (adapter *GroupRepositoryAdapter) RemoveAll() error {
	adapter.groups = make([]*entities.Group, 0)
	return nil
}
