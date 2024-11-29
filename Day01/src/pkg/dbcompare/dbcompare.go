// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbcompare

import (
	"fmt"

	"github.com/emmonbear/APG1/Day01.git/pkg/dbreader"
)

// Comparer is a struct that provides methods to compare recipes.
type Comparer struct{}

// NewComparer creates a new instance of Comparer.
func NewComparer() *Comparer {
	return &Comparer{}
}

// CompareRecipes compares two sets of recipes and prints the differences.
func (c *Comparer) CompareRecipes(oldDB, newDB dbreader.Recipes) {
	oldCakeMap := c.createCakeMap(oldDB.Cakes)
	newCakeMap := c.createCakeMap(newDB.Cakes)

	c.compareCakes(oldCakeMap, newCakeMap)
}

// CreateCakeMap creates a map of cakes from a slice of cakes.
func (c *Comparer) createCakeMap(cakes []dbreader.Cake) map[string]dbreader.Cake {
	cakeMap := make(map[string]dbreader.Cake, len(cakes))
	for _, cake := range cakes {
		cakeMap[cake.Name] = cake

	}

	return cakeMap
}

// CompareCakes compares two maps of cakes and prints the differences.
func (c *Comparer) compareCakes(oldCakeMap, newCakeMap map[string]dbreader.Cake) {
	for name := range newCakeMap {
		if _, ok := oldCakeMap[name]; !ok {
			fmt.Printf("ADDED cake \"%s\"\n", name)
		}
	}

	for name := range oldCakeMap {
		if _, ok := newCakeMap[name]; !ok {
			fmt.Printf("REMOVED cake \"%s\"\n", name)
		}
	}

	for name, oldCake := range oldCakeMap {
		if newCake, ok := newCakeMap[name]; ok {
			c.compareCake(oldCake, newCake)
		}
	}

}

// CompareCake compares two cakes and prints the differences.
func (c *Comparer) compareCake(oldCake, newCake dbreader.Cake) {
	if oldCake.Time != newCake.Time {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldCake.Name, newCake.Time, oldCake.Time)
	}

	oldIngredientsMap := c.createIngredientsMap(oldCake.Ingredients)
	newIngredientsMap := c.createIngredientsMap(newCake.Ingredients)

	c.compareIngredients(oldIngredientsMap, newIngredientsMap, oldCake.Name)
}

// CreateIngredientsMap creates a map of ingredients from a slice of ingredients.
func (c *Comparer) createIngredientsMap(ingredients []dbreader.Ingredients) map[string]dbreader.Ingredients {
	ingredientMap := make(map[string]dbreader.Ingredients, len(ingredients))
	for _, ingredient := range ingredients {
		ingredientMap[ingredient.Name] = ingredient
	}

	return ingredientMap
}


// CompareIngredients compares two maps of ingredients and prints the differences.
func (c *Comparer) compareIngredients(oldIngredientsMap, newIngredientsMap map[string]dbreader.Ingredients, cakeName string) {
	for name := range newIngredientsMap {
		if _, ok := oldIngredientsMap[name]; !ok {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", name, cakeName)
		}
	}

	for name := range oldIngredientsMap {
		if _, ok := newIngredientsMap[name]; !ok {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", name, cakeName)
		}
	}

	for name, oldIngredient := range oldIngredientsMap {
		if newIngredient, ok := newIngredientsMap[name]; ok {
			c.compareIngredient(oldIngredient, newIngredient, cakeName)
		}
	}

}

// CompareIngredient compares two ingredients and prints the differences.
func (c *Comparer) compareIngredient(oldIngredient, newIngredient dbreader.Ingredients, cakeName string) {
	if oldIngredient.Unit != newIngredient.Unit && newIngredient.Unit != "" {
		fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIngredient.Name, cakeName, newIngredient.Unit, oldIngredient.Unit)
	}

	if oldIngredient.Count != newIngredient.Count && oldIngredient.Unit == newIngredient.Unit {
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIngredient.Name, cakeName, newIngredient.Count, oldIngredient.Count)
	}

	if oldIngredient.Unit != "" && newIngredient.Unit == "" {
		fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", oldIngredient.Unit, oldIngredient.Name, cakeName)
	}
}
