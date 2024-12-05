// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/emmonbear/APG1/Day02/src/pkg/wc"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	fs := flag.NewFlagSet("main", flag.ContinueOnError)
	options := wc.NewWCFlags()
	err := options.ParseFlags(fs, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(options)
}
