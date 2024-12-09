package rotate

import (
	"flag"
	"fmt"
)

type Config struct {
	FileNames  []string
	ArchiveDir string
}

func ParseFlags(args []string) (*Config, error) {
	fs := flag.NewFlagSet("rotate", flag.ContinueOnError)

	var archiveDir string
	fs.StringVar(&archiveDir, "a", ".", "Directory to store archives")
	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	fileNames := fs.Args()

	if len(fileNames) == 0 {
		return nil, fmt.Errorf("no log files specified")
	}

	return &Config{
		ArchiveDir: archiveDir,
		FileNames:  fileNames,
	}, nil
}
