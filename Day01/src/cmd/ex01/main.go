// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"

	"github.com/emmonbear/APG1/Day01.git/pkg/dbcompare"
	"github.com/emmonbear/APG1/Day01.git/pkg/dbreader"
)

type DBFiles struct {
	oldFile string
	newFile string
}

func main() {
	files := parseFlags()
	oldReader := dbreader.GetDBReader(files.oldFile)
	newReader := dbreader.GetDBReader(files.newFile)
	oldRecipes, err := oldReader.Read(files.oldFile)
	if err != nil {
		fmt.Printf("Failed to read old file: %v\n", err)
		return
	}
	newRecipes, err := newReader.Read(files.newFile)
	if err != nil {
		fmt.Printf("Failed to read new file: %v\n", err)
		return
	}

	comparer := dbcompare.NewComparer()
	comparer.CompareRecipes(oldRecipes, newRecipes)

}

func parseFlags() DBFiles {
	oldFile := flag.String("old", "", "Path to old file")
	newFile := flag.String("new", "", "Path to new file")
	flag.Parse()
	return DBFiles{*oldFile, *newFile}
}
