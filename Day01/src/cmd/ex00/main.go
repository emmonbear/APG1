// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"os"

	"github.com/emmonbear/APG1/Day01.git/pkg/dbprinter"
	"github.com/emmonbear/APG1/Day01.git/pkg/dbreader"
)

func main() {
	filename := parseFlags()

	reader := dbreader.GetDBReader(filename)
	recipes, err := reader.Read(filename)

	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	printer := dbprinter.GetDBPrinter(filename)
	err = printer.Print(os.Stdout, recipes)

	if err != nil {
		log.Fatalf("Failed to print data: %v", err)
	}
}

func parseFlags() string {
	filename := flag.String("f", "", "Path to file")
	flag.Parse()
	return *filename
}
