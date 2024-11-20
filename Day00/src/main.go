package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/emmonbear/APG1/Day00/src/iutils"
)

type Flags struct {
	Mean   bool
	Median bool
	Mode   bool
	SD     bool
}

func main() {
	// flags := parseFlags()
	numbers, err := iutils.ReadInput()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Read input error: ", err)
		return
	}

	if len(numbers) == 0 {
		fmt.Println("There are no valid numbers to process.")
		return
	}

}

func parseFlags() Flags {
	var flags Flags

	flag.BoolVar(&flags.Mean, "mean", true, "Print mean")
	flag.BoolVar(&flags.Median, "median", true, "Print median")
	flag.BoolVar(&flags.Mode, "mode", true, "Print mode")
	flag.BoolVar(&flags.SD, "sd", true, "Print standart deviation")

	flag.Parse()

	return flags
}
