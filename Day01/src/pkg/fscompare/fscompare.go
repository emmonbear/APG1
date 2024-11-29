package fscompare

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type FSComparer struct{}

func NewFSComparer() *FSComparer {
	return &FSComparer{}
}

func (c *FSComparer) CompareDumps(oldFile, newFile string) error {
	oldSet, err := c.readFile(oldFile)
	if err != nil {
		log.Printf("read file error: %v", err)
		return err
	}

	newSet, err := c.readFile(newFile)
	if err != nil {
		log.Printf("read file error: %v", err)
		return err
	}

	for path := range oldSet {
		if _, ok := newSet[path]; !ok {
			fmt.Printf("ADDED %s\n", path)
		}
	}

	for path := range newSet {
		if _, ok := oldSet[path]; !ok {
			fmt.Printf("REMOVED %s\n", path)
		}
	}
	return nil
}

func (c *FSComparer) readFile(filename string) (map[string]struct{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileSet := make(map[string]struct{})
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fileSet[line] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return fileSet, nil
}
