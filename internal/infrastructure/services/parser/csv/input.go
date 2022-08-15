package csv

import (
	"challenge2019/internal/domain/input"
	"github.com/gocarina/gocsv"
	"os"
)

func ParseInput(FileName string) ([]*input.Model, error) {
	file, err := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	arr := make([]*input.Model, 0)

	if err := gocsv.UnmarshalFile(file, &arr); err != nil { // Load clients from file
		panic(err)
	}

	return arr, nil
}
