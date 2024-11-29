package fscompare

import (
	"bytes"
	"os"
	"testing"
)

func TestCompareDumps(t *testing.T) {
	oldFile := "../../snapshot1.txt"
	newFile := "../../snapshot2.txt"
	comparer := NewFSComparer()

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	comparer.CompareDumps(newFile, oldFile)

	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	buf.ReadFrom(r)

	expectedOutput := `ADDED /etc/systemd/system/very_important/stash_location.jpg
REMOVED /var/log/browser_history.txt
`

	if buf.String() != expectedOutput {
		t.Errorf("expected output \n%q\n, got \n%q", expectedOutput, buf.String())
	}
}

func TestCompareDumpsFile1NotFound(t *testing.T) {
	oldFile := "snapshot1.txt"
	newFile := "snapshot2.txt"
	comparer := NewFSComparer()
	err := comparer.CompareDumps(oldFile, newFile)
	if err == nil {
		t.Fatalf("expected err, got nil")
	}
}

func TestCompareDumpsFile2NotFound(t *testing.T) {
	oldFile := "../../snapshot1.txt"
	newFile := "snapshot2.txt"
	comparer := NewFSComparer()
	err := comparer.CompareDumps(oldFile, newFile)
	if err == nil {
		t.Fatalf("expected err, got nil")
	}
}
