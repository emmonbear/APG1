package rotate

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

func Rotate(config *Config) error {
	results := make(chan error)
	for _, file := range config.FileNames {
		go func(f string) {
			results <- rotateFile(f, config.ArchiveDir)
		}(file)
	}

	for range config.FileNames {
		if err := <-results; err != nil {
			return err
		}
	}

	return nil
}

func rotateFile(file string, archiveDir string) error {
	info, err := os.Stat(file)
	if err != nil {
		return fmt.Errorf("failed to stat file %s: %v", file, err)
	}

	if info.IsDir() {
		return fmt.Errorf("path %s is a directory, not a file", file)
	}

	timestamp := info.ModTime().Unix()
	tmp := fmt.Sprintf("%s_%d.tar.gz", filepath.Base(file), timestamp)
	archiveName := filepath.Join(archiveDir, tmp)

	if err := createTarGz(file, archiveName, info); err != nil {
		return fmt.Errorf("failed to create archive file %s: %v", file, err)
	}

	return nil
}

func createTarGz(file string, archiveName string, info os.FileInfo) error {
	archiveFile, err := os.Create(archiveName)
	if err != nil {
		return fmt.Errorf("failed to create archive file %s: %v", file, err)
	}
	defer archiveFile.Close()

	gzipWriter := gzip.NewWriter(archiveFile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	sourceFile, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", file, err)
	}
	defer sourceFile.Close()

	header := &tar.Header{
		Name:    filepath.Base(file),
		Size:    info.Size(),
		Mode:    int64(info.Mode()),
		ModTime: info.ModTime(),
	}

	if err := tarWriter.WriteHeader(header); err != nil {
		return fmt.Errorf("failed to write header for file %s: %v", file, err)
	}

	if _, err := io.Copy(tarWriter, sourceFile); err != nil {
		return fmt.Errorf("failed to write file content to archive file %s: %v", file, err)
	}

	return nil
}
