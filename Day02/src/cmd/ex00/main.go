// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/emmonbear/APG1/Day02/src/pkg/finder"
)

type FindFlags struct {
	PrintSymlinks    bool
	PrintDirectories bool
	PrintFiles       bool
	Extension        string
}

func main() {
	flags, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	if flag.NArg() < 1 {
		log.Fatal("Usage: ./myfind [options] <path>")
	}

	root := flag.Arg(0)
	files, _ := finder.Find(root, finder.Options{
		IncludeFiles:       flags.PrintFiles,
		IncludeDirectories: flags.PrintDirectories,
		IncludeSymlinks:    flags.PrintSymlinks,
		ExtensionFilter:    flags.Extension,
	})
	for _, file := range files {
		fmt.Println(finder.FormatEntry(file))
	}
}

func parseFlags() (FindFlags, error) {
	var flags FindFlags

	flag.BoolVar(&flags.PrintFiles, "f", false, "Show files")
	flag.BoolVar(&flags.PrintDirectories, "d", false, "Show directories")
	flag.BoolVar(&flags.PrintSymlinks, "sl", false, "Show symlinks")
	flag.StringVar(&flags.Extension, "ext", "", "Show files with specific extension (work only with -f)")

	flag.Parse()

	if !flags.PrintDirectories && !flags.PrintFiles && !flags.PrintSymlinks && flags.Extension == "" {
		flags = FindFlags{true, true, true, ""}
	}

	if flags.Extension != "" && !flags.PrintFiles {
		return flags, fmt.Errorf("the -ext option can only be used with the -f option")
	}

	return flags, nil
}
