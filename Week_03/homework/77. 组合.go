package main

import "fmt"

type Result struct {
	data [][]int
}

func (r *Result) search(i, n  ,k int, res []int) {
	//T
	if len(res) == k {
		tmp := make([]int, k, k)
		copy(tmp, res)
		r.data = append(r.data, tmp)
		return
	}

	//D
	for i := i+1; i <= n; i++ {
		res = append(res, i)
		r.search(i, n, k, res)
		//R
		res = res[:len(res)-1]
	}
}

func combine(n int, k int) [][]int {
	r := Result{make([][]int, 0)}
	r.search(0, n, k, []int{})
	return r.data
}

func main( ) {
	testCase := [][]int{
		//{4,2},
		{5,4},
	}
	for _,c := range testCase{
		r := combine(c[0],c[1])
		fmt.Println("Result:", r)
	}

}
