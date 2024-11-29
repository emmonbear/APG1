// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbcompare

import (
	"bytes"
	"os"
	"testing"

	"github.com/emmonbear/APG1/Day01.git/pkg/dbreader"
)

func TestCompareRecipes(t *testing.T) {
	oldDB := "../../original_database.xml"
	newDB := "../../stolen_database.json"

	oldReader := dbreader.GetDBReader(oldDB)
	newReader := dbreader.GetDBReader(newDB)

	oldRecipes, err := oldReader.Read(oldDB)
	if err != nil {
		t.Fatalf("failed to read old database: %v", err)
	}

	newRecipes, err := newReader.Read(newDB)
	if err != nil {
		t.Fatalf("failed to read new database: %v", err)
	}

	comparer := NewComparer()

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	comparer.CompareRecipes(oldRecipes, newRecipes)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)

	expectedOutput := `ADDED cake "Moonshine Muffin"
REMOVED cake "Blueberry Muffin Cake"
CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "45 min" instead of "40 min"
ADDED ingredient "Coffee Beans" for cake "Red Velvet Strawberry Cake"
REMOVED ingredient "Vanilla extract" for cake "Red Velvet Strawberry Cake"
CHANGED unit for ingredient "Flour" for cake "Red Velvet Strawberry Cake" - "mugs" instead of "cups"
CHANGED unit count for ingredient "Strawberries" for cake "Red Velvet Strawberry Cake" - "8" instead of "7"
REMOVED unit "pieces" for ingredient "Cinnamon" for cake "Red Velvet Strawberry Cake"
`
	if buf.Len() != len(expectedOutput) {
		t.Errorf("expected output \n%d\n, got \n%d", buf.Len(), len(expectedOutput))
	}

}
