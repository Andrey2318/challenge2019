package partner

import (
	"challenge2019/internal/domain/partner"
	"github.com/patrickmn/go-cache"
	"sync"
)

const table = "partners"

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
	Theatre     string
	PartnerID   string
	MinSizeSlab uint32
	MaxSizeSlab uint32
	MinimumCost uint32
	CostPerGB   uint32
}

func (m *model) Get() *partner.Model {
	return &partner.Model{
		Theatre:     m.Theatre,
		PartnerID:   m.PartnerID,
		MinSizeSlab: m.MinSizeSlab,
		MaxSizeSlab: m.MaxSizeSlab,
		MinimumCost: m.MinimumCost,
		CostPerGB:   m.CostPerGB,
	}
}

func (m *model) Set(data *partner.Model) *model {
	m.Theatre = data.Theatre
	m.PartnerID = data.PartnerID
	m.MinSizeSlab = data.MinSizeSlab
	m.MaxSizeSlab = data.MaxSizeSlab
	m.MinimumCost = data.MinimumCost
	m.CostPerGB = data.CostPerGB
	return m
}

func domainToModel(data []*partner.Model) []*model {
	arr := make([]*model, len(data))
	for i := range arr {
		arr[i] = (&model{}).Set(data[i])
	}
	return arr
}

func modelToDomain(data []*model) []*partner.Model {
	arr := make([]*partner.Model, len(data))
	for i := range arr {
		arr[i] = data[i].Get()
	}
	return arr
}
