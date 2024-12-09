// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main is the entry point for the myfind utility.
// It parses command-line arguments, applies search options,
// and prints the results of the search operation.
//
// Usage:
//
//	./myfind [options] <path>
//
// Options and functionality are defined in the finder package.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/emmonbear/APG1/Day02/src/pkg/finder"
)

// main is the entry point of the myfind utility. It:
// - Parses command-line flags using a custom flag set.
// - Validates the input arguments and ensures a search path is provided.
// - Calls finder.Find to perform a file search with the specified options.
// - Prints the results of the search to standard output or reports errors to standard error.
func main() {
	fs := flag.NewFlagSet("main", flag.ContinueOnError)
	options := finder.NewOptions()
	err := options.ParseFlags(fs, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if fs.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Usage: ./myfind [options] <path>")
		os.Exit(1)
	}

	root := fs.Arg(0)
	files, err := finder.Find(root, *options)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(&file)
	}
}
