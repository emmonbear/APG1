// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package wc

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"unicode/utf8"
)

// WCFlags represents the flags used to specify which type of counts to perform.
type WCFlags struct {
	Lines bool // Flag to count the number of lines
	Chars bool // Flag to count the number of characters
	Words bool // Flag to count the number of words
}

// NewWCFlags initializes a new WCFlags structure with default values (no flags enabled).
func NewWCFlags() *WCFlags {
	return &WCFlags{false, false, false}
}

// ParseFlags parses the command-line flags and arguments.
// It expects flags for line count (-l), character count (-m), and word count (-w),
// and checks for conflicts between the flags.
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

	if (wc.Lines && wc.Chars) || (wc.Lines && wc.Words) || (wc.Words && wc.Lines) {
		return fmt.Errorf("you can specify only one of the flags: -l, -c, -w")
	}

	if !(wc.Lines || wc.Chars || wc.Words) {
		wc.Words = true
	}

	return nil
}

// WC reads the contents of a file and performs the requested counting based on the provided flags.
// It returns the count of lines, words, or characters depending on the flag set.
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
