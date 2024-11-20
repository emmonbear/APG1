// Package main provides a command-line tool to calculate and print various statistical measures
// (mean, median, mode, and standard deviation) for a list of numbers read from standard input.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/emmonbear/APG1/Day00/src/anscombe"
	"github.com/emmonbear/APG1/Day00/src/iutils"
)

// Flags structure to hold command-line flags.
type Flags struct {
	Mean   bool // Flag to indicate if mean should be calculated.
	Median bool // Flag to indicate if median should be calculated.
	Mode   bool // Flag to indicate if mode should be calculated.
	SD     bool // Flag to indicate if standard deviation should be calculated.
}

func main() {
	numbers, err := iutils.ReadInput()
	flags := parseFlags()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Read input error: %v\n", err)
		return
	}

	if len(numbers) == 0 {
		fmt.Println("There are no valid numbers to process.")
		return
	}

	processFlags(numbers, flags)
}

// parseFlags parses the command-line flags and returns a Flags structure.
func parseFlags() Flags {
	var flags Flags

	flag.BoolVar(&flags.Mean, "mean", false, "Print mean")
	flag.BoolVar(&flags.Median, "median", false, "Print median")
	flag.BoolVar(&flags.Mode, "mode", false, "Print mode")
	flag.BoolVar(&flags.SD, "sd", false, "Print standart deviation")

	flag.Parse()

	// If no flags are set, default to calculating all measures.
	if !flags.Mean && !flags.Median && !flags.Mode && !flags.SD {
		flags.Mean = true
		flags.Median = true
		flags.Mode = true
		flags.SD = true
	}

	return flags
}

// processFlags processes the numbers based on the provided flags and prints the results.
func processFlags(numbers []int, flags Flags) {
	if flags.Mean {
		mean := anscombe.CalculateMean(numbers)
		fmt.Printf("Mean: %.2f\n", mean)
	}

	if flags.Median {
		median := anscombe.CalculateMedian((numbers))
		fmt.Printf("Median: %.2f\n", median)
	}

	if flags.Mode {
		mode := anscombe.CalculateMode((numbers))
		fmt.Printf("Mode: %d\n", mode)
	}

	if flags.SD {
		sd := anscombe.CalculateSD((numbers))
		fmt.Printf("SD: %.2f\n", sd)
	}
}
