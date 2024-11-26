package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/emmonbear/APG1/Day01.git/pkg/dbprinter"
	"github.com/emmonbear/APG1/Day01.git/pkg/dbreader"
)

func main() {
	filename := parseFlags()

	reader := getDBReader(filename)
	recipes, err := reader.Read(filename)

	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	printer := getDBPrinter(filename)
	err = printer.Print(os.Stdout, recipes)

	if err != nil {
		log.Fatalf("Failed to print data: %v", err)
	}
}

func getDBReader(filename string) dbreader.DBreader {
	ext := filepath.Ext(filename)
	switch ext {
	case ".xml":
		return &dbreader.XMLReader{}
	case ".json":
		return &dbreader.JSONReader{}
	default:
		log.Fatalf("Unsupported file extension: %s", ext)
		return nil
	}
}

func getDBPrinter(filename string) dbprinter.DBPrinter {
	ext := filepath.Ext(filename)
	switch ext {
	case ".xml":
		return &dbprinter.XMLPrinter{}
	case ".json":
		return &dbprinter.JSONPrinter{}
	default:
		log.Fatalf("Unsupported file extension: %s", ext)
		return nil
	}
}

func parseFlags() string {
	filepath := flag.String("f", "", "Path to file")
	flag.Parse()
	return *filepath
}
