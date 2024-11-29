// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbprinter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"path/filepath"
)

// DBPrinter is an interface for printing data to a writer.
type DBPrinter interface {
	// Print writes the data to the provided writer in a specific format.
	Print(w io.Writer, data interface{}) error
}

// XMLPrinter is a struct that implements the DBPrinter interface for XML format.
type XMLPrinter struct{}

// Print writes the data to the provided writer in XML format.
func (p *XMLPrinter) Print(w io.Writer, data interface{}) error {
	output, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	fmt.Fprintln(w, string(output))
	return nil
}

// JSONPrinter is a struct that implements the DBPrinter interface for JSON format.
type JSONPrinter struct{}

// Print writes the data to the provided writer in JSON format.
func (p *JSONPrinter) Print(w io.Writer, data interface{}) error {
	output, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	fmt.Fprintln(w, string(output))
	return nil
}

// GetDBPrinter returns a DBPrinter implementation based on the file extension.
func GetDBPrinter(filename string) DBPrinter {
	ext := filepath.Ext(filename)
	switch ext {
	case ".xml":
		return &XMLPrinter{}
	case ".json":
		return &JSONPrinter{}
	default:
		log.Printf("Unsupported file extension: %s", ext)
		return nil
	}
}
