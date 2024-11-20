package anscombe

import (
	"math"
	"testing"
)

func TestCalculateMean1(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := 3.
	result := CalculateMean(numbers)

	if result != expected {
		t.Errorf("CalculateMean(%v) = %f; want %f", numbers, result, expected)
	}
}

func TestCalculateMean2(t *testing.T) {
	numbers := []int{11, 2, 3, 8, 9, 1, 99}
	expected := 19.
	result := CalculateMean(numbers)

	if result != expected {
		t.Errorf("CalculateMean(%v) = %f; want %f", numbers, result, expected)
	}
}

func TestCalculateMedianOdd(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := 3.
	result := CalculateMedian(numbers)

	if result != expected {
		t.Errorf("CalculateMean(%v) = %f; want %f", numbers, result, expected)
	}
}

func TestCalculateMedianEven(t *testing.T) {
	numbers := []int{2, 5, 1, 3}
	expected := 2.5
	result := CalculateMedian(numbers)

	if result != expected {
		t.Errorf("CalculateMean(%v) = %f; want %f", numbers, result, expected)
	}
}

func TestCalculateMode1(t *testing.T) {
	numbers := []int{1, 3, 4, 5, 6, 7, 8, 9}
	expected := 1
	result := CalculateMode(numbers)

	if result != expected {
		t.Errorf("CalculateMean(%v) = %d; want %d", numbers, result, expected)
	}
}

func TestCalculateMode2(t *testing.T) {
	numbers := []int{1, 3, 4, 5, 6, 7, 3, 9}
	expected := 3
	result := CalculateMode(numbers)

	if result != expected {
		t.Errorf("CalculateMean(%v) = %d; want %d", numbers, result, expected)
	}
}

func TestCalculateMode3(t *testing.T) {
	numbers := []int{1, 3, 4, 5, 6, 7, 3, 9, 1}
	expected := 1
	result := CalculateMode(numbers)

	if result != expected {
		t.Errorf("CalculateMean(%v) = %d; want %d", numbers, result, expected)
	}
}

func TestCalculateSD(t *testing.T) {
	numbers := []int{2, 5, 1, 3, 6, 7}
	expected := 2.1602467
	result := CalculateSD(numbers)
	tolerance := 1e-6
	if math.Abs(result-expected) > tolerance {
		t.Errorf("CalculateMean(%v) = %f; want %f", numbers, result, expected)
	}
}
