package main

import "fmt"

type Result struct {
	data []string
}

func (r *Result) search(s, digits string, i int, sMap map[uint8]string) {
	if i == len(digits) {
		r.data = append(r.data, s)
		return
	}
	letters, _ := sMap[digits[i]]
	for j := 0; j < len(letters); j++ {
		r.search(s+string(letters[j]), digits, i+1, sMap)
	}
}

func letterCombinations(digits string) []string {
	if len(digits) < 2 {
		return nil
	}
	r := Result{}
	sMap := map[uint8]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	r.search("", digits, 0, sMap)
	return r.data
}

func main() {
	testCase := []string{
		"23",
	}
	for _, c := range testCase{
		r := letterCombinations(c)
		fmt.Println(r)
	}
}
