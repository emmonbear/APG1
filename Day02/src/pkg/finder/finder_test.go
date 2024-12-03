package finder

import (
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
		{Path: "../../cmd/ex00/main_test.go", Type: File, Link: ""},
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
		{Path: "../../test", Type: Directory, Link: ""},
		{Path: "../../test/broken_link", Type: Symlink, Link: "[broken]"},
		{Path: "../../test/correct_link", Type: Symlink, Link: "../cmd/ex00/main.go"},
	}

	if len(result) != len(expected) {
		t.Fatalf("Lengths do not match: got %d, expected %d", len(result), len(expected))
	}

	for i, str := range result {
		_ = FormatEntry(str)
		if !reflect.DeepEqual(result[i], expected[i]) {
			t.Errorf("Mismatch at index %d:\nGot:      %+v\nExpected: %+v", i, result[i], expected[i])
		}
	}

}
