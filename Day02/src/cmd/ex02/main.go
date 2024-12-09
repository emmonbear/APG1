// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main is the entry point for the myXargs utility.
// It reads input arguments, processes them, and executes the specified commands.
// It handles errors gracefully and provides appropriate feedback for each step.
//
// Usage:
//
//	./myXargs [options] [arguments...]
//
// The functionality for parsing and executing commands is provided by the xargs package.
package main

import (
	"fmt"
	"os"

	"github.com/emmonbear/APG1/Day02/src/pkg/xargs"
)

// Package main is the entry point for the myXargs utility.
// It reads input arguments, processes them, and executes the specified commands.
// It handles errors gracefully and provides appropriate feedback for each step.
//
// Usage:
//
//	./myXargs [options] [arguments...]
//
// The functionality for parsing and executing commands is provided by the xargs package.
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./myXargs [options] [arguments...]")
		os.Exit(1)
	}

	myXargs := xargs.New(os.Args[1], os.Args[2:])

	if err := myXargs.ParseCommandLine(os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "Error during input processing: %v\n", err)
		os.Exit(1)
	}

	if err := myXargs.Execute(os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "A command execution error: %v\n", err)
		os.Exit(1)
	}
}
