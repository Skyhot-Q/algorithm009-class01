package main

import (
	"fmt"
)

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	for c := 0; c < n; c++ {
//		nums1[m+c] = nums2[c]
//	}
//
//	for c := 0; c < m+n; c++ {
//		for j := c+1; j < m+n; j++ {
//			if nums1[c] > nums1[j]{
//				temp := nums1[c]
//				nums1[c] = nums1[j]
//				nums1[j] = temp
//			}
//		}
//	}
//}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1
	for p2 >= 0 {
		if p1 >=0&&nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
}

type mergeTestData struct {
	nums1    []int
	nums1Len int
	nums2    []int
	nums2Len int
}

func main() {
	testCase := []mergeTestData{
		{nums1: []int{1, 2, 3, 0, 0, 0}, nums1Len: 3, nums2: []int{2, 5, 6}, nums2Len: 3},
		{nums1: []int{2,0}, nums1Len: 1, nums2: []int{1}, nums2Len: 1},
	}
	for _, t := range testCase {
		merge(t.nums1, t.nums1Len, t.nums2, t.nums2Len)
		fmt.Println("Result:", t.nums1)
	}
}
