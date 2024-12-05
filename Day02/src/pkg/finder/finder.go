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
type EntryType int

const (
	// File represents a regular file.
	File EntryType = iota
	// Directory represents a directory.
	Directory
	// Symlink represents a symbolic link.
	Symlink
)

type Entry struct {
	Path string
	Type EntryType
	Link string
}

type Options struct {
	IncludeFiles       bool
	IncludeDirectories bool
	IncludeSymlinks    bool
	ExtensionFilter    string
}

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

func NewOptions() *Options {
	return &Options{false, false, false, ""}
}

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
