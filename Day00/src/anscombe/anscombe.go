// Package anscombe provides functions for calculating statistical measures such as
// mean, median, mode, and standard deviation.
package anscombe

import (
	"math"
	"sort"
)

// CalculateMean calculates the mean (average) of a slice of integers.
// It returns the mean as a float64.
func CalculateMean(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return float64(sum) / float64(len(numbers))
}

// CalculateMedian calculates the median of a slice of integers.
// It returns the median as a float64.
func CalculateMedian(numbers []int) float64 {
	sort.Ints(numbers)
	length := len(numbers)

	if length%2 == 0 {
		return float64(numbers[length/2-1]+numbers[length/2]) / 2.
	}

	return float64(numbers[length/2])
}

// CalculateMode calculates the mode (most frequent value) of a slice of integers.
// It returns the mode as an int. If there are multiple modes, it returns the smallest one.
func CalculateMode(numbers []int) int {
	countMap := make(map[int]int)
	maxCount := 0
	var mode int

	for _, num := range numbers {
		countMap[num]++
		if countMap[num] > maxCount {
			maxCount = countMap[num]
			mode = num
		} else if countMap[num] == maxCount {
			if num < mode {
				mode = num
			}
		}
	}

	return mode
}

// CalculateSD calculates the standard deviation of a slice of integers.
// It returns the standard deviation as a float64.
func CalculateSD(numbers []int) float64 {
	mean := CalculateMean(numbers)
	varianceSum := .0

	for _, num := range numbers {
		varianceSum += math.Pow(float64(num)-mean, 2)
	}

	variance := varianceSum / float64(len(numbers))
	standartDeviation := math.Sqrt(variance)

	return standartDeviation
}
