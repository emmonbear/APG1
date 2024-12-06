package wc

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type WCFlags struct {
	Lines bool
	Chars bool
	Words bool
}

func NewWCFlags() *WCFlags {
	return &WCFlags{false, false, false}
}

func (wc *WCFlags) ParseFlags(fs *flag.FlagSet, args []string) error {
	fs.BoolVar(&wc.Lines, "l", false, "Count the numbers of lines")
	fs.BoolVar(&wc.Chars, "m", false, "Count the numbers of characters")
	fs.BoolVar(&wc.Words, "w", false, "Count the numbers of words")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if fs.NArg() < 1 {
		return fmt.Errorf("you need to specify a txt file")
	}

	for _, filename := range fs.Args() {
		if filepath.Ext(filename) != ".txt" {
			return fmt.Errorf("the file %s must have a .txt extension", filename)
		}
	}

	if (wc.Lines && wc.Chars) || (wc.Lines && wc.Words) || (wc.Words && wc.Lines) {
		return fmt.Errorf("you can specify only one of the flags: -l, -c, -w")
	}

	if !(wc.Lines || wc.Chars || wc.Words) {
		wc.Words = true
	}

	return nil
}

func WC(filename string, options *WCFlags) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("could not open file %s: %v", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var count int
	switch {
	case options.Lines:
		for scanner.Scan() {
			count++
		}
	}
	fmt.Printf("%d\t%s\n", count, filename)
	return count, nil
}
