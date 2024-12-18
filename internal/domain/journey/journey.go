package journey

import (
	"ev-pooling-test-api/internal/domain/group"
	"ev-pooling-test-api/internal/domain/vehicle"
)

type Journey struct {
	Vehicle *vehicle.Vehicle
	Group   *group.Group
}
