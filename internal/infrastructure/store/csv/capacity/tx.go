package capacity

import (
	"challenge2019/internal/domain/capacity"
	"challenge2019/pkg/other"
	"context"
	"github.com/patrickmn/go-cache"
)

func (r *Repository) Save(ctx context.Context, data ...*capacity.Model) error {
	mutex.Lock()
	defer mutex.Unlock()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	db := r.db()
	s := domainToModel(data)

rep:
	for i := range s {
		_, obj := other.Exist(db, func(d *model) bool {
			if s[i].PartnerID == d.PartnerID {
				return true
			}
			return false
		})
		if obj != nil {
			obj.Capacity = s[i].Capacity
			s = other.Remove(s, i)
			goto rep
		}

	}

	if len(s) > 0 {
		db = append(db, s...)
	}

	r._db.Set(table, db, cache.DefaultExpiration)
	return nil
}
