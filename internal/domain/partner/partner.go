package partner

import "context"

type Model struct {
	Theatre     string
	PartnerID   string
	MinSizeSlab uint32
	MaxSizeSlab uint32
	MinimumCost uint32
	CostPerGB   uint32
}

type Condition struct {
	Theatres   []string
	PartnerIDs []string
	Limit      uint64
	Offset     uint64
}

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	First(ctx context.Context, condition *Condition) (*Model, error)
	Find(ctx context.Context, condition *Condition) ([]*Model, error)
	Count(ctx context.Context, condition *Condition) (uint64, error)
	Save(ctx context.Context, data ...*Model) error
}
