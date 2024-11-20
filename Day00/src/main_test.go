package main

import (
	"flag"
	"os"
	"testing"
)

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestParseFlagsDefault(t *testing.T) {
	os.Args = []string{"cmd"}
	flags := parseFlags()
	if !flags.Mean || !flags.Median || !flags.Mode || !flags.SD {
		t.Errorf("Expected all flags to be true, got Mean: %v, Median: %v, Mode: %v, SD: %v", flags.Mean, flags.Median, flags.Mode, flags.SD)
	}
}

func TestParseFlagsAll(t *testing.T) {
	resetFlags()
	os.Args = []string{"cmd", "-mean", "-median", "-mode", "-sd"}
	flags := parseFlags()
	if !flags.Mean || !flags.Median || !flags.Mode || !flags.SD {
		t.Errorf("Expected all flags to be true, got Mean: %v, Median: %v, Mode: %v, SD: %v", flags.Mean, flags.Median, flags.Mode, flags.SD)
	}
}

func TestParseFlagsSingle(t *testing.T) {
	resetFlags()
	os.Args = []string{"cmd", "-mean"}
	flags := parseFlags()
	if !flags.Mean || flags.Median || flags.Mode || flags.SD {
		t.Errorf("Expected only Mean to be true, got Mean: %v, Median: %v, Mode: %v, SD: %v", flags.Mean, flags.Median, flags.Mode, flags.SD)
	}
}
