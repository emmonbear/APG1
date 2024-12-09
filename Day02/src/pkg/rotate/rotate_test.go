package rotate

import (
	"reflect"
	"testing"
)

type ParserTest struct {
	name     string
	args     []string
	expected *Config
}

type ParserTestEdge struct {
	name string
	args []string
}

var parserTests = []ParserTest{
	{
		name: "./myRotate -a path log.log",
		args: []string{"-a", "path", "log.log"},
		expected: &Config{
			FileNames:  []string{"log.log"},
			ArchiveDir: "path"},
	},
	{
		name: "./myRotate log.log",
		args: []string{"log.log"},
		expected: &Config{
			FileNames:  []string{"log.log"},
			ArchiveDir: ""},
	},
	{
		name: "./myRotate path log1.log log2.log",
		args: []string{"path", "log1.log", "log2.log"},
		expected: &Config{
			FileNames:  []string{"path", "log1.log", "log2.log"},
			ArchiveDir: ""},
	},
}

var parserEdgeTests = []ParserTestEdge{
	{
		name: "./myRotate",
		args: nil,
	},
	{
		name: "./myRotate -s log.log",
		args: []string{"-s", "log.log"},
	},
}

func TestParserFlags(t *testing.T) {
	for _, tt := range parserTests {
		t.Run(tt.name, func(t *testing.T) {
			testParserFlags(t, tt.args, tt.expected)
		})
	}
}

func testParserFlags(t *testing.T, args []string, expected *Config) {
	result, err := ParseFlags(args)
	if err != nil {
		t.Fatalf("no errors expected, got %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestParserFlagsEdge(t *testing.T) {
	for _, tt := range parserEdgeTests {
		t.Run(tt.name, func(t *testing.T) {
			testParserFlagsEdge(t, tt.args)
		})
	}
}

func testParserFlagsEdge(t *testing.T, args []string) {
	_, err := ParseFlags(args)
	if err == nil {
		t.Fatalf("expected %v, got nil", err)
	}
}
