package converters

import (
	"ev-pooling-test-api/internal/application/dtos"
	"ev-pooling-test-api/internal/domain/vehicle"
)

func ToVehicleDTO(vehicle *vehicle.Vehicle) *dtos.VehicleDTO {
	return &dtos.VehicleDTO{
		ID:    vehicle.ID,
		Seats: vehicle.Seats,
	}
}

func ToVehicleEntity(vehicleDTO *dtos.VehicleDTO) *vehicle.Vehicle {
	return &vehicle.Vehicle{
		ID:    vehicleDTO.ID,
		Seats: vehicleDTO.Seats,
	}
}

func ToVehicleEntities(vehiclesDTO []*dtos.VehicleDTO) []*vehicle.Vehicle {
	vehicles := make([]*vehicle.Vehicle, len(vehiclesDTO))
	for i, vehicleDTO := range vehiclesDTO {
		vehicles[i] = ToVehicleEntity(vehicleDTO)
	}
	return vehicles
}
