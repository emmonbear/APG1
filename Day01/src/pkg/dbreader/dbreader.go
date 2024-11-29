// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
	"path/filepath"
)

// Ingredients represents an ingredient in a recipe.
type Ingredients struct {
	Name  string `xml:"itemname" json:"ingredient_name"`
	Count string `xml:"itemcount" json:"ingredient_count"`
	Unit  string `xml:"itemunit" json:"ingredient_unit"`
}

// Cake represents a cake recipe.
type Cake struct {
	Name        string        `xml:"name" json:"name"`
	Time        string        `xml:"stovetime" json:"time"`
	Ingredients []Ingredients `xml:"ingredients>item" json:"ingredients"`
}

// Recipes represents a collection of cake recipes.
type Recipes struct {
	Cakes []Cake `xml:"cake" json:"cake"`
}

// DBReader is an interface for reading recipe data from a file.
type DBReader interface {
	// Read reads the recipe data from the specified file.
	Read(filename string) (Recipes, error)
}

// XMLReader is a struct that implements the DBReader interface for XML files.
type XMLReader struct{}

// Read reads the recipe data from the specified XML file.
func (r *XMLReader) Read(filename string) (Recipes, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return Recipes{}, err
	}

	var recipes Recipes
	err = xml.Unmarshal(data, &recipes)

	if err != nil {
		return Recipes{}, err
	}

	return recipes, nil
}

// JSONReader is a struct that implements the DBReader interface for JSON files.
type JSONReader struct{}

// Read reads the recipe data from the specified JSON file.
func (r *JSONReader) Read(filename string) (Recipes, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return Recipes{}, err
	}

	var recipes Recipes
	err = json.Unmarshal(data, &recipes)

	if err != nil {
		return Recipes{}, err
	}

	return recipes, nil
}

// GetDBReader returns a DBReader implementation based on the file extension.
func GetDBReader(filename string) DBReader {
	ext := filepath.Ext(filename)
	switch ext {
	case ".xml":
		return &XMLReader{}
	case ".json":
		return &JSONReader{}
	default:
		log.Printf("Unsupported file extension: %s", ext)
		return nil
	}
}
