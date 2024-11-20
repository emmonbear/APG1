package iutils

import "testing"

func TestValidateInput123(t *testing.T) {
	input := "123"
	expected := 123
	num, err := validateInput(input)
	if err != nil {
		t.Errorf("validateInput(%s) error = %v, wantErr %v", input, err, false)
	}
	if num != expected {
		t.Errorf("validateInput(%s) = %d, want %d", input, num, expected)
	}
}

func TestValidateInputNegative123(t *testing.T) {
	input := "-123"
	expected := -123
	num, err := validateInput(input)
	if err != nil {
		t.Errorf("validateInput(%s) error = %v, wantErr %v", input, err, false)
	}
	if num != expected {
		t.Errorf("validateInput(%s) = %d, want %d", input, num, expected)
	}
}

func TestValidateInput100000(t *testing.T) {
	input := "100000"
	expected := 100000
	num, err := validateInput(input)
	if err != nil {
		t.Errorf("validateInput(%s) error = %v, wantErr %v", input, err, false)
	}
	if num != expected {
		t.Errorf("validateInput(%s) = %d, want %d", input, num, expected)
	}
}

func TestValidateInputNegative100000(t *testing.T) {
	input := "-100000"
	expected := -100000
	num, err := validateInput(input)
	if err != nil {
		t.Errorf("validateInput(%s) error = %v, wantErr %v", input, err, false)
	}
	if num != expected {
		t.Errorf("validateInput(%s) = %d, want %d", input, num, expected)
	}
}

func TestValidateInput100001(t *testing.T) {
	input := "100001"
	expected := 0
	num, err := validateInput(input)
	if err == nil {
		t.Errorf("validateInput(%s) error = %v, wantErr %v", input, err, true)
	}
	if num != expected {
		t.Errorf("validateInput(%s) = %d, want %d", input, num, expected)
	}
}

func TestValidateInputNegative100001(t *testing.T) {
	input := "-100001"
	expected := 0
	num, err := validateInput(input)
	if err == nil {
		t.Errorf("validateInput(%s) error = %v, wantErr %v", input, err, true)
	}
	if num != expected {
		t.Errorf("validateInput(%s) = %d, want %d", input, num, expected)
	}
}

func TestValidateInputABC(t *testing.T) {
	input := "abc"
	expected := 0
	num, err := validateInput(input)
	if err == nil {
		t.Errorf("validateInput(%s) error = %v, wantErr %v", input, err, true)
	}
	if num != expected {
		t.Errorf("validateInput(%s) = %d, want %d", input, num, expected)
	}
}

func TestValidateInputEmpty(t *testing.T) {
	input := ""
	expected := 0
	num, err := validateInput(input)
	if err == nil {
		t.Errorf("validateInput(%s) error = %v, wantErr %v", input, err, true)
	}
	if num != expected {
		t.Errorf("validateInput(%s) = %d, want %d", input, num, expected)
	}
}
