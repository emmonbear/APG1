package iutils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func validateInput(input string) (int, error) {
	num, err := strconv.Atoi(input)

	if err != nil {
		return 0, fmt.Errorf("Error: incorrect entry %s. An integer is expected.", input)
	}

	if num < -100000 || num > 100000 {
		return 0, fmt.Errorf("Error: number %d is out of valid range.", num)
	}

	return num, err
}
