// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package xargs

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

type ParcerCommandLineTest struct {
	name     string
	input    string
	expected []string
}

type ExecuteTest struct {
	name      string
	command   string
	baseArgs  []string
	inputArgs string
	expected  string
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

var executeTests = []ExecuteTest{
	{
		name:     "find -type f -name \"*.txt\" ./ | ./myXargs wc -m",
		command:  "wc",
		baseArgs: []string{"-m"},
		inputArgs: strings.Join([]string{
			"../../test/files/test_3.txt",
			"../../test/files/test_4.txt",
			"../../test/files/test_6.txt",
			"../../test/files/test_5.txt",
			"../../test/files/test_2.txt",
			"../../test/files/test_1.txt",
		}, " "),
		expected: "  718 ../../test/files/test_3.txt\n" +
			" 1196 ../../test/files/test_4.txt\n" +
			"   25 ../../test/files/test_6.txt\n" +
			" 6557 ../../test/files/test_5.txt\n" +
			" 1793 ../../test/files/test_2.txt\n" +
			" 2876 ../../test/files/test_1.txt\n" +
			"13165 total\n",
	},
	{
		name:     "find -type f -name \"*.txt\" ./ | ./myXargs wc -l",
		command:  "wc",
		baseArgs: []string{"-l"},
		inputArgs: strings.Join([]string{
			"../../test/files/test_3.txt",
			"../../test/files/test_4.txt",
			"../../test/files/test_6.txt",
			"../../test/files/test_5.txt",
			"../../test/files/test_2.txt",
			"../../test/files/test_1.txt",
		}, " "),
		expected: "    8 ../../test/files/test_3.txt\n" +
			"   18 ../../test/files/test_4.txt\n" +
			"   25 ../../test/files/test_6.txt\n" +
			"   26 ../../test/files/test_5.txt\n" +
			"   25 ../../test/files/test_2.txt\n" +
			"   39 ../../test/files/test_1.txt\n" +
			"  141 total\n",
	},
	{
		name:     "find -type f -name \"*.txt\" ./ | ./myXargs wc -w",
		command:  "wc",
		baseArgs: []string{"-w"},
		inputArgs: strings.Join([]string{
			"../../test/files/test_3.txt",
			"../../test/files/test_4.txt",
			"../../test/files/test_6.txt",
			"../../test/files/test_5.txt",
			"../../test/files/test_2.txt",
			"../../test/files/test_1.txt",
		}, " "),
		expected: "  143 ../../test/files/test_3.txt\n" +
			"  171 ../../test/files/test_4.txt\n" +
			"    0 ../../test/files/test_6.txt\n" +
			" 1143 ../../test/files/test_5.txt\n" +
			"  270 ../../test/files/test_2.txt\n" +
			"  404 ../../test/files/test_1.txt\n" +
			" 2131 total\n",
	},
}

func TestParcerCommandLine(t *testing.T) {
	for _, tt := range parcerCommandLineTests {
		t.Run(tt.name, func(t *testing.T) {
			testParcerCommandLine(t, tt.input, tt.expected)
		})
	}
}

func TestExecute(t *testing.T) {
	for _, tt := range executeTests {
		t.Run(tt.name, func(t *testing.T) {
			testExecute(t, tt.command, tt.baseArgs, tt.inputArgs, tt.expected)
		})
	}
}

func testParcerCommandLine(t *testing.T, in string, expected []string) {
	result := New("", nil)
	err := result.ParseCommandLine(strings.NewReader(in))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(result.InputArgs, expected) {
		t.Fatalf("expected %v, got %v", expected, result.InputArgs)
	}
}

func testExecute(t *testing.T, command string, baseArgs []string, inputArgs string, expected string) {
	args := New(command, baseArgs)

	if err := args.ParseCommandLine(strings.NewReader(inputArgs)); err != nil {
		t.Fatalf("Expected not error, got %v", err)
	}

	var result bytes.Buffer
	args.Execute(&result)

	output := result.String()

	if output != expected {
		t.Fatalf("Expected \n%v\n, got \n%v\n", expected, output)
	}
}
