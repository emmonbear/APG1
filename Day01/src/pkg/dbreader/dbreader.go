package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"os"
)

type Ingredients struct {
	Name  string `xml:"itemname" json:"ingredient_name"`
	Count string `xml:"itemcount" json:"ingredient_count"`
	Unit  string `xml:"itemunit" json:"ingredient_unit"`
}

type Cake struct {
	Name        string        `xml:"name" json:"name"`
	Time        string        `xml:"stovetime" json:"time"`
	Ingredients []Ingredients `xml:"ingredients>item" json:"ingredients"`
}

type Recipes struct {
	Cakes []Cake `xml:"cake" json:"cake"`
}

type DBreader interface {
	Read(filename string) (Recipes, error)
}

type XMLReader struct{}

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

type JSONReader struct{}

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
