// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main is the entry point for the wc utility.
// It parses command-line arguments, performs word count operations
// on specified files concurrently, and outputs the results.
//
// Usage:
//
//	./mywc [options] <file1> <file2> ...
//
// Options and functionality are defined in the wc package.
package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/emmonbear/APG1/Day02/src/pkg/wc"
)

// Person represents an individual with a name and an age.
// This struct can be used to demonstrate basic struct usage, but is not
// utilized within the main program.
type Person struct {
	Name string // Name represents the name of the person
	Age  int    // Age represents the age of the person
}

// main is the entry point for the wc utility. It:
// - Parses command-line flags using a custom flag set.
// - For each provided file, it concurrently counts the words using the wc.WC function.
// - Outputs the word count for each file, formatted as: "count <filename>".
// - Handles errors gracefully and reports them to standard error.
func main() {
	fs := flag.NewFlagSet("main", flag.ContinueOnError)

	options := wc.NewWCFlags()
	err := options.ParseFlags(fs, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	var wg sync.WaitGroup
	for _, filename := range fs.Args() {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			count, err := wc.WC(file, options)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			fmt.Printf("%d\t%s\n", count, filename)
		}(filename)
	}

	wg.Wait()
}
