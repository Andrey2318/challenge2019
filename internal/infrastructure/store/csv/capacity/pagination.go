package capacity

import (
	"challenge2019/internal/domain/capacity"
)

func pagination(data []*model, condition *capacity.Condition) []*model {
	if condition == nil {
		return data
	}
	if len(data) == 0 {
		return data
	}
	if condition.Limit == 0 && condition.Offset == 0 {
		return data
	}

	limit := condition.Limit
	if condition.Limit > uint64(len(data)) {
		condition.Limit = uint64(len(data))
	}
	if condition.Offset > uint64(len(data)) {
		return nil
	}

	return data[condition.Offset:limit]
}
