package parcer

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Record struct {
	Name      string  `csv:"Name"`
	Address   string  `csv:"Address"`
	Phone     string  `csv:"Phone"`
	Longitude float64 `csv:"Longitude"`
	Latitude  float64 `csv:"Latitude"`
}

func Parcer(filePath string) ([]Record, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var result []Record
	for _, row := range records[1:] {
		longitude, _ := strconv.ParseFloat(row[4], 64)
		latitude, _ := strconv.ParseFloat(row[5], 64)

		record := Record{
			Name:      row[1],
			Address:   row[2],
			Phone:     row[3],
			Longitude: longitude,
			Latitude:  latitude,
		}
		result = append(result, record)
	}
	return result, nil
}
