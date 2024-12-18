package journey

type JourneyRepository interface {
	Add(journey *Journey) error
	FindByGroupId(groupID int) (*Journey, error)
	FindAll() ([]*Journey, error)
	RemoveByGroupId(groupID int) error
	RemoveAll() error
}
