// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/emmonbear/APG1/Day02/src/pkg/finder"
)

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
