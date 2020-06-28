package main

func search(nums []int, target int) int {
	//二分查找法，判断方式采用mid、左右边界对比，判断target所处那边进行查找
	nLen := len(nums)
	if nLen == 0 {
		return -1
	}
	left, right := 0, nLen-1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else {
			if target <= nums[right] && target >= nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
