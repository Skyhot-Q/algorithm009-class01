package main

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	var countS, countG [10]int
	bulls, cows := 0,0
	for i := range secret{
		ns := int(secret[i] - '0')
		ng := int(guess[i] - '0')
		if ng == ns {
			bulls++
			continue
		}

		if countG[ns] > 0 {
			cows++
			countG[ns]--
		} else {
			countS[ns]++
		}

		if countS[ng] > 0 {
			cows++
			countS[ng]--
		} else {
			countG[ng]++
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)

}

func main() {
	testCase := [][]string{
		{"1807", "7810"},
		{"1123", "0111"},
		{"1122", "1222"},
	}
	for _, t := range testCase {
		result := getHint(t[0], t[1])
		fmt.Println("Result:", result)
	}
}
