package adapters

import (
	"errors"
	"ev-pooling-test-api/api/v1/models/entities"
	"ev-pooling-test-api/api/v1/repositories"
	"sort"
)

type VehicleRepositoryAdapter struct {
	vehicles []*entities.Vehicle
}

func NewVehicleRepositoryAdapter() repositories.VehicleRepository {
	return &VehicleRepositoryAdapter{
		vehicles: make([]*entities.Vehicle, 0),
	}
}

func (adapter *VehicleRepositoryAdapter) Add(vehicle *entities.Vehicle) error {
	adapter.vehicles = append(adapter.vehicles, vehicle)
	sort.Slice(adapter.vehicles, func(i, j int) bool {
		return adapter.vehicles[i].Seats < adapter.vehicles[j].Seats
	})
	return nil
}

func (adapter *VehicleRepositoryAdapter) AddAll(vehicles []*entities.Vehicle) error {
	adapter.vehicles = append(adapter.vehicles, vehicles...)
	sort.Slice(adapter.vehicles, func(i, j int) bool {
		return adapter.vehicles[i].Seats < adapter.vehicles[j].Seats
	})
	return nil
}

func (adapter *VehicleRepositoryAdapter) FindById(id int) (*entities.Vehicle, error) {
	for _, vehicle := range adapter.vehicles {
		if vehicle.ID == id {
			return vehicle, nil
		}
	}

	return nil, nil
}

func (adapter *VehicleRepositoryAdapter) FindAll() ([]*entities.Vehicle, error) {
	return adapter.vehicles, nil
}

func (adapter *VehicleRepositoryAdapter) RemoveByID(id int) error {
	for i, vehicle := range adapter.vehicles {
		if vehicle.ID == id {
			adapter.vehicles = append(adapter.vehicles[:i], adapter.vehicles[i+1:]...)
			return nil
		}
	}

	return errors.New("vehicle not found")
}

func (adapter *VehicleRepositoryAdapter) RemoveAll() error {
	adapter.vehicles = make([]*entities.Vehicle, 0)
	return nil
}
