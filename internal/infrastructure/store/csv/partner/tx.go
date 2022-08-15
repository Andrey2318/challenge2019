package partner

import (
	"challenge2019/internal/domain/partner"
	"context"
	"github.com/patrickmn/go-cache"
)

func (r *Repository) Save(ctx context.Context, data ...*partner.Model) error {
	mutex.Lock()
	defer mutex.Unlock()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	db := r.db()
	db = append(db, domainToModel(data)...)
	r._db.Set(table, db, cache.DefaultExpiration)
	return nil
}
