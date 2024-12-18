package services

import (
	"ev-pooling-test-api/internal/domain/vehicle"
	"sync"
)

type VehicleService struct {
	repository vehicle.VehicleRepository
	mutex      *sync.Mutex
}

func NewVehicleService(repository vehicle.VehicleRepository) *VehicleService {
	return &VehicleService{
		repository: repository,
		mutex:      &sync.Mutex{},
	}
}

func (service *VehicleService) GetAllVehicles() []*vehicle.Vehicle {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	vehicles, err := service.repository.FindAll()
	if err != nil {
		return nil
	}

	return vehicles
}

func (service *VehicleService) AddVehicle(vehicle *vehicle.Vehicle) bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.Add(vehicle)
	return err == nil
}

func (service *VehicleService) AddVehicles(vehicles []*vehicle.Vehicle) bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.AddAll(vehicles)
	return err == nil
}

func (service *VehicleService) RemoveVehicleByID(id int) bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.RemoveByID(id)
	return err == nil
}

func (service *VehicleService) RemoveAll() bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.RemoveAll()
	return err == nil
}
