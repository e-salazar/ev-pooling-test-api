package converters

import (
	"ev-pooling-test-api/api/v1/models/dtos"
	"ev-pooling-test-api/api/v1/models/entities"
)

func ToGroupDTO(group *entities.Group) *dtos.GroupDTO {
	return &dtos.GroupDTO{
		ID:     group.ID,
		People: group.People,
	}
}

func ToGroupEntity(groupDTO *dtos.GroupDTO) *entities.Group {
	return &entities.Group{
		ID:     groupDTO.ID,
		People: groupDTO.People,
	}
}

func ToGroupEntities(groupsDTO []*dtos.GroupDTO) []*entities.Group {
	groups := make([]*entities.Group, len(groupsDTO))
	for i, groupDTO := range groupsDTO {
		groups[i] = ToGroupEntity(groupDTO)
	}
	return groups
}
