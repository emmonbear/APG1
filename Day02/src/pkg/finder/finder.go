package finder

import (
	"io/fs"
	"os"
	"path/filepath"
)

type EntryType int

const (
	File EntryType = iota
	Directory
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

func FormatEntry(entry Entry) string {
	switch entry.Type {
	case Directory:
		return entry.Path
	case File:
		return entry.Path
	case Symlink:
		if entry.Link == "[broken]" {
			return entry.Path + " -> [broken]"
		}
		return entry.Path + " -> " + entry.Link
	default:
		return entry.Path
	}
}
