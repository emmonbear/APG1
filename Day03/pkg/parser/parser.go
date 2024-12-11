// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package parcer provides functionality to parse CSV files into a structured format.
package parser

import (
	"encoding/csv"
	"os"
	"strconv"
)

// Record represents a single entry in the CSV file.
type Record struct {
	Name      string  `csv:"Name"`
	Address   string  `csv:"Address"`
	Phone     string  `csv:"Phone"`
	Longitude float64 `csv:"Longitude"`
	Latitude  float64 `csv:"Latitude"`
}

// Parcer reads a CSV file from the given file path and returns a slice of Record structs.
// The CSV file is expected to have a header row and use tab-separated values.
// It returns an error if the file cannot be opened or read.
func Parser(filePath string) ([]Record, error) {
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
