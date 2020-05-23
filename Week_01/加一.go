package main

import "fmt"

func plusOne(digits []int) []int {
	dLen := len(digits)
	for i := dLen - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

func main() {
	testCase := [][]int{
		{1, 2, 3},
		{3, 2, 4},
		{4, 3, 2, 1},
		{5, 9, 9, 9},
		{9, 9, 9, 9},
	}

	for _, d := range testCase {
		result := plusOne(d)
		fmt.Println("Result: ", result)
	}
}
