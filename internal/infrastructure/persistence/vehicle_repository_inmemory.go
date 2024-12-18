package persistence

import (
	"errors"
	"ev-pooling-test-api/internal/domain/vehicle"
	"sort"
)

type VehicleRepositoryInMemory struct {
	vehicles []*vehicle.Vehicle
}

func NewVehicleRepositoryInMemory() vehicle.VehicleRepository {
	return &VehicleRepositoryInMemory{
		vehicles: make([]*vehicle.Vehicle, 0),
	}
}

func (repository *VehicleRepositoryInMemory) Add(vehicle *vehicle.Vehicle) error {
	repository.vehicles = append(repository.vehicles, vehicle)
	sort.Slice(repository.vehicles, func(i, j int) bool {
		return repository.vehicles[i].Seats < repository.vehicles[j].Seats
	})
	return nil
}

func (repository *VehicleRepositoryInMemory) AddAll(vehicles []*vehicle.Vehicle) error {
	repository.vehicles = append(repository.vehicles, vehicles...)
	sort.Slice(repository.vehicles, func(i, j int) bool {
		return repository.vehicles[i].Seats < repository.vehicles[j].Seats
	})
	return nil
}

func (repository *VehicleRepositoryInMemory) FindById(id int) (*vehicle.Vehicle, error) {
	for _, vehicle := range repository.vehicles {
		if vehicle.ID == id {
			return vehicle, nil
		}
	}

	return nil, nil
}

func (repository *VehicleRepositoryInMemory) FindAll() ([]*vehicle.Vehicle, error) {
	return repository.vehicles, nil
}

func (repository *VehicleRepositoryInMemory) RemoveByID(id int) error {
	for i, vehicle := range repository.vehicles {
		if vehicle.ID == id {
			repository.vehicles = append(repository.vehicles[:i], repository.vehicles[i+1:]...)
			return nil
		}
	}

	return errors.New("vehicle not found")
}

func (repository *VehicleRepositoryInMemory) RemoveAll() error {
	repository.vehicles = make([]*vehicle.Vehicle, 0)
	return nil
}
