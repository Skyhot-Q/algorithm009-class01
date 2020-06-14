package template

func twoSplit() {
	left, right = 0, len(array)-1
	for left <= right {
		mid = (left + right) / 2
		if array[mid] == target {
			// find the target
			break || return result
		} else if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
