package repositories

import "ev-pooling-test-api/api/v1/models/entities"

type GroupRepository interface {
	Add(group *entities.Group) error
	FindById(id int) (*entities.Group, error)
	FindAll() ([]*entities.Group, error)
	RemoveByID(id int) error
	RemoveAll() error
}
