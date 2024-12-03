package main

import (
	"flag"
	"os"
	"testing"
)

func TestParseFlagsDefault(t *testing.T) {
	flags, _ := parseFlags()
	expected := FindFlags{true, true, true, ""}

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
	setupFlags([]string{"myFind", "-f", "-ext", "txt", "./path"})
	flags, _ := parseFlags()
	expected := FindFlags{false, false, true, "txt"}
	if flags != expected {
		t.Fatalf("Expected %+v, got %+v", expected, flags)
	}
}

func TestParseFlagsWithExtWithoutFile(t *testing.T) {
	setupFlags([]string{"myFind", "-ext", "txt", "./path"})
	flags, err := parseFlags()
	if err == nil {
		t.Fatalf("Expected %v, got nil", err)
	}

	expected := FindFlags{false, false, false, "txt"}
	if flags != expected {
		t.Fatalf("Expected %+v, got %+v", expected, flags)
	}

}

func TestParseFlagsMultiply1(t *testing.T) {
	setupFlags([]string{"myFind", "-f", "-ext", "go", "-d", "./path"})
	flags, _ := parseFlags()
	expected := FindFlags{false, true, true, "go"}
	if flags != expected {
		t.Fatalf("Expected %+v, got %+v", expected, flags)
	}

}

func TestParseFlagsMultiply2(t *testing.T) {
	setupFlags([]string{"myFind", "-f", "-sl", "-ext", "go", "-d", "./path"})
	flags, _ := parseFlags()
	expected := FindFlags{true, true, true, "go"}
	if flags != expected {
		t.Fatalf("Expected %+v, got %+v", expected, flags)
	}

}

func TestMain(t *testing.T) {
	setupFlags([]string{"myFind", "-f", "./test"})
	main()
}

func setupFlags(args []string) {
	os.Args = args
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}
