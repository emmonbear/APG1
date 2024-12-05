package wc

import (
	"flag"
	"fmt"
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

	if (wc.Lines && wc.Chars) || (wc.Lines && wc.Words) || (wc.Words && wc.Lines) {
		return fmt.Errorf("you can specify only one of the flags: -l, -c, -w")
	}

	if !(wc.Lines || wc.Chars || wc.Words) {
		wc.Words = true
	}

	return nil
}
