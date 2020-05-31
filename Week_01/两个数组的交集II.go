package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	for _, v := range nums1 {
		m1[v] ++
	}
	res := make([]int, 0)
	for _, v := range nums2 {
		if j, ok := m1[v]; ok && j != 0 {
			res = append(res, v)
			m1[v]--
		}
	}
	return res
}

func main() {
	testCase := [][][]int{
		{{1, 2, 2, 1}, {2, 2}},
		{{4, 9, 5}, {9, 4, 9, 8, 4}},
	}
	for _, c := range testCase {
		r := intersect(c[0], c[1])
		fmt.Println(r)
	}
}
