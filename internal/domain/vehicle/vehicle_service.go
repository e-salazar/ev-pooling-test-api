package vehicle

type VehicleService interface {
	GetAllVehicles() []*Vehicle
	AddVehicle(vehicle *Vehicle) bool
	AddVehicles(vehicles []*Vehicle) bool
	RemoveVehicleByID(id int) bool
	RemoveAll() bool
}
