package csv

import (
	"challenge2019/internal/domain/partner"
	"github.com/gocarina/gocsv"
	"os"
	"strconv"
	"strings"
)

type partnerModel struct {
	Theatre        string `csv:"Theatre"`
	PartnerID      string `csv:"Partner ID"`
	MinMaxSizeSlab string `csv:"Size Slab (in GB)"`
	MinimumCost    uint32 `csv:"Minimum cost"`
	CostPerGB      uint32 `csv:"Cost Per GB"`
}

func ParsePartners(FileName string) ([]*partner.Model, error) {
	file, err := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	arr := make([]*partnerModel, 0)

	if err := gocsv.UnmarshalFile(file, &arr); err != nil { // Load clients from file
		panic(err)
	}

	d := make([]*partner.Model, len(arr))

	for i, m := range arr {
		d[i] = &partner.Model{
			Theatre:     strings.TrimSpace(m.Theatre),
			PartnerID:   strings.TrimSpace(m.PartnerID),
			MinimumCost: m.MinimumCost,
			CostPerGB:   m.CostPerGB,
		}
		SizeSlab := strings.Split(m.MinMaxSizeSlab, "-")
		if len(SizeSlab) == 2 {
			min, err := strconv.ParseUint(strings.TrimSpace(SizeSlab[0]), 10, 32)
			if err != nil {
				return nil, err
			}

			d[i].MinSizeSlab = uint32(min)

			max, err := strconv.ParseUint(strings.TrimSpace(SizeSlab[1]), 10, 32)
			if err != nil {
				return nil, err
			}

			d[i].MaxSizeSlab = uint32(max)
		}
	}

	return d, nil
}
