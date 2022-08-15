package csv

import (
	"challenge2019/internal/domain/capacity"
	"github.com/gocarina/gocsv"
	"os"
	"strings"
)

type capacityModel struct {
	PartnerID string `csv:"Partner ID"`
	Capacity  uint32 `csv:"Capacity (in GB)"`
}

func ParseCapacities(FileName string) ([]*capacity.Model, error) {
	file, err := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	arr := make([]*capacityModel, 0)

	if err := gocsv.UnmarshalFile(file, &arr); err != nil { // Load clients from file
		panic(err)
	}

	d := make([]*capacity.Model, len(arr))

	for i, m := range arr {
		d[i] = &capacity.Model{
			PartnerID: strings.TrimSpace(m.PartnerID),
			Capacity:  m.Capacity,
		}
	}

	return d, nil
}
