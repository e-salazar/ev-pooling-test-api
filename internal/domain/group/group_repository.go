package group

type GroupRepository interface {
	Add(group *Group) error
	FindById(id int) (*Group, error)
	FindAll() ([]*Group, error)
	RemoveByID(id int) error
	RemoveAll() error
}
