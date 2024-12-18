package group

type GroupService interface {
	GetAllGroups() []*Group
	AddGroup(group *Group) bool
	RemoveGroupByID(id int) bool
	RemoveAll() bool
	GetGroupByID(id int) *Group
}
