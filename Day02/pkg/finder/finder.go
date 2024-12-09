// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package finder

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// EntryType defines the type of a file system entry (File, Directory, Symlink).
// It represents the possible types of file system entries found during the search.
type EntryType int

const (
	// File represents a regular file.
	File EntryType = iota
	// Directory represents a directory.
	Directory
	// Symlink represents a symbolic link.
	Symlink
)

// Entry represents a file system entry (file, directory, or symlink).
// It contains the path, type, and for symlinks, the target link.
type Entry struct {
	Path string    // Path of the entry
	Type EntryType // Type of the entry (File, Directory, or Symlink)
	Link string    // Target link for symlinks, or empty if not a symlink
}

type Options struct {
	IncludeFiles       bool   // If true, include regular files in the search results
	IncludeDirectories bool   // If true, include directories in the search results
	IncludeSymlinks    bool   // If true, include symbolic links in the search results
	ExtensionFilter    string // Filter files by extension (works only with -f option)
}

// Find searches the file system starting from the root directory, and returns the matching entries
// based on the specified options. It will traverse all subdirectories recursively.
func Find(root string, options Options) ([]Entry, error) {
	var results []Entry

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			if os.IsNotExist(err) {
				return err
			}
			return nil
		}

		if d.IsDir() && options.IncludeDirectories {
			results = append(results, Entry{Path: path, Type: Directory})
		}

		if d.Type().IsRegular() && options.IncludeFiles {
			if options.ExtensionFilter == "" || filepath.Ext(d.Name()) == "."+options.ExtensionFilter {
				results = append(results, Entry{Path: path, Type: File})
			}
		}

		if d.Type()&os.ModeSymlink != 0 && options.IncludeSymlinks {
			linkTarget, err := os.Readlink(path)
			if err != nil {
				linkTarget = "[broken]"
			} else {
				absLinkTarget := filepath.Join(filepath.Dir(path), linkTarget)
				if _, err := os.Stat(absLinkTarget); err != nil {
					linkTarget = "[broken]"
				}
			}
			results = append(results, Entry{Path: path, Type: Symlink, Link: linkTarget})
		}
		return nil

	})

	return results, err
}

// ParseFlags parses command-line flags to configure the search options.
// It ensures that flags are set correctly and returns an error if the flags are invalid.
func (o *Options) ParseFlags(fs *flag.FlagSet, args []string) error {
	fs.BoolVar(&o.IncludeFiles, "f", false, "Show files")
	fs.BoolVar(&o.IncludeDirectories, "d", false, "Show directories")
	fs.BoolVar(&o.IncludeSymlinks, "sl", false, "Show symlinks")
	fs.StringVar(&o.ExtensionFilter, "ext", "", "Show files with specific extension (work only with -f)")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if !o.IncludeDirectories && !o.IncludeFiles && !o.IncludeSymlinks && o.ExtensionFilter == "" {
		*o = Options{true, true, true, ""}
	}

	if o.ExtensionFilter != "" && !o.IncludeFiles {
		return fmt.Errorf("the -ext option can only be used with the -f option")
	}

	return nil
}

// NewOptions creates and returns a new instance of Options with default values.
func NewOptions() *Options {
	return &Options{false, false, false, ""}
}

// String returns a human-readable string representation of the Entry, formatted based on its type.
// For symlinks, the target link is included, and if broken, it indicates "[broken]".
func (e *Entry) String() string {
	formattedPath := e.Path

	if !filepath.IsAbs(e.Path) && !strings.HasPrefix(e.Path, "./") {
		formattedPath = "./" + e.Path
	}

	if e.Type == Directory && !strings.HasSuffix(formattedPath, "/") {
		formattedPath += "/"
	}

	if e.Type == Symlink {
		if e.Link == "[broken]" {
			return formattedPath + " -> [broken]"
		}
		return formattedPath + " -> " + e.Link
	}

	return formattedPath
}
