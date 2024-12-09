// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package finder

import (
	"flag"
	"reflect"
	"testing"
)

func TestFindWithoutFlags(t *testing.T) {
	root := "../../"
	options := Options{true, true, true, ""}
	result, _ := Find(root, options)
	expected := []Entry{
		{Path: "../../", Type: Directory, Link: ""},
		{Path: "../../.gitignore", Type: File, Link: ""},
		{Path: "../../.gitkeep", Type: File, Link: ""},
		{Path: "../../Makefile", Type: File, Link: ""},
		{Path: "../../c.out", Type: File, Link: ""},
		{Path: "../../cmd", Type: Directory, Link: ""},
		{Path: "../../cmd/ex00", Type: Directory, Link: ""},
		{Path: "../../cmd/ex00/main.go", Type: File, Link: ""},
		{Path: "../../cmd/ex01", Type: Directory, Link: ""},
		{Path: "../../cmd/ex01/main.go", Type: File, Link: ""},
		{Path: "../../cmd/ex02", Type: Directory, Link: ""},
		{Path: "../../cmd/ex02/main.go", Type: File, Link: ""},
		{Path: "../../cmd/ex03", Type: Directory, Link: ""},
		{Path: "../../cmd/ex03/main.go", Type: File, Link: ""},
		{Path: "../../go.mod", Type: File, Link: ""},
		{Path: "../../pkg", Type: Directory, Link: ""},
		{Path: "../../pkg/finder", Type: Directory, Link: ""},
		{Path: "../../pkg/finder/finder.go", Type: File, Link: ""},
		{Path: "../../pkg/finder/finder_test.go", Type: File, Link: ""},
		{Path: "../../pkg/rotate", Type: Directory, Link: ""},
		{Path: "../../pkg/rotate/rotate.go", Type: File, Link: ""},
		{Path: "../../pkg/rotate/rotate_test.go", Type: File, Link: ""},
		{Path: "../../pkg/wc", Type: Directory, Link: ""},
		{Path: "../../pkg/wc/wc.go", Type: File, Link: ""},
		{Path: "../../pkg/wc/wc_test.go", Type: File, Link: ""},
		{Path: "../../pkg/xargs", Type: Directory, Link: ""},
		{Path: "../../pkg/xargs/xargs.go", Type: File, Link: ""},
		{Path: "../../pkg/xargs/xargs_test.go", Type: File, Link: ""},
		{Path: "../../test", Type: Directory, Link: ""},
		{Path: "../../test/broken_link", Type: Symlink, Link: "[broken]"},
		{Path: "../../test/correct_link", Type: Symlink, Link: "../cmd/ex00/main.go"},
		{Path: "../../test/files", Type: Directory, Link: ""},
		{Path: "../../test/files/test_1.txt", Type: File, Link: ""},
		{Path: "../../test/files/test_2.txt", Type: File, Link: ""},
		{Path: "../../test/files/test_3.txt", Type: File, Link: ""},
		{Path: "../../test/files/test_4.txt", Type: File, Link: ""},
		{Path: "../../test/files/test_5.txt", Type: File, Link: ""},
		{Path: "../../test/files/test_6.txt", Type: File, Link: ""},
	}

	if len(result) != len(expected) {
		t.Fatalf("Lengths do not match: got %d, expected %d", len(result), len(expected))
	}

	for i := range result {
		if !reflect.DeepEqual(result[i], expected[i]) {
			t.Errorf("Mismatch at index %d:\nGot:      %+v\nExpected: %+v", i, result[i], expected[i])
		}
	}

}

func TestParseFlagsWithExtWithoutFile(t *testing.T) {
	args := []string{"-ext", "txt", "./path"}
	fs := flag.NewFlagSet("TestParseFlagsWithExtAndFile", flag.ContinueOnError)
	options := NewOptions()
	err := options.ParseFlags(fs, args)
	if err == nil {
		t.Fatalf("Error during flag parsing: %v", err)
	}
}

func TestParseFlags(t *testing.T) {
	tests := []struct {
		fileSetName string
		args        []string
		expected    *Options
	}{
		{
			fileSetName: "default",
			args:        []string{"./"},
			expected: &Options{
				IncludeFiles:       true,
				IncludeDirectories: true,
				IncludeSymlinks:    true,
			},
		},
		{
			fileSetName: "1",
			args:        []string{"-f", "./path"},
			expected: &Options{
				IncludeFiles: true,
			},
		},
		{
			fileSetName: "2",
			args:        []string{"-d", "./path"},
			expected: &Options{
				IncludeDirectories: true,
			},
		},
		{
			fileSetName: "3",
			args:        []string{"-f", "-ext", "txt", "./path"},
			expected: &Options{
				IncludeFiles:    true,
				ExtensionFilter: "txt",
			},
		},
		{
			fileSetName: "4",
			args:        []string{"-f", "-ext", "go", "-d", "./path"},
			expected: &Options{
				IncludeFiles:       true,
				ExtensionFilter:    "go",
				IncludeDirectories: true,
			},
		},
		{
			fileSetName: "5",
			args:        []string{"-f", "-sl", "-ext", "go", "-d", "./path"},
			expected: &Options{
				IncludeFiles:       true,
				ExtensionFilter:    "go",
				IncludeDirectories: true,
				IncludeSymlinks:    true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.fileSetName, func(t *testing.T) {
			fs := flag.NewFlagSet(tt.fileSetName, flag.ContinueOnError)
			options := NewOptions()
			err := options.ParseFlags(fs, tt.args)
			if err != nil {
				t.Fatalf("Expected non-nil reader, got nil")
			}
			if !reflect.DeepEqual(options, tt.expected) {
				t.Fatalf("Expected %+v, got %+v", tt.expected, options)
			}
		})
	}

}
