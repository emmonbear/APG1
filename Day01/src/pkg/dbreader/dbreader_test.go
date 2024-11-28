package dbreader

import (
	"os"
	"testing"
)

func TestReadJSON(t *testing.T) {
	filename := "../../stolen_database.json"
	reader := GetDBReader(filename)
	recipes, err := reader.Read(filename)
	if err != nil {
		t.Fatalf("read failed: %v", err)
	}

	expectedRecipes := Recipes{
		Cakes: []Cake{
			{
				Name: "Red Velvet Strawberry Cake",
				Time: "45 min",
				Ingredients: []Ingredients{
					{Name: "Flour", Count: "2", Unit: "mugs"},
					{Name: "Strawberries", Count: "8"},
					{Name: "Coffee Beans", Count: "2.5", Unit: "tablespoons"},
					{Name: "Cinnamon", Count: "1"},
				},
			},
			{
				Name: "Moonshine Muffin",
				Time: "30 min",
				Ingredients: []Ingredients{
					{Name: "Brown sugar", Count: "1", Unit: "mug"},
					{Name: "Blueberries", Count: "1", Unit: "mug"},
				},
			},
		},
	}

	if len(recipes.Cakes) != len(expectedRecipes.Cakes) {
		t.Fatalf("Expected %d recipes, got %d", len(expectedRecipes.Cakes), len(recipes.Cakes))
	}

	for i := range recipes.Cakes {
		if recipes.Cakes[i].Name != expectedRecipes.Cakes[i].Name {
			t.Fatalf("Expected recipe name %s, got %s", expectedRecipes.Cakes[i].Name, recipes.Cakes[i].Name)
		}

		if recipes.Cakes[i].Time != expectedRecipes.Cakes[i].Time {
			t.Fatalf("Expected recipe time %s, got %s", expectedRecipes.Cakes[i].Time, recipes.Cakes[i].Time)
		}

		if len(recipes.Cakes[i].Ingredients) != len(expectedRecipes.Cakes[i].Ingredients) {
			t.Fatalf("Expected %d ingredients, got %d", len(expectedRecipes.Cakes[i].Ingredients), len(recipes.Cakes[i].Ingredients))
		}

		for j := range recipes.Cakes[i].Ingredients {
			if recipes.Cakes[i].Ingredients[j].Name != expectedRecipes.Cakes[i].Ingredients[j].Name {
				t.Fatalf("Expected ingredient name %s, got %s", expectedRecipes.Cakes[i].Ingredients[j].Name, recipes.Cakes[i].Ingredients[j].Name)
			}

			if recipes.Cakes[i].Ingredients[j].Count != expectedRecipes.Cakes[i].Ingredients[j].Count {
				t.Fatalf("Expected ingredient count %s, got %s", expectedRecipes.Cakes[i].Ingredients[j].Count, recipes.Cakes[i].Ingredients[j].Count)
			}

			if recipes.Cakes[i].Ingredients[j].Unit != expectedRecipes.Cakes[i].Ingredients[j].Unit {
				t.Fatalf("Expected ingredient unit %s, got %s", expectedRecipes.Cakes[i].Ingredients[j].Unit, recipes.Cakes[i].Ingredients[j].Unit)
			}
		}
	}

}

func TestReadXML(t *testing.T) {
	filename := "../../original_database.xml"
	reader := GetDBReader(filename)
	recipes, err := reader.Read(filename)
	if err != nil {
		t.Fatalf("read failed: %v", err)
	}

	expectedRecipes := Recipes{
		Cakes: []Cake{
			{
				Name: "Red Velvet Strawberry Cake",
				Time: "40 min",
				Ingredients: []Ingredients{
					{Name: "Flour", Count: "3", Unit: "cups"},
					{Name: "Vanilla extract", Count: "1.5", Unit: "tablespoons"},
					{Name: "Strawberries", Count: "7", Unit: ""},
					{Name: "Cinnamon", Count: "1", Unit: "pieces"},
				},
			},
			{
				Name: "Blueberry Muffin Cake",
				Time: "30 min",
				Ingredients: []Ingredients{
					{Name: "Baking powder", Count: "3", Unit: "teaspoons"},
					{Name: "Brown sugar", Count: "0.5", Unit: "cup"},
					{Name: "Blueberries", Count: "1", Unit: "cup"},
				},
			},
		},
	}

	if len(recipes.Cakes) != len(expectedRecipes.Cakes) {
		t.Fatalf("Expected %d recipes, got %d", len(expectedRecipes.Cakes), len(recipes.Cakes))
	}

	for i := range recipes.Cakes {
		if recipes.Cakes[i].Name != expectedRecipes.Cakes[i].Name {
			t.Fatalf("Expected recipe name %s, got %s", expectedRecipes.Cakes[i].Name, recipes.Cakes[i].Name)
		}

		if recipes.Cakes[i].Time != expectedRecipes.Cakes[i].Time {
			t.Fatalf("Expected recipe time %s, got %s", expectedRecipes.Cakes[i].Time, recipes.Cakes[i].Time)
		}

		if len(recipes.Cakes[i].Ingredients) != len(expectedRecipes.Cakes[i].Ingredients) {
			t.Fatalf("Expected %d ingredients, got %d", len(expectedRecipes.Cakes[i].Ingredients), len(recipes.Cakes[i].Ingredients))
		}

		for j := range recipes.Cakes[i].Ingredients {
			if recipes.Cakes[i].Ingredients[j].Name != expectedRecipes.Cakes[i].Ingredients[j].Name {
				t.Fatalf("Expected ingredient name %s, got %s", expectedRecipes.Cakes[i].Ingredients[j].Name, recipes.Cakes[i].Ingredients[j].Name)
			}

			if recipes.Cakes[i].Ingredients[j].Count != expectedRecipes.Cakes[i].Ingredients[j].Count {
				t.Fatalf("Expected ingredient count %s, got %s", expectedRecipes.Cakes[i].Ingredients[j].Count, recipes.Cakes[i].Ingredients[j].Count)
			}

			if recipes.Cakes[i].Ingredients[j].Unit != expectedRecipes.Cakes[i].Ingredients[j].Unit {
				t.Fatalf("Expected ingredient unit %s, got %s", expectedRecipes.Cakes[i].Ingredients[j].Unit, recipes.Cakes[i].Ingredients[j].Unit)
			}
		}
	}

}

func TestGetDBReader(t *testing.T) {
	tests := []struct {
		filename string
		expected DBReader
	}{
		{
			filename: "test.xml",
			expected: &XMLReader{},
		},
		{
			filename: "test.json",
			expected: &JSONReader{},
		},
		{
			filename: "test.txt",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			reader := GetDBReader(tt.filename)
			if reader == nil && tt.expected != nil {
				t.Fatalf("Expected non-nil reader, got nil")
			}
			if reader != nil && tt.expected == nil {
				t.Fatalf("Expected nil reader, got non-nil")
			}
			if reader != nil && tt.expected != nil {
				switch reader.(type) {
				case *XMLReader:
					if _, ok := tt.expected.(*XMLReader); !ok {
						t.Fatalf("Expected %T, got %T", tt.expected, reader)
					}
				case *JSONReader:
					if _, ok := tt.expected.(*JSONReader); !ok {
						t.Fatalf("Expected %T, got %T", tt.expected, reader)
					}
				default:
					t.Fatalf("Unexpected reader type: %T", reader)
				}
			}
		})
	}
}

func TestReadJSONNotFound(t *testing.T) {
	filename := "test.json"
	reader := GetDBReader(filename)

	_, err := reader.Read(filename)

	if err == nil {
		t.Fatalf("Expected error for non-existent file, got nil")
	}
}

func TestReadXMLNotFound(t *testing.T) {
	filename := "test.xml"
	reader := GetDBReader(filename)

	_, err := reader.Read(filename)

	if err == nil {
		t.Fatalf("Expected error for non-existent file, got nil")
	}
}

func TestReadInvalidJSON(t *testing.T) {
	tempFile, err := os.CreateTemp("", "invalid_json.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	invalidJSON := `"invalid": "json"`
	if _, err := tempFile.Write([]byte(invalidJSON)); err != nil {
		t.Fatalf("Failed to write invalid JSON to temp file: %v", err)
	}
	tempFile.Close()

	reader := &JSONReader{}
	_, err = reader.Read(tempFile.Name())
	if err == nil {
		t.Fatalf("Expected error for invalid JSON, got nil")
	}
}

func TestReadInvalidXML(t *testing.T) {
	tempFile, err := os.CreateTemp("", "invalid_xml.xml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	invalidXML := `"invalid": "xml"`
	if _, err := tempFile.Write([]byte(invalidXML)); err != nil {
		t.Fatalf("Failed to write invalid XML to temp file: %v", err)
	}
	tempFile.Close()

	reader := &XMLReader{}
	_, err = reader.Read(tempFile.Name())
	if err == nil {
		t.Fatalf("Expected error for invalid XML, got nil")
	}
}
