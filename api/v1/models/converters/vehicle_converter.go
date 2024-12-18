package converters

import (
	"ev-pooling-test-api/api/v1/models/dtos"
	"ev-pooling-test-api/api/v1/models/entities"
)

func ToVehicleDTO(vehicle *entities.Vehicle) *dtos.VehicleDTO {
	return &dtos.VehicleDTO{
		ID:    vehicle.ID,
		Seats: vehicle.Seats,
	}
}

func ToVehicleEntity(vehicleDTO *dtos.VehicleDTO) *entities.Vehicle {
	return &entities.Vehicle{
		ID:    vehicleDTO.ID,
		Seats: vehicleDTO.Seats,
	}
}

func ToVehicleEntities(vehiclesDTO []*dtos.VehicleDTO) []*entities.Vehicle {
	vehicles := make([]*entities.Vehicle, len(vehiclesDTO))
	for i, vehicleDTO := range vehiclesDTO {
		vehicles[i] = ToVehicleEntity(vehicleDTO)
	}
	return vehicles
}
