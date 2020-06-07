package main

import (
	"fmt"
)

type flow struct {
	result []string
}
func (d *flow)generate(x []byte, left, right int){
	//T
	if len(x) == left + right {
		d.result = append(d.result, string(x[:]))
		return
	}
	//P
	if left < len(x)/2 {
		x[left+right] = '('
		d.generate(x, left+1, right)
	}

	if right < left  {
		x[left+right] = ')'
		d.generate(x, left, right+1)
	}
}


func generateParenthesis(n int) []string {
	d := flow{make([]string, 0)}

	d.generate(make([]byte, n*2), 0, 0)
	return d.result
}

func main() {
	x := generateParenthesis(3)
	fmt.Println(x)
}