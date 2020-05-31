package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i < n+1; i++ {
		s := strings.Builder{}
		if i%3 == 0 {
			s.WriteString( "Fizz")
		}
		if i%5 == 0 {
			s.WriteString("Buzz")
		}

		if s.Len() ==0 {
			s.WriteString(strconv.Itoa(i))
		}
		res[i-1] = s.String()
	}
	return res
}


func main() {
	testCase := []int{
		15,
	}

	for _, c := range testCase{
		r := fizzBuzz(c)
		fmt.Println("Result: ", r)
	}
}