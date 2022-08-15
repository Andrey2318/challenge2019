package partner

import (
	"challenge2019/internal/domain/partner"
	"context"
)

func (r *Repository) First(ctx context.Context, condition *partner.Condition) (*partner.Model, error) {
	mutex.Lock()
	defer mutex.Unlock()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	data := where(r.db(), condition)

	if len(data) > 0 {
		return data[0].Get(), nil
	}

	return nil, nil
}

func (r *Repository) Find(ctx context.Context, condition *partner.Condition) ([]*partner.Model, error) {
	mutex.Lock()
	defer mutex.Unlock()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	data := pagination(where(r.db(), condition), condition)

	return modelToDomain(data), nil
}

func (r *Repository) Count(ctx context.Context, condition *partner.Condition) (uint64, error) {
	mutex.Lock()
	defer mutex.Unlock()
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}

	data := where(r.db(), condition)

	return uint64(len(data)), nil
}
