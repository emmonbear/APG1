package wc

import (
	"flag"
	"reflect"
	"testing"
)

type ParserTest struct {
	name     string
	options  []string
	expected *WCFlags
}

type WCTest struct {
	name     string
	filename string
	options  *WCFlags
	expected int
}

var parserTests = []ParserTest{
	{name: "Without flags", options: []string{"test.txt"}, expected: &WCFlags{Words: true}},
	{name: "With flag -l", options: []string{"-l", "test.txt"}, expected: &WCFlags{Lines: true}},
	{name: "With flag -w", options: []string{"-w", "test.txt"}, expected: &WCFlags{Words: true}},
	{name: "With flag -m", options: []string{"-m", "test.txt"}, expected: &WCFlags{Chars: true}},
	{name: "With flag -m and many files", options: []string{"-m", "test1.txt", "test2.txt", "test3.txt"}, expected: &WCFlags{Chars: true}},
}

var parserEdgeTests = []ParserTest{
	{name: "With -l and -m", options: []string{"-l", "-m", "test.txt"}},
	{name: "With incorrect flag", options: []string{"-ls", "test.txt"}},
	{name: "Without txt file", options: []string{"-l"}},
}

var wcTests = []WCTest{
	{name: "count lines in test_1.txt", filename: "../../test/files/test_1.txt", options: &WCFlags{Lines: true}, expected: 40},
	{name: "count words in test_1.txt", filename: "../../test/files/test_1.txt", options: &WCFlags{Words: true}, expected: 404},
	{name: "count chars in test_1.txt", filename: "../../test/files/test_1.txt", options: &WCFlags{Chars: true}, expected: 2876},
	{name: "count lines in test_2.txt", filename: "../../test/files/test_2.txt", options: &WCFlags{Lines: true}, expected: 26},
	{name: "count words in test_2.txt", filename: "../../test/files/test_2.txt", options: &WCFlags{Words: true}, expected: 270},
	{name: "count chars in test_2.txt", filename: "../../test/files/test_2.txt", options: &WCFlags{Chars: true}, expected: 1793},
	{name: "count lines in test_3.txt", filename: "../../test/files/test_3.txt", options: &WCFlags{Lines: true}, expected: 9},
	{name: "count words in test_3.txt", filename: "../../test/files/test_3.txt", options: &WCFlags{Words: true}, expected: 143},
	{name: "count chars in test_3.txt", filename: "../../test/files/test_3.txt", options: &WCFlags{Chars: true}, expected: 718},
	{name: "count lines in test_4.txt", filename: "../../test/files/test_4.txt", options: &WCFlags{Lines: true}, expected: 19},
	{name: "count words in test_4.txt", filename: "../../test/files/test_4.txt", options: &WCFlags{Words: true}, expected: 171},
	{name: "count chars in test_4.txt", filename: "../../test/files/test_4.txt", options: &WCFlags{Chars: true}, expected: 1196},
	{name: "count lines in test_5.txt", filename: "../../test/files/test_5.txt", options: &WCFlags{Lines: true}, expected: 27},
	{name: "count words in test_5.txt", filename: "../../test/files/test_5.txt", options: &WCFlags{Words: true}, expected: 1143},
	{name: "count chars in test_5.txt", filename: "../../test/files/test_5.txt", options: &WCFlags{Chars: true}, expected: 6557},
	{name: "count lines in test_6.txt", filename: "../../test/files/test_6.txt", options: &WCFlags{Lines: true}, expected: 25},
	{name: "count words in test_6.txt", filename: "../../test/files/test_6.txt", options: &WCFlags{Words: true}, expected: 0},
	{name: "count chars in test_6.txt", filename: "../../test/files/test_6.txt", options: &WCFlags{Chars: true}, expected: 24},
}

func TestParserWCFlags(t *testing.T) {
	for _, tt := range parserTests {
		t.Run(tt.name, func(t *testing.T) {
			testParcerWCFlags(t, tt.name, tt.options, tt.expected)
		})
	}
}

func TestParserWCFlagsEdge(t *testing.T) {
	for _, tt := range parserEdgeTests {
		t.Run(tt.name, func(t *testing.T) {
			testParcerWCFlagEdges(t, tt.name, tt.options)
		})
	}
}

func TestWC(t *testing.T) {
	for _, tt := range wcTests {
		t.Run(tt.name, func(t *testing.T) {
			testWC(t, tt.filename, tt.options, tt.expected)
		})
	}
}

func testParcerWCFlags(t *testing.T, name string, options []string, expected *WCFlags) {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	result := NewWCFlags()
	err := result.ParseFlags(fs, options)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func testParcerWCFlagEdges(t *testing.T, name string, options []string) {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	result := NewWCFlags()
	err := result.ParseFlags(fs, options)
	if err == nil {
		t.Fatal(err)
	}

}

func testWC(t *testing.T, filename string, options *WCFlags, expected int) {
	result, err := WC(filename, options)
	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
