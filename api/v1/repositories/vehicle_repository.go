package repositories

import "ev-pooling-test-api/api/v1/models/entities"

type VehicleRepository interface {
	Add(vehicle *entities.Vehicle) error
	AddAll(vehicles []*entities.Vehicle) error
	FindById(id int) (*entities.Vehicle, error)
	FindAll() ([]*entities.Vehicle, error)
	RemoveByID(id int) error
	RemoveAll() error
}
