package journey

type JourneyService interface {
	GetAllJourneys() []*Journey
	AddJourney(journey *Journey) bool
	GetJourneyByGroupID(groupID int) *Journey
	RemoveJourneyByGroupID(groupID int) bool
	RemoveAll() bool
	CreateJourneys()
	GetInfoByGroupID(groupID int) (bool, *Journey)
}
