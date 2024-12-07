package xargs

import (
	"reflect"
	"strings"
	"testing"
)

type ParcerCommandLineTest struct {
	name     string
	input    string
	expected []string
}

var parcerCommandLineTests = []ParcerCommandLineTest{
	{name: "echo -e \"1\n2\n3\"", input: "1\n2\n3", expected: []string{"1", "2", "3"}},
	{name: "echo -e \"1 2 3\"", input: "1 2 3", expected: []string{"1", "2", "3"}},
	{name: "echo -e \"/a\n/b\n/c\"", input: "/a\n/b\n/c", expected: []string{"/a", "/b", "/c"}},
	{name: "echo -e \"a b c d e\"", input: "a b c d e", expected: []string{"a", "b", "c", "d", "e"}},
	{name: "echo -e \"1 2\n3 4 5\"", input: "1 2\n3 4 5", expected: []string{"1", "2", "3", "4", "5"}},
	{name: "echo -e \"a\n\nb\nc\"", input: "a\n\nb\nc", expected: []string{"a", "b", "c"}},
	{name: "echo -e \"1 2 3 4 5\"", input: "1 2 3 4 5", expected: []string{"1", "2", "3", "4", "5"}},
	{name: "echo -e \"1 2\n3 4\n5\"", input: "1 2\n3 4\n5", expected: []string{"1", "2", "3", "4", "5"}},
	{name: "echo -e \"abc def ghi\"", input: "abc def ghi", expected: []string{"abc", "def", "ghi"}},
	{name: "echo -e \"abc\ndef\nghi\"", input: "abc\ndef\nghi", expected: []string{"abc", "def", "ghi"}},
	{name: "echo -e \"1\n2\n3\n4\n5\"", input: "1\n2\n3\n4\n5", expected: []string{"1", "2", "3", "4", "5"}},
	{name: "echo -e \"1\n\n2\n3\"", input: "1\n\n2\n3", expected: []string{"1", "2", "3"}},
	{name: "echo -e \"a b\nc d e\"", input: "a b\nc d e", expected: []string{"a", "b", "c", "d", "e"}},
	{name: "echo -e \"a\nb\nc\n\"", input: "a\nb\nc\n", expected: []string{"a", "b", "c"}},
	{name: "echo -e \"a b c\n\n  d e f\"", input: "a b c\n\n  d e f", expected: []string{"a", "b", "c", "d", "e", "f"}},
	{name: "echo -e \"1\n   2\n3\"", input: "1\n   2\n3", expected: []string{"1", "2", "3"}},
	{name: "echo -e \"abc\n\n   def ghi\"", input: "abc\n\n   def ghi", expected: []string{"abc", "def", "ghi"}},
	{name: "echo -e \"x y z\n123\"", input: "x y z\n123", expected: []string{"x", "y", "z", "123"}},
	{name: "echo -e \"hello world\nfoo bar\nbaz\"", input: "hello world\nfoo bar\nbaz", expected: []string{"hello", "world", "foo", "bar", "baz"}},
	{name: "echo -e \"one\ntwo\nthree four\"", input: "one\ntwo\nthree four", expected: []string{"one", "two", "three", "four"}},
	{name: "echo -e \"1 2 3\n4 5\"", input: "1 2 3\n4 5", expected: []string{"1", "2", "3", "4", "5"}},
	{name: "echo -e \"alpha beta gamma\n\ndelta\"", input: "alpha beta gamma\n\ndelta", expected: []string{"alpha", "beta", "gamma", "delta"}},
	{name: "echo -e \"123 abc 456\n789 def\"", input: "123 abc 456\n789 def", expected: []string{"123", "abc", "456", "789", "def"}},
	{name: "echo -e \"line1\nline2\nline3 line4\"", input: "line1\nline2\nline3 line4", expected: []string{"line1", "line2", "line3", "line4"}},
}

func TestParcerCommandLine(t *testing.T) {
	for _, tt := range parcerCommandLineTests {
		t.Run(tt.name, func(t *testing.T) {
			testParcerCommandLine(t, tt.input, tt.expected)
		})
	}
}

func testParcerCommandLine(t *testing.T, in string, expected []string) {
	result := New()
	err := result.ParseCommandLine(strings.NewReader(in))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(result.InputArgs, expected) {
		t.Fatalf("expected %v, got %v", expected, result.InputArgs)
	}
}
