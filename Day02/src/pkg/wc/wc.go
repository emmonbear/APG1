package wc

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"unicode/utf8"
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

	var lines int
	var chars int
	var words int
	wordRegexp := regexp.MustCompile(`\S+`)

	for scanner.Scan() {
		lines++
		line := scanner.Text()
		chars += utf8.RuneCountInString(line)
		words += len(wordRegexp.FindAllString(line, -1))
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file %s: %v", filename, err)
	}

	switch {
	case options.Lines:
		return lines, nil
	case options.Chars:
		if lines == 0 {
			return chars, nil
		}
		return chars + lines - 1, nil
	case options.Words:
		return words, nil
	default:
		return 0, nil
	}
}
