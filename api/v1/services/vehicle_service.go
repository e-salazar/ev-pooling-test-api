package services

import (
	"sync"

	"ev-pooling-test-api/api/v1/models/entities"
	"ev-pooling-test-api/api/v1/repositories"
)

type VehicleService interface {
	GetAllVehicles() []*entities.Vehicle
	AddVehicle(vehicle *entities.Vehicle) bool
	AddVehicles(vehicles []*entities.Vehicle) bool
	RemoveVehicleByID(id int) bool
	RemoveAll() bool
}

type VehicleServiceImpl struct {
	repository repositories.VehicleRepository
	mutex      *sync.Mutex
}

func NewVehicleService(repository repositories.VehicleRepository) *VehicleServiceImpl {
	return &VehicleServiceImpl{
		repository: repository,
		mutex:      &sync.Mutex{},
	}
}

func (service *VehicleServiceImpl) GetAllVehicles() []*entities.Vehicle {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	vehicles, err := service.repository.FindAll()
	if err != nil {
		return nil
	}

	return vehicles
}

func (service *VehicleServiceImpl) AddVehicle(vehicle *entities.Vehicle) bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.Add(vehicle)
	return err == nil
}

func (service *VehicleServiceImpl) AddVehicles(vehicles []*entities.Vehicle) bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.AddAll(vehicles)
	return err == nil
}

func (service *VehicleServiceImpl) RemoveVehicleByID(id int) bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.RemoveByID(id)
	return err == nil
}

func (service *VehicleServiceImpl) RemoveAll() bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.RemoveAll()
	return err == nil
}
