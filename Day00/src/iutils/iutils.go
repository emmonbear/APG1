// Package iutils provides utility functions for reading and validating integer input from standard input.
package iutils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	minValue = -100000
	maxValue = 100000
)

// ReadInput reads integers from standard input, validates them, and returns a slice of valid integers.
// It returns an error if there is an issue with reading the input.
func ReadInput() ([]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()
		num, err := validateInput(line)

		if err != nil {
			fmt.Println(err)
			continue
		}

		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

// validateInput validates that the input string is a valid integer within the range [-100000, 100000].
// It returns the integer value and an error if the input is invalid.
func validateInput(input string) (int, error) {
	num, err := strconv.Atoi(input)

	if err != nil {
		return 0, fmt.Errorf("incorrect entry %s. An integer is expected", input)
	}

	if num < minValue || num > maxValue {
		return 0, fmt.Errorf("number %d is out of valid range", num)
	}

	return num, err
}
