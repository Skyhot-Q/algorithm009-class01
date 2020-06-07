package main

import (
	"fmt"
	"math"
)

func getKthMagicNumber(k int) int {
	result := make([]int, k, k)

	p3, p5, p7 := 0, 0, 0

	result[0] = 1
	for i := 1; i < k; i++ {
		result[i] = int(
			math.Min(
				float64(result[p3]*3),
				math.Min(float64(result[p5]*5), float64(result[p7]*7)),
			))

		if result[i] == result[p3]*3 {
			p3++
		}
		if result[i] == result[p5]*5 {
			p5++
		}
		if result[i] == result[p7]*7 {
			p7++
		}
	}
	fmt.Println(result)
	return result[k-1]
}

func main(){
	testCase := []int{
		7,
	}
	for _, c:= range testCase{
		r := getKthMagicNumber(c)
		fmt.Println("Result:", r)
	}
}
