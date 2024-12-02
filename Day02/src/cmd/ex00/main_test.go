package main

import (
	"flag"
	"os"
	"testing"
)

func TestParseFlagsDefault(t *testing.T) {
	flags, _ := parseFlags()
	expected := FindFlags{false, false, false, ""}

	if flags != expected {
		t.Fatalf("Expected %+v, got %+v", expected, flags)
	}
}

func TestParseFlagsWithFiles(t *testing.T) {
	setupFlags([]string{"myFind", "-f", "./path"})
	flags, _ := parseFlags()
	expected := FindFlags{false, false, true, ""}
	if flags != expected {
		t.Fatalf("Expected %+v, got %+v", expected, flags)
	}
}

func TestParseFlagsWithSymlinks(t *testing.T) {
	setupFlags([]string{"myFind", "-sl", "./path"})
	flags, _ := parseFlags()

	expected := FindFlags{true, false, false, ""}
	if flags != expected {
		t.Fatalf("Expected %+v, got %+v", expected, flags)
	}
}

func TestParseFlagsWithDirectories(t *testing.T) {
	setupFlags([]string{"myFind", "-d", "./path"})
	flags, _ := parseFlags()
	expected := FindFlags{false, true, false, ""}
	if flags != expected {
		t.Fatalf("Expected %+v, got %+v", expected, flags)
	}
}

func TestParseFlagsWithExtAndFile(t *testing.T) {
	setupFlags([]string{"myFind", "-f", "./path", "-ext", "txt"})
	flags, _ := parseFlags()
	expected := FindFlags{false, false, true, "txt"}
	if flags != expected {
		t.Fatalf("Expected %+v, got %+v", expected, flags)
	}
}
func setupFlags(args []string) {
	os.Args = args
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}
