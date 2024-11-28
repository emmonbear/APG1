package main

import (
	"flag"

	"github.com/emmonbear/APG1/Day01.git/pkg/fscompare"
)

type DumpFiles struct {
	oldFile string
	newFile string
}

func main() {
	files := parseFlags()
	fscompare.NewFSComparer().CompareDumps(files.newFile, files.oldFile)

}

func parseFlags() DumpFiles {
	oldFile := flag.String("old", "", "Path to old file")
	newFile := flag.String("new", "", "Path to new file")
	flag.Parse()
	return DumpFiles{*oldFile, *newFile}
}
