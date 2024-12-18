package converters

import (
	"ev-pooling-test-api/internal/application/dtos"
	"ev-pooling-test-api/internal/domain/group"
)

func ToGroupDTO(group *group.Group) *dtos.GroupDTO {
	return &dtos.GroupDTO{
		ID:     group.ID,
		People: group.People,
	}
}

func ToGroupEntity(groupDTO *dtos.GroupDTO) *group.Group {
	return &group.Group{
		ID:     groupDTO.ID,
		People: groupDTO.People,
	}
}

func ToGroupEntities(groupsDTO []*dtos.GroupDTO) []*group.Group {
	groups := make([]*group.Group, len(groupsDTO))
	for i, groupDTO := range groupsDTO {
		groups[i] = ToGroupEntity(groupDTO)
	}
	return groups
}
