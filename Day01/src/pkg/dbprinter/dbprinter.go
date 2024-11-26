package dbprinter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
)

type DBPrinter interface {
	Print(w io.Writer, data interface{}) error
}

type XMLPrinter struct{}

func (p *XMLPrinter) Print(w io.Writer, data interface{}) error {
	output, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	fmt.Fprintln(w, string(output))
	return nil
}

type JSONPrinter struct{}

func (p *JSONPrinter) Print(w io.Writer, data interface{}) error {
	output, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	fmt.Fprintln(w, string(output))
	return nil
}
