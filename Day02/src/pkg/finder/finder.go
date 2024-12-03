package finder

import (
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

// Entry represents a file system entry with its path, type (file, directory, or symlink),
// and for symlinks, the target of the symlink.
type Entry struct {
	Path string
	Type EntryType
	Link string
}

// Options holds the configuration for the `Find` function, specifying which types
// of entries should be included in the results and any file extension filter.
type Options struct {
	IncludeFiles       bool
	IncludeDirectories bool
	IncludeSymlinks    bool
	ExtensionFilter    string
}

// Find recursively traverses the file system starting at `root` and returns a list of entries
// based on the specified `Options` (which types of entries to include, and optional extension filter).
// It skips any errors encountered during traversal (e.g., permission errors) but continues the search.
func Find(root string, options Options) ([]Entry, error) {
	var results []Entry

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
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

// FormatEntry formats an Entry as a string based on its type.
// If the entry is a symlink, it appends the target link or "[broken]" if the link is broken.
// It also adds "./" for relative paths.
func FormatEntry(entry Entry) string {

	formattedPath := entry.Path

	if !filepath.IsAbs(entry.Path) && !strings.HasPrefix(entry.Path, "./") {
		formattedPath = "./" + entry.Path
	    }
    
	switch entry.Type {
	case Directory:
	    return formattedPath
	case File:
	    return formattedPath
	case Symlink:
	    if entry.Link == "[broken]" {
		return formattedPath + " -> [broken]"
	    }
	    return formattedPath + " -> " + entry.Link
	default:
	    return formattedPath
	}
    }
