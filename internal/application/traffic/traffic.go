package traffic

import (
	"challenge2019/internal/domain/capacity"
	"challenge2019/internal/domain/partner"
	"context"
	"sort"
)

type UseCase struct {
	partnerRep  partner.Repository
	capacityRep capacity.Repository
}

func New(partnerRep partner.Repository, capacityRep capacity.Repository) *UseCase {
	return &UseCase{partnerRep: partnerRep, capacityRep: capacityRep}
}

func (uc *UseCase) Statement1(ctx context.Context, req []*StatementRequest) ([]*StatementResponse, error) {
	data := make([]*StatementResponse, len(req))
	for i := range req {
		mdata, err := uc.partnerRep.Find(ctx, &partner.Condition{
			Theatres: []string{req[i].Theatre},
		})
		if err != nil {
			return nil, err
		}
		data[i] = &StatementResponse{
			DeliveryID: req[i].DeliveryID,
			Status:     false,
			Cost:       0,
		}
		cost := uint32(0)
		for j := range mdata {
			if req[i].Size >= mdata[j].MinSizeSlab && req[i].Size <= mdata[j].MaxSizeSlab {
				if mdata[j].CostPerGB < cost || cost == 0 {
					cost = mdata[j].CostPerGB
					data[i].Cost = mdata[j].CostPerGB * req[i].Size
					if data[i].Cost < mdata[j].MinimumCost {
						data[i].Cost = mdata[j].MinimumCost
					}
					data[i].Status = true
					data[i].PartnerID = mdata[j].PartnerID
				}
			}
		}
	}

	return data, nil
}

type StatementRequest struct {
	DeliveryID string
	Size       uint32
	Theatre    string
}

type StatementResponse struct {
	DeliveryID string
	Status     bool
	PartnerID  string
	Cost       uint32
}

func (uc *UseCase) Statement2(ctx context.Context, req []*StatementRequest) ([]*StatementResponse, error) {
	data := make([]*StatementResponse, len(req))
	sort.Slice(req, func(i, j int) bool {
		if req[i].Size > req[j].Size {
			return true
		}
		return false
	})
	for i := range req {
		mdata, err := uc.partnerRep.Find(ctx, &partner.Condition{
			Theatres: []string{req[i].Theatre},
		})
		if err != nil {
			return nil, err
		}

		sort.Slice(mdata, func(i, j int) bool {
			if mdata[i].CostPerGB < mdata[j].CostPerGB {
				return true
			}
			return false
		})

		data[i] = &StatementResponse{
			DeliveryID: req[i].DeliveryID,
			Status:     false,
			Cost:       0,
		}
		partners := make([]struct {
			part     *partner.Model
			capacity *capacity.Model
		}, 0)
		minCost := uint32(0)
		avgCost := float64(0)
		if len(mdata) > 0 {
			minCost = mdata[0].CostPerGB
		} else {
			continue
		}

		for j := range mdata {
			if req[i].Size >= mdata[j].MinSizeSlab && req[i].Size <= mdata[j].MaxSizeSlab {
				model, err := uc.capacityRep.First(ctx, &capacity.Condition{PartnerIDs: []string{mdata[j].PartnerID}})
				if err != nil {
					return nil, err
				}
				if model == nil {
					continue
				}
				if model.Capacity < req[i].Size {
					continue
				}
				if mdata[j].CostPerGB < minCost {
					minCost = mdata[j].CostPerGB
				}

				avgCost += float64(mdata[j].CostPerGB)

				partners = append(partners, struct {
					part     *partner.Model
					capacity *capacity.Model
				}{part: mdata[j], capacity: model})
			}
		}

		if len(partners) == 0 {
			continue
		}

		avgCost = avgCost / float64(len(partners))

		maxSize := req[i].Size
		sumSize := req[i].Size
		avgSize := float64(0)
		for j := range req {
			avgSize += float64(req[j].Size)
			if j != i {
				if req[i].Theatre == req[j].Theatre {
					if maxSize < req[j].Size {
						maxSize = req[j].Size
					}
					if maxSize < req[j].Size {
						maxSize = req[j].Size
					}
					sumSize += req[j].Size
				}
			}
		}

		avgSize = avgSize / float64(len(req))

		index := 0
		for j := range partners {
			par := partners[j]
			if sumSize <= par.capacity.Capacity && minCost == par.part.CostPerGB {
				index = j
				break
			}
			if req[i].Size == maxSize && par.capacity.Capacity >= maxSize && minCost == par.part.CostPerGB {
				index = j
				break
			} else {
				if req[i].Size <= par.capacity.Capacity && uint32(avgCost) >= par.part.CostPerGB {
					index = j
					break
				}
				index = j
			}
		}

		data[i].Cost = partners[index].part.CostPerGB * req[i].Size
		if data[i].Cost < partners[index].part.MinimumCost {
			data[i].Cost = partners[index].part.MinimumCost
		}
		data[i].Status = true
		data[i].PartnerID = partners[index].part.PartnerID
		partners[index].capacity.Capacity -= req[i].Size
		if err := uc.capacityRep.Save(ctx, partners[index].capacity); err != nil {
			return nil, err
		}
	}

	sort.Slice(data, func(i, j int) bool {
		if data[i].DeliveryID < data[j].DeliveryID {
			return true
		}
		return false
	})

	return data, nil
}
