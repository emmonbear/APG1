package main

import (
	"flag"
	"log"

	"github.com/emmonbear/APG1/Day01.git/pkg/fscompare"
)

type DumpFiles struct {
	oldFile string
	newFile string
}

func main() {
	files := parseFlags()
	err := fscompare.NewFSComparer().CompareDumps(files.newFile, files.oldFile)
	if err != nil {
		log.Fatalf("failed to compare dumps: %v", err)
	}

}

func parseFlags() DumpFiles {
	oldFile := flag.String("old", "", "Path to old file")
	newFile := flag.String("new", "", "Path to new file")
	flag.Parse()
	return DumpFiles{*oldFile, *newFile}
}
