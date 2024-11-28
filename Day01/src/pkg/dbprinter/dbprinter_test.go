package dbprinter

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"testing"
)

func TestGetDBPrinter(t *testing.T) {
	tests := []struct {
		filename string
		expected DBPrinter
	}{
		{
			filename: "test.xml",
			expected: &XMLPrinter{},
		},
		{
			filename: "test.json",
			expected: &JSONPrinter{},
		},
		{
			filename: "test.txt",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			reader := GetDBPrinter(tt.filename)
			if reader == nil && tt.expected != nil {
				t.Fatalf("Expected non-nil reader, got nil")
			}
			if reader != nil && tt.expected == nil {
				t.Fatalf("Expected nil reader, got non-nil")
			}
			if reader != nil && tt.expected != nil {
				switch reader.(type) {
				case *XMLPrinter:
					if _, ok := tt.expected.(*XMLPrinter); !ok {
						t.Fatalf("Expected %T, got %T", tt.expected, reader)
					}
				case *JSONPrinter:
					if _, ok := tt.expected.(*JSONPrinter); !ok {
						t.Fatalf("Expected %T, got %T", tt.expected, reader)
					}
				default:
					t.Fatalf("Unexpected reader type: %T", reader)
				}
			}
		})
	}
}

func TestPrintJSON(t *testing.T) {
	printer := &JSONPrinter{}
	data := map[string]string{"key": "value"}
	expectedOutput := `{
    "key": "value"
}
`
	var buf bytes.Buffer
	err := printer.Print(&buf, data)
	if err != nil {
		t.Fatalf("expected no error %v", err)
	}

	if buf.String() != expectedOutput {
		t.Errorf("expected output %q, got %q", expectedOutput, buf.String())
	}
}

func TestPrintJSONMarshalError(t *testing.T) {
	printer := &JSONPrinter{}
	data := make(chan int)

	var buf bytes.Buffer
	err := printer.Print(&buf, data)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var jsonErr *json.UnsupportedTypeError
	if !errors.As(err, &jsonErr) {
		t.Errorf("expected json.UnsupportedTypeError, got %v", err)
	}
}

type Data struct {
	Key string `xml:"key"`
}

func TestPrintXML(t *testing.T) {
	printer := &XMLPrinter{}
	data := Data{Key: "value"}
	expectedOutput := `<Data>
    <key>value</key>
</Data>
`
	var buf bytes.Buffer
	err := printer.Print(&buf, data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if buf.String() != expectedOutput {
		t.Errorf("expected output %q, got %q", expectedOutput, buf.String())
	}
}

func TestPrintXMLMarshalError(t *testing.T) {
	printer := &XMLPrinter{}
	data := make(chan int)

	var buf bytes.Buffer
	err := printer.Print(&buf, data)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var xmlErr *xml.UnsupportedTypeError
	if !errors.As(err, &xmlErr) {
		t.Errorf("expected xml.UnsupportedTypeError, got %v", err)
	}
}
