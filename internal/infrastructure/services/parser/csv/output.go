package csv

import (
	"challenge2019/internal/domain/output"
	"github.com/gocarina/gocsv"
	"os"
)

func SaveOutput(fileName string, data []*output.Model) error {
	csvContent, err := gocsv.MarshalString(&data)
	if err != nil {
		return err
	}

	if err := os.WriteFile(fileName, []byte(csvContent), os.ModePerm); err != nil {
		return err
	}

	return nil
}
