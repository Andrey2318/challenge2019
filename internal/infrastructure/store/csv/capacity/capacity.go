package capacity

import (
	"challenge2019/internal/domain/capacity"
	"github.com/patrickmn/go-cache"
	"sync"
)

const table = "capacities"

var mutex sync.Mutex

type Repository struct {
	_db *cache.Cache
}

func New(_db *cache.Cache) *Repository {
	data := make([]*model, 0)
	_db.Set(table, data, cache.NoExpiration)
	return &Repository{_db: _db}
}

func (r *Repository) db() []*model {
	data, ok := r._db.Get(table)
	if !ok {
		return nil
	}
	n := make([]*model, len(data.([]*model)))
	copy(n, data.([]*model))
	return n
}

type model struct {
	PartnerID string
	Capacity  uint32
}

func (m *model) Get() *capacity.Model {
	return &capacity.Model{
		Capacity:  m.Capacity,
		PartnerID: m.PartnerID,
	}
}

func (m *model) Set(data *capacity.Model) *model {
	m.Capacity = data.Capacity
	m.PartnerID = data.PartnerID
	return m
}

func domainToModel(data []*capacity.Model) []*model {
	arr := make([]*model, len(data))
	for i := range arr {
		arr[i] = (&model{}).Set(data[i])
	}
	return arr
}

func modelToDomain(data []*model) []*capacity.Model {
	arr := make([]*capacity.Model, len(data))
	for i := range arr {
		arr[i] = data[i].Get()
	}
	return arr
}
