// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main is the entry point for the log rotation utility.
// It parses command-line arguments, performs log rotation, and handles errors gracefully.
//
// Usage:
//
//	./myRotate [options] <log_file1> <log_file2> ...
//
// The functionality for parsing flags and rotating log files is provided by the rotate package.
package main

import (
	"fmt"
	"os"

	"github.com/emmonbear/APG1/Day02/src/pkg/rotate"
)

// main is the entry point for the log rotation utility. It:
// - Parses command-line flags using the `rotate.ParseFlags` function.
// - Calls the `rotate.Rotate` function to perform log file rotation.
// - Reports any errors encountered during flag parsing or rotation to standard error.
//
// Usage:
//
//	./myRotate [options] <log_file1> <log_file2> ...
func main() {
	config, err := rotate.ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "flag handling error: %v\n", err)
		os.Exit(1)
	}
	err = rotate.Rotate(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "rotate error: %v\n", err)
		os.Exit(1)
	}
}
