package partner

import (
	"challenge2019/internal/domain/partner"
	"challenge2019/pkg/other"
)

func where(data []*model, condition *partner.Condition) []*model {
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

	if len(condition.Theatres) > 0 {
		data = other.Filter[*model](data, func(data *model) bool {
			for i := range condition.Theatres {
				if condition.Theatres[i] == data.Theatre {
					return true
				}
			}
			return false
		})
	}

	return data
}
