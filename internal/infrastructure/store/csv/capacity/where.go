package capacity

import (
	"challenge2019/internal/domain/capacity"
	"challenge2019/pkg/other"
)

func where(data []*model, condition *capacity.Condition) []*model {
	if condition == nil {
		return nil
	}

	if len(condition.PartnerIDs) > 0 {
		data = other.Filter[*model](data, func(data *model) bool {
			for i := range condition.PartnerIDs {
				if condition.PartnerIDs[i] == data.PartnerID {
					return true
				}
			}
			return false
		})
	}

	return data
}
