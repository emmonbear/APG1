package main

import (
	"flag"
	"fmt"
)

type DBFiles struct {
	oldFile string
	newFile string
}

func main() {
	files := parseFlags()
	fmt.Printf("Old file: %s\n", files.oldFile)
	fmt.Printf("New file: %s\n", files.newFile)
}

func parseFlags() DBFiles {
	oldFile := flag.String("old", "", "Path to old file")
	newFile := flag.String("new", "", "Path to new file")
	flag.Parse()
	return DBFiles{*oldFile, *newFile}
}
