package vehicle

type VehicleRepository interface {
	Add(vehicle *Vehicle) error
	AddAll(vehicles []*Vehicle) error
	FindById(id int) (*Vehicle, error)
	FindAll() ([]*Vehicle, error)
	RemoveByID(id int) error
	RemoveAll() error
}
