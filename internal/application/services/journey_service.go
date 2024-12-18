package services

import (
	"ev-pooling-test-api/internal/domain/group"
	"ev-pooling-test-api/internal/domain/journey"
	"ev-pooling-test-api/internal/domain/vehicle"
	"fmt"
	"sync"
)

type JourneyService struct {
	repository     journey.JourneyRepository
	vehicleService vehicle.VehicleService
	groupService   group.GroupService
	mutex          *sync.Mutex
}

func NewJourneyService(repository journey.JourneyRepository, vehicleService vehicle.VehicleService, groupService group.GroupService) *JourneyService {
	return &JourneyService{
		repository:     repository,
		vehicleService: vehicleService,
		groupService:   groupService,
		mutex:          &sync.Mutex{},
	}
}

func (service *JourneyService) Status() {
	// Print available vehicles
	fmt.Println("------------------------------")
	fmt.Println("Available vehicles:", len(service.vehicleService.GetAllVehicles()))
	for _, vehicle := range service.vehicleService.GetAllVehicles() {
		fmt.Println(vehicle)
	}

	// Print waiting groups
	fmt.Println("Waiting groups:", len(service.groupService.GetAllGroups()))
	for _, group := range service.groupService.GetAllGroups() {
		fmt.Println(group)
	}

	// Print current journeys
	fmt.Println("Current journeys:", len(service.GetAllJourneys()))
	for _, journey := range service.GetAllJourneys() {
		fmt.Println("Vehicle:", journey.Vehicle, "Group:", journey.Group)
	}
}

func (service *JourneyService) GetAllJourneys() []*journey.Journey {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	journeys, err := service.repository.FindAll()
	if err != nil {
		return nil
	}

	return journeys
}

func (service *JourneyService) AddJourney(journey *journey.Journey) bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.Add(journey)
	return err == nil
}

func (service *JourneyService) GetJourneyByGroupID(groupID int) *journey.Journey {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	journey, err := service.repository.FindByGroupId(groupID)
	if err != nil {
		return nil
	}

	return journey
}

func (service *JourneyService) RemoveJourneyByGroupID(groupID int) bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.RemoveByGroupId(groupID)
	return err == nil
}

func (service *JourneyService) RemoveAll() bool {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	err := service.repository.RemoveAll()
	return err == nil
}

func (service *JourneyService) CreateJourneys() {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	// Get all waiting waitingGroups from groupservice
	for i := 0; i < len(service.groupService.GetAllGroups()); i++ {
		// Check if there is an available vehicle for the group
		var availableVehicle *vehicle.Vehicle
		for _, vehicle := range service.vehicleService.GetAllVehicles() {
			if vehicle.Seats >= service.groupService.GetAllGroups()[i].People {
				availableVehicle = vehicle
				service.vehicleService.RemoveVehicleByID(vehicle.ID)
				break
			}
		}

		// Assign vehicle and group to a new journey
		if availableVehicle != nil {
			journey := &journey.Journey{
				Vehicle: availableVehicle,
				Group:   service.groupService.GetAllGroups()[i],
			}
			service.repository.Add(journey)
			service.groupService.RemoveGroupByID(service.groupService.GetAllGroups()[i].ID)
			i-- // Adjust index after removal
		}
	}
}

func (service *JourneyService) GetInfoByGroupID(groupID int) (bool, *journey.Journey) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	waitingGroups := service.groupService.GetAllGroups()
	journeys, err := service.repository.FindAll()
	if err != nil {
		return false, nil
	}

	for _, journey := range journeys {
		if journey.Group.ID == groupID {
			return true, journey
		}
	}

	for _, group := range waitingGroups {
		if group.ID == groupID {
			return true, nil
		}
	}

	return false, nil
}
