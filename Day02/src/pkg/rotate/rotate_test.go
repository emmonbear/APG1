// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package rotate

import (
	"os"
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

type RotateTest struct {
	name          string
	fileNames     []string
	archiveDir    string
	expectedError bool
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
			ArchiveDir: "."},
	},
	{
		name: "./myRotate path log1.log log2.log",
		args: []string{"path", "log1.log", "log2.log"},
		expected: &Config{
			FileNames:  []string{"path", "log1.log", "log2.log"},
			ArchiveDir: "."},
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

var rotateTests = []RotateTest{
	{
		name:          "single file rotation",
		fileNames:     []string{"file1.log"},
		archiveDir:    "./archives",
		expectedError: false,
	},
}

func TestParserFlags(t *testing.T) {
	for _, tt := range parserTests {
		t.Run(tt.name, func(t *testing.T) {
			testParserFlags(t, tt.args, tt.expected)
		})
	}
}

func TestParserFlagsEdge(t *testing.T) {
	for _, tt := range parserEdgeTests {
		t.Run(tt.name, func(t *testing.T) {
			testParserFlagsEdge(t, tt.args)
		})
	}
}

func TestRotate(t *testing.T) {
	for _, tt := range rotateTests {
		t.Run(tt.name, func(t *testing.T) {
			testRotate(t, tt)
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

func testParserFlagsEdge(t *testing.T, args []string) {
	_, err := ParseFlags(args)
	if err == nil {
		t.Fatalf("expected %v, got nil", err)
	}
}

func testRotate(t *testing.T, tt RotateTest) {
	var tempFiles []string
	for range tt.fileNames {
		filePath := createTempFile(t, "test content")
		tempFiles = append(tempFiles, filePath)
	}
	defer func() {
		for _, file := range tempFiles {
			os.Remove(file)
		}
		os.RemoveAll(tt.archiveDir)
	}()
	if err := os.MkdirAll(tt.archiveDir, 0755); err != nil {
		t.Fatalf("failed to create archive dir: %v", err)
	}

	config := &Config{
		FileNames:  tempFiles,
		ArchiveDir: tt.archiveDir,
	}

	err := Rotate(config)

	if (err != nil) != tt.expectedError {
		t.Fatalf("unexpected error: %v", err)
	}

}

func createTempFile(t *testing.T, content string) string {
	t.Helper()
	tmpFile, err := os.CreateTemp("", "test_file_*.log")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer tmpFile.Close()

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	return tmpFile.Name()
}
