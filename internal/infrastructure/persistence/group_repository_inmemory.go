package persistence

import (
	"errors"
	"ev-pooling-test-api/internal/domain/group"
)

type GroupRepositoryInMemory struct {
	groups []*group.Group
}

func NewGroupRepositoryInMemory() group.GroupRepository {
	return &GroupRepositoryInMemory{
		groups: make([]*group.Group, 0),
	}
}

func (repository *GroupRepositoryInMemory) Add(group *group.Group) error {
	repository.groups = append(repository.groups, group)
	return nil
}

func (repository *GroupRepositoryInMemory) FindById(id int) (*group.Group, error) {
	for _, group := range repository.groups {
		if group.ID == id {
			return group, nil
		}
	}

	return nil, nil
}

func (repository *GroupRepositoryInMemory) FindAll() ([]*group.Group, error) {
	return repository.groups, nil
}

func (repository *GroupRepositoryInMemory) RemoveByID(id int) error {
	for i, group := range repository.groups {
		if group.ID == id {
			repository.groups = append(repository.groups[:i], repository.groups[i+1:]...)
			return nil
		}
	}

	return errors.New("group not found")
}

func (repository *GroupRepositoryInMemory) RemoveAll() error {
	repository.groups = make([]*group.Group, 0)
	return nil
}
